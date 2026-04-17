DELETE FROM books b
USING authors a, categories c
WHERE b.author_id = a.id
  AND b.category_id = c.id
  AND (
      (b.title = 'Clean Code' AND a.name = 'Robert C. Martin' AND c.name = 'Programming') OR
      (b.title = 'Refactoring' AND a.name = 'Martin Fowler' AND c.name = 'Refactoring') OR
      (b.title = 'The Pragmatic Programmer' AND a.name = 'Andrew Hunt' AND c.name = 'Career') OR
      (b.title = 'Designing Data-Intensive Applications' AND a.name = 'Martin Kleppmann' AND c.name = 'Databases') OR
      (b.title = 'The Art of Computer Programming' AND a.name = 'Donald E. Knuth' AND c.name = 'Algorithms') OR
      (b.title = 'Introduction to Algorithms' AND a.name = 'Thomas H. Cormen' AND c.name = 'Computer Science') OR
      (b.title = 'Building Microservices' AND a.name = 'Sam Newman' AND c.name = 'Microservices') OR
      (b.title = 'The Phoenix Project' AND a.name = 'Gene Kim' AND c.name = 'DevOps') OR
      (b.title = 'Test-Driven Development' AND a.name = 'Kent Beck' AND c.name = 'Testing') OR
      (b.title = 'Dune' AND a.name = 'Frank Herbert' AND c.name = 'Fiction')
  );
