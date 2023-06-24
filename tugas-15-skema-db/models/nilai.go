package models

import "time"

type Nilai struct {
	ID           uint      `json:"id"`
	Indeks       string    `json:"indeks"`
	Skor         int       `json:"skor"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	MahasiswaId  int       `json:"mahasiswa_id"`
	MataKuliahId int       `json:"mata_kuliah_id"`
}
