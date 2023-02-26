# Kafka: управление консумер-группами

## Создать новую консумер-группу Kafka

Консумер-группы не требуется создавать явно. Соответствующая группа будет создана автоматически при верной настройке консумера на работу с группами.

## Распечатать список консумер-групп Kafka

### Kafka CLI

Подключитесь к брокеру Kafka и выполните команду, чтобы вывести на экран список консумер-групп.

```
bin/kafka-consumer-groups.sh --command-config $HOME/config.properties \
                             --bootstrap-server $(hostname):9092 \
                             --list
```

## Посмотреть конфигурацию консумер-группы Kafka

### Kafka CLI

Подключитесь к брокеру Kafka и выполните команду, чтобы вывести на экран конфигурацию выбранной консумер-группы и её состояния.

```
bin/kafka-consumer-groups.sh --command-config $HOME/config.properties \
                             --bootstrap-server $(hostname):9092 \
                             --group "group-name" \
                             --state \
                             --describe
```

## Изменить или удалить консумер-группы Kafka

Любое изменение состава или конфигурации консумер-группы проводится с полным выходом участников из неё. Обратите внимание, что в момент остановки и запуска приложения данные не будут прочитаны.

Для изменения:

- Остановите приложение.
- Измените консумер-группу.
- Запустите приложение.

### Kafka CLI

Переместить оффсет для консумер-группы. Например, на 1 января 2022 года на 00:00 MSK.

```
bin/kafka-consumer-groups.sh --command-config $HOME/config.properties \
                             --bootstrap-server $(hostname):9092 \
                             --topic "topic-name" \
                             --group "group-name" \
                             --reset-offsets \
                             --to-datetime 2022-01-01T00:00:00.000+0300 \
                             --execute
```

Переместить оффсет всех партиций на пять записей вперёд.

```
bin/kafka-consumer-groups.sh --command-config $HOME/config.properties \
                             --bootstrap-server $(hostname):9092 \
                             --topic "topic-name:0" \
                             --group "group-name" \
                             --reset-offsets \
                             --shift-to 5 \
                             --execute
```

Переместить оффсет партиции №0 на самый ранний (начать читать с начала).

```
bin/kafka-consumer-groups.sh --command-config $HOME/config.properties \
                             --bootstrap-server $(hostname):9092 \
                             --topic "topic-name:0" \
                             --group "group-name" \
                             --reset-offsets \
                             --to-earliest \
                             --execute
```

Переместить оффсет для партиций №3 и №4 на самый крайний (пропустить все сообщения и начать с конца).

```
bin/kafka-consumer-groups.sh --command-config $HOME/config.properties \
                             --bootstrap-server $(hostname):9092 \
                             --topic "topic-name:3,4" \
                             --group "group-name" \
                             --reset-offsets \
                             --to-latest \
                             --execute
```

Распечатать текущую позицию оффсетов консумер-группы. Предварительно необходимо остановить приложение-консумер.

```
bin/kafka-consumer-groups.sh --command-config $HOME/config.properties \
                             --bootstrap-server $(hostname):9092 \
                             --topic "topic-name" \
                             --group "group-name" \
                             --reset-offsets \
                             --to-current \
                             --export \
                             --execute
```

Восстановить оффсеты консумер-группы из CSV-файла offsets.csv.

```
bin/kafka-consumer-groups.sh --command-config $HOME/config.properties \
                             --bootstrap-server $(hostname):9092 \
                             --topic "topic-name" \
                             --group "group-name" \
                             --reset-offsets \
                             --from-file offsets.csv \
                             --execute
```

Удалить оффсет консумер-группы

```
bin/kafka-consumer-groups.sh --command-config $HOME/config.properties \
                             --bootstrap-server $(hostname):9092 \
                             --topic "topic-name" \
                             --group "group-name" \
                             --delete-offsets
```

Удалить консумер-группу.

```
bin/kafka-consumer-groups.sh --command-config $HOME/config.properties \
                             --bootstrap-server $(hostname):9092 \
                             --group "group-name" \
                             --delete
```