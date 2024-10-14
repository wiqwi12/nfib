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
	"github.com/gofiber/fiber/v2/log"
)

func IsExsist(db *sql.DB, bookRequest *models.Book) bool {
	var conditions []string

	if bookRequest.Author_name != "" {
		conditions = append(conditions, fmt.Sprintf("author_name = '%s'", bookRequest.Author_name))
	}

	if bookRequest.Book_title != "" {
		conditions = append(conditions, fmt.Sprintf("book_title = '%s'", bookRequest.Book_title))
	}

	query := "SELECT EXISTS (SELECT 1 FROM books"

	if len(conditions) != 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += ")"

	var exists bool
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		log.Error("Query err", err)
	}

	return exists
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

func Create(db *sql.DB, lg logger.MyLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.Book

		if err := c.BodyParser(&request); err != nil {
			lg.Error("builder pars err", err)
			return c.Status(http.StatusBadRequest).SendString("build parser broken")
		}

		if IsExsist(db, &request) {
			lg.Info("Already exsist")
			return c.Status(http.StatusConflict).SendString("Bruh book already exsists")

		}

		query := queryBuilder.CreateBuilder(&request)

		if query == "" {
			lg.Error("somerthind went wrong")
			c.Status(http.StatusBadRequest).SendString("invalid request")
		}

		if query == "" {
			lg.Error("empty request")
			return c.Status(http.StatusBadRequest).SendString("empty request")
		}

		if _, err := db.Query(query); err != nil {
			lg.Error("Create Query error", err)
		} else {
			lg.Info("successfully added")
			return c.Status(http.StatusOK).SendString("successfully added")
		}

		return nil
	}
}

func Select(db *sql.DB, lg logger.MyLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.Book
		if err := c.BodyParser(&request); err != nil {
			lg.Error("builder pars err", err)
			return c.Status(http.StatusBadRequest).SendString("build parser broken")
		}

		query := queryBuilder.SelectBuilder(&request)

		if query == "" {
			lg.Error("somerthind went wrong")
			c.Status(http.StatusBadRequest).SendString("invalid request")
		}

		rows, err := db.Query(query)
		if err != nil {
			lg.Error("query err", err)
			return err
		}
		defer rows.Close()

		var result []string

		for rows.Next() {
			var id int
			var authorName, bookTitle string
			var releaseYear int

			err := rows.Scan(&id, &authorName, &bookTitle, &releaseYear)
			if err != nil {
				lg.Error("Rows scan err:", err)
				return err
			}

			rowString := fmt.Sprintf("ID: %d, Author: %s, Title: %s, Year: %d", id, authorName, bookTitle, releaseYear)
			result = append(result, rowString)
		}

		if err = rows.Err(); err != nil {
			lg.Error("Rows iteration err:", err)
			return err
		}

		return c.JSON(result)
	}

}

func Update(db *sql.DB, lg logger.MyLogger) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request models.Book

		if err := c.BodyParser(&request); err != nil {
			lg.Error("Update Parse Error:", err)
			c.Status(http.StatusBadRequest).SendString("Parser broken")
		}

		query := queryBuilder.UpdateBuilder(&request)

		if query == "" {
			lg.Error("somerthind went wrong")
			c.Status(http.StatusBadRequest).SendString("invalid request")
		}

		if _, err := db.Query(query); err != nil {
			lg.Error("update query error", err)
		} else {
			lg.Info("successfully updated")
			c.Status(http.StatusOK).SendString("successfully updated")
		}
		return nil
	}
}
