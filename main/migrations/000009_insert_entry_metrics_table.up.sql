INSERT INTO entry_metric (metric_type_id, diary_entry_id)
SELECT
    (SELECT id FROM dictionary_item WHERE type = 'METRIC_NAME' AND label = 'Activity minutes'),
    (SELECT id FROM diary_entry WHERE description = 'Morning walk before lectures');

INSERT INTO entry_metric (metric_type_id, diary_entry_id)
SELECT
    (SELECT id FROM dictionary_item WHERE type = 'METRIC_NAME' AND label = 'Activity minutes'),
    (SELECT id FROM diary_entry WHERE description = 'Quick workout and shower');

INSERT INTO entry_metric (metric_type_id, diary_entry_id)
SELECT
    (SELECT id FROM dictionary_item WHERE type = 'METRIC_NAME' AND label = 'Water intake'),
    (SELECT id FROM diary_entry WHERE description = 'Lunch break and hydration');

INSERT INTO entry_metric (metric_type_id, diary_entry_id)
SELECT
    (SELECT id FROM dictionary_item WHERE type = 'METRIC_NAME' AND label = 'Activity minutes'),
    (SELECT id FROM diary_entry WHERE description = 'Evening football with friends');

INSERT INTO entry_metric (metric_type_id, diary_entry_id)
SELECT
    (SELECT id FROM dictionary_item WHERE type = 'METRIC_NAME' AND label = 'Water intake'),
    (SELECT id FROM diary_entry WHERE description = 'Tea break and stretching');
