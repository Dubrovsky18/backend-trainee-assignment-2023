# Segments Service
Segments Service представляет собой веб-сервис, который позволяет управлять сегментами пользователей и получать историю изменений.

## Запуск
### Требования
Docker 

Docker Compose

## Шаги
1. #### Склонируйте репозиторий:


    git clone https://github.com/Dubrovsky18/backend-trainee-assignment-2023.git

    cd segments-service

2. #### Создайте файл .env в корневой папке проекта и укажите следующие переменные окружения:


    POSTGRES_HOST=localhost
    POSTGRES_PORT=2345
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=postgres
    POSTGRES_NAME=postgres
    
    HTTP_SERVER_ADDR=":4443"
    GIN_MODE=release
    
    SERVICE_NAME=segments

3. #### Запустите службу с помощью Docker Compose:

    
    docker-compose up --build


Сервис будет доступен по адресу http://localhost:4443

## Методы

### Создать сегмент
POST /api/v1/slug/create

Создает новый сегмент.

Пример запроса:


    curl -X POST -H "Content-Type: application/json" -d '{"name_slug": "new-segment"}' http://localhost:4443/api/v1/slug/create


### Удалить сегмент
DELETE /api/v1/slug/delete/{name_slug}

Удаляет указанный сегмент.

Пример запроса:


    curl -X DELETE http://localhost:4443/api/v1/slug/delete/segment

### Получить все сегменты
GET /api/v1/slug/get_all

Получает список всех сегментов.

Пример запроса:


    curl http://localhost:4443/api/v1/slug/get_all

### Создать пользователя
POST /api/v1/user/create/{user_id}

Создает нового пользователя с указанным идентификатором.

Пример запроса:


    curl -X POST -H "Content-Type: application/json" -d '{"id": 123, "slugs": ["segment-1", "segment-2"]}' http://localhost:4443/api/v1/user/create/123

### Удалить пользователя
DELETE /user/delete/{user_id}

Удаляет пользователя с указанным идентификатором.

Пример запроса:


        curl -X DELETE http://localhost:4443/api/v1/user/delete/123

### Получить сегменты пользователя
GET /api/v1/user/get_slugs/{user_id}

Получает список сегментов пользователя.

Пример запроса:


        curl http://localhost:4443/api/v1/user/get_slugs/123

### Добавить/удалить сегменты у пользователя
POST /api/v1/user/add_del_slug/{user_id}

Добавляет или удаляет сегменты у пользователя.

Пример запроса:

        curl -X POST -H "Content-Type: application/json" -d '{"add_segments": ["new-segment"], "remove_segments": ["old-segment"]}' http://localhost:4443/api/v1/user/add_del_slug/123

### Получить историю сегментов пользователя
POST /api/v1/user/extra/history/{user_id}

Получает историю изменения сегментов пользователя за указанный период.

Пример запроса:


    curl -X POST -H "Content-Type: application/json" -d '{"user_id": 123, "year_start": 2023, "year_finish": 2023, "month_start": 8, "month_finish": 8}' http://localhost:4443/api/v1/user/extra/history/123


