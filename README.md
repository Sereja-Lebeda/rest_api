1. Установка Fiber:
  "go get github.com/gofiber/fiber/v2"

2. Уставновка database\sql пакета:
   "go get github.com/lib/pq"

3. Установка Homebrew, если его нет:
  "/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

4. Установка Postgres:
  "brew install postgresql"
4.1. Запуск службы:
   "brew services start postgresql"
4.2. Зайти Postgres:
   "postgres -D /usr/local/var/postgres"
4.3. Подключение к базе данных
   "/usr/local/pgsql/bin/psql -U postgres"
4.3. Создание базы данных:
   "CREATE DATABASE tasksdb;"
4.4. Подключение к базе данных:
   "\c tasksdb"
4.5. Создание таблицы:
   "CREATE TABLE tasks (
      id SERIAL PRIMARY KEY,
      title VARCHAR(50) NOT NULL,
      description VARCHAR(100) NOT NULL.
      status VARCHAR(50) NOT NULL, 
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
   ); "
4.6. Проверка, что таблица создалась:
   "\d tasksdb"

5. Переход в директорию, где хранится main.go

6. Запуск main файла:
   "go run main.go"

7. Проверка работопособности API, например через Postman (см. Readme_Instructions)
