CREATE TABLE Events (
    e_id bigserial PRIMARY KEY,
    c_id bigserial NOT NULL,
    start_time varchar NOT NULL,
    end_time varchar NOT NULL,
    title varchar NOT NULL,
    color varchar NOT NULL,
    FOREIGN KEY (c_id) REFERENCES Course (c_id) ON UPDATE CASCADE ON DELETE CASCADE
);