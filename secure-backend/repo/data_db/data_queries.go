package datadb

var GetSuid = "SELECT suid FROM snakes WHERE uuid = ? AND sid = ?"
var AddSnakeFeed = "INSERT INTO feeding(suid, sid, feed_date, prey_type, prey_size, notes) VALUES (?, ?, ?, ?, ?, ?)"
var UpdateSnakeFeed = "UPDATE feeding SET %s WHERE snake_uuid = ? AND feed_date = ?"
var DeleteSnakeFeed = "DELETE FROM feeding WHERE suid = ? AND feed_date = ?"
