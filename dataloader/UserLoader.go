package dbl

import (
	"time"
	"fmt"
	
	"github.com/nu7hatch/gouuid"
	"github.com/goldmoment/manager"
)

func ValidateUser(username string, password string) string {
    userid := "nil"
	db.Database.QueryRow("SELECT id FROM users WHERE (username=? AND password=?)", username, password).Scan(&userid)
	return userid
}

func RegisterUser(username string, password string) bool {
	stmt, err := db.Database.Prepare(`INSERT INTO users(username, password, id) VALUES(?,?,?)`)
	if err != nil {
		return false
	}
	defer stmt.Close()
	
	id, err := uuid.NewV4()
    if  err != nil {
		return false
    }
	if _, err := stmt.Exec(username, password, id.String()); err != nil {
		return false
    }
    
    return true
}

func UpdateUser(userID string, token string, expired time.Time, role string) bool {
	stmt, err := db.Database.Prepare(`UPDATE users SET token = ?, expired = ?, role = ?, modified_at = ? WHERE id = ?`)
	if err != nil {
		return false
	}
	defer stmt.Close()
	
	// Execute sql query
	if _, err := stmt.Exec(token, expired, role, time.Now(), userID); err != nil {
		return false
		fmt.Print(err)
    }
    return true
}