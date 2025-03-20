# Введение
По материалам [Мониторинг приложения Golang с Prometheus, Grafana, New Relic и Sentry](https://nuancesprog.ru/p/25427/)
Исходник: https://github.com/mertcakmak2/golang-new-relic-sentry-prometheus/blob/main/mocks/mockUserUsecase.go

Цели:
- В промежуточном ПО с помощью идентификатора запросов отправить в New Relic журнал запросов и журнал ответов.
- Отправить в Sentry с помощью идентификатора запросов ошибку, которой перехвачено исключение.
- Имитация базы данных, репозитория, варианта применения.
- Модульный тест.
- Мониторинг службы с Prometheus и Grafana.
- CRUD-операции.

Технологический стек:
- Golang;
- Prometheus;
- Grafana;
- New Relic;
- Sentry.

# Установка

Запуск `docker-compose`:
```bash
docker-compose up -d
```

Установка пакетов:
```bash
go get github.com/gin-gonic/gin
go get github.com/prometheus/client_golang/prometheus
go get github.com/prometheus/client_golang/prometheus/promauto
go get github.com/prometheus/client_golang/prometheus/promhttp
go get github.com/getsentry/sentry-go
go get github.com/newrelic/go-agent/v3/integrations/nrgin
go get go.uber.org/zap
go get gorm.io/driver/postgres
go get gorm.io/gorm
go get github.com/sethvargo/go-envconfig
go get github.com/stretchr/testify
```

Установка генератора моков:
```bash
go install github.com/golang/mock/mockgen@v1.6.0
```

Генерация моков:
```bash
go generate .
```
или
```bash
go generate ./..
```
В папке `mocks` должен появиться результат - файлы `mockUserRepository.go` и `mockUserUsecase.go`


Запуск тестов:
```bash
go test -v .
```

Чтобы посмотреть тестовое покрытие, запускаем такие команды:
```bash
go test -covermode=count -coverpkg=./... -coverprofile coverage.out -v ./...
go tool cover -html coverage.out
```

Файл `coverage.out` с тестовым покрытием генерируется в корневом пути и открывается в браузере.

