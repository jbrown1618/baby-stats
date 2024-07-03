INSERT OR IGNORE INTO version_history (version) 
VALUES (1);

INSERT OR IGNORE INTO user (user_id, username)
VALUES (1, 'testuser');

INSERT OR IGNORE INTO baby (baby_id, user_id, name, birth_date)
VALUES (1, 1, 'Miles', DATETIME('2024-01-20'));

INSERT OR IGNORE INTO event (event_id, baby_id, type, start_time)
VALUES (1, 1, 'feeding', DATETIME('2024-01-20 10:00:00')),
       (2, 1, 'feeding', DATETIME('2024-01-20 13:00:00'));

INSERT OR IGNORE INTO measurement (measurement_id, event_id, amount, unit)
VALUES (1, 1, 9.5, 'oz'),
       (2, 2, 8.5, 'oz');