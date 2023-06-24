package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
	config "tugas15/Config"
	"tugas15/models"
)

const (
	tableMahasiswa = "mahasiswa"
	// layoutMatkul = "2006-01-02 15:04:05"
)

func GetAllMahasiswa(ctx context.Context) ([]models.Mahasiswa, error) {
	var mahasiswas []models.Mahasiswa
	//ini untuk instance mysql
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to mySQL", err)
	}

	//ini untuk query get data
	query := fmt.Sprintf("SELECT * FROM %v ORDER BY created_at DESC", tableMahasiswa)
	rowQuery, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	//ini harus urut sesuai yang ada di tabel
	for rowQuery.Next() {
		var item models.Mahasiswa
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&item.ID,
			&item.Nama,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		item.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		item.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		mahasiswas = append(mahasiswas, item)
	}

	return mahasiswas, nil
}

func InsertMahasiswa(ctx context.Context, mahasiswa models.Mahasiswa) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to mySQL", err)
	}
	mahasiswa.ID = uint(rand.Intn(1000))
	query := fmt.Sprintf("INSERT INTO %v (nama, created_at, updated_at) values ('%v',NOW(), NOW())", tableMahasiswa, mahasiswa.Nama)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func UpdateMahasiswa(ctx context.Context, mahasiswa models.Mahasiswa, idMahasiswa int) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("cant connect to mysql", err)
	}
	query := fmt.Sprintf("UPDATE %v set nama = '%v', updated_at = NOW() where id = %v", tableMahasiswa, mahasiswa.Nama, idMahasiswa)

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMahasiswa(ctx context.Context, idMahasiswa string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("cant connect to mysql", err)
	}
	query := fmt.Sprintf("DELETE FROM %v where id = %s", tableMahasiswa, idMahasiswa)

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
