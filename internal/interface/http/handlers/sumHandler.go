package handlers

import (
	"database/sql"
	"fib/internal/interface/http/queryBuilder"
	"fib/internal/models"
	"fib/pkg/logger"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SumHandler(c *fiber.Ctx) error {
	var r models.Request
	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	sum := 0
	for _, num := range r.Nums {
		sum += num
	}

	return c.JSON(fiber.Map{"sum": sum})
}

func Delete(db *sql.DB, lg logger.MyLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.Book

		if err := c.BodyParser(&request); err != nil {
			lg.Error("delete Parse Error:", err)
			c.Status(http.StatusBadRequest).SendString("Parser broken")
		}

		query := queryBuilder.UpdateBuilder(&request)

		if query == "" {
			lg.Error("somerthind went wrong")
			c.Status(http.StatusBadRequest).SendString("invalid request")
		}

		if _, err := db.Query(query); err != nil {
			lg.Error("delete query error", err)
		} else {
			lg.Info("successfully deleted")
			c.Status(http.StatusOK).SendString("successfully deleted")
		}
		return nil
	}
}

func UpdateBuilder(bookRequest *models.Book) string {
	var setConditions []string

	if bookRequest.Author_name != "" {
		setConditions = append(setConditions, fmt.Sprintf("author_name = '%s'", bookRequest.Author_name))
	}

	if bookRequest.Book_title != "" {
		setConditions = append(setConditions, fmt.Sprintf("book_title = '%s'", bookRequest.Book_title))
	}

	if bookRequest.Release_year != 0 {
		setConditions = append(setConditions, fmt.Sprintf("release_year = %d", bookRequest.Release_year))
	}

	if len(setConditions) == 0 {
		return ""
	}

	query := fmt.Sprintf("UPDATE books SET %s WHERE id = %d",
		strings.Join(setConditions, ", "), bookRequest.Id)

	return query
}
