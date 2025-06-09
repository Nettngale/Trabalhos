CREATE DATABASE IF NOT EXISTS cacau_show;
USE cacau_show;

CREATE TABLE clientes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    endereco VARCHAR(255) NOT NULL,
    cartao_credito VARCHAR(20) NOT NULL,
    cvv VARCHAR(5) NOT NULL,
    produto VARCHAR(100) NOT NULL,
	usuario VARCHAR(100) NOT NULL,
    senha VARCHAR(100) NOT NULL
);
