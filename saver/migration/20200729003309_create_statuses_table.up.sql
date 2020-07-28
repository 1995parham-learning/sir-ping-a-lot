CREATE TABLE IF NOT EXISTS statuses
(
    id serial PRIMARY KEY ,
    url_id INT NOT NULL ,
    clock TIMESTAMP NOT NULL ,
    status_code INT NOT NULL ,
    FOREIGN KEY (url_id) REFERENCES urls (id)
);