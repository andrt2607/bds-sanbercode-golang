package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	config "quiz3/Config"
	"quiz3/models"
	"time"
)

const (
	tableBook = "book"
	// layoutTime    = "2006-01-02 15:04:05"
)

func GetAllBooks(ctx context.Context) ([]models.Book, error) {
	var books []models.Book
	//ini untuk instance mysql
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to mySQL", err)
	}

	//ini untuk query get data
	query := fmt.Sprintf("SELECT * FROM %v ORDER BY created_at DESC", tableBook)
	rowQuery, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	//ini harus urut sesuai yang ada di tabel
	for rowQuery.Next() {
		var item models.Book
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&item.Id,
			&item.Title,
			&item.ImageUrl,
			&item.ReleaseYear,
			&item.Price,
			&item.TotalPage,
			&item.Thickness,
			&createdAt,
			&updatedAt,
			&item.CategoryId,
			&item.Description,
		); err != nil {
			return nil, err
		}

		item.CreatedAt, err = time.Parse(layoutTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		item.UpdatedAt, err = time.Parse(layoutTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		books = append(books, item)
	}

	return books, nil
}

func InsertBook(ctx context.Context, book models.Book) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to mySQL", err)
	}
	switch {
	case book.TotalPage <= 100:
		book.Thickness = "tipis"
	case book.TotalPage >= 101 && book.TotalPage <= 200:
		book.Thickness = "sedang"
	case book.TotalPage >= 100 && book.TotalPage <= 201:
		book.Thickness = "tebal"
	}
	// category.Id = uint(rand.Intn(1000))
	query := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, category_id, created_at, updated_at, thickness) values ('%v', '%v', '%v', '%v', '%v' , %v , %v ,NOW(), NOW(), '%v')",
		tableBook, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.CategoryId, book.Thickness)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func UpdateBook(ctx context.Context, book models.Book, idBook int) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("cant connect to mysql", err)
	}
	query := fmt.Sprintf("UPDATE %v set title = '%v', description = '%v', image_url = '%v', release_year = '%v', price = '%v', total_page = %v, category_id = %v, updated_at = NOW() where id = %v",
		tableBook, book.Title, book.Description, book.ImageUrl, book.ReleaseYear, book.Price, book.TotalPage, book.CategoryId, idBook)

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func DeleteBook(ctx context.Context, idBook int) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("cant connect to mysql", err)
	}
	query := fmt.Sprintf("DELETE FROM %v where id = %v", tableBook, idBook)

	s, err := db.ExecContext(ctx, query)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}
	return nil
}
