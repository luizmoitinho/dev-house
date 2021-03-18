CREATE DATABASE IF NOT EXISTS devhouse;

USE devhouse;

DROP TABLE IF EXISTS tb_users;

CREATE TABLE tb_users(
  user_id INT auto_increment PRIMARY KEY,
  name VARCHAR(250) NOT NULL,
  login VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(100) NOT NULL UNIQUE,
  password TEXT not null,
  created_at timestamp default now()
)ENGINE=INNODB 
