DELETE FROM dictionary_item
WHERE (type, label) IN (
    ('METRIC_NAME', 'Activity minutes'),
    ('METRIC_NAME', 'Water intake'),
    ('METRIC_UNIT', 'minutes'),
    ('METRIC_UNIT', 'milliliters'),
    ('METRIC_UNIT', 'glasses')
);
