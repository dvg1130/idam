package authdb

//existing user

var UserExists = "SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)"

var RegisterUser = "INSERT INTO users (username, password) VALUES (?, ?)"

var LoginUser = "SELECT password, role, uuid FROM users Where username = ?"

var AdminGetUsers = "SELECT uuid, username, role FROM users"
