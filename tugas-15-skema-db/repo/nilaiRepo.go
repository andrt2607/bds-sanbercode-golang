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
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAllNilai(ctx context.Context) ([]models.Nilai, error) {
	var nilais []models.Nilai
	//ini untuk instance mysql
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to mySQL", err)
	}

	//ini untuk query get data
	query := fmt.Sprintf("SELECT * FROM %v ORDER BY created_at DESC", table)
	rowQuery, err := db.QueryContext(ctx, query)

	if err != nil {
		log.Fatal(err)
	}

	//ini harus urut sesuai yang ada di tabel
	for rowQuery.Next() {
		var nilai models.Nilai
		var createdAt, updatedAt string
		if err = rowQuery.Scan(
			&nilai.ID,
			&nilai.Nama,
			&nilai.MataKuliah,
			&nilai.IndeksNilai,
			&nilai.Nilai,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		nilai.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		nilai.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		nilais = append(nilais, nilai)
	}

	return nilais, nil
}

func Insert(ctx context.Context, nilai models.Nilai) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to mySQL", err)
	}
	switch {
	case nilai.Nilai >= 80:
		nilai.IndeksNilai = "A"
	case nilai.Nilai >= 70 && nilai.Nilai < 80:
		nilai.IndeksNilai = "B"
	case nilai.Nilai >= 60 && nilai.Nilai < 70:
		nilai.IndeksNilai = "C"
	case nilai.Nilai >= 50 && nilai.Nilai < 60:
		nilai.IndeksNilai = "D"
	case nilai.Nilai < 80:
		nilai.IndeksNilai = "E"
	}
	nilai.ID = uint(rand.Intn(1000))
	query := fmt.Sprintf("INSERT INTO %v (id, name, mata_kuliah, indeks_nilai, nilai, created_at, updated_at) values ('%v','%v','%v','%v','%v', NOW(), NOW())", table, nilai.ID, nilai.Nama, nilai.MataKuliah, nilai.IndeksNilai, nilai.Nilai)
	_, err = db.ExecContext(ctx, query)

	if err != nil {
		return err
	}
	return nil
}

func UpdateNilai(ctx context.Context, nilai models.Nilai, id int) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("cant connect to mysql", err)
	}
	query := fmt.Sprintf("UPDATE %v set name = '%v', mata_kuliah = '%v', nilai = %v , updated_at = NOW() where id = %v", table, nilai.Nama, nilai.MataKuliah, nilai.Nilai, id)

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	return nil
}

func DeleteNilai(ctx context.Context, id string) error {
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("cant connect to mysql", err)
	}
	query := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

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
