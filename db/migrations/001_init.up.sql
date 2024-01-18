CREATE TABLE users (
    id SERIAL, 
    jwt TEXT, 
    login TEXT NOT NULL UNIQUE, 
    password TEXT NOT NULL);

CREATE TABLE goods (
    id SERIAL, 
    description TEXT, 
    price INT NOT NULL DEFAULT 0);
    

\copy goods(description, price)  FROM '/home/igor/Desktop/code/gophermartuser/db/goods.csv' DELIMITER ';' CSV HEADER;
