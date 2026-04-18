CREATE TABLE entry_metric_value (
    id BIGSERIAL PRIMARY KEY,
    entry_metric_id BIGINT NOT NULL REFERENCES entry_metric(id) ON DELETE CASCADE,
    unit_id BIGINT NOT NULL REFERENCES dictionary_item(id),
    value INTEGER NOT NULL
);
