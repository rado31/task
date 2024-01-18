package api

var get_one_query string = `
    SELECT
        u.id,
        u.name,
        u.surname,
        u.patronymic,
        u.age,
        g.title AS gender,
        u.nationality
    FROM users u
    INNER JOIN genders g ON g.id = u.genders_id
    WHERE u.id = $1
`

var get_all_query string = `
    SELECT
        u.id,
        u.name,
        u.surname,
        u.patronymic,
        u.age,
        g.title AS gender,
        u.nationality
    FROM users u
    INNER JOIN genders g ON g.id = u.genders_id
`

var create_query string = `
    INSERT INTO users (name, surname, patronymic, age, genders_id, nationality)
    VALUES ($1, $2, $3, $4, $5, $6)
`

var update_query string = `
    UPDATE users SET name = $1, surname = $2, patronymic = $3, age = $4,
    genders_id = $5, nationality = $6
    WHERE id = $7
`

var delete_query string = `DELETE FROM users WHERE id = $1`
