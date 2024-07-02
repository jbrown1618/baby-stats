CREATE TABLE IF NOT EXISTS welcome (
  id      ROWID,
  message TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS feeding_event (
  id ROWID,
  end_time INTEGER,
  amount REAL,
  notes TEXT
);