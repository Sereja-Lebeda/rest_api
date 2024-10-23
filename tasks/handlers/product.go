package handlers

import (
	"rest_api/tasks/database"
	"rest_api/tasks/models"

	"github.com/gofiber/fiber/v2"
)

// getting all tasks
func GetProducts(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT * FROM products") //id, title, description, status, created_at FROM products")
	if err != nil {
		return c.Status(500).SendString("Ошибка выполнения запроса к базе данных" + err.Error())
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Status, &product.Created_at)
		if err != nil {
			return c.Status(500).SendString("Ошибка сканирования данных")
		}
		products = append(products, product)
	}

	return c.JSON(products)
}

// creating new task
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	if product.Title == "" {
		return c.Status(400).SendString("Поле Title не должно быть пустым")
	}

	if product.Status == "" {
		return c.Status(400).SendString("Поле Status не должно быть пустым")
	}

	_, err := database.DB.Exec("INSERT INTO products (title, description, status) VALUES ($1, $2, $3)", //, $4)",
		product.Title, product.Description, product.Status) //, product.Created_at)

	if err != nil {
		// Логируем ошибку для получения детальной информации
		return c.Status(500).SendString("Ошибка вставки данных в базу данных: " + err.Error())

		//	return c.Status(500).SendString("Ошибка вставки данных в базу данных")
	}

	return c.Status(201).SendString("Продукт успешно создан")
}

// getting task by ID
func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	row := database.DB.QueryRow("SELECT id, title, description, status, created_at FROM products WHERE id = $1", id)

	var product models.Product
	err := row.Scan(&product.ID, &product.Title, &product.Description, &product.Status, &product.Created_at)
	if err != nil {
		return c.Status(404).SendString("Продукт не найден")
	}

	return c.JSON(product)
}

// updating task
func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(400).SendString("Неверный формат запроса")
	}

	if product.Title == "" {
		return c.Status(400).SendString("Поле Title не должно быть пустым")
	}

	if product.Status == "" {
		return c.Status(400).SendString("Поле Status не должно быть пустым")
	}

	_, err := database.DB.Exec("UPDATE products SET title = $1, description = $2, status = $3 WHERE id = $4",
		product.Title, product.Description, product.Status, id)

	if err != nil {
		return c.Status(500).SendString("Ошибка обновления данных")
	}

	return c.SendString("Продукт успешно обновлен")
}

// deleting task
func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := database.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return c.Status(500).SendString("Ошибка удаления продукта")
	}

	return c.SendString("Продукт успешно удален")
}
