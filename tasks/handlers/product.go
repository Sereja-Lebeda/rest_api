package handlers

import (
	"rest_api/tasks/database"
	"rest_api/tasks/models"

	"github.com/gofiber/fiber/v2"
)

// getting all tasks
func GetTasks(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT * FROM tasks")
	if err != nil {
		return c.Status(500).SendString("Ошибка выполнения запроса к базе данных")
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created_at)
		if err != nil {
			return c.Status(500).SendString("Ошибка сканирования данных")
		}
		tasks = append(tasks, task)
	}

	return c.JSON(tasks)
}

// creating new task
func CreateTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	if task.Title == "" {
		return c.Status(400).SendString("Поле Title не должно быть пустым")
	}

	if task.Status == "" {
		return c.Status(400).SendString("Поле Status не должно быть пустым")
	}

	_, err := database.DB.Exec("INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3)",
		task.Title, task.Description, task.Status)

	if err != nil {
		// Логируем ошибку для получения детальной информации
		return c.Status(500).SendString("Ошибка вставки данных в базу данных: ")
	}

	return c.Status(201).SendString("Задача успешно создана")
}

// getting task by ID
func GetTask(c *fiber.Ctx) error {
	id := c.Params("id")
	row := database.DB.QueryRow("SELECT id, title, description, status, created_at FROM tasks WHERE id = $1", id)

	var task models.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.Created_at)
	if err != nil {
		return c.Status(404).SendString("Задача с таким ID не найдена")
	}

	return c.JSON(task)
}

// updating task
func UpdateTask(c *fiber.Ctx) error {
	id := c.Params("id")
	task := new(models.Task)

	if err := c.BodyParser(task); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	if task.Title == "" {
		return c.Status(400).SendString("Поле Title не должно быть пустым")
	}

	if task.Status == "" {
		return c.Status(400).SendString("Поле Status не должно быть пустым")
	}

	result, err := database.DB.Exec("UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4",
		task.Title, task.Description, task.Status, id)
	if err != nil {
		return c.Status(500).SendString("Ошибка обновления данных")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(500).SendString("Ошибка получения данных о затронутых строках")
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("Задача с таким ID не найдена")
	}

	return c.SendString("Задача успешно обновлена")
}

// deleting task
func DeleteTask(c *fiber.Ctx) error {
	id := c.Params("id")

	result, err := database.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return c.Status(500).SendString("Ошибка удаления задачи")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return c.Status(500).SendString("Ошибка получения данных о затронутых строках")
	}

	if rowsAffected == 0 {
		return c.Status(404).SendString("Задача с таким ID не найдена")
	}

	return c.SendString("Задача успешно удалена")
}
