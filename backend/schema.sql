-- Estrutura das tabelas
CREATE TABLE IF NOT EXISTS clientes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    telefone VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS barbeiros (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    telefone VARCHAR(20)
);

CREATE TABLE IF NOT EXISTS agendamentos (
    id SERIAL PRIMARY KEY,
    cliente_id INTEGER NOT NULL REFERENCES clientes(id),
    barbeiro_id INTEGER NOT NULL REFERENCES barbeiros(id),
    data_hora TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'agendado'
);

-- Dados de exemplo (seed)
INSERT INTO clientes (nome, telefone) VALUES
    ('João Silva', '35999998888'),
    ('Pedro Paulo', '35988889999'),
    ('Ana Costa', '35977776666');

INSERT INTO barbeiros (nome, telefone) VALUES
    ('Robson José', '35900001111'),
    ('Carlos Xavier', '35911110000');