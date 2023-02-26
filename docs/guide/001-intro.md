# Подготовка

Для работы вам необходимо:

- [Docker](https://www.docker.com/products/docker-desktop/) _(20.10.22 или старше)_
- [Docker Compose](https://docs.docker.com/compose/) _(v2.15.1 или старше)_
- Доступ в Интернет.

Мы тестировали курс только на Linux и Mac OS, но, скорее всего, команды без проблем заработают и под Windows.

Рекомендуем заранее загрузить Docker-образы на хорошем Интернет-подключении (1.3GB).

```bash
docker-compose --profile app pull
docker-compose --profile app build
```

_На соединении в 8Mbps вам потребуется около 20-25 минут на загрузку образов_.

---

✅ Готово. Переходите к [запуску кластера](./002-getting-started.md).