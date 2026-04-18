CREATE TABLE entry_metric (
    id BIGSERIAL PRIMARY KEY,
    metric_type_id BIGINT NOT NULL REFERENCES dictionary_item(id),
    diary_entry_id BIGINT NOT NULL REFERENCES diary_entry(id) ON DELETE CASCADE
);
