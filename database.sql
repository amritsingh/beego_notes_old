/* As root user */                                                                                                                                                                                                 
CREATE USER 'beego_notes'@'localhost' IDENTIFIED BY 'tmp_pwd';
CREATE DATABASE beego_notes;
GRANT ALL PRIVILEGES ON beego_notes . * TO 'beego_notes'@'localhost';
 
