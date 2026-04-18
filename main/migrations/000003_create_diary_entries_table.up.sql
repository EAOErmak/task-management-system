CREATE TABLE diary_entry (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    user_id BIGINT NOT NULL REFERENCES users(id),
    when_started TIMESTAMPTZ NOT NULL,
    when_ended TIMESTAMPTZ NOT NULL,
    duration INTEGER NOT NULL,
    mood SMALLINT NULL,
    description VARCHAR(1000) NOT NULL
);

CREATE INDEX idx_diary_entry_user_id ON diary_entry (user_id);
