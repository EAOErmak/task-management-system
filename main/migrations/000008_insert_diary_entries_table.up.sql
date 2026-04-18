INSERT INTO diary_entry (user_id, when_started, when_ended, duration, mood, description)
VALUES (
    (SELECT id FROM users WHERE username = 'seed_alina'),
    '2026-04-01 06:30:00+00',
    '2026-04-01 07:15:00+00',
    45,
    5,
    'Morning walk before lectures'
);

INSERT INTO diary_entry (user_id, when_started, when_ended, duration, mood, description)
VALUES (
    (SELECT id FROM users WHERE username = 'seed_bekzat'),
    '2026-04-02 08:00:00+00',
    '2026-04-02 08:20:00+00',
    20,
    4,
    'Quick workout and shower'
);

INSERT INTO diary_entry (user_id, when_started, when_ended, duration, mood, description)
VALUES (
    (SELECT id FROM users WHERE username = 'seed_dana'),
    '2026-04-03 12:10:00+00',
    '2026-04-03 12:40:00+00',
    30,
    3,
    'Lunch break and hydration'
);

INSERT INTO diary_entry (user_id, when_started, when_ended, duration, mood, description)
VALUES (
    (SELECT id FROM users WHERE username = 'seed_ermek'),
    '2026-04-04 18:00:00+00',
    '2026-04-04 18:50:00+00',
    50,
    5,
    'Evening football with friends'
);

INSERT INTO diary_entry (user_id, when_started, when_ended, duration, mood, description)
VALUES (
    (SELECT id FROM users WHERE username = 'seed_madina'),
    '2026-04-05 21:00:00+00',
    '2026-04-05 21:15:00+00',
    15,
    4,
    'Tea break and stretching'
);
