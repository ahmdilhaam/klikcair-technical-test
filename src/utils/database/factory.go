package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (conn *Connection) List() []interface{} {
	db, dbErr := conn.Open()

	if dbErr != nil {
		fmt.Printf("error when open db = %s\n", dbErr.Error())
	}

	var data []interface{}

	results, err := db.Query("SELECT id, name FROM users")
	conn.Close(db)

	if err != nil {
		fmt.Printf("error when get user list = %s\n", err.Error())
		return data
	}

	for results.Next() {
		var user User
		// for each row, scan the result into our user composite object
		err = results.Scan(&user.ID, &user.Name)
		if err != nil {
			fmt.Printf("error when pointing = %s\n", err.Error())
			continue
		}
		// and then print out the user's Name attribute
		data = append(data, map[string]interface{}{
			"ID":   user.ID,
			"Name": user.Name,
		})
	}

	return data
}

func (conn *Connection) CheckBalance(phoneNumber int) interface{} {
	db, dbErr := conn.Open()

	if dbErr != nil {
		fmt.Printf("error when open db = %s\n", dbErr.Error())
	}

	var user User

	err := db.QueryRow("SELECT balance FROM users where phone_number = ?", phoneNumber).Scan(&user.Balance)
	conn.Close(db)

	if err != nil {
		fmt.Printf("error when get user detail = %s\n", err.Error())
		return nil
	}

	return map[string]interface{}{
		"balance": user.Balance,
	}
}

func (conn *Connection) Withdrawal(phoneNumber int, amount float64) string {
	db, dbErr := conn.Open()

	if dbErr != nil {
		fmt.Printf("error when open db = %s\n", dbErr.Error())
	}

	var user User

	err := db.QueryRow("SELECT id, balance FROM users where phone_number = ?", phoneNumber).Scan(&user.ID, &user.Balance)

	if err != nil {
		fmt.Printf("error when get user detail = %s\n", err.Error())
		return "error"
	}

	total := user.Balance - amount
	if total < 0 {
		return "insufficient balance"
	}

	_, err = db.Exec("insert into transcations(type, amount, user_id) values (?, ?, ?)", "withdrawal", amount, user.ID)
	if err != nil {
		fmt.Printf("error when store to transactions = %s\n", err.Error())
		return "error"
	}

	_, err = db.Exec("update users set balance = ? where id = ?", total, user.ID)
	if err != nil {
		fmt.Printf("error when update balance = %s\n", err.Error())
		return "error"
	}

	conn.Close(db)

	return "success to balance withdrawal"
}
