package api

import (
	"context"
	db "task/database"
	"task/src/api/schemas"
	"task/src/utils"

	"github.com/georgysavva/scany/v2/pgxscan"
)

func get_one_(id int) schemas.User {
	var user schemas.User

	db.DB.QueryRow(context.Background(), get_one_query, id).Scan(
		&user.ID, &user.Name, &user.Surname, &user.Patronymic,
		&user.Age, &user.Gender, &user.Nationality,
	)

	return user
}

func get_all_(
	offset int, limit int, order_by string, desc string,
) ([]schemas.User, error) {
	var users []schemas.User
	var by string
	var pagination = " OFFSET $1 LIMIT $2"

	if order_by == "id" {
		by = " ORDER BY u.id " + desc
	} else if order_by == "name" {
		by = " ORDER BY u.name " + desc
	} else if order_by == "surname" {
		by = " ORDER BY u.surname " + desc
	} else {
		by = " ORDER BY u.age " + desc
	}

	scan_err := pgxscan.Select(
		context.Background(), db.DB, &users, get_all_query+by+pagination,
		offset, limit,
	)

	if scan_err != nil {
		return nil, scan_err
	}

	return users, nil
}

func create_(user utils.Result) error {
	var gender_id int

	if user.Gender == "male" {
		gender_id = 1
	} else {
		gender_id = 2
	}

	_, err := db.DB.Exec(
		context.Background(), create_query,
		user.Name, user.Surname, user.Patronymic,
		user.Age, gender_id, user.Nationality,
	)

	if err != nil {
		return err
	}

	return nil
}

func update_(user schemas.Update) error {
	var gender_id int

	if user.Gender == "male" {
		gender_id = 1
	} else {
		gender_id = 2
	}

	_, err := db.DB.Exec(
		context.Background(), update_query,
		user.Name, user.Surname, user.Patronymic,
		user.Age, gender_id, user.Nationality, user.ID,
	)

	if err != nil {
		return err
	}

	return nil
}

func remove_(id int) (int64, error) {
	res, err := db.DB.Exec(context.Background(), delete_query, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
