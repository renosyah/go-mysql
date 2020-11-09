CREATE TABLE user(
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name TEXT,
    email TEXT,
    created_at DATETIME NOT NULL DEFAULT now(),
    flag_status INT
);