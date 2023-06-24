
CREATE TABLE nilai (
  id INT auto_increment primary key,
  indeks VARCHAR(255) not null,
  skor int not null,
  created_at DATETIME not null,
  updated_at DATETIME not null
)

ALTER TABLE nilai
ADD mahasiswa_id INT

ALTER TABLE nilai
ADD mata_kuliah_id INT

ALTER TABLE nilai
ADD FOREIGN KEY (mahasiswa_id) REFERENCES mahasiswa(id)

ALTER TABLE nilai
ADD FOREIGN KEY (mata_kuliah_id) REFERENCES mata_kuliah(id)

CREATE TABLE mahasiswa (
  id INT auto_increment primary key,
  nama VARCHAR(255) not null,
  created_at DATETIME not null,
  updated_at DATETIME not null
)

CREATE TABLE mata_kuliah (
  id INT auto_increment primary key,
  nama VARCHAR(255) not null,
  created_at DATETIME not null,
  updated_at DATETIME not null
)
