INSERT INTO books (title, author_id, category_id, price)
SELECT 'Clean Code', a.id, c.id, 34.90
FROM authors a
JOIN categories c ON c.name = 'Programming'
WHERE a.name = 'Robert C. Martin'
  AND NOT EXISTS (
      SELECT 1
      FROM books b
      WHERE b.title = 'Clean Code'
        AND b.author_id = a.id
        AND b.category_id = c.id
  );

INSERT INTO books (title, author_id, category_id, price)
SELECT 'Refactoring', a.id, c.id, 39.50
FROM authors a
JOIN categories c ON c.name = 'Refactoring'
WHERE a.name = 'Martin Fowler'
  AND NOT EXISTS (
      SELECT 1
      FROM books b
      WHERE b.title = 'Refactoring'
        AND b.author_id = a.id
        AND b.category_id = c.id
  );

INSERT INTO books (title, author_id, category_id, price)
SELECT 'The Pragmatic Programmer', a.id, c.id, 29.99
FROM authors a
JOIN categories c ON c.name = 'Career'
WHERE a.name = 'Andrew Hunt'
  AND NOT EXISTS (
      SELECT 1
      FROM books b
      WHERE b.title = 'The Pragmatic Programmer'
        AND b.author_id = a.id
        AND b.category_id = c.id
  );

INSERT INTO books (title, author_id, category_id, price)
SELECT 'Designing Data-Intensive Applications', a.id, c.id, 44.00
FROM authors a
JOIN categories c ON c.name = 'Databases'
WHERE a.name = 'Martin Kleppmann'
  AND NOT EXISTS (
      SELECT 1
      FROM books b
      WHERE b.title = 'Designing Data-Intensive Applications'
        AND b.author_id = a.id
        AND b.category_id = c.id
  );

INSERT INTO books (title, author_id, category_id, price)
SELECT 'The Art of Computer Programming', a.id, c.id, 89.99
FROM authors a
JOIN categories c ON c.name = 'Algorithms'
WHERE a.name = 'Donald E. Knuth'
  AND NOT EXISTS (
      SELECT 1
      FROM books b
      WHERE b.title = 'The Art of Computer Programming'
        AND b.author_id = a.id
        AND b.category_id = c.id
  );

INSERT INTO books (title, author_id, category_id, price)
SELECT 'Introduction to Algorithms', a.id, c.id, 54.90
FROM authors a
JOIN categories c ON c.name = 'Computer Science'
WHERE a.name = 'Thomas H. Cormen'
  AND NOT EXISTS (
      SELECT 1
      FROM books b
      WHERE b.title = 'Introduction to Algorithms'
        AND b.author_id = a.id
        AND b.category_id = c.id
  );

INSERT INTO books (title, author_id, category_id, price)
SELECT 'Building Microservices', a.id, c.id, 37.80
FROM authors a
JOIN categories c ON c.name = 'Microservices'
WHERE a.name = 'Sam Newman'
  AND NOT EXISTS (
      SELECT 1
      FROM books b
      WHERE b.title = 'Building Microservices'
        AND b.author_id = a.id
        AND b.category_id = c.id
  );

INSERT INTO books (title, author_id, category_id, price)
SELECT 'The Phoenix Project', a.id, c.id, 24.90
FROM authors a
JOIN categories c ON c.name = 'DevOps'
WHERE a.name = 'Gene Kim'
  AND NOT EXISTS (
      SELECT 1
      FROM books b
      WHERE b.title = 'The Phoenix Project'
        AND b.author_id = a.id
        AND b.category_id = c.id
  );

INSERT INTO books (title, author_id, category_id, price)
SELECT 'Test-Driven Development', a.id, c.id, 31.20
FROM authors a
JOIN categories c ON c.name = 'Testing'
WHERE a.name = 'Kent Beck'
  AND NOT EXISTS (
      SELECT 1
      FROM books b
      WHERE b.title = 'Test-Driven Development'
        AND b.author_id = a.id
        AND b.category_id = c.id
  );

INSERT INTO books (title, author_id, category_id, price)
SELECT 'Dune', a.id, c.id, 18.75
FROM authors a
JOIN categories c ON c.name = 'Fiction'
WHERE a.name = 'Frank Herbert'
  AND NOT EXISTS (
      SELECT 1
      FROM books b
      WHERE b.title = 'Dune'
        AND b.author_id = a.id
        AND b.category_id = c.id
  );
