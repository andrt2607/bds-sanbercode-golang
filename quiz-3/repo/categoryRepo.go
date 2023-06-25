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
	tableCategory = "category"
	layoutTime    = "2006-01-02 15:04:05"
)

func GetAllCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	//ini untuk instance mysql
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to mySQL", err)
	}

	//ini untuk query get data
	query := fmt.Sprintf("SELECT * FROM %v ORDER BY created_at DESC", tableCategory)
	rowQuery, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	//ini harus urut sesuai yang ada di tabel
	for rowQuery.Next() {
		var item models.Category
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&item.Id,
			&item.Name,
			&createdAt,
			&updatedAt,
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

		categories = append(categories, item)
	}

	return categories, nil
}

func Insertcategory(ctx context.Context, category models.Category) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to mySQL", err)
	}
	// category.Id = uint(rand.Intn(1000))
	query := fmt.Sprintf("INSERT INTO %v (name, created_at, updated_at) values ('%v',NOW(), NOW())", tableCategory, category.Name)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func Updatecategory(ctx context.Context, category models.Category, idcategory int) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("cant connect to mysql", err)
	}
	query := fmt.Sprintf("UPDATE %v set name = '%v', updated_at = NOW() where id = %v", tableCategory, category.Name, idcategory)

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func Deletecategory(ctx context.Context, idcategory int) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("cant connect to mysql", err)
	}
	query := fmt.Sprintf("DELETE FROM %v where id = %v", tableCategory, idcategory)

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
