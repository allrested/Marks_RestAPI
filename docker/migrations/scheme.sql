-- migrations/migration.sql

CREATE TABLE IF NOT EXISTS records (
    id INT AUTO_INCREMENT PRIMARY KEY,
    names VARCHAR(255) NOT NULL,
    marks JSON,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);