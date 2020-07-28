CREATE TABLE IF NOT EXISTS statuses
(
    id serial PRIMARY KEY ,
    url_id INT NOT NULL ,
    clock TIMESTAMP NOT NULL ,
    status_code INT NOT NULL ,
    FOREIGN KEY (url_id) REFERENCES urls (id)
);

CREATE OR REPLACE FUNCTION delete_expired_row()
RETURNS TRIGGER AS
    $BODY$
        BEGIN
        DELETE FROM statuses WHEN clock < NOW() - INTERVAL '2 days';
        RETURN NULL;
        END;
    $BODY$
    LANGUAGE plpgsql;
    CREATE TRIGGER delete_expired_rows
    AFTER INSERT
    ON statuses
    FOR EACH ROW
    EXECUTE PROCEDURE delete_expired_row();