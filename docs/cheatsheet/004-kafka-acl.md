# Kafka: управление доступами к топикам

## Создать нового пользователя

### Kafka CLI

Подключитесь к брокеру и выполните команду. Укажите в параметре password случайно сгенерированный пароль.

```
bin/kafka-configs.sh --alter \
                     --entity-type users \
                     --entity-name "john-doe" \
                     --bootstrap-server $(hostname):9092 \
                     --command-config /etc/kafka/config.properties \
                     --add-config 'SCRAM-SHA-256=[iterations=8192,password=keepinsecret],SCRAM-SHA-512=[password=keepinsecret]'
```

## Распечатать список пользователей

### Kafka CLI

Подключитесь к брокеру и выполните команду.

```
bin/kafka-configs.sh --describe \
                     --entity-type users \
                     --bootstrap-server $(hostname):9092 \
                     --command-config /etc/kafka/config.properties
```

## Удалить пользователя

### Kafka CLI

Подключитесь к брокеру и выполните команду для выдачи прав на чтение топика консумеру.

```
bin/kafka-configs.sh --delete \
                     --entity-type users \
                     --entity-name "john-doe" \
                     --bootstrap-server $(hostname):9092 \
                     --command-config /etc/kafka/config.properties
```

## Выдать права консумер-группе

### Kafka CLI

Подключитесь к брокеру и выполните команду для выдачи прав на чтение топика консумером.

```
bin/kafka-acls.sh --add \
                  --bootstrap-server $(hostname):9092 \
                  --allow-principal User:producer-user \
                  --topic "topic-name" \
                  --operation Write \
                  --command-config /etc/kafka/config.properties
```

## Выдать права продюсеру

### Kafka CLI

Подключитесь к брокеру и выполните команду для выдачи прав на запись в топик продюсеру.

```
bin/kafka-acls.sh --add \
                  --bootstrap-server $(hostname):9092 \
                  --allow-principal User:producer-user \
                  --topic "topic-name" \
                  --operation Write \
                  --command-config /etc/kafka/config.properties
```

## Распечатать список прав

### Kafka CLI

Подключитесь к брокеру и выполните команду.

```
bin/kafka-acls.sh --list \
                  --bootstrap-server $(hostname):9092 \
                  --command-config /etc/kafka/config.properties
```

## Забрать права у продюсера

### Kafka CLI

Подключитесь к брокеру и выполните команду, чтобы отобрать права.

```
bin/kafka-acls.sh --remove \
                  --bootstrap-server $(hostname):9092 \
                  --allow-principal User:producer-user \
                  --topic "topic-name" \
                  --operation Write \
                  --command-config /etc/kafka/config.properties
```

## Забрать права у консумера

### Kafka CLI

Подключитесь к брокеру и выполните команду, чтобы отобрать права.

```
bin/kafka-acls.sh --remove \
                  --bootstrap-server $(hostname):9092 \
                  --allow-principal User:consumer-user \
                  --topic "topic-name" \
                  --operation Read \
                  --command-config /etc/kafka/config.properties
```

## Удаление пользователя

### Kafka CLI

Чтобы удалить пользователя, достаточно забрать все права со всех топиков, которые ему назначены.