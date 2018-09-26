package db

import (
	"log"
	"testing"
)

const (
	DB_URL = "szlocaldb:szcrf123@tcp(10.193.1.19:3306)/test?charset=utf8&parseTime=True&loc=Local"
)

//func TestUpdate_ErrCount(t *testing.T) {
//	valMap := map[string]interface{}{"user_name": "hel11lo"}
//	Update_ErrCount("user_id=?", "3232", valMap, DB_URL)
//}

//func TestGet_User(t *testing.T) {
//	user, _ := Get_User("admin1", DB_URL)
//	v, ok := user.(*comm.Wait_User)
//	log.Println(ok)

//	if ok {
//		log.Println(v.Login_name)
//	}
//}

func TestGet_UserList(t *testing.T) {
	check_code, _ := Get_UserList(DB_URL)
	log.Println(check_code)
}
