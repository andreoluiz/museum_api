CREATE TABLE if not exists managers (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    active BOOLEAN DEFAULT true
);
