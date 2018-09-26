package comm

import (
	"time"
)

var ConfigNode CONFIG_NODE

type ERROR_NODE struct {
	Status_code string `json:"status_code"`
	Status_msg  string `json:"status_msg"`
}

const (
	//GORM配置
	GormLogMode       bool   = true
	GormSingularTable bool   = true
	ParnerMonthKey    string = "ParnerMonth"
	ParnerMaxExeIdKey string = "ParnerMaxExeId"
	SecretKey                = "welcome to wangshubo's blog"
	DEFAULT_USER_TYPE        = 1

	//密码错误次数
	PWD_MAX_ERR_COUNT = 5
	//校验码失效时间M
	CHECK_CODE_EXPIRED_MINUTE = 5
	//短信注册验证码类型
	CHECK_CODE_REGISTER = "r"
	//忘记密码验证码类型
	CHECK_CODE_RESETPWD = "f"
	//验证码间隔时间
	CHECK_CODE_MAX_INTERVAL = 60
	//短信注册验证码类型
	CHECK_CODE_ENABLED = "e"
	//用户锁定状态
	USER_LOCKED = 1

	//上传图片最小SIZE
	PIC_MIN_SIZE = 1024
	//上传图片最大SIZE
	PIC_MAX_SIZE = 1024 * 1024
	//图片默认目录
	PIC_DIRECTORY = "d:\\images\\"

	BEGIN_TAG = "======>"
	END_TAG   = "<======"
)

var (
	ParnerSwitch   bool   = false
	ParnerMonth    string = string(time.Now().Format("20060102150405")[0:6])
	ParnerMaxExeId int64  = 0

	RESP_SUCC            = ERROR_NODE{"0000", "执行成功"}
	RESP_NET_ERROR       = ERROR_NODE{"9010", "网络出现错误..."}
	RESP_JSON_ERROR      = ERROR_NODE{"9020", "JSON解析出现错误..."}
	RESP_DB_ERROR        = ERROR_NODE{"9030", "DB错误..."}
	RESP_DB_UPDATE_ERROR = ERROR_NODE{"9040", "DB更新错误..."}
	RESP_USER_EXIST      = ERROR_NODE{"1000", "用户已经存在"}
	RESP_USER_LOCK       = ERROR_NODE{"1600", "用户已经被锁定"}
	RESP_USER_NOEXIST    = ERROR_NODE{"1100", "用户不存在"}
	RESP_CHKCODE_ERROR   = ERROR_NODE{"2200", "校验码错误"}
	RESP_PWD_ERROR       = ERROR_NODE{"1300", "用户名或密码错误"}
	RESP_TYPE_ERROR      = ERROR_NODE{"1500", "类型转换错苏"}
	RESP_OLDPWD_ERROR    = ERROR_NODE{"1700", "旧密码输入错误"}
	RESP_CODE_BUSY       = ERROR_NODE{"2400", "获取验证码太频繁"}
	RESP_CODE_TYPE       = ERROR_NODE{"2500", "验证码类型错误"}
	RESP_UPLOAD_ERROR    = ERROR_NODE{"4001", "上传图片错误"}
	RESP_UPLOAD_DBERROR  = ERROR_NODE{"4040", "上传图片,更新数据库错误"}
)

type CONFIG_NODE struct {
	ListenPort   string
	RedisUrl     string
	RedisPass    string
	DbUrl        string
	ScanInterval int
}
