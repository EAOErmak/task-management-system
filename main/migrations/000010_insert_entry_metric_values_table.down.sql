DELETE FROM entry_metric_value
WHERE entry_metric_id IN (
    SELECT em.id
    FROM entry_metric em
    JOIN diary_entry de ON de.id = em.diary_entry_id
    WHERE de.description IN (
        'Morning walk before lectures',
        'Quick workout and shower',
        'Lunch break and hydration',
        'Evening football with friends',
        'Tea break and stretching'
    )
);
