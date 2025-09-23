CREATE TABLE IF NOT EXISTS users (
    uuid     CHAR(36) NOT NULL PRIMARY KEY DEFAULT (UUID()),
    username VARCHAR(255) NOT NULL PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    role     VARCHAR(255) NOT NULL DEFAULT 'basic'

);

CREATE TABLE IF NOT EXISTS users (
    uuid     CHAR(36) NOT NULL PRIMARY KEY DEFAULT (UUID()),
    ->     username VARCHAR(255) NOT NULL PRIMARY KEY,
    ->     password VARCHAR(255) NOT NULL,
    ->     role     VARCHAR(255) NOT NULL DEFAULT 'basic'
    -> );