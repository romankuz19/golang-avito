# REST API Avito-proj

### Для запуска приложения через docker-compose:

1. Добавить .env файл в корень проекта с полем DB_PASSWORD = "MySuperPassword"

2. Выполнить команду

```
docker-compose up --build avito-proj 
```

3. Подождать пока инициализируется контейнер с БД, затем выполнить

``` 
docker-compose up avito-proj
```

4. Если приложение запускается впервые, необходимо применить миграционный файл к базе данных:

```
avito-proj-db-tables.sql
```

### Альтернативный способ запуска приложения

1. Изменить host: "db" на host: "localhost" в файле configs/config.yml
2. Поднять базу в соответствии с параметрами config.yml
3. Добавить .env файл в корень проекта с полем DB_PASSWORD = "MySuperPassword"
4. Применить миграционный файл к БД avito-proj-db-tables.sql
5. Выполнить команды go mod tidy && go run cmd/main.go


### Примеры запросов/ответов

1. Можно посмотреть с помощью swagger'a, запустив приложение и перейдя по адресу 
http://localhost:8000/swagger/index.html
2. Можно импортнуть файл AvitoTechBackendTrainee.postman_collection.json из корня проекта  в Postman


### О приложении

В приложении реализованы 5 методов:
1. Создание сегмента
2. Удаление сегмента
3. Добавление сегментов пользователю
4. Получение сегментов пользователя 
5. Создание пользователя

Добавлена небольшая проверка (поле не может быть пустым) при создании сегмента или пользователя

При попытке добавить или удалить несуществующий сегмент у пользователя будет возвращено сообщение об ошибке, при этом никаких изменений в БД соответственно не произойдет

Был добавлен swagger файл. Перейти на страницу swaggera можно по адресу 
http://localhost:8000/swagger/index.html

Автор: Кузнецов Роман
Email: roman.ksv1@mail.ru