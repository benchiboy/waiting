package business

import (
	"testing"
	"time"
)

func TestGet_UserList(t *testing.T) {
	check_code, _ := Get_UserList(DB_URL)
	log.Println(check_code)
}
