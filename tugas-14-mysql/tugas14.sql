CREATE TABLE mahasiswa (
  id INT auto_increment primary key,
  name VARCHAR(255) not null,
  mata_kuliah VARCHAR(255) not null,
  indeks_nilai VARCHAR(255) not null,
  nilai int not null,
  created_at DATETIME not null,
  updated_at DATETIME not null
)