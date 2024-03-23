CREATE DATABASE my_db;

CREATE TABLE IF NOT EXISTS authors (
    id INTEGER PRIMARY KEY,
    first_name VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS books (
  id INTEGER PRIMARY KEY,
  author_id INTEGER NOT NULL,
  title VARCHAR(1024) NOT NULL,
  FOREIGN KEY (author_id) REFERENCES authors (id)
);

CREATE TABLE IF NOT EXISTS readers (
    id INTEGER PRIMARY KEY,
    login VARCHAR(255) NOT NULL,
    CONSTRAINT unique_login UNIQUE (login)
);

CREATE TABLE IF NOT EXISTS reader_books (
    book_id INTEGER NOT NULL,
    reader_id INTEGER NOT NULL,
    FOREIGN KEY (book_id) REFERENCES books (id),
    FOREIGN KEY (reader_id) REFERENCES readers (id),
    CONSTRAINT unique_book_reader UNIQUE (book_id, reader_id)
);