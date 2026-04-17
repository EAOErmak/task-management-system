INSERT INTO authors (name)
VALUES
    ('Robert C. Martin'),
    ('Martin Fowler'),
    ('Andrew Hunt'),
    ('Martin Kleppmann'),
    ('Donald E. Knuth'),
    ('Thomas H. Cormen'),
    ('Sam Newman'),
    ('Gene Kim'),
    ('Kent Beck'),
    ('Frank Herbert')
ON CONFLICT (name) DO NOTHING;
