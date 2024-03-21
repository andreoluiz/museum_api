CREATE SCHEMA IF NOT EXISTS virtualArtExplore;

CREATE TABLE IF NOT EXISTS gerente (
    id SERIAL PRIMARY KEY,
    Nome VARCHAR(100) NOT NULL,
    data_nascimento DATE,
    cpf VARCHAR(14) UNIQUE
);

CREATE TABLE IF NOT EXISTS museu (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    localizacao VARCHAR(200),
    descricao TEXT,
    gerente_id INTEGER REFERENCES gerente(id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS obras (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    autor VARCHAR(100) NOT NULL,
    descricao TEXT,
    data_criacao DATE,
    museu_id INTEGER REFERENCES museu(id) ON DELETE SET NULL
);
