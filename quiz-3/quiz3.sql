CREATE TABLE category (
  id INT auto_increment primary key,
  name VARCHAR(255) not null,
  created_at DATETIME not null,
  updated_at DATETIME not null
)

ALTER TABLE book
ADD FOREIGN KEY (category_id) REFERENCES category(id)

CREATE TABLE book (
  id INT auto_increment primary key,
  title VARCHAR(255) not null,
  description varchar(255) not null,
  image_url varchar(255) not null,
    release_year int not null,
    price varchar(255) not null,
    total_page int not null,
    thickness varchar(255) not null,
  created_at DATETIME not null,
  updated_at DATETIME not null,
    category_id int
)
