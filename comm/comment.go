package comm

/*
	===============================================================================
							用户评论内容数据表结构
	===============================================================================
*/

//////////////////////////////////////////////////////////////////////////////////
/*
	用户评论表结构
*/
type Wait_Users_Comment struct {
	Auto_id      int64  `gorm:"column:auto_id;primary_key:auto_id;AUTO_INCREMENT;"`
	User_id      string `gorm:"column:user_id"`
	Comment_no   string `gorm:"column:comment_no"`
	Device_type  string `gorm:"column:device_type"`
	Device_ip    string `gorm:"column:device_ip"`
	Lng          string `gorm:"column:lng"`
	Lat          string `gorm:"column:lat"`
	Comment_type int64  `gorm:"column:comment_type"`
	Geo_hash     string `gorm:"column:geohash"`
	Comment_msg  string `gorm:"column:comment_msg"`
	Insert_time  int64  `gorm:"column:insert_time"`
	Update_time  int64  `gorm:"column:update_time"`
	Version      int    `gorm:"column:version"`
}

/*
	PostComment request
*/
type Wait_PostComment_Request struct {
	User_id     string `json:"user_id"`
	Device_type string `json:"device_type"`
	Device_ip   string `json:"device_ip"`
	Lng         string `json:"lng"`
	Lat         string `json:"lat"`
	Comment_msg string `json:"comment_msg"`
}

/*
	PostComment response
*/
type Wait_PostComment_Response struct {
	Status_code string `json:"status_code"`
	Status_msg  string `json:"status_msg"`
	Comment_no  string `json:"comment_no"`
}

//////////////////////////////////////////////////////////////////////////////////
/*
	GetCommentList request
*/
type Wait_GetCommentList_Request struct {
	User_id  string `json:"user_id"`
	Curr_lng string `json:"curr_lng"`
	Curr_lat string `json:"curr_lat"`
}

/*
	GetCommentList response
*/
type Wait_GetCommentList_Response struct {
	Status_code  string              `json:"status_code"`
	Status_msg   string              `json:"status_msg"`
	Comment_list []Wait_Comment_Info `json:"comment_list"`
}

//////////////////////////////////////////////////////////////////////////////////
/*
	GetCommentInfo request
*/
type Wait_GetComment_Request struct {
	Comment_no string `json:"comment_no"`
}

/*
	GetCommentInfo response
*/
type Wait_GetComment_Response struct {
	Status_code  string            `json:"status_code"`
	Status_msg   string            `json:"status_msg"`
	Comment_info Wait_Comment_Info `json:"comment_info"`
}

/*
	CommentInfo Node
*/
type Wait_Comment_Info struct {
	User_id     string `json:"user_id"`
	Device_type string `json:"device_type"`
	Device_ip   string `json:"device_ip"`
	Lng         string `json:"lng"`
	Lat         string `json:"lat"`
	Comment_msg string `json:"comment_msg"`
}

func (Wait_Comment_Info) TableName() string {
	return "wait_users_comments"
}

//////////////////////////////////////////////////////////////////////////////////
