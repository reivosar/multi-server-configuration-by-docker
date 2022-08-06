USE mysql;

/** CREATE TABLE **/
CREATE TABLE IF NOT EXISTS users (
    id          int AUTO_INCREMENT PRIMARY KEY,
    username    VARCHAR ( 50 ) UNIQUE NOT NULL,
    password    VARCHAR ( 50 ) NOT NULL,
    email VARCHAR ( 255 ) UNIQUE NOT NULL,
    created_on TIMESTAMP NOT NULL
);