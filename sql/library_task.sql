CREATE DATABASE library_db;

CREATE SCHEMA IF NOT EXISTS library;

-- Author entity
CREATE TABLE IF NOT EXISTS library.authors (
    id SERIAL PRIMARY KEY,
    first_name      VARCHAR(100),
    last_name       VARCHAR(100),
    patronymic      VARCHAR(100),
    fullname        text generated always as (last_name || ' ' || first_name || ' ' || patronymic) stored,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Book entity
CREATE TABLE IF NOT EXISTS library.books (
  id SERIAL PRIMARY KEY,
  author_id INTEGER NOT NULL REFERENCES library.authors (id),
  genre VARCHAR(1024) NOT NULL,
  title VARCHAR(1024) NOT NULL,
  a_copies INTEGER DEFAULT 0,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Reader entity
CREATE TABLE IF NOT EXISTS library.readers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL unique,
    login VARCHAR(100) NOT NULL unique,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Reader's books entity
CREATE TABLE IF NOT EXISTS library.reader_books (
    id SERIAL PRIMARY KEY,
    reader_id INTEGER NOT NULL REFERENCES library.readers (id),
    book_id INTEGER NOT NULL REFERENCES library.books (id),
    borrow_at date,

    UNIQUE (book_id, reader_id)
);

INSERT INTO library.authors
(first_name, last_name, patronymic) VALUES
('Alexandr', 'Pushkin', 'Sergeevich');

INSERT INTO library.books
(author_id, title, genre) VALUES
(1, 'Ruslan and Ludmila', 'Pulp fiction');

INSERT INTO library.readers
(name, email, login) VALUES
('Example', 'example@mail.ru', 'example');

INSERT INTO library.reader_books
(reader_id, book_id, borrow_at) VALUES
(1, 1, current_date);

-- Get list of books borrowed by specific reader
SELECT * FROM library.readers r JOIN library.reader_books rb ON r.id = rb.reader_id JOIN library.books b on b.id = rb.book_id WHERE name LIKE '%Example';

-- Find readers who borrowed specific book
SELECT * FROM library.reader_books rb JOIN library.books b on b.id = rb.book_id join library.readers r on r.id = rb.reader_id WHERE b.title = 'Ruslan and Ludmila';

-- Update number of available copies of book
UPDATE library.books SET a_copies = a_copies - 1 WHERE id = 1;

-- Write a query to find all books that are currently borrowed and overdue
-- (i.e., not returned within 14 days from the borrow date). Display the book titles and the names of members who borrowed them.
SELECT * FROM library.reader_books rb JOIN library.readers r on r.id = rb.reader_id JOIN library.books b on b.id = rb.book_id WHERE date_part('day', current_date) - date_part('day', rb.borrow_at::date) > 14;

-- Find the most popular genre in the library.
-- Display the genre with the highest total number of books borrowed.
SELECT title, genre, count(*) as total_borrow FROM library.books b JOIN library.reader_books rb on b.id = rb.book_id GROUP BY title, genre ORDER BY total_borrow LIMIT 1;



