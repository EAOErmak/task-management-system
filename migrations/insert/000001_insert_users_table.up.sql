-- Seeded users password: SeedPass123!
INSERT INTO users (username, password, role)
VALUES
    ('alice_admin', '$2a$10$761Zg19auqWKrGw9rcnb/eCsG8KVB91cU9vdfkRZyHR.sgdYOLdlq', 'admin'),
    ('boris_user', '$2a$10$761Zg19auqWKrGw9rcnb/eCsG8KVB91cU9vdfkRZyHR.sgdYOLdlq', 'user'),
    ('clara_user', '$2a$10$761Zg19auqWKrGw9rcnb/eCsG8KVB91cU9vdfkRZyHR.sgdYOLdlq', 'user'),
    ('daniel_user', '$2a$10$761Zg19auqWKrGw9rcnb/eCsG8KVB91cU9vdfkRZyHR.sgdYOLdlq', 'user'),
    ('elena_user', '$2a$10$761Zg19auqWKrGw9rcnb/eCsG8KVB91cU9vdfkRZyHR.sgdYOLdlq', 'user'),
    ('felix_user', '$2a$10$761Zg19auqWKrGw9rcnb/eCsG8KVB91cU9vdfkRZyHR.sgdYOLdlq', 'user'),
    ('greta_user', '$2a$10$761Zg19auqWKrGw9rcnb/eCsG8KVB91cU9vdfkRZyHR.sgdYOLdlq', 'user'),
    ('hugo_user', '$2a$10$761Zg19auqWKrGw9rcnb/eCsG8KVB91cU9vdfkRZyHR.sgdYOLdlq', 'user'),
    ('irene_user', '$2a$10$761Zg19auqWKrGw9rcnb/eCsG8KVB91cU9vdfkRZyHR.sgdYOLdlq', 'user'),
    ('jack_user', '$2a$10$761Zg19auqWKrGw9rcnb/eCsG8KVB91cU9vdfkRZyHR.sgdYOLdlq', 'user')
ON CONFLICT (username) DO NOTHING;
