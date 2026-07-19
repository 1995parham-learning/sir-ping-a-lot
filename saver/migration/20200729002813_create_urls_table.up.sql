CREATE TABLE IF NOT EXISTS urls
(
    id serial PRIMARY KEY ,
    user_id INT NOT NULL ,
    url VARCHAR (250) NOT NULL ,
    period INT NOT NULL ,
    FOREIGN KEY (user_id) REFERENCES users (id)
);