# Kafka: чтение и запись

## Прочитать топик Kafka

### Redpanda Console

Для просмотра сообщений в топике Kafka, необходимо найти его в секции Topics, а далее во вкладке Messages отобразить список сообщений. При нажатии на значок "плюс" вы можете также просмотреть структуру отдельного сообщения: key, headers и value.

### Kafka CLI

Подключитесь к брокеру Kafka и введите команду, чтобы прочитать топик с самого начала.

```
bin/kafka-console-consumer.sh --consumer.config $HOME/config.properties \
                              --bootstrap-server $(hostname):9092 \
                              --topic "topic" \
                              --from-beginning
```

Показать в выводе время, ключ и значение.

```
bin/kafka-console-consumer.sh --consumer.config $HOME/config.properties \
                              --bootstrap-server $(hostname):9092 \
                              --topic "topic" \
                              --from-beginning \
                              --formatter kafka.tools.DefaultMessageFormatter \
                              --property print.timestamp=true \
                              --property print.key=true \
                              --property print.value=true
```

Прочитать сообщения в составе консумер-группы consumer-group.

```
bin/kafka-console-consumer.sh --consumer.config $HOME/config.properties \
                              --bootstrap-server $(hostname):9092 \
                              --topic "topic" \
                              --group "consumer-group" \
                              --to-latest
```

## Записать в топик Kafka

### Kafka CLI
Подключитесь к брокеру Kafka и введите команду, чтобы записать данные в топик из stdin.

```
bin/kafka-console-producer.sh --producer.config $HOME/producer.properties \
                              --broker-list $(hostname):9092 \
                              --topic "$(hostname):9092"
```

Записать данные в формате ключ-значение с разделителем-двоеточием.

```
bin/kafka-console-producer.sh --producer.config $HOME/producer.properties \
                              --broker-list $(hostname):9092 \
                              --topic "$(hostname):9092" \
                              --property parse.key=true \
                              --property key.separator=:
```