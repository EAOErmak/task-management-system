INSERT INTO entry_metric_value (entry_metric_id, unit_id, value)
SELECT
    em.id,
    (SELECT id FROM dictionary_item WHERE type = 'METRIC_UNIT' AND label = 'minutes'),
    45
FROM entry_metric em
JOIN diary_entry de ON de.id = em.diary_entry_id
JOIN dictionary_item di ON di.id = em.metric_type_id
WHERE de.description = 'Morning walk before lectures'
  AND di.type = 'METRIC_NAME'
  AND di.label = 'Activity minutes';

INSERT INTO entry_metric_value (entry_metric_id, unit_id, value)
SELECT
    em.id,
    (SELECT id FROM dictionary_item WHERE type = 'METRIC_UNIT' AND label = 'minutes'),
    20
FROM entry_metric em
JOIN diary_entry de ON de.id = em.diary_entry_id
JOIN dictionary_item di ON di.id = em.metric_type_id
WHERE de.description = 'Quick workout and shower'
  AND di.type = 'METRIC_NAME'
  AND di.label = 'Activity minutes';

INSERT INTO entry_metric_value (entry_metric_id, unit_id, value)
SELECT
    em.id,
    (SELECT id FROM dictionary_item WHERE type = 'METRIC_UNIT' AND label = 'milliliters'),
    750
FROM entry_metric em
JOIN diary_entry de ON de.id = em.diary_entry_id
JOIN dictionary_item di ON di.id = em.metric_type_id
WHERE de.description = 'Lunch break and hydration'
  AND di.type = 'METRIC_NAME'
  AND di.label = 'Water intake';

INSERT INTO entry_metric_value (entry_metric_id, unit_id, value)
SELECT
    em.id,
    (SELECT id FROM dictionary_item WHERE type = 'METRIC_UNIT' AND label = 'minutes'),
    50
FROM entry_metric em
JOIN diary_entry de ON de.id = em.diary_entry_id
JOIN dictionary_item di ON di.id = em.metric_type_id
WHERE de.description = 'Evening football with friends'
  AND di.type = 'METRIC_NAME'
  AND di.label = 'Activity minutes';

INSERT INTO entry_metric_value (entry_metric_id, unit_id, value)
SELECT
    em.id,
    (SELECT id FROM dictionary_item WHERE type = 'METRIC_UNIT' AND label = 'glasses'),
    2
FROM entry_metric em
JOIN diary_entry de ON de.id = em.diary_entry_id
JOIN dictionary_item di ON di.id = em.metric_type_id
WHERE de.description = 'Tea break and stretching'
  AND di.type = 'METRIC_NAME'
  AND di.label = 'Water intake';
