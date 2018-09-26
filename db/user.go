package db

import (
	"fmt"
	"log"
	"time"
	"waiting/comm"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
	公共连接数据库方法
*/
func connect_db() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", comm.ConfigNode.DbUrl)
	if err != nil {
		log.Println("open database error:", err)
		return nil, fmt.Errorf(err.Error())
	}
	db.LogMode(true)
	return db, nil
}

/*
	新增注册用户
*/
func Create_User(user comm.Wait_User) error {
	db, err := connect_db()
	if err != nil {
		return err
	}
	defer db.Close()
	ret := db.Create(&user)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

/*
	根据登录账号查询用户信息
*/
func Get_User(loginName string) (interface{}, error) {
	db, err := connect_db()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	user := new(comm.Wait_User)
	rows := db.Where("login_name = ?", loginName).First(&user)
	if rows.Error != nil {
		if rows.RecordNotFound() {
			return nil, nil
		} else {
			return nil, rows.Error
		}
	}
	return user, nil
}

/*
	获取用户列表，账号名称、用户姓名、图片URL
*/
func Get_UserList() ([]comm.Wait_User, error) {
	db, err := connect_db()
	if err != nil {
		return nil, err
	}

	defer db.Close()
	userList := make([]comm.Wait_User, 10, 10)
	rows := db.Select("login_name,user_name,pic_full").Find(&userList)
	if rows.Error != nil {
		if rows.RecordNotFound() {
			return nil, nil
		} else {
			return nil, rows.Error
		}
	}
	return userList, nil
}

/*
	获取用户列表
*/
func Get_UserList2(dbUrl string) ([]interface{}, error) {
	db, err := connect_db()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	userList := make([]interface{}, 10, 10)
	rows := db.Select("login_name").Find(&userList)
	if rows.Error != nil {
		if rows.RecordNotFound() {
			return nil, nil
		} else {
			return nil, rows.Error
		}
	}
	return userList, nil
}

/*
	根据手机得到生效的验证码
*/
func Get_CheckCode(phoneNo string, checkCode string, checkType string) (interface{}, error) {
	db, err := connect_db()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	check_code := new(comm.Wait_Checkcode)

	nowTime := time.Now().Format("2006-01-02 15:04:05")

	rows := db.Where("user_phone = ? and check_code=?  and check_type=? and ? BETWEEN valid_btime and valid_etime and code_status='e' ", phoneNo, checkCode, checkType, nowTime).First(&check_code)
	if rows.Error != nil {
		return nil, rows.Error
	}
	return check_code, nil
}

/*
	生成短信验证码
*/
func Create_CheckCode(check_code comm.Wait_Checkcode) error {
	db, err := connect_db()
	if err != nil {
		return err
	}
	defer db.Close()
	tx := db.Begin()
	ret := db.Create(&check_code)
	if ret.Error != nil {
		tx.Rollback()
		return ret.Error
	}
	tx.Commit()
	return nil
}

/*
	得到最近一次发送验证码的时间
*/

func Get_LastCheckCodeTime(phoneNo string) (time.Time, error) {
	db, err := connect_db()
	if err != nil {
		return time.Now(), err
	}
	defer db.Close()
	check_code := new(comm.Wait_Checkcode)
	rows := db.Where("user_phone = ? and  code_status='e' ", phoneNo).Last(&check_code)
	if rows.Error != nil {
		if rows.RecordNotFound() {
			return time.Now().Add(time.Duration(-time.Minute * comm.CHECK_CODE_EXPIRED_MINUTE)), nil
		} else {
			return time.Now(), rows.Error
		}
	}
	return check_code.Insert_time, nil
}

/*
	根据用户账号修改用户密码
*/

func Change_User_Pwd(loginName string, newPwd string) error {
	db, err := connect_db()
	if err != nil {
		return err
	}
	defer db.Close()
	var tmp_user comm.Wait_User
	tx := db.Begin()
	if err = tx.Model(&tmp_user).Where("login_name=?", loginName).Update(map[string]interface{}{"login_pwd": newPwd}).Error; err != nil {
		log.Println("ERROR", err.Error())
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

/*
	如果用户密码错误，更新账户表的密码错误次数
*/

func Update_User_PwdErr(user comm.Wait_User) error {
	db, err := connect_db()
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.Begin()
	pwderr_count := user.Pwderr_count + 1
	if pwderr_count > comm.PWD_MAX_ERR_COUNT {
		if err = tx.Model(&user).Where("login_name=?", user.Login_name).Update(map[string]interface{}{"user_islock": 1, "pwderr_count": pwderr_count}).Error; err != nil {
			log.Println("ERROR", err.Error())
			tx.Rollback()
		}
	} else {
		if err = tx.Model(&user).Where("login_name=?", user.Login_name).Update(map[string]interface{}{"pwderr_count": pwderr_count}).Error; err != nil {
			log.Println("ERROR", err.Error())
			tx.Rollback()
		}
	}
	tx.Commit()
	return nil
}

/*
	更新用户上传的图片URL
*/

func Update_User_HeadPic(user comm.Wait_User) error {
	db, err := connect_db()
	if err != nil {
		return err
	}
	defer db.Close()
	tx := db.Begin()
	if err = tx.Model(&user).Where("login_name=?", user.Login_name).Update(map[string]interface{}{"pic_head": user.Pic_head}).Error; err != nil {
		log.Println("ERROR", err.Error())
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

/*
 更新用户最新的一次图形验证码
*/

func Update_User_PicCode(user comm.Wait_User) error {
	db, err := connect_db()
	if err != nil {
		return err
	}
	defer db.Close()
	tx := db.Begin()
	if err = tx.Model(&user).Where("login_name=?", user.Login_name).Update(map[string]interface{}{"last_pic_code": user.Last_pic_code}).Error; err != nil {
		log.Println("ERROR", err.Error())
		tx.Rollback()
	}
	tx.Commit()
	return nil
}
