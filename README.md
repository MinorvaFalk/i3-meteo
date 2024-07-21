# I3 Meteo Project Test

Aplikasi terdiri dari **API** dan **Scheduling service** yang berfungsi untuk mengambil data dari [Meteosource](https://www.meteosource.com). Untuk API sendiri menggunakkan framework [gin](https://github.com/gin-gonic/gin), sedangkan untuk scheduler menggunakan bantuan library [cron](https://github.com/robfig/cron).

## Application Flow

![app diagram](docs/diagram.png)

## Requirements

- Go >= 1.22.2
- Docker
- Make

## How to Run

Environment variables

```bash
ENV=development|production
PORT=8080
DSN=postgresql://user:password@host:post/database?sslmode=disable
REDIS_ADDR=host:port
# empty if not using redis auth
REDIS_USERNAME=
REDIS_PASSWORD=
# refer https://www.meteosource.com/ dashboard
METEOSOURCE_API_KEY=
# change url if you are using paid account
METEOSOURCE_URL=https://www.meteosource.com/api/v1/free
```

Running in docker compose
```bash
# using make
make compose

# or

docker compose -p i3-meteo -f infra/docker-compose.yml up -d
```

Running instances manually

```bash
# using makefile

make postgres
make redis
make migrate-up
make run-api
```