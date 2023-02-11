# Kafka: настройка клиентов

## Kafka CLI

Создайте конфигурационный файл `$HOME/config.properties`. Для SASL-аутентификации с механизмом SCRAM поверх открытого текста.

```
bootstrap.servers=broker1:9094,broker2:9094,broker3:9094
security.protocol=SASL_PLAINTEXT
sasl.mechanism=SCRAM-SHA-512
sasl.jaas.config=org.apache.kafka.common.security.scram.ScramLoginModule required username="admin" password="keepinsecret";
```

## kafkactl

Создайте конфигурационный файл `$HOME/.config/kafkactl/config.yml`.

```yaml
contexts:
  dev:
    brokers:
      - broker-1:9092
      - broker-2:9092
      - broker-3:9092
    kafkaversion: 2.7.2
    producer:
      maxmessagebytes: 1000000
      partitioner: hash
      requiredacks: WaitForAll
    requesttimeout: 10s
    sasl:
      enabled: true
      mechanism: scram-sha512
      username: admin
      password: keepinsecret

current-context: dev
```