# API-сервис для вопросов и ответов

REST API сервис для создания вопросов и ответов к ним. 

## Основной функционал
- Создание, получение и удаление вопросов
- Добавление ответов к вопросам. Получение и удаление ответа по ID
- Получение вопроса со всеми его ответами
- Валидация входящих данных через swagger
- Логирование через обёртку `slog`
- Работа с PostgreSQL через `GORM`, миграции через `goose`
- `docker-compose` для запуска локально

## Запуск
### Через Docker:
1. Установить docker
2. В корне создать конфиг `.env` и заполнить (пример в `.env-example`):
```dotenv
POSTGRES_HOST=postgres
POSTGRES_DB=questions
POSTGRES_USER=questions
POSTGRES_PASSWORD=0000
POSTGRES_PORT=5432
```
В `docker-compose.yml` эти переменные пробрасываются в сервис `api`

3. Сгенерировать swagger:
```shell
make gen-swagger
```

4. Запустить:
```shell
docker-compose up -d
```
5. Проверить:
```shell
docker-compose ps
docker-compose logs
```
6. Сервис будет доступен по `http://localhost:8080`

### Через Makefile:
1. Создать БД в PostgreSQL:
```sql
CREATE DATABASE questions;
```
2. Настроить окружение:
- `.env`:
```dotenv
POSTGRES_HOST=localhost
POSTGRES_DB=questions
POSTGRES_USER=questions
POSTGRES_PASSWORD=0000
POSTGRES_PORT=5432
```

- `config/local.json` (или `main.json`):
```json
{
  "mode": "local",
  "httpServer": {
    "port": 8080,
    "readTimeout": "5s",
    "writeTimeout": "5s",
    "maxHeaderBytes": 1048576
  },
  "handler": {
    "requestTimeout": "5s",
    "streamTimeout": "5s"
  },
  "service": {}
}
```

3. Сгенерировать swagger:
```shell
make gen-swagger
```

4. Поднять PostgreSQL и выполнить миграции:
```shell
make postges.start
make migrations.up
```

5. Запустить app:
```shell
make run
```

## API endpoints

Базовый путь: `/`.
Схема и модели в `api/swagger/swagger.yml`.

| Questions | Path            | INFO                                                             |
|-----------|-----------------|------------------------------------------------------------------|
| `GET`     | /questions      | Получить список всех вопросов                                    |
| `POST`    | /questions      | Создать новый вопрос                                             |
| `GET`     | /questions/{id} | Получить вопрос по ID со всеми его ответами                      |
| `DELETE`  | /questions/{id} | Удалить вопрос по ID                                             |

| Answers  | Path                    | INFO                     |
|----------|-------------------------|--------------------------|
| `POST`   | /questions/{id}/answers | Добавить ответ к вопросу |
| `GET`    | /answers/{id}           | Получить ответ по ID     |
| `DELETE` | /answers/{id}           | Удалить ответ по ID      |

Подробнее:
1. __Questions__
- `GET /questions`

Получить список всех вопросов. 

- `POST /questions`

Тело запроса (`application/json`):
```json
{
  "text": "Текст вопроса"
}
```
Успешный ответ `200 OK`/`201 Created`
```json
{
  "id": 1,
  "text": "Текст вопроса",
  "created_at": "2025-01-01T12:00:00Z"
}
```

- `GET /questions/{id}`

Успешный ответ `200 OK`:
```json
{
  "id": 1,
  "text": "Текст вопроса",
  "created_at": "2025-01-01T12:00:00Z",
  "answers": [
    {
      "id": 10,
      "question_id": 1,
      "user_id": "uuid",
      "text": "Ответ на вопрос",
      "created_at": "2025-01-01T13:00:00Z"
    }
  ]
}
```

- `DELETE /questions/{id}`

Успешный ответ: `204 No Content`

2. __Answers__

- `POST /questions/{id}/answers`

Обязательный заголовок: `useruuid: <UUID юзера`.

Тело запроса:
```json
{
  "text": "Текст ответа",
  "user_id": "UUID юзера"
}

```

Успешный ответ `200 OK`/`201 Created`: 
```json
{
  "id": 10,
  "question_id": 1,
  "user_id": "uuid",
  "text": "Текст ответа",
  "created_at": "2025-01-01T13:00:00Z"
}
```

- `GET /answers/{id}`

- `DELETE /answers/{id}`
Обязательный заголовок: `useruuid: <UUID юзера`.

Только владелец ответа может его удалить! 

3. __Формат ошибок__

`ErrorResponse`:
```json
{
  "code": 400,
  "message": "описание ошибки",
  "detail": "детали ошибки"
}
```

4. __Коды ошибок__
- `400 Bad Request`– неверные параметры или тело запроса, ошибки валидации.
- `403 Forbidden` – попытка удалить ответ (не владельцем), либо отсутствие заголовка useruuid.
- `404 Not Found` – вопрос/ответ не найден.
- `500 Internal Server Error` – внутренняя ошибка сервера/БД.

## Архитектура
### Слои:
- `cmd/app` - точка входа
- `internal/`
  - `config` - загрузка конфига из JSON и `.env`
  - `infrastructure/db/postgres` - подключение к БД через GORM
  - `entity` - сущности и маппинг в swagger-модели
  - `repository` - репы
  - `usecase` - бизнес-логика
  - `handler/http` - HTTP-слой, роутер, хэндлеры
- `apigen/swagger` - swagger и сгенерированные модели
- `migrations` - миграции
- `pkg/`
  - `logger` - инициализация логгера
  - `utils` - вспомогательные функции (`ToPtr`, `FromPtr`)

### Модели данных

1. __Questions__ - вопрос

- `id`:_int64_ - id вопроса
- `text`:_string_ = текст вопроса
- `created_at`:_datetime_ - дата создания вопроса

2. __Answers__ - ответ на вопрос
- `id`:_int64_ - id ответа
- `question_id`:_int64_ - связь с вопросом по id
- `user_id`:_string_ - uuid юзера
- `text`:_string_ - текст ответа
- `created_at`:_datetime_ - дата создания ответа


## Стек
- Language: Go 1.25
- Web: `net/http`, `gorilla/mux`
- ORM: GORM
- Migrations: `goose`
- swagger
- DB: PostgreSQL
- Tests: `testing`, `gomock`, `testify`
- Logs: `tint` + `slog`
- Configs: `viper`, `.env`, JSON-config