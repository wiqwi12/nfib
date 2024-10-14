package queryBuilder

import (
	"fib/internal/models"
	"fmt"
	"strings"
)

func SelectBuilder(bookRequest *models.Book) string {

	var conditions []string

	if bookRequest.Id != 0 {
		conditions = append(conditions, fmt.Sprintf("id = %d", bookRequest.Id))
	}

	if bookRequest.Author_name != "" {
		conditions = append(conditions, fmt.Sprintf("author_name = '%s'", bookRequest.Author_name))
	}

	if bookRequest.Book_title != "" {
		conditions = append(conditions, fmt.Sprintf("book_title = '%s'", bookRequest.Book_title))
	}

	if bookRequest.Release_year != 0 {
		conditions = append(conditions, fmt.Sprintf("release_year = %d", bookRequest.Release_year))
	}

	query := "SELECT * FROM books"

	query += " WHERE " + strings.Join(conditions, " AND ")

	if len(conditions) == 0 {
		return ""
	}
	return query

}

func CreateBuilder(createRequest *models.Book) string {

	var conditions []string
	var insertConditions []string

	if createRequest.Author_name != "" {
		conditions = append(conditions, fmt.Sprintf("'%s'", createRequest.Author_name))
		insertConditions = append(insertConditions, "author_name")
	}

	if createRequest.Book_title != "" {
		conditions = append(conditions, fmt.Sprintf("'%s'", createRequest.Book_title))
		insertConditions = append(insertConditions, "book_title")
	}

	if createRequest.Release_year != 0 {
		conditions = append(conditions, fmt.Sprintf("%d", createRequest.Release_year))
		insertConditions = append(insertConditions, "release_year")
	}

	if len(insertConditions) == 0 {
		return ""
	}

	query := fmt.Sprintf("INSERT INTO books (%s) VALUES (%s)",
		strings.Join(insertConditions, ", "),
		strings.Join(conditions, ", "),
	)

	return query
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

func DeleteBuilder(bookRequest *models.Book) string {

	var conditions []string

	if bookRequest.Id != 0 {
		conditions = append(conditions, fmt.Sprintf("id = %d", bookRequest.Id))
	}

	if bookRequest.Author_name != "" {
		conditions = append(conditions, fmt.Sprintf("author_name = '%s'", bookRequest.Author_name))
	}

	if bookRequest.Book_title != "" {
		conditions = append(conditions, fmt.Sprintf("book_title = '%s'", bookRequest.Book_title))
	}

	if bookRequest.Release_year != 0 {
		conditions = append(conditions, fmt.Sprintf("release_year = %d", bookRequest.Release_year))
	}

	if len(conditions) == 0 {
		return ""
	}

	query := fmt.Sprintf("DELETE FROM books WHERE %s", strings.Join(conditions, " AND "))

	return query

}
