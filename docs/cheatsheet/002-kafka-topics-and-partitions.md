# Kafka: управление топиками и партициями

## Создать топик Kafka

### Kafka CLI
Подключитесь к брокеру Kafka и выполните команду для создания топика.

```
bin/kafka-topics.sh --command-config $HOME/config.properties \
                    --topic "topic-name" \
                    --create \
                    --zookeeper localhost:2181/kafka \
                    --replication-factor 2 \
                    --partitions 10
```

## Изменить настройки топика Kafka

### Kafka CLI

Подключитесь к брокеру Kafka и примените конфигурацию для топика. В примере указание политики устаревания по времени (48 часов).

```
bin/kafka-topics.sh --command-config $HOME/config.properties \
                    --topic "topic-name" \
                    --alter \
                    --zookeeper localhost:2181/kafka \
                    --config retention.ms=$((48 * 60 * 60 * 1000))
```

Определить новое число партиций, например, с целью масштабирования до 4 штук.

```
bin/kafka-topics.sh --command-config $HOME/config.properties \
                    --topic "topic-name" \
                    --alter \
                    --zookeeper localhost:2181/kafka \
                    --partitions 4
```

Определить при этом брокеры для размещения партиций.

```
bin/kafka-topics.sh --command-config $HOME/config.properties \
                    --topic "topic-name" \
                    --alter \
                    --zookeeper localhost:2181/kafka \
                    --replica-assignment 0:1:2,0:1:2,0:1:2,2:1:0
                    --partitions 4
```

## Распечатать список топиков Kafka

### Kafka CLI

Подключитесь к брокеру Kafka и выполните следующую команду, чтобы увидеть список топиков со назначенной ему конфигурацией.

```
bin/kafka-topics.sh --command-config $HOME/config.properties \
                    --zookeeper localhost:2181/kafka \
                    --describe \
                    --topics-with-overrides
```

Без системных топиков (например, `__consumer_offsets`).

```
bin/kafka-topics.sh --command-config $HOME/config.properties \
                    --zookeeper localhost:2181/kafka \
                    --describe \
                    --topics-with-overrides \
                    --exclude-internal
```

Показать список топиков и его нереплицированные партиции.

```
bin/kafka-topics.sh --command-config $HOME/config.properties \
                    --zookeeper localhost:2181/kafka \
                    --describe \
                    --under-replicated-partitions
```

Показать список топиков и список партиций без активных лидеров.

```
bin/kafka-topics.sh --command-config $HOME/config.properties \
                    --zookeeper localhost:2181/kafka \
                    --describe \
                    --unavailable-partitions
```

## Посмотреть конфигурацию топика Kafka

### Kafka CLI

Подключитесь к брокеру Kafka и выполните следующую команду, чтобы увидеть конфигурацию топика.

```
bin/kafka-configs.sh --command-config $HOME/config.properties \
                     --zookeeper localhost:2181/kafka \
                     --describe \
                     --entity-type topics \
                     --entity-name "topic-name"
```

## Очистить топик Kafka

### Kafka CLI

В настоящий момент не существует команды, что удаляла бы сообщение в Kafka-топике. Для этого необходимо установить низкий retention, а после — вернуть значение на место. Подключитесь к брокеру Kafka и укажите retention размером в минуту (60000ms).

```
bin/kafka-configs.sh --command-config $HOME/config.properties \
                     --zookeeper localhost:2181/kafka
                     --alter \
                     --entity-type topics \
                     --entity-name "topic-name" \
                     --add-config retention.ms=60000
```

Подождите минуту, чтобы Kafka-брокеры успели удалить связанные сегменты данных, а после верните значение retention назад. Например, для 3 суток.

```
bin/kafka-configs.sh --command-config $HOME/config.properties \
                     --zookeeper localhost:2181/kafka
                     --alter \
                     --entity-type topics \
                     --entity-name "topic-name" \
                     --add-config retention.ms=$((72 * 60 * 60 * 1000))
```

## Удалить топик Kafka

Перед удалением необходимо убедиться, что в топике нет записи и нет чтения. Для этого в Redpanda UI во вкладке Consumers должно быть пусто, а во вкладке Messages не должно появляться новых сообщений.

### Kafka CLI
Подключитесь к брокеру, остановите консумер-группы и удалите топик.

```
bin/kafka-topics.sh --command-config $HOME/config.properties \
                    --topic "topic-name" \
                    --delete \
                    --bootstrap-server $(hostname):9094
```

## Распечатать оффсеты партиций топика Kafka

### Kafka CLI

Подключитесь к брокеру и выполните команду, чтобы увидеть список партиций в топике с оффсетами в формате `topic-name:partition-id:offset`.

```
bin/kafka-run-class.sh kafka.tools.GetOffsetShell \
                       --command-config $HOME/config.properties \
                       --broker-list $(hostname):9094 \
                       --topic "topic-name"
```

Показать самый крайний оффсет (latest).

```
bin/kafka-run-class.sh kafka.tools.GetOffsetShell \
                       --command-config $HOME/config.properties \
                       --broker-list $(hostname):9094 \
                       --topic "topic-name" \
                       --time -1
```

Показать самый ранний оффсет (earliest).

```
bin/kafka-run-class.sh kafka.tools.GetOffsetShell \
                       --command-config $HOME/config.properties \
                       --broker-list $(hostname):9094 \
                       --topic "topic-name" \
                       --time -2
```

Посмотреть оффсет партиции топика Kafka.

```
bin/kafka-run-class.sh kafka.tools.GetOffsetShell \
                       --command-config $HOME/config.properties \
                       --broker-list $(hostname):9094 \
                       --topic "topic-name" \
                       --partitions "partition-id1, partition-id2"
```

## Увеличить фактор репликации топика Kafka

### Kafka CLI

Подготовьте план переназначения партиции в формате JSON. В примере мы используем фактор репликации 3.

```
/tmp/reassignment.json
{
  "version": 1,
  "partitions": [
    { "topic": "topic-name", "partition": 0, "replicas": [0, 1, 2] },
    { "topic": "topic-name", "partition": 1, "replicas": [0, 1, 2] },
    { "topic": "topic-name", "partition": 2, "replicas": [0, 1, 2] }
  ]
}
```

Скопируйте конфигурацию на брокер и выполните следующую команду:

```
bin/kafka-reassign-partitions.sh --bootstrap-server $(hostname):9094 \
                                 --command-config $HOME/config.properties \
                                 --reassignment-json-file /tmp/reassignment.json \
                                 --execute
```

## Перенести партиции топика Kafka на другой брокер

### Kafka CLI

Подготовьте план переназначения партиции в формате JSON. В поле replicas укажите желаемые идентификаторы брокеров.

```
/tmp/reassignment.json
{
  "version": 1,
  "partitions": [
    { "topic": "topic-name", "partition": 0, "replicas": [0, 1] },
    { "topic": "topic-name", "partition": 1, "replicas": [1, 2] },
    { "topic": "topic-name", "partition": 2, "replicas": [0, 2] }
  ]
}
```

Скопируйте конфигурацию на брокер и выполните следующую команду.

```
bin/kafka-reassign-partitions.sh --bootstrap-server $(hostname):9094 \
                                 --command-config $HOME/config.properties \
                                 --reassignment-json-file /tmp/reassignment.json \
                                 --execute
```