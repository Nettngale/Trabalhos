CREATE DATABASE cadastro_cliente;

USE cadastro_cliente;

CREATE TABLE clientes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100),
    telefone VARCHAR(20),
    renda VARCHAR(100),
    data_nascimento DATE,
    cep VARCHAR(10)
);