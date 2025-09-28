package datadb

// snakes
var GetSuid = "SELECT suid FROM snakes WHERE uuid = ? AND sid = ?"

var GetSnakes = "SELECT sid, species FROM snakes WHERE uuid = ?"

var GetSnake = "SELECT sid, species, sex, age, genes, notes FROM snakes WHERE sid = ? AND owner_uuid = ?"

var AddSnake = "INSERT INTO snakes (owner_uuid, sid, species, sex, age, genes, notes) VALUES (?, ?, ?, ?, ?, ?, ?)"

var UpdateSnake = "UPDATE snakes SET %s WHERE sid = ? AND owner_uuid = ?"

var DeleteSnake = "DELETE FROM snakes WHERE sid = ? AND uuid = ?"

var SnakeExists = "SELECT EXISTS(SELECT 1 FROM snakes WHERE sid = ? AND owner_uuid = ?)"

// feed
var AddSnakeFeed = "INSERT INTO feeding(suid, sid, feed_date, prey_type, prey_size, notes) VALUES (?, ?, ?, ?, ?, ?)"

var UpdateSnakeFeed = "UPDATE feeding SET %s WHERE snake_uuid = ? AND feed_date = ?"

var DeleteSnakeFeed = "DELETE FROM feeding WHERE suid = ? AND feed_date = ?"

// health
var GetSnakeHealth = "SELECT check_date, weight, length, topic, notes FROM health WHERE suid = ? ORDER BY check_date DESC"

var PostSnakeHealth = "(suid, uuid, sid, check_date, weight, length, topic, notes) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

var UpdateSnakeHealth = "UPDATE health SET %s WHERE suid = ? AND check_date = ?"

var DeleteSnakeHealth = "DELETE FROM health WHERE suid = ? AND check_date = ?"
