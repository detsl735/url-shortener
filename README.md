# url-shortener

[![Go](https://img.shields.io/badge/Go-1.25.4-blue.svg)]()
[![Gin](https://img.shields.io/badge/Gin-HTTP%20Framework-green.svg)]()
[![Redis](https://img.shields.io/badge/Redis-Storage-red.svg)]()

Лёгкий, быстрый и удобный сервис сокращения ссылок, написанный на Go.  
Использует Gin для HTTP, Redis как высокопроизводительное хранилище и base58 для генерации коротких человекочитаемых идентификаторов.

Проект спроектирован модульно:  
- handler/ — HTTP API
- shortener/ — логика генерации ключей
- store/ — интерфейсы и реализации хранилища
- main.go — инициализация и запуск сервера

---

## Функциональность

- Сокращение длинных URL с выдачей короткого идентификатора
- Перенаправление с короткого URL на оригинальный
- Redis как основное хранилище
- Безопасная генерация коротких ссылок через base58 (URL-friendly символы)
- Конфигурация через .env

---

## Технологии

| Компонент        | Библиотека                    |
|------------------|-----------------------------------------------|
| HTTP сервер      | github.com/gin-gonic/gin                    |
| Хранилище        | github.com/go-redis/redis/v8                |
| Генерация ключей | github.com/itchyny/base58-go                |
| Конфигурация     | github.com/joho/godotenv                    |
| Тестирование     | github.com/stretchr/testify                 |

## ⚙️ Конфигурация

Создайте `.env` в корне проекта:

```env
REDIS_PASSWORD=
REDIS_HOST=
HOST=
```

## Запуск
```bash
git clone https://github.com/Detsl735/url-shortener.git
cd url-shortener

go mod download
go run main.go
```
