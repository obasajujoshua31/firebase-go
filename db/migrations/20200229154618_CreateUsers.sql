
-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id SERIAL,
    uuid varchar(255) NOT NULL,
    email varchar(255) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    PRIMARY KEY(id)
);

CREATE INDEX idx_id ON users(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_id ON users;
DROP TABLE users;

-- +goose StatementEnd
