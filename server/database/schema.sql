CREATE TABLE IF NOT EXISTS version_history (
  version_history_id INTEGER PRIMARY KEY,
  version            INTEGER NOT NULL,
  as_of_date         TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS user (
  user_id  INTEGER PRIMARY KEY,
  username TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS baby (
  baby_id    INTEGER PRIMARY KEY,
  user_id    INTEGER NOT NULL,
  name       TEXT NOT NULL,
  birth_date INTEGER NOT NULL,

  FOREIGN KEY(user_id) REFERENCES user(user_id)
);

CREATE TABLE IF NOT EXISTS event (
  event_id    INTEGER PRIMARY KEY,
  baby_id     INTEGER NOT NULL,
  type        TEXT NOT NULL,
  start_time  INTEGER NOT NULL,
  end_time    INTEGER,
  notes       TEXT,

  FOREIGN KEY(baby_id) REFERENCES baby(baby_id)
);

CREATE TABLE IF NOT EXISTS measurement (
  measurement_id INTEGER PRIMARY KEY,
  event_id       INTEGER NOT NULL,
  amount         REAL NOT NULL,
  unit           TEXT NOT NULL,

  FOREIGN KEY(event_id) REFERENCES event(event_id)
);
