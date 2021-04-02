CREATE DATABASE IF NOT EXISTS devhouse;

USE devhouse;

DROP TABLE IF EXISTS tb_users;
DROP TABLE IF EXISTS tb_followers;


CREATE TABLE tb_users(
  user_id INT auto_increment PRIMARY KEY,
  name VARCHAR(250) NOT NULL,
  login VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(100) NOT NULL UNIQUE,
  password TEXT not null,
  created_at timestamp default now()
)ENGINE=INNODB; 

CREATE TABLE tb_followers(
 user_id INT NOT NULL, 
 follower_id INT NOT NULL
)ENGINE=INNODB;

ALTER TABLE tb_followers ADD CONSTRAINT FK_followers_users_1
  FOREIGN KEY (user_id) REFERENCES tb_users(user_id) ON DELETE CASCADE;


ALTER TABLE tb_followers ADD CONSTRAINT FK_followers_users_2
  FOREIGN KEY (follower_id) REFERENCES tb_users(user_id) ON DELETE CASCADE;

ALTER TABLE tb_followers ADD CONSTRAINT PK_followers
  PRIMARY KEY(user_id, follower_id);
