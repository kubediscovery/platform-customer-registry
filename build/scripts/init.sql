SELECT 'CREATE DATABASE desafio'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'desafio');

CREATE USER clean WITH ENCRYPTED PASSWORD 'architecture';
GRANT ALL PRIVILEGES ON DATABASE desafio TO clean;

\c desafio;

CREATE TABLE "order" (
  id VARCHAR(36) PRIMARY KEY,
  name VARCHAR(100),
  description TEXT,
  stock INTEGER,
  price REAL,
  amount INTEGER,
  category VARCHAR(255)
);

INSERT INTO "order" (id, name, description, stock, price, amount, category) VALUES ('b34f408d-7067-4b84-8782-3c8e5b2f893d','item1', 'item 1', 4, 10.25, 3, 'categoria1');
INSERT INTO "order" (id, name, description, stock, price, amount, category) VALUES ('ed97a67d-76ac-4558-adc8-0942b1779df4','item2', 'item 1', 5, 10.26, 4, 'categoria2');
INSERT INTO "order" (id, name, description, stock, price, amount, category) VALUES ('f07dd0a9-4af4-46f6-9ca0-fbdee853e753','item3', 'item 2', 6, 10.27, 5, 'categoria3');
INSERT INTO "order" (id, name, description, stock, price, amount, category) VALUES ('31c027e4-692a-45eb-ac4e-e861fd8eb0ca','item4', 'item 3', 7, 10.28, 6, 'categoria4');