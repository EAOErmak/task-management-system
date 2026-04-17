CREATE TABLE favorite_books (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    user_id BIGINT NOT NULL,
    book_id BIGINT NOT NULL,
    CONSTRAINT fk_favorite_books_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_favorite_books_book FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX udx_favorite_books_user_book ON favorite_books (user_id, book_id);
