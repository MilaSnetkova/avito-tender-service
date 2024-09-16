-- Таблица для тендеров
CREATE TABLE tenders (
    id SERIAL PRIMARY KEY,
    organization_id INT REFERENCES organization(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'CREATED',
    version INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Таблица для предложений
CREATE TABLE offers (
    id SERIAL PRIMARY KEY,
    tender_id INT REFERENCES tenders(id) ON DELETE CASCADE,
    author_id INT REFERENCES employee(id) ON DELETE CASCADE,
    status VARCHAR(20) NOT NULL DEFAULT 'CREATED',
    version INT NOT NULL DEFAULT 1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);