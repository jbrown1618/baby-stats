INSERT OR IGNORE INTO version_history (version) 
VALUES (1);

INSERT INTO user (username)
VALUES ('testuser');

INSERT INTO baby (user_id, name, birth_date)
VALUES ((SELECT user_id FROM user WHERE username = 'testuser'), 'Miles', DATETIME('2024-01-20'));

INSERT INTO event (baby_id, type, start_time)
VALUES ((SELECT baby_id FROM baby WHERE name = 'Miles'), 'feeding', 1719968822),
       ((SELECT baby_id FROM baby WHERE name = 'Miles'), 'feeding', 1719968810);

INSERT INTO measurement (event_id, amount, unit)
VALUES ((SELECT event_id FROM event WHERE start_time = 1719968822), 9.5, 'oz'),
       ((SELECT event_id FROM event WHERE start_time = 1719968810), 8.5, 'oz')