# Теория. Ментальная модель Kafka.

## Предисловие

В СберМаркете мы широко используем Kafka в качестве шины данных для микросервисов и знаем непонаслышке, что работать с Kafka нужно уметь. Воркшоп состоит из двух частей: теоретическая и практическая. В теоретической части мы:

- Обсудим сценарии использования
- Узнаем что такое консумер, продюсер и брокер
- Поймём как связаны топики, партиции и сегменты
- Посмотрим на формат сообщения Kafka
- Поймём зачем нужен лидер партиции и как реплицируются данные
- Зачем нужно партицирование
- Узнаем какие есть гарантии доставки сообщений
- Затронем идемпотентность обработки событий
- Определим что такое консумер-группа
- Посмотрим на ребалансировку консумеров

## Предпосылки к использованию

Долгое время инженеры разрабатывали программы, оперирующие объектами реального мира, сохраняя состояние о них в базы данных. Будь то, например, пользователи, заказы или товары. Представление о вещах мира как об объектах широко распространено среди ИТ-разработки (будь то парадигма ООП), однако сейчас больше компаний и технических команд всё чаще предпочитают думать не о самих объектах, а событиях, которые они порождают — то бишь, изменениях объектов во времени. 

Популярность событийно-ориентированного подхода вызвана стремлением компаний снизить связность сервисов друг с другом (что крайне важно при микросервисной трансформации) и улучшить устойчивость приложений к сбоям за счёт изоляции поставщиков данных и их потребителей.

События, проходящие в системах, как и объекты, также можно хранить в традиционных реляционных базах данных, однако это достаточно громоздко и неэффективно. Вместо этого мы используем структуру под названием лог.

## Устройство Apache Kafka

Лог — это упорядоченный поток событий во времени. Некоторое событие происходит и попадает всегда в конец лога, оставаясь неизменным.

Apache Kafka — это система по управлению такими логами и платформа, призванная соединить поставщиков данных и их потребителей, предоставив возможность получать упорядоченный поток событий в реальном времени.

### Продюсеры

Чтобы записать события в кластер Kafka, есть продюсеры. Продюсер — это приложение, которое вы разрабатываете. Программа-продюсер записывает сообщение в Kafka, а Kafka сохраняет события, возвращает подтверждение (acknowledgement) о записи, продюсер получает его и начинает следующую запись. И так далее.

### Брокеры

Сам же кластер Kafka состоит из брокеров. Представьте себе дата-центр и серверы в нём. В первом приближении думайте о Kafka-брокере как о компьютере: это процесс в операционной системе с доступом к своему локальному диску. Все брокеры соединены друг с другом сетью, действуя сообща как единый Kafka-кластер. Когда продюсеры пишут события в Kafka-кластер, они работают с брокерами в нём. 

В облачной среде Kafka-кластер не обязательно работает на выделенных серверах, а может быть виртуальными машинами или контейнерами в Kubernetes. Главное — каждый кластер Kafka состоит из брокеров.

### Консумеры

События, записанные продюсерами на локальные диски брокеров, могут быть прочитаны консумерами. Консумер — это также приложение, которое вы разрабатываете. В этом случае по-прежнему кластер Kafka — это нечто, обслуживаемое инфраструктурой, но что делаете вы как пользователь — пишете продюсер и консумер. Программа-консумер подписывается на события (поллит) и получает данные в ответ. И так по кругу.

Консумером может быть программа, подбирающая кандидата на основе координат партнёра, или при появлении заказа — инициирующая новую сборку. При этом консумер также может быть и продюсером, создавая новые события для размещения в Kafka для других сервисов.

Итого основы Kafka: продюсер, брокер и консумер.

## Архитектура Kafka

Итак, давайте посмотрим на архитектуру Kafka внимательнее. Слева есть продюсеры, в середине брокеры, а справа — консумеры. Kafka же представляет собой группу брокеров, связанных с Zookeeper-кворумом. Kafka использует Zookeeper для достижения консенсуса состояния в распределённой системе: есть несколько вещей, с которыми должен быть «согласен» каждый брокер и Zookeeper помогает достичь этого «согласия» внутри кластера.

> _Начиная с Kafka 3.4 необходимость в использовании Zookeeper отпала: для арбитража используется собственный протокол KRaft, решающий те же задачи, но на уровне брокеров, однако для простоты мы пока остановимся на традиционной схеме_.

Так вот, Zookeeper представляет собой выделенный кластер серверов для образования кворума и поддержки внутренних процессов Kafka. Благодаря Zookeeper, кластером Kafka можно управлять: добавлять пользователей, топики и задавать им настройки. 

Zookeeper также помогает при обнаружении сбоя в мастере, провести выборы нового и сохранить работоспособность кластера Kafka. И, что немаловажно, Zookeeper хранит в себе все авторизационные данные и ограничения (Access Control Lists) при работе консумеров и продюсеров с брокерами.

Подводя промежуточный итог, кластер Kafka позволяет изолировать консумеры и продюсеры друг от друга. Продюсер ничего не знает о консумерах при записи данных в брокер, а консумер — ничего не знает о продюсере данных. 

Скажем, если консумер станет работать медленнее продюсера, то это никак не влияет на запись этих событий. То же и с добавлением, удалением или даже сбоем: изменения консумеров не оказывают никакого влияния на продюсеры.

## Устройство брокеров

Теперь поговорим отдельно о брокерах. Наверняка вы несколько раз слышали про какие-то топики, теперь коротко о том, что это такое. События в Kafka-брокерах образуют топики. 

### Топики

Топик — это логическое представление категорий сообщений в группы. Например, события по статусам заказов, координат партнёров, маршрутных листов и так далее. 

Ключевое слово здесь — **логическое**. Мы создаём топики для событий общей группы и стараемся не смешивать их друг с другом. Например, координаты партнёров не должны быть в том же топике, что и статусы заказов, а обновлённые статусы по заказам — не быть вперемешку с обновлением регистраций пользователей. 

О топике стоит думать как о логе: вы пишете событие в конец, не разрушая при этом старые события. Один продюсер может писать в один или несколько топиков, в один топик могут писать один или более продюсеров, как и много консумеров могут читать из одного топика, как и один консумер может читать сразу несколько топиков.

Теоретически, нет никаких ограничений на число этих топиков, но практически это ограничено числом партиций, о которых мы поговорим позднее.

### Партиции и сегменты

Топиков в кластере Kafka может быть много и нет ограничений на их создание. Однако рано или поздно, компьютер, выполняя операции на процессоре и вводе-выводе, упирается в свой предел. К сожалению, мы не можем увеличивать мощность и производительность компьютеров бесконечно, а потому топик следует делить на части.

В Kafka эти части называются партиции. Каждый Kafka-топик состоит из одной или более партиций, каждая из которых может быть размещена на разных брокерах. Это то, благодаря чему Kafka может масштабироваться: пользователь может создать топик, разделить его на партиции и разместить каждую из них на своём брокере.

Формально партиция — это и есть строго упорядоченный лог сообщений. Каждое сообщение в партиции добавлено в конец без возможности изменить его в будущем и как-то повлиять на уже записанные сообщения. При этом сам топик вцелом не имеет никакого порядка, но порядок сообщений всегда есть в одной из его партиций.

Партиции же представлены физически на дисках в виде сегментов — отдельных файлов, что могут быть созданы, отротированы или удалены в соответствии с настройкой устаревания данных в них. Обычно, если вы не администрируете кластер Kafka, вам не приходится много думать о сегментах партиций, но вы должны помнить о партициях, как о модели хранения данных в топиках Kafka.



