package dao

import (
	"read-api/infrastructure/dao"
	"read-api/model/user"
)

func selectAll() []user {
    db := db.Connect()
    defer db.Close()

    rows, err := db.Query("SELECT * FROM user")
    if err != nil {
        panic(err.Error())
    }

    results := make([]User, 0)
    for rows.Next() {
        var user user
        err = rows.Scan(&user.id, &user.username, &user.email)
        if err != nil {
            panic(err.Error())
        }
        results = append(results, user)
    }
    return results
}