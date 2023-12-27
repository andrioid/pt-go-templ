-- +goose Up
CREATE TABLE user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	passwd TEXT NOT NULL,
	is_admin INTEGER
);

CREATE TABLE user_token (
	token TEXT PRIMARY KEY NOT NULL, -- uuid
	user_id INTEGER NOT NULL,
	created_on DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	expires_on DATETIME NOT NULL, -- Long lived, can be renewed
	client_hash TEXT, -- To prevent session hijacking
	FOREIGN KEY(user_id) REFERENCES user(id)	
);

CREATE TABLE sessions (
	token TEXT PRIMARY KEY,
	data BLOB NOT NULL,
	expiry REAL NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions(expiry);

-- +goose Down
DROP TABLE user;
DROP TABLE user_token;
DROP TABLE sessions;