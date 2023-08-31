# REST API Avito-proj

### Для запуска приложения через docker-compose:

1. Добавить .env файл в корень проекта с полем 

```
DB_PASSWORD = "MySuperPassword"
```

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
в корне проекта avito-proj-db-tables.sql
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

Примеры запросов: 

GET http://localhost:8000/api/sections/users/3 
Ответ: [
    "NEW_SECTION_TEST_2",
    "NEW_SECTION_TEST_10001",
    "NEW_SECTION_TEST_3"
]

GET http://localhost:8000/api/users/6/history
body { "Date": "2023-08" }
Ответ: [
    {
        "UserId": 6,
        "SectionId": 18,
        "OperationType": "addition",
        "OperationDate": "2023-08-31T18:10:52.565604Z",
        "SectionName": "NEW_SECTION_TEST_99"
    },
    {
        "UserId": 6,
        "SectionId": 18,
        "OperationType": "deletion",
        "OperationDate": "2023-08-31T18:14:34.639106Z",
        "SectionName": "NEW_SECTION_TEST_99"
    }
]

PUT http://localhost:8000/api/sections/users/3
body {
    "SectionsAdd": ["new-section-test-3"],
    "SectionsDelete": []
}
Ответ: {
    "message": "Success"
}

DELETE http://localhost:8000/api/sections/
body{
    "Slug": 999
}
Ответ: {"message":"Incorrect input"}
### О приложении

В приложении реализованы 5 методов:
1. Создание сегмента
2. Удаление сегмента
3. Добавление сегментов пользователю
4. Получение сегментов пользователя 
5. Создание пользователя
6. Получение истории добавления/удаления сегментов у пользователя (без ссылки на csv файл)

Добавлена небольшая проверка (поле не может быть пустым) при создании сегмента или пользователя

При попытке добавить или удалить несуществующий сегмент у пользователя будет возвращено сообщение об ошибке, при этом никаких изменений в БД соответственно не произойдет

Был добавлен swagger файл. Перейти на страницу swaggera можно по адресу 
http://localhost:8000/swagger/index.html

Наполовину сделано доп. задание 1 на получение истории добавления/удаления сегментов у пользователя

Автор: Кузнецов Роман
Email: roman.ksv1@mail.ru
