package main

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func getAllUsers() []User {
	db := getDBConnection()
	defer db.Close()

	rows, err := db.Query("SELECT id, username, email FROM users")
	if err != nil {
		panic(err.Error())
	}

	var results []User
	results = make([]User, 0)
	for rows.Next() {
		var user User
		err = rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		results = append(results, user)
	}
	return results
}
