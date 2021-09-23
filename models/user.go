package models

import (
	"github.com/ferriantitus/rest-echo/db"
	"net/http"
	"strings"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Roles    int    `json:"roles"`
}

func FetchAllUserOrID(id string) (Response, error) {
	var obj User
	var arrobj []User
	var res Response

	if id != "" {
		id = strings.Trim(id, "/")
	}

	con := db.Createconf()

	if id != "" {
		sqlStatement := "SELECT * FROM user WHERE id=?"

		rows, err := con.Query(sqlStatement, id)
		defer rows.Close()
		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.Id, &obj.Username, &obj.Email, &obj.Roles)
			if err != nil {
				return res, err
			}
			arrobj = append(arrobj, obj)
		}

		res.Status = http.StatusOK
		res.Message = "success"
		res.Data = arrobj

		return res, nil

	} else {

		sqlStatement := "SELECT * FROM user"

		rows, err := con.Query(sqlStatement)
		defer rows.Close()

		if err != nil {
			return res, err
		}

		for rows.Next() {
			err = rows.Scan(&obj.Id, &obj.Username, &obj.Email, &obj.Roles)
			if err != nil {
				return res, err
			}

			arrobj = append(arrobj, obj)
		}

		res.Status = http.StatusOK
		res.Message = "success"
		res.Data = arrobj

		return res, nil
	}
}

func StoreUser(username string, email string, roles int) (Response, error) {
	var res Response

	con := db.Createconf()

	sqlStatement := "INSERT user (username, email, roles) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, email, roles)
	if err != nil {
		return res, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusCreated
	res.Message = "success"
	res.Data = map[string]int64{
		"last_inserted_id" : lastInsertedID,
	}

	return res, nil
}

func UpdateUser(id int, username string, email string, roles int) (Response, error){
	var res Response

	con := db.Createconf()

	sqlStatement := "UPDATE user SET username = ?, email = ?, roles = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(username, email, roles, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteUser(id int) (Response, error) {
	var res Response

	con := db.Createconf()

	sqlStatement := "DELETE FROM user WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}