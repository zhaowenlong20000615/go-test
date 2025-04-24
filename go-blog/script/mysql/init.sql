create database goblog;

USE goblog;

create table IF NOT EXISTS `user` (
uid INT AUTO_INCREMENT PRIMARY KEY,
username VARCHAR(50) NOT NULL,
avatar VARCHAR(100) NOT NULL,
passwd VARCHAR(100) NOT NULL,
create_at DATE,
update_at DATE
)
