package business

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
	"waiting/comm"
	"waiting/db"

	"github.com/mojocn/base64Captcha"
)

type Token struct {
	Token string `json:"token"`
}

func Write_Response(response interface{}, w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Max-Age", "1728000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "content-type,Action, Module,Authorization") //有使用自定义头 需要这个,Action, Module是例子
	fmt.Fprintf(w, string(json))
}

/*
	1、验证用户是否存在
	2、验证密码是否正确
	3、如果密码错误一定次数，锁定用户
	注明：如果密码锁定后，密码解锁后，需要把密码错误次数设置为0
*/
func UserLogin(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "Begin Login......")
	t1 := time.Now()

	log.Println("======>", r.Method)
	var login comm.Wait_Login_Request
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		Write_Response(comm.RESP_JSON_ERROR, w, r)
		log.Println(comm.END_TAG, "Json Error......", time.Since(t1))
		return
	}
	log.Println(login)
	defer r.Body.Close()
	v, err := db.Get_User(login.Login_name)
	if err != nil {
		Write_Response(comm.RESP_DB_ERROR, w, r)
		log.Println(comm.END_TAG, "DB Error......", time.Since(t1))
		return
	}
	if v == nil {
		Write_Response(comm.RESP_USER_NOEXIST, w, r)
		log.Println(comm.END_TAG, "User NoExist......", time.Since(t1))
		return
	}

	user, ok := v.(*comm.Wait_User)
	if login.Pic_code != "" {
		verifyResult := base64Captcha.VerifyCaptcha(user.Last_pic_code, login.Pic_code)
		if !verifyResult {
			Write_Response(comm.RESP_PWD_ERROR, w, r)
			log.Println("图形校验失败.....")
			return
		}
	}

	if ok {
		if user.Islock == comm.USER_LOCKED {
			Write_Response(comm.RESP_USER_LOCK, w, r)
			log.Println(comm.END_TAG, "User Locked......", time.Since(t1))
			return
		}
		if user.Login_pwd != login.Login_pwd {

			err = db.Update_User_PwdErr(*user)
			if err != nil {
				Write_Response(comm.RESP_DB_UPDATE_ERROR, w, r)
			} else {
				Write_Response(comm.RESP_PWD_ERROR, w, r)
			}
			log.Println(comm.END_TAG, "Password Error......", time.Since(t1))
			return
		}
	}

	user_resp := comm.Wait_Login_Response{Status_code: comm.RESP_SUCC.Status_code,
		Status_msg:   comm.RESP_SUCC.Status_msg,
		User_Phone:   user.Login_name,
		User_Balance: user.User_balance,
		User_Points:  user.User_points,
		User_avatar:  user.Pic_head,
		User_nick:    user.Nick_name,
	}

	Write_Response(user_resp, w, r)
	log.Println(user_resp)
	log.Println(comm.END_TAG, "End Successful......", time.Since(t1))
	return
}

/*
	1、检查要注册用户是否存在
	2、向用户发送验证码
	3、校验验证码无误后，创建新用户
*/
func RegUser(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "RegUser......")
	t1 := time.Now()
	var reguser_req comm.Wait_RegUser_Request
	err := json.NewDecoder(r.Body).Decode(&reguser_req)
	if err != nil {
		log.Println(err.Error())
		Write_Response(comm.RESP_JSON_ERROR, w, r)
		return
	}
	defer r.Body.Close()
	v, err := db.Get_User(reguser_req.Login_name)
	//出错
	if err != nil {
		Write_Response(comm.RESP_DB_ERROR, w, r)
		log.Println(comm.END_TAG, "RegUser......", time.Since(t1))
		return
	}
	//此用户已经存在
	if v != nil {
		Write_Response(comm.RESP_USER_EXIST, w, r)
		log.Println(comm.END_TAG, "RegUser......", time.Since(t1))
		return
	}
	create_user := comm.Wait_User{
		User_id:         time.Now().UnixNano(),
		Last_login_time: time.Now(),
		User_type:       comm.DEFAULT_USER_TYPE,
		Nick_name:       reguser_req.Nick_Name,
		Login_name:      reguser_req.Login_name,
		Login_pwd:       reguser_req.Login_pwd,
		Insert_time:     time.Now(),
		Update_time:     time.Now()}
	err = db.Create_User(create_user)
	if err != nil {
		Write_Response(comm.RESP_DB_ERROR, w, r)
		log.Println(comm.END_TAG, "RegUser......", time.Since(t1))
		return
	}
	Write_Response(comm.RESP_SUCC, w, r)
	log.Println(comm.END_TAG, "RegUser......", time.Since(t1))
	return
}

/*
	1、上传用户头像图片
	2、如果上传成功，更新DB中存储文件名称
*/
func UploadPics(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "UploadPics......")
	t1 := time.Now()
	var uppics_req comm.Wait_UploadPics_Request
	err := json.NewDecoder(r.Body).Decode(&uppics_req)
	if err != nil {
		Write_Response(comm.RESP_JSON_ERROR, w, r)
		log.Println(comm.END_TAG, "UploadPics json Error:......", err, time.Since(t1))
		return
	}
	defer r.Body.Close()
	if len(uppics_req.Pic_data) < comm.PIC_MIN_SIZE {
		log.Println(comm.END_TAG, "PicSize too small......", time.Since(t1))
		return
	}
	log.Println("user_id", uppics_req.User_id, "pic_data:", uppics_req.Pic_data[23:], "len:", len(uppics_req.Pic_data))
	picBuf, err := base64.StdEncoding.DecodeString(uppics_req.Pic_data[23:])
	if err != nil {
		log.Println(comm.END_TAG, "DecodeString Error......", time.Since(t1))
		return
	}
	picName := fmt.Sprintf("%s%s-%d.jpg", comm.PIC_DIRECTORY, uppics_req.User_id, time.Now().UnixNano())
	log.Println("pic name:", picName)
	err = ioutil.WriteFile(picName, picBuf, 0777)
	if err != nil {
		Write_Response(comm.RESP_UPLOAD_ERROR, w, r)
		log.Println(comm.END_TAG, "Upload pics error......", time.Since(t1))
	}

	currUser := comm.Wait_User{
		Login_name: uppics_req.User_id,
		Pic_head:   picName,
	}
	err = db.Update_User_HeadPic(currUser)
	if err != nil {
		Write_Response(comm.RESP_UPLOAD_DBERROR, w, r)
		log.Println(comm.END_TAG, "Upload pics error......", time.Since(t1))
	}
	Write_Response(comm.RESP_SUCC, w, r)
	log.Println(comm.END_TAG, "UploadPics Successful......", time.Since(t1))
	return
}

/*
	1、上传用户头像图片
	2、如果上传成功，更新DB中存储文件名称
*/
func UploadPics_input(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "UploadPics......")
	t1 := time.Now()
	r.FormFile("fileName")
	file, handle, err := r.FormFile("fileName")
	if err != nil {
		log.Println(err.Error())
		err.Error()
		return
	}
	r.FormFile("userName")

	userName := r.FormValue("userName")
	fileName := userName + "-" + handle.Filename

	f, err := os.OpenFile("./images/"+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	io.Copy(f, file)
	if err != nil {
		err.Error()
	}
	defer f.Close()
	defer file.Close()
	fmt.Println("upload success")

	currUser := comm.Wait_User{
		Login_name: userName,
		Pic_head:   "http://localhost:8089/" + fileName,
	}

	err = db.Update_User_HeadPic(currUser)

	pic_resp := comm.Wait_UploadPics_Response{
		Status_code: comm.RESP_SUCC.Status_code,
		Status_msg:  comm.RESP_SUCC.Status_msg,
		User_picurl: "http://localhost:8089/" + fileName,
	}
	Write_Response(pic_resp, w, r)
	log.Println(comm.END_TAG, "UploadPics Successful......", time.Since(t1))
	return
}

/*
	1、根据旧密码修改新密码
*/

func ChangePwd(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "ChangePwd......")
	t1 := time.Now()
	var changepwd_req comm.Wait_ChangePwd_Request
	err := json.NewDecoder(r.Body).Decode(&changepwd_req)
	if err != nil {
		Write_Response(comm.RESP_JSON_ERROR, w, r)
		log.Println(comm.END_TAG, "ResetPwd......", time.Since(t1))
		return
	}
	defer r.Body.Close()
	v, err := db.Get_User(changepwd_req.Login_name)
	if err != nil {
		Write_Response(comm.RESP_DB_ERROR, w, r)
		log.Println(comm.END_TAG, "DB Error......", time.Since(t1))
		return
	}
	if v == nil {
		Write_Response(comm.RESP_USER_NOEXIST, w, r)
		log.Println(comm.END_TAG, "User NoExist......", time.Since(t1))
		return
	}
	user, ok := v.(*comm.Wait_User)
	if ok {
		if user.Login_pwd != changepwd_req.Old_pwd {
			Write_Response(comm.RESP_OLDPWD_ERROR, w, r)
			log.Println(comm.END_TAG, "Old Password Error......", time.Since(t1))
			return
		}
	}
	err = db.Change_User_Pwd(changepwd_req.Login_name, changepwd_req.New_pwd)
	if err != nil {
		log.Println(err.Error())
		Write_Response(comm.RESP_DB_ERROR, w, r)
		return
	}
	Write_Response(comm.RESP_SUCC, w, r)
	log.Println(comm.END_TAG, "ChangePwd......", time.Since(t1))
}

/*
	1、根据手机号码和短信验证码校验，通过修改密码
	2、修改密码成功后，把短信验证码的状态设置为失效
*/

func ForgetPwd(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "ForgetPwd......")
	t1 := time.Now()
	var forgetpwd_req comm.Wait_ForgetPwd_Request
	err := json.NewDecoder(r.Body).Decode(&forgetpwd_req)
	if err != nil {
		Write_Response(comm.RESP_JSON_ERROR, w, r)
		log.Println(comm.END_TAG, "ForgetPwd......", time.Since(t1))
		return
	}
	defer r.Body.Close()
	//校验短信验证码是否正确
	v, err := db.Get_CheckCode(forgetpwd_req.User_phone, forgetpwd_req.Check_code, comm.CHECK_CODE_RESETPWD)
	if err != nil || v == nil {
		log.Println("ForgetPwd...")
		Write_Response(comm.RESP_CHKCODE_ERROR, w, r)
		return
	}
	err = db.Change_User_Pwd(forgetpwd_req.Login_name, forgetpwd_req.New_pwd)
	if err != nil {
		log.Println("=========>", err.Error())
		Write_Response(comm.RESP_DB_ERROR, w, r)
		return
	}
	Write_Response(comm.RESP_SUCC, w, r)
	log.Println(comm.END_TAG, "ForgetPwd......", time.Since(t1))
}

/*
	1、检查上次验证码发送时间
	2、按请求的类型，产生验证码，发送给指定手机
*/

func CheckCode(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "Checkcode......")
	t1 := time.Now()
	var checkcode_req comm.Wait_CheckCode_Request
	err := json.NewDecoder(r.Body).Decode(&checkcode_req)
	if err != nil {
		Write_Response(comm.RESP_JSON_ERROR, w, r)
		log.Println(comm.END_TAG, "ResetPwd......", time.Since(t1))
		return
	}
	log.Println(checkcode_req)
	defer r.Body.Close()
	//检查校验码类型
	if checkcode_req.Code_type != comm.CHECK_CODE_REGISTER && checkcode_req.Code_type != comm.CHECK_CODE_RESETPWD {
		Write_Response(comm.RESP_CODE_TYPE, w, r)
		return
	}
	//查看上次发送时间
	t, err := db.Get_LastCheckCodeTime(checkcode_req.User_Phone)
	if err != nil {
		Write_Response(comm.RESP_DB_ERROR, w, r)
		log.Println(err.Error())
		return
	}
	if time.Now().Sub(t).Seconds() < comm.CHECK_CODE_MAX_INTERVAL {
		Write_Response(comm.RESP_CODE_BUSY, w, r)
		return
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	check_code := comm.Wait_Checkcode{User_phone: checkcode_req.User_Phone, Check_code: vcode, Check_type: checkcode_req.Code_type}
	//增加调用第三方接口
	check_code.Send_code = "0"
	check_code.Insert_time = time.Now()
	check_code.Code_status = comm.CHECK_CODE_ENABLED
	check_code.Valid_btime = time.Now().Add(time.Duration(-time.Minute * comm.CHECK_CODE_EXPIRED_MINUTE))
	check_code.Valid_etime = time.Now().Add(time.Duration(time.Minute * comm.CHECK_CODE_EXPIRED_MINUTE))
	check_code.Update_time = time.Now()
	err = db.Create_CheckCode(check_code)
	if err != nil {
		Write_Response(comm.RESP_DB_ERROR, w, r)
		log.Println(err.Error())
		return
	}
	//resp := comm.Wait_CheckCode_Response{Status_code: comm.RESP_SUCC.Status_code, Status_msg: comm.RESP_SUCC.Status_msg, Check_code: vcode}
	Write_Response(comm.RESP_SUCC, w, r)
	log.Println(comm.END_TAG, "Send_Checkcode......", time.Since(t1))
}

/*
	获取图形验证码
*/
func GetCaptchas(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "GetCaptchas......")

	t1 := time.Now()
	var cap_req comm.Wait_Capture_Request
	err := json.NewDecoder(r.Body).Decode(&cap_req)
	if err != nil {
		Write_Response(comm.RESP_JSON_ERROR, w, r)
		return
	}
	defer r.Body.Close()

	//数字验证码配置
	var configD = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      200,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}

	idKeyD, capD := base64Captcha.GenerateCaptcha("", configD)
	//以base64编码
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	fmt.Println(idKeyD, base64stringD)

	user := comm.Wait_User{Login_name: cap_req.Login_name, Last_pic_code: idKeyD}
	err = db.Update_User_PicCode(user)
	if err != nil {
		log.Println("=========>", err.Error())
		Write_Response(comm.RESP_DB_ERROR, w, r)
		return
	}
	cap_resp := comm.Wait_Capture_Response{comm.RESP_SUCC.Status_code, comm.RESP_SUCC.Status_msg, base64stringD}
	Write_Response(cap_resp, w, r)
	log.Println(comm.END_TAG, "GetCaptchas......", time.Since(t1))
}

/*
	1、 得到注册用户列表
*/
func GetUserList(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "GetUserList......")
	t1 := time.Now()
	var userlist_req comm.Wait_GetUserList_Request
	err := json.NewDecoder(r.Body).Decode(&userlist_req)
	if err != nil {
		log.Println(err.Error())
		Write_Response(comm.RESP_JSON_ERROR, w, r)
		return
	}
	defer r.Body.Close()
	userList, err := db.Get_UserList()
	if err != nil {
		Write_Response(comm.RESP_DB_ERROR, w, r)
		log.Println(comm.END_TAG, "GetUserList......", time.Since(t1))
		return
	}
	userPics := make([]comm.Wait_GetUserPic, len(userList), len(userList))
	for i, v := range userList {
		userPics[i].Login_name = v.Login_name
		userPics[i].Pic_full = v.Pic_full
	}
	userlist_resp := comm.Wait_GetUserList_Response{comm.RESP_SUCC.Status_code, comm.RESP_SUCC.Status_msg, userPics}
	Write_Response(userlist_resp, w, r)
	log.Println(comm.END_TAG, "GetUserList......", time.Since(t1))
}

/*
	1、得到用户的基础信息
		1、包括用户手机、昵称
		2、账户余额
		3、积分情况
*/
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	log.Println(comm.BEGIN_TAG, "GetUserInfo......")
	t1 := time.Now()

	var user_req comm.Wait_GetUser_Request
	err := json.NewDecoder(r.Body).Decode(&user_req)
	if err != nil {
		log.Println(err.Error())
		Write_Response(comm.RESP_JSON_ERROR, w, r)
		return
	}
	defer r.Body.Close()

	v, err := db.Get_User(user_req.Login_name)
	if err != nil {
		Write_Response(comm.RESP_DB_ERROR, w, r)
		log.Println(comm.END_TAG, "GetUserInfo......", time.Since(t1))
		return
	}

	userInfo, ok := v.(*comm.Wait_User)
	if !ok {
		Write_Response(comm.RESP_USER_NOEXIST, w, r)
		log.Println(comm.END_TAG, "GetUserInfo......", time.Since(t1))
		return
	}

	user_resp := comm.Wait_GetUser_Response{Status_code: comm.RESP_SUCC.Status_code,
		Status_msg:    comm.RESP_SUCC.Status_msg,
		User_Phone:    userInfo.Login_name,
		User_Balance:  userInfo.User_balance,
		User_Points:   userInfo.User_points,
		User_ImageUrl: userInfo.Pic_head,
	}

	Write_Response(user_resp, w, r)
	log.Println(comm.END_TAG, "GetUserInfo......", time.Since(t1))
}
