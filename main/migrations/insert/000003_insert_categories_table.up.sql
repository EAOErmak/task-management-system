INSERT INTO categories (name)
VALUES
    ('Programming'),
    ('Refactoring'),
    ('Career'),
    ('Databases'),
    ('Algorithms'),
    ('Computer Science'),
    ('Microservices'),
    ('DevOps'),
    ('Testing'),
    ('Fiction')
ON CONFLICT (name) DO NOTHING;
