CREATE TABLE if not exists museums (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    image BYTEA,
    category1 VARCHAR(100),
    category2 VARCHAR(100),
    link VARCHAR(255),
    address VARCHAR(255),
    cep VARCHAR(20),
    city VARCHAR(100),
    state VARCHAR(100),
    information TEXT,
    manager_id INTEGER NOT NULL,
    active BOOLEAN DEFAULT true,
    CONSTRAINT fk_manager FOREIGN KEY (manager_id) REFERENCES managers(id)
);
