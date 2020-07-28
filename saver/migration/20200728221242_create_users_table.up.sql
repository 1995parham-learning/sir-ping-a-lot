CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY ,
    email VARCHAR (250) NOT NULL ,
    password VARCHAR (250) NOT NULL
);