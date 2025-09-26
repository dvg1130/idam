package authdb

//existing user

var UserExists = "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"

var RegisterUser = "INSERT INTO users (username, password) VALUES (?, ?)"
