package db

import (
	"log"
	"waiting/comm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
	保存用户的评论内容
*/
func Create_comment(comment comm.Wait_Users_Comment) error {
	db, err := connect_db()
	if err != nil {
		log.Println("error", err.Error())
		return err
	}
	defer db.Close()
	ret := db.Create(&comment)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

/*
	获取用户评论列表
*/
func Get_CommentList(dbUrl string) ([]comm.Wait_Comment_Info, error) {
	db, err := connect_db()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	commentList := make([]comm.Wait_Comment_Info, 10, 10)
	rows := db.Select("comment_no,lng,lat,comment_msg").Find(&commentList)
	if rows.Error != nil {
		if rows.RecordNotFound() {
			return nil, nil
		} else {
			return nil, rows.Error
		}
	}
	return commentList, nil
}

/*
 根据评论编号得到评论信息
*/
func Get_CommentByNo(commentNo string) (*comm.Wait_Comment_Info, error) {
	db, err := connect_db()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	comment := new(comm.Wait_Comment_Info)
	rows := db.Where("comment_no = ?", commentNo).First(&comment)
	if rows.Error != nil {
		if rows.RecordNotFound() {
			return nil, nil
		} else {
			return nil, rows.Error
		}
	}
	return comment, nil
}
