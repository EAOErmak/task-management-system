DELETE FROM favorite_books fb
USING users u, books b, authors a, categories c
WHERE fb.user_id = u.id
  AND fb.book_id = b.id
  AND b.author_id = a.id
  AND b.category_id = c.id
  AND (
      (u.username = 'alice_admin' AND b.title = 'Clean Code' AND a.name = 'Robert C. Martin' AND c.name = 'Programming') OR
      (u.username = 'boris_user' AND b.title = 'Refactoring' AND a.name = 'Martin Fowler' AND c.name = 'Refactoring') OR
      (u.username = 'clara_user' AND b.title = 'The Pragmatic Programmer' AND a.name = 'Andrew Hunt' AND c.name = 'Career') OR
      (u.username = 'daniel_user' AND b.title = 'Designing Data-Intensive Applications' AND a.name = 'Martin Kleppmann' AND c.name = 'Databases') OR
      (u.username = 'elena_user' AND b.title = 'The Art of Computer Programming' AND a.name = 'Donald E. Knuth' AND c.name = 'Algorithms') OR
      (u.username = 'felix_user' AND b.title = 'Introduction to Algorithms' AND a.name = 'Thomas H. Cormen' AND c.name = 'Computer Science') OR
      (u.username = 'greta_user' AND b.title = 'Building Microservices' AND a.name = 'Sam Newman' AND c.name = 'Microservices') OR
      (u.username = 'hugo_user' AND b.title = 'The Phoenix Project' AND a.name = 'Gene Kim' AND c.name = 'DevOps') OR
      (u.username = 'irene_user' AND b.title = 'Test-Driven Development' AND a.name = 'Kent Beck' AND c.name = 'Testing') OR
      (u.username = 'jack_user' AND b.title = 'Dune' AND a.name = 'Frank Herbert' AND c.name = 'Fiction')
  );
