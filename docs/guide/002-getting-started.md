# Начало работы

## Запуск кластера

Скопируйте конфигурационный файл:

```bash
cp .env.example .env
```

Запустите кластер:

```bash
docker-compose up -d
```

Если вы всё выполнили верно, то в списке запущенных контейнеров вы увидите Kafka, Zookeeper и Redpanda Console UI.

```bash
docker-compose ps
```

```text
NAME                IMAGE                                           COMMAND                   SERVICE             CREATED             STATUS              PORTS
grafana             docker.io/grafana/grafana:latest                "/run.sh"                 grafana             9 seconds ago       Up 7 seconds        0.0.0.0:3000->3000/tcp
kafka-1             kafka-workshop-kafka-1                          "/opt/bitnami/script…"    kafka-1             8 seconds ago       Up 7 seconds        9092/tcp
kafka-2             kafka-workshop-kafka-2                          "/opt/bitnami/script…"    kafka-2             8 seconds ago       Up 7 seconds        9092/tcp
kafka-3             kafka-workshop-kafka-3                          "/opt/bitnami/script…"    kafka-3             8 seconds ago       Up 7 seconds        9092/tcp
kafka-exporter      docker.io/bitnami/kafka-exporter:latest         "kafka_exporter --ka…"    kafka-exporter      9 seconds ago       Up 2 seconds        9308/tcp
prometheus          quay.io/prometheus/prometheus:latest            "/bin/prometheus --c…"    prometheus          9 seconds ago       Up 7 seconds        0.0.0.0:9090->9090/tcp
ui                  docker.redpanda.com/vectorized/console:v2.1.1   "/bin/sh -c 'echo \"$…"   ui                  8 seconds ago       Up 6 seconds        0.0.0.0:8080->8080/tcp
zookeeper           docker.io/bitnami/zookeeper:3.8                 "/opt/bitnami/script…"    zookeeper           9 seconds ago       Up 7 seconds        2181/tcp, 2888/tcp, 3888/tcp, 8080/tcp
```

## Состав кластера

Стенд Kafka предназначен для локальных экспериментов при изучении Apache Kafka и состоит из следующих групп программ:

| **Название**      | **Контейнеры** | **Описание**     |
| ----------------- | ---------------- | ---------------- |
| **Кластер**                      |
| Kafka             | `kafka-1`, `kafka-2`, `kafka-3` | Основной кластер |
| Zookeeper         | `zookeeper` | Координатор и менеджер кворума |
| Redpanda Console  | `ui` | UI для управления Kafka |
| **Приложения**                       |
| Консумер | `consumer-1` | Пример программы на Go для чтения данных из Kafka |
| Продюсер | `producer` | Пример программы на Go для чтения данных из Kafka |
| **Наблюдаемость**                    |
| Kafka Exporter | `kafka-exporter` | Экспортер метрик Kafka в формате PromQL |
| Prometheus | `prometheus` | Сервер метрик в формате PromQL |
| Grafana | `grafana` | Сервер визуализации метрик |

## Интерфейсы

Кластер представляет следующие публичные интерфейсы:

- [Prometheus UI](http://localhost:9090/) (`http://localhost:9090`)
- [Веб-интерфейс UI (Redpanda Console)](http://localhost:8080/) (`http://localhost:8080`)
- [Grafana UI](http://localhost:3000/) (`http://localhost:3000`)

## Доступ

Для демонстрации возможностей мы не используем авторизацию, за исключением Grafana.

* Имя пользователя: `admin`
* Пароль: `admin`

✅ Готово. Переходите к [созданию топика](./003-topics-and-partitions.md).
