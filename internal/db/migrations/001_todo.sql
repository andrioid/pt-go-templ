-- +goose Up
CREATE TABLE todo (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	finished INTEGER DEFAULT 0
);

-- +goose Down
DROP TABLE todo