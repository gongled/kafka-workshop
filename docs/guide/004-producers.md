# Работа продюсера

Мы заранее подготовили пример программы-продюсера на Go. Для простоты, представьте себе приложение, что собирает координаты сборщиков и курьеров СберМаркет (кандидатов), а затем отправляет их в один топик `example` для будущих потребителей. Продюсер записывает сообщения в формате JSON, передавая идентификатор кандидата (`id`), широту (`lat`) и долготу (`lon`).

```json
{
  "id": 1,
  "lat": -17.000132,
  "lon": 28.587008
}
```

Посмотрите пример записываемых событий:

```
docker-compose --profile app logs -f producer
```

```
producer-1  | 2023/02/25 18:51:36 message written at topic example: 6 = {"id":6,"lat":-56.998822,"lon":-120.432007}
producer-1  | 2023/02/25 18:51:38 message written at topic example: 1 = {"id":1,"lat":-46.10566,"lon":-21.857902}
producer-1  | 2023/02/25 18:51:40 message written at topic example: 3 = {"id":3,"lat":73.333584,"lon":-113.570201}
producer-1  | 2023/02/25 18:51:42 message written at topic example: 1 = {"id":1,"lat":40.295088,"lon":-78.754235}
producer-1  | 2023/02/25 18:51:44 message written at topic example: 5 = {"id":5,"lat":-16.964839,"lon":40.875567}
producer-1  | 2023/02/25 18:51:46 message written at topic example: 5 = {"id":5,"lat":37.125313,"lon":-99.247528}
producer-1  | 2023/02/25 18:51:48 message written at topic example: 4 = {"id":4,"lat":43.214251,"lon":116.305726}
producer-1  | 2023/02/25 18:51:50 message written at topic example: 1 = {"id":1,"lat":55.444785,"lon":170.116819}
producer-1  | 2023/02/25 18:51:52 message written at topic example: 3 = {"id":3,"lat":83.039649,"lon":-136.749644}
producer-1  | 2023/02/25 18:51:54 message written at topic example: 3 = {"id":3,"lat":-17.167125,"lon":157.868446}
producer-1  | 2023/02/25 18:51:56 message written at topic example: 4 = {"id":4,"lat":81.208194,"lon":121.136919}
```

Откройте [исходный код продюсера](../../examples/producer/main.go). Обратите внимание на настройки `kafka.Writer`:

```go
w := &kafka.Writer{
    Addr:     kafka.TCP(addrs...),
    Topic:    topic,
    Balancer: &kafka.Hash{},
}
```

Опция `Balancer` определяет стратегию балансировки событий между партициями. В нашем примере мы используем хеш-функцию от передаваемого ключа (`key`). В качестве ключа используем идентификатор кандидата, чтобы события одного и того же пользователя сохранялись в верном порядке.

```go
payload := kafka.Message{
    Key:   []byte(strconv.Itoa(candidate_id)),
    Value: []byte(msg),
}

err = w.WriteMessages(context.Background(), payload)

if err != nil {
    log.Fatal("failed to write messages:", err)
}
```

Откройте [веб-интерфейс Redpanda Console](http://localhost:8080/topics/example#messages) и посмотрите содержимое любого сообщения. Убедитесь, что поля ключ и тело заполнены так, как ожидается.

---

✅ Готово. Переходите к [работе с консумерами](./005-consumers.md).