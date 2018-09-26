package business

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"waiting/comm"
	"waiting/db"

	"github.com/mmcloughlin/geohash"
)

/*
	1、保存用户的评论内容
*/
func Create_Comment(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "Comment......")
	t1 := time.Now()

	if r.Method == "OPTIONS" {
	}

	var post_comment comm.Wait_PostComment_Request
	err := json.NewDecoder(r.Body).Decode(&post_comment)
	log.Println(r.Body)
	if err != nil {
		log.Println("Error===>", err.Error())
		//Write_Response(comm.RESP_JSON_ERROR, w, r)
		//return
	}
	log.Println(post_comment)
	defer r.Body.Close()

	lng, _ := strconv.ParseFloat(post_comment.Lng, 64)
	lat, _ := strconv.ParseFloat(post_comment.Lat, 64)
	getval := geohash.EncodeInt(lat, lng)

	//log.Println("geohash=======>", geohash.)
	create_comment := comm.Wait_Users_Comment{
		User_id:     post_comment.User_id,
		Comment_no:  strconv.FormatInt(time.Now().UnixNano(), 10),
		Device_type: post_comment.Device_type,
		Device_ip:   post_comment.Device_ip,
		Lng:         post_comment.Lng,
		Lat:         post_comment.Lat,
		Geo_hash:    fmt.Sprintf("%d", getval),
		Comment_msg: post_comment.Comment_msg,
		Insert_time: time.Now().Unix(),
		Update_time: time.Now().Unix()}
	err = db.Create_comment(create_comment)
	if err != nil {
		Write_Response(comm.RESP_DB_ERROR, w, r)
		log.Println(comm.END_TAG, "Comment Error......", time.Since(t1))
		return
	}
	Write_Response(comm.RESP_SUCC, w, r)
	log.Println(comm.END_TAG, "Comment Succ......", time.Since(t1))
	return
}

/*
	1、 获取用户评论列表
*/

func Get_CommentList(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "GetCommentList......")
	t1 := time.Now()
	var getcomment_req comm.Wait_GetCommentList_Request
	log.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&getcomment_req)
	if err != nil {
		log.Println(err.Error())
		Write_Response(comm.RESP_JSON_ERROR, w, r)
		return
	}
	defer r.Body.Close()
	commentList, err := db.Get_CommentList("")
	if err != nil {
		Write_Response(comm.RESP_DB_ERROR, w, r)
		log.Println(comm.END_TAG, "GetCommentList......", time.Since(t1))
		return
	}
	getCommentList_resp := comm.Wait_GetCommentList_Response{comm.RESP_SUCC.Status_code, comm.RESP_SUCC.Status_msg, commentList}
	log.Println("Get_CommentList==", getCommentList_resp)
	Write_Response(getCommentList_resp, w, r)
	log.Println(comm.END_TAG, "GetCommentList......", time.Since(t1))
}

/*
	1、 根据评论编号获取评论内容
*/

func Get_CommentByNo(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "Get_CommentByNo......")
	t1 := time.Now()
	var comment_req comm.Wait_GetComment_Request
	err := json.NewDecoder(r.Body).Decode(&comment_req)
	if err != nil {
		log.Println(err.Error())
		Write_Response(comm.RESP_JSON_ERROR, w, r)
		return
	}
	defer r.Body.Close()

	log.Println("reqIno", comment_req)
	v, err := db.Get_CommentByNo(comment_req.Comment_no)
	if err != nil {
		Write_Response(comm.RESP_DB_ERROR, w, r)
		log.Println(comm.END_TAG, "Get_CommentByNo......", time.Since(t1))
		return
	}
	comment_resp := comm.Wait_GetComment_Response{comm.RESP_SUCC.Status_code, comm.RESP_SUCC.Status_msg, *v}
	log.Println("Get_CommentByNo==", comment_resp)
	Write_Response(comment_resp, w, r)
	log.Println(comm.END_TAG, "Get_CommentByNo......", time.Since(t1))
}

/*
	1、 根据评论编号获取评论内容
*/

func Get_Test(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "Get_CommentByNo......")
	log.Println(comm.END_TAG, "Get_CommentByNo......")
	log.Println(r.Method)
	log.Println("Header---->", r.Header)
	log.Println("Authorization---->", r.Header.Get("Authorization"))
	log.Println("Accept------>", r.Header.Get("Accept"))
	log.Println("Accept------>", r.Header.Get("Module"))
	getCommentList_resp := comm.Wait_GetCommentList_Response{comm.RESP_SUCC.Status_code, comm.RESP_SUCC.Status_msg, nil}
	Write_Response(getCommentList_resp, w, r)
}
