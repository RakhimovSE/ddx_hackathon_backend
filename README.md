# DDX Hackathon Backend

## Описание

DDX Hackathon Backend - это серверная часть приложения, разработанная на Go. Она предоставляет API для управления тренировочными планами, тренировками и общением с тренерами.

## Технологический стек

- **Go**: 1.22.3
- **Gin**: v1.10.0 - веб-фреймворк для создания API
- **GORM**: v1.9.16 - ORM для работы с базой данных

## Основные зависимости

- **gin-gonic/gin**: v1.10.0
- **go-faker/faker**: v4.4.2
- **jinzhu/gorm**: v1.9.16
- **joho/godotenv**: v1.5.1
- **stretchr/testify**: v1.9.0
- **golang.org/x/crypto**: v0.23.0

## Установка и настройка

1. Установите Go, следуя инструкциям на [официальном сайте Go](https://golang.org/doc/install).
2. Клонируйте репозиторий:
   ```sh
   git clone <URL репозитория>
   ```
3. Перейдите в директорию проекта:
   ```sh
   cd ddx_hackathon_backend
   ```
4. Установите зависимости:
   ```sh
   go mod tidy
   ```
5. Переименуйте файлы `.env.debug.example` и `.env.release.example` в `.env.debug` и `.env.release` соответственно, и подставьте в них свои значения.

6. Создайте базу данных в PostgreSQL:

   - На Mac или Linux:
     ```sh
     sudo -u postgres psql
     CREATE DATABASE ddx_hackathon;
     CREATE USER yourusername WITH ENCRYPTED PASSWORD 'yourpassword';
     GRANT ALL PRIVILEGES ON DATABASE ddx_hackathon TO yourusername;
     ```

7. Укажите данные для подключения к базе данных в файле `.env.debug` либо `.env.release`.

8. Для переключения режима запуска между `debug` и `release`, задайте переменную окружения `GIN_MODE=debug` либо `GIN_MODE=release`.

## Запуск приложения

1. Загрузите начальные данные:
   ```sh
   go run main.go load_data
   ```
2. Сгенерируйте тестовые данные:
   ```sh
   go run main.go seed
   ```
3. Запустите приложение:
   ```sh
   go run main.go
   ```

## Доступные команды

Все команды выполняются через `go run main.go <команда>`:

- `load_data`: Загрузить данные из файла.
- `seed`: Сгенерировать тестовые данные (тренеры, тренировочные планы, тренеры клиентов, тренировочные планы клиентов).
- `seed_trainers`: Сгенерировать тестовые данные для тренеров.
- `seed_training_plans`: Сгенерировать тестовые данные для тренировочных планов.
- `seed_client_trainers`: Сгенерировать тестовые данные для тренеров клиентов.
- `seed_client_training_plans`: Сгенерировать тестовые данные для тренировочных планов клиентов.

## Структура проекта

- `main.go`: Основной файл для запуска приложения.
- `database/`: Миграции базы данных и подключение.
- `handlers/`: Обработчики HTTP-запросов.
- `models/`: Определения моделей данных.
- `routes/`: Определения маршрутов API.
- `scripts/`: Скрипты для загрузки и генерации данных.

## Маршруты API

### Клиенты (client_routes.go)

- `GET /clients/:client_id`: Получить информацию о клиенте.
- `PUT /clients/:client_id`: Обновить информацию о клиенте.

### Тренировки клиентов (client_workout_routes.go)

- `GET /clients/:client_id/workouts`: Получить тренировки клиента.
- `GET /training_plans/:training_plan_id/workouts`: Получить тренировки по ID тренировочного плана.
- `GET /client_workouts/:client_workout_id`: Получить упражнения и подходы для тренировки клиента.

### Упражнения (exercise_routes.go)

- `GET /exercises`: Получить список упражнений.
- `GET /exercises/:id`: Получить информацию об упражнении.

### Тренеры (trainer_routes.go)

- `GET /trainers`: Получить список тренеров.
- `GET /trainers/:id`: Получить информацию о тренере.

### Тренировочные планы (training_plan_routes.go)

- `GET /training_plans`: Получить список тренировочных планов.
- `GET /training_plans/:id`: Получить информацию о тренировочном плане.

### Пользователи (user_routes.go)

- `GET /users`: Получить список пользователей.
- `GET /users/:id`: Получить информацию о пользователе.
- `POST /users`: Создать нового пользователя.
- `PUT /users/:id`: Обновить информацию о пользователе.
- `DELETE /users/:id`: Удалить пользователя.

### Тренировки (workout_routes.go)

- `GET /workouts`: Получить список тренировок.
- `GET /workouts/:id`: Получить информацию о тренировке.

## Ссылка на репозиторий фронтенда

[Frontend Repository](https://github.com/RakhimovSE/ddx_hackathon_frontend)
