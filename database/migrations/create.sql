CREATE TABLE genders (
    id SERIAL PRIMARY KEY,
    title VARCHAR
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    surname VARCHAR,
    patronymic VARCHAR,
    age INTEGER,
    genders_id INTEGER,
    nationality VARCHAR,

    CONSTRAINT genders_id
        FOREIGN KEY (genders_id)
            REFERENCES genders (id)
);
