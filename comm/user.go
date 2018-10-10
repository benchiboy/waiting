package comm

import (
	"time"
)

/*
	===============================================================================
							用户注册及认证数据表结构
	===============================================================================
*/

/*
	用户注册表
*/
type Wait_User struct {
	Auto_id         int64     `gorm:"column:auto_id;primary_key:auto_id;AUTO_INCREMENT;"`
	User_id         int64     `gorm:"column:user_id"`
	User_type       int32     `gorm:"column:user_type"`
	Nick_name       string    `gorm:"column:nick_name"`
	Login_name      string    `gorm:"column:login_name"`
	Login_pwd       string    `gorm:"column:login_pwd"`
	Is_admin        int32     `gorm:"column:is_admin"`
	Islock          int32     `gorm:"column:is_lock"`
	Pwderr_count    int32     `gorm:"column:pwderr_count"`
	Last_pic_code   string    `gorm:"column:last_pic_code"`
	Last_login_time time.Time `gorm:"column:last_login_time"`
	Insert_time     time.Time `gorm:"column:insert_time"`
	Update_time     time.Time `gorm:"column:update_time"`
	Pic_head        string    `gorm:"column:pic_head"`
	Pic_full        string    `gorm:"column:pic_full"`
	User_memo       string    `gorm:"column:user_memo"`
	User_balance    string    `gorm:"column:user_balance"`
	User_points     string    `gorm:"column:user_points"`
	Version         int       `gorm:"column:version"`
}

/*
	验证码发送表
*/
type Wait_Checkcode struct {
	Auto_id      int64     `gorm:"column:auto_id;primary_key:auto_id;AUTO_INCREMENT;"`
	User_id      int64     `gorm:"column:user_id"`
	User_phone   string    `gorm:"column:user_phone"`
	Check_code   string    `gorm:"column:check_code"`
	Check_type   string    `gorm:"column:check_type"`
	Send_code    string    `gorm:"column:send_code"`
	Send_msg     string    `gorm:"column:send_msg"`
	Code_status  string    `gorm:"column:code_status"`
	Verify_times int32     `gorm:"column:verify_times"`
	Valid_btime  time.Time `gorm:"column:valid_btime"`
	Valid_etime  time.Time `gorm:"column:valid_etime"`
	Insert_time  time.Time `gorm:"column:insert_time"`
	Update_time  time.Time `gorm:"column:update_time"`
}

/*
	用户支付订单表
*/
type Wait_Users_Charge_orders struct {
	Auto_id      int64     `gorm:"column:auto_id;primary_key:auto_id;AUTO_INCREMENT;"`
	User_id      string    `gorm:"column:user_id"`
	Mct_type     string    `gorm:"column:mct_type"`
	Trxn_type    string    `gorm:"column:trxn_type"`
	Old_order_no string    `gorm:"column:old_order_no"`
	Order_no     string    `gorm:"column:order_no"`
	Order_type   string    `gorm:"column:order_type"`
	Order_amt    float64   `gorm:"column:order_amt"`
	Order_date   time.Time `gorm:"column:order_date"`
	Status_code  string    `gorm:"column:status_code"`
	Status_msg   string    `gorm:"column:status_msg"`
	Memo         string    `gorm:"column:memo"`
	User_bal_amt float64   `gorm:"column:user_bal_amt"`
	Insert_time  time.Time `gorm:"column:insert_time"`
	Update_time  time.Time `gorm:"column:update_time"`
	Version      int       `gorm:"column:version"`
}

/*
	用户支付明细表
*/
type Wait_Users_Charge_orders_detail struct {
	Auto_id     int64     `gorm:"column:auto_id;primary_key:auto_id;AUTO_INCREMENT;"`
	User_id     string    `gorm:"column:user_id"`
	Order_no    string    `gorm:"column:order_no"`
	Goods_no    string    `gorm:"column:goods_no"`
	Goods_amt   float64   `gorm:"column:goods_amt"`
	Goods_type  string    `gorm:"column:goods_type"`
	Insert_time time.Time `gorm:"column:insert_time"`
	Update_time time.Time `gorm:"column:update_time"`
	Version     int       `gorm:"column:version"`
}

/*
	用户图片列表
*/
type Wait_Users_pics struct {
	Auto_id     int64     `gorm:"column:auto_id;primary_key:auto_id;AUTO_INCREMENT;"`
	User_id     string    `gorm:"column:user_id"`
	Pic_id      int64     `gorm:"column:pic_id"`
	Pic_type    int       `gorm:"column:pic_type"`
	Pic_name    string    `gorm:"column:pic_name"`
	Insert_time time.Time `gorm:"column:insert_time"`
	Update_time time.Time `gorm:"column:update_time"`
	Version     int       `gorm:"column:version"`
}

/*
	用户图片评论表
*/
type Wait_Users_Pics_Comments struct {
	Auto_id         int64     `gorm:"column:auto_id;primary_key:auto_id;AUTO_INCREMENT;"`
	User_id         string    `gorm:"column:user_id"`
	Comment_user_id string    `gorm:"column:comment_user_id"`
	Pic_id          int       `gorm:"column:pic_id"`
	Comment_type    string    `gorm:"column:comment_type"`
	Comment_msg     string    `gorm:"column:comment_msg"`
	Insert_time     time.Time `gorm:"column:insert_time"`
	Update_time     time.Time `gorm:"column:update_time"`
	Version         int       `gorm:"column:version"`
}

/*
	===============================================================================
							用户注册及认证报文定义
	===============================================================================
*/

/*
	Waiting login
*/
type Wait_Login_Request struct {
	Login_name string `json:"login_name"`
	Login_pwd  string `json:"login_pwd"`
	Pic_code   string `json:"pic_code"`
	Timestamp  string `json:"Timestamp"`
}

/*
	Waiting api union response
*/
type Wait_Login_Response struct {
	Status_code   string `json:"status_code"`
	Status_msg    string `json:"status_msg"`
	User_nick     string `json:"user_nick,omitempty"`
	User_ImageUrl string `json:"user_imageurl,omitempty"`
	User_Phone    string `json:"user_phone,omitempty"`
	User_Balance  string `json:"user_balance,omitempty"`
	User_Points   string `json:"user_points,omitempty"`
	User_avatar   string `json:"user_avatar,omitempty"`
}

/*
	Get User request
*/
type Wait_GetUser_Request struct {
	Login_name string `json:"login_name"`
}

/*
	Get User response
*/
type Wait_GetUser_Response struct {
	Status_code   string `json:"status_code"`
	Status_msg    string `json:"status_msg"`
	User_nick     string `json:"user_nick,omitempty"`
	User_ImageUrl string `json:"user_imageurl,omitempty"`
	User_Phone    string `json:"user_phone,omitempty"`
	User_Balance  string `json:"user_balance,omitempty"`
	User_Points   string `json:"user_points,omitempty"`
}

/*
	Get UserList Request
*/
type Wait_GetUserList_Request struct {
	RegDate string `json:"reg_date"`
}

/*
	Get UserList Response
*/
type Wait_GetUserList_Response struct {
	Status_code string            `json:"status_code"`
	Status_msg  string            `json:"status_msg"`
	User_list   []Wait_GetUserPic `json:"user_list"`
}

/*
	Get UserList Sub Entity
*/
type Wait_GetUserPic struct {
	Login_name string `json:"login_name"`
	User_name  string `json:"user_name"`
	Pic_full   string `json:"pic_full"`
}

/*
	Register user request
*/
type Wait_RegUser_Request struct {
	Login_name string `json:"login_name"`
	Login_pwd  string `json:"confirmPassWord"`
	Nick_Name  string `json:"nick_name"`
}

/*
	Register user response
*/
type Wait_RegUser_Response struct {
	Status_code string `json:"status_code"`
	Status_msg  string `json:"status_msg"`
}

/*
	Forget Password request
*/
type Wait_ForgetPwd_Request struct {
	Login_name string `json:"login_name"`
	User_phone string `json:"user_phone"`
	Check_code string `json:"check_code"`
	New_pwd    string `json:"new_pwd"`
}

/*
	Forget Password response
*/
type Wait_ForgetPwd_Response struct {
	Status_code string `json:"status_code"`
	Status_msg  string `json:"status_msg"`
}

/*
	Change Password request
*/
type Wait_ChangePwd_Request struct {
	Login_name string `json:"login_name"`
	Old_pwd    string `json:"oldPassWord"`
	New_pwd    string `json:"confirmPassWord"`
}

/*
	Change Password response
*/
type Wait_ChangePwd_Response struct {
	Status_code string `json:"status_code"`
	Status_msg  string `json:"status_msg"`
}

/*
	Send CheckCode request
*/
type Wait_CheckCode_Request struct {
	User_Phone string `json:"user_phone"`
	Code_type  string `json:"type"`
}

/*
	Send CheckCode response
*/
type Wait_CheckCode_Response struct {
	Status_code string `json:"status_code"`
	Status_msg  string `json:"status_msg"`
	Check_code  string `json:"check_code"`
}

/*
	User Upload pics request
*/
type Wait_UploadPics_Request struct {
	User_id  string `json:"user_id"`
	Pic_data string `json:"pic_data"`
}

/*
	User Upload pics response
*/
type Wait_UploadPics_Response struct {
	Status_code string `json:"status_code"`
	Status_msg  string `json:"status_msg"`
	User_picurl string `json:"user_picurl"`
}

/*
	获取图形验证码
*/
type Wait_Capture_Request struct {
	Login_name string `json:"login_name"`
}

type Wait_Capture_Response struct {
	Status_code string `json:"status_code"`
	Status_msg  string `json:"status_msg"`
	PicBase64   string `json:"pic_base64"`
}
