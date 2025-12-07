Запустите сервис: go run cmd/server/main.go и убедитесь по логам в терминале, что сервис запустился
```
2025/12/08 01:22:38 [app] store initialized successfully
2025/12/08 01:22:38 [app] service initialized successfully
2025/12/08 01:22:38 [app] router initialized successfully
2025/12/08 01:22:38 [app] server running on :8484
```
Введите в адресной строке Postman адрес в соответствии с Вашим cfg (например http://localhost:8484)

1. POST /create_event
```
Method: POST
Body → raw → JSON:
{
  "user_id": 1,
  "title": "Meeting",
  "notice": "Discuss project",
  "date": "2025-12-08T10:00:00Z"
}
```
Ожидаемый результат:
```
{
    "id": 1,
    "result": "event created"
}
```
2. POST /update_event
```
Method: POST
Body → raw → JSON:
{
  "event_id": 1,
  "user_id": 1,
  "title": "Updated Meeting",
  "notice": "Updated notice",
  "date": "2025-12-08T11:00:00Z"
}
```
Ожидаемый результат
```
{
    "result": "event updated"
}
```
3. POST /delete_event
```
Method: POST
Body → raw → JSON:
{
  "event_id": 1
}
```
Ожидаемый результат
```
{
  "result": "event deleted"
}
```
4. GET /events_for_day
URL: http://localhost:8484/events_for_day?user_id=1&date=2025-12-08
```
Method: GET
Ожидаемый результат: массив событий пользователя на указанный день
{
    "result": [
        {
            "event_id": 3,
            "user_id": 1,
            "title": "dfdfddffd",
            "notice": "Discuss project",
            "date": "2025-12-08T10:20:00Z"
        },
        {
            "event_id": 4,
            "user_id": 1,
            "title": "dfd",
            "notice": "Discuss project",
            "date": "2025-12-08T10:22:00Z"
        }
    ]
}
```
5. GET /events_for_week
http://localhost:8484/events_for_week?user_id=1&date=2025-12-08
```
Method: GET
Ожидаемый результат: массив событий пользователя на указанную неделю
{
    "result": [
        {
            "event_id": 3,
            "user_id": 1,
            "title": "dfdfddffd",
            "notice": "Discuss project",
            "date": "2025-12-08T10:20:00Z"
        },
        {
            "event_id": 4,
            "user_id": 1,
            "title": "dfd",
            "notice": "Discuss project",
            "date": "2025-12-08T10:22:00Z"
        },
        {
            "event_id": 5,
            "user_id": 1,
            "title": "fuck cska",
            "notice": "Discuss ",
            "date": "2025-12-09T10:22:00Z"
        },
        {
            "event_id": 6,
            "user_id": 1,
            "title": "fuck zenit",
            "notice": "football ",
            "date": "2025-12-10T11:00:00Z"
        }
    ]
}
```
6. GET /events_for_month
http://localhost:8484/events_for_month?user_id=1&date=2025-12-08
```
Method: GET
Ожидаемый результат: массив событий пользователя на месяц
{
    "result": [
        {
            "event_id": 3,
            "user_id": 1,
            "title": "dfdfddffd",
            "notice": "Discuss project",
            "date": "2025-12-08T10:20:00Z"
        },
        {
            "event_id": 4,
            "user_id": 1,
            "title": "dfd",
            "notice": "Discuss project",
            "date": "2025-12-08T10:22:00Z"
        },
        {
            "event_id": 5,
            "user_id": 1,
            "title": "fuck cska",
            "notice": "Discuss ",
            "date": "2025-12-09T10:22:00Z"
        },
        {
            "event_id": 6,
            "user_id": 1,
            "title": "fuck zenit",
            "notice": "football ",
            "date": "2025-12-10T11:00:00Z"
        }
    ]
}
```
