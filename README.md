# Events_Calendar
HTTP-сервис для работы с календарём событий.
Реализован на Go с использованием чистой архитектуры, маршрутизатора ginext, in-memory хранилища, фоновых воркеров и асинхронного логирования.

ТЗ на проект (TASK.md) находится в корневой директории

## Возможности
Создание, обновление и удаление событий  
Получение событий:  
на день  
на неделю  
на месяц  
Асинхронное логирование запросов  
Фоновая очистка старых событий  
Хранение данных в памяти (in-memory)  

## Структура проекта
```
├── cmd/server/main.go  Точка входа 
├── internal/
│ ├── app/              Инициализация зависимостей 
│ ├── handlers/         HTTP-хендлеры 
│ ├── middleware/       Middleware (логирование) 
│ ├── model/            Модели данных 
│ ├── service/          Бизнес-логика 
│ ├── storage/          Абстракция хранилища 
│ │ └── inmemory/       Реализация InMemory 
│ └── worker/           Фоновый воркер
├── go.mod 
├── POSTMAN_TESTS.md    Кейс тестов для Postman
├── TASK.md             ТЗ
└── README.md 
```

## Технологии
- Go 1.25  
- Gin (для HTTP API между узлами)  
- Стандартная библиотека Go  
- Структурированное логирование  
- Асинхронный логгер  

## Подготовка
Скопировать репозиторий, выполнив команду в терминале:
```
 git clone https://github.com/Vladimirmoscow84/Events_Calendar.git
 ```
Настроить .env файл (указать порт)
Запустить сервер
```
go run cmd/server/main.go
```
По логам убедиться, что сервер запустился на настроенном хосте
```
[app] store initialized successfully
[app] service initialized successfully
[app] router initialized successfully
[app] server running on :(ваш порт)
```
В браузере перейти на страницу проекта.

Пимер работы и кейс тестов находятся в корневой директории в файле POSTMAN_TESTS.md

Автор  
Разработчик: Vladimirmoscow84  
Контакт: ccr1@yandex.ru  
GitHub: github.com/Vladimirmoscow84  