-- +goose Up
CREATE TABLE user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	passwd TEXT NOT NULL,
	hash_type TEXT NOT NULL,
	is_email_verified INTEGER NOT NULL default 0,
	is_admin INTEGER NOT NULL default 0
);

-- +goose Down
DROP TABLE user;