INSERT INTO favorite_books (user_id, book_id)
SELECT u.id, b.id
FROM users u
JOIN books b ON b.title = 'Clean Code'
JOIN authors a ON b.author_id = a.id
JOIN categories c ON b.category_id = c.id
WHERE u.username = 'alice_admin'
  AND a.name = 'Robert C. Martin'
  AND c.name = 'Programming'
ON CONFLICT (user_id, book_id) DO NOTHING;

INSERT INTO favorite_books (user_id, book_id)
SELECT u.id, b.id
FROM users u
JOIN books b ON b.title = 'Refactoring'
JOIN authors a ON b.author_id = a.id
JOIN categories c ON b.category_id = c.id
WHERE u.username = 'boris_user'
  AND a.name = 'Martin Fowler'
  AND c.name = 'Refactoring'
ON CONFLICT (user_id, book_id) DO NOTHING;

INSERT INTO favorite_books (user_id, book_id)
SELECT u.id, b.id
FROM users u
JOIN books b ON b.title = 'The Pragmatic Programmer'
JOIN authors a ON b.author_id = a.id
JOIN categories c ON b.category_id = c.id
WHERE u.username = 'clara_user'
  AND a.name = 'Andrew Hunt'
  AND c.name = 'Career'
ON CONFLICT (user_id, book_id) DO NOTHING;

INSERT INTO favorite_books (user_id, book_id)
SELECT u.id, b.id
FROM users u
JOIN books b ON b.title = 'Designing Data-Intensive Applications'
JOIN authors a ON b.author_id = a.id
JOIN categories c ON b.category_id = c.id
WHERE u.username = 'daniel_user'
  AND a.name = 'Martin Kleppmann'
  AND c.name = 'Databases'
ON CONFLICT (user_id, book_id) DO NOTHING;

INSERT INTO favorite_books (user_id, book_id)
SELECT u.id, b.id
FROM users u
JOIN books b ON b.title = 'The Art of Computer Programming'
JOIN authors a ON b.author_id = a.id
JOIN categories c ON b.category_id = c.id
WHERE u.username = 'elena_user'
  AND a.name = 'Donald E. Knuth'
  AND c.name = 'Algorithms'
ON CONFLICT (user_id, book_id) DO NOTHING;

INSERT INTO favorite_books (user_id, book_id)
SELECT u.id, b.id
FROM users u
JOIN books b ON b.title = 'Introduction to Algorithms'
JOIN authors a ON b.author_id = a.id
JOIN categories c ON b.category_id = c.id
WHERE u.username = 'felix_user'
  AND a.name = 'Thomas H. Cormen'
  AND c.name = 'Computer Science'
ON CONFLICT (user_id, book_id) DO NOTHING;

INSERT INTO favorite_books (user_id, book_id)
SELECT u.id, b.id
FROM users u
JOIN books b ON b.title = 'Building Microservices'
JOIN authors a ON b.author_id = a.id
JOIN categories c ON b.category_id = c.id
WHERE u.username = 'greta_user'
  AND a.name = 'Sam Newman'
  AND c.name = 'Microservices'
ON CONFLICT (user_id, book_id) DO NOTHING;

INSERT INTO favorite_books (user_id, book_id)
SELECT u.id, b.id
FROM users u
JOIN books b ON b.title = 'The Phoenix Project'
JOIN authors a ON b.author_id = a.id
JOIN categories c ON b.category_id = c.id
WHERE u.username = 'hugo_user'
  AND a.name = 'Gene Kim'
  AND c.name = 'DevOps'
ON CONFLICT (user_id, book_id) DO NOTHING;

INSERT INTO favorite_books (user_id, book_id)
SELECT u.id, b.id
FROM users u
JOIN books b ON b.title = 'Test-Driven Development'
JOIN authors a ON b.author_id = a.id
JOIN categories c ON b.category_id = c.id
WHERE u.username = 'irene_user'
  AND a.name = 'Kent Beck'
  AND c.name = 'Testing'
ON CONFLICT (user_id, book_id) DO NOTHING;

INSERT INTO favorite_books (user_id, book_id)
SELECT u.id, b.id
FROM users u
JOIN books b ON b.title = 'Dune'
JOIN authors a ON b.author_id = a.id
JOIN categories c ON b.category_id = c.id
WHERE u.username = 'jack_user'
  AND a.name = 'Frank Herbert'
  AND c.name = 'Fiction'
ON CONFLICT (user_id, book_id) DO NOTHING;
