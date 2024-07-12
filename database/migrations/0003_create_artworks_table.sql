CREATE TABLE if not exists artworks (
    id SERIAL PRIMARY KEY,
    museum_id INTEGER NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    image BYTEA,
    author VARCHAR(100),
    active BOOLEAN DEFAULT true,
    CONSTRAINT fk_museum FOREIGN KEY (museum_id) REFERENCES museums(id)
);
