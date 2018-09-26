// waiting project main.go
package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"waiting/business"
	"waiting/comm"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"

	//"github.com/codegangsta/negroni"
	goconf "github.com/pantsing/goconf"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	http_srv *http.Server
)

//var sessionManager = scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")

func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(comm.SecretKey), nil
		})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}

}

func init() {
	//日志配置

	log.Println("init session...")
	//	sessionManager.Persist(true)
	//sessionManager.Lifetime(time.Second * 10)

	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
	log.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "crfcredit_parser.log",
		MaxSize:    500, // megabytes
		MaxBackups: 50,
		MaxAge:     90, //days
	}))
	//读取配置文件
	envConf := flag.String("env", "config-ci.json", "select a environment config file")
	flag.Parse()
	log.Println("config file ==", *envConf)
	c, err := goconf.New(*envConf)
	if err != nil {
		log.Fatalln("读配置文件出错", err)
	}
	//填充配置文件
	c.Get("/config/LISTEN_PORT", &comm.ConfigNode.ListenPort)
	c.Get("/config/DB_URL", &comm.ConfigNode.DbUrl)

}

func go_WebServer() {
	//开启http服务
	log.Println("Service start listen ", comm.ConfigNode.ListenPort)
	http.Handle("/", http.FileServer(http.Dir("./html/")))
	//busi process
	//http.HandleFunc("/WeiboCancelCallback", business.WeiboCancelCallback)
	//http.HandleFunc("/WeiboCallback", business.WeiboCallback)
	//	http.Handle("/user/v1/getuser", negroni.New(
	//		negroni.HandlerFunc(ValidateTokenMiddleware),
	//		negroni.Wrap(http.HandlerFunc(business.GetUser)),
	//	))

	//用户管理模块

	//mux := http.NewServeMux()
	//mux.HandleFunc("/put", business.Session_test)
	//mux.HandleFunc("/get", getHandler)

	http.HandleFunc("/user/v1/reguser", business.RegUser)
	http.HandleFunc("/user/v1/login", business.UserLogin)
	http.HandleFunc("/user/v1/changepwd", business.ChangePwd)
	http.HandleFunc("/user/v1/forgetpwd", business.ForgetPwd)
	http.HandleFunc("/user/v1/checkcode", business.CheckCode)
	http.HandleFunc("/user/v1/captchas", business.GetCaptchas)

	http.HandleFunc("/user/v1/getuserlist", business.GetUserList)
	http.HandleFunc("/user/v1/uploadpics", business.UploadPics)

	http.HandleFunc("/user/v1/uploadpics_input", business.UploadPics_input)
	http.HandleFunc("/user/v1/uploadpics_form", business.UploadPics_form)

	http.HandleFunc("/comment/v1/postcomment", business.Create_Comment)
	http.HandleFunc("/comment/v1/getcomments", business.Get_CommentList)
	http.HandleFunc("/comment/v1/getcommentbyno", business.Get_CommentByNo)
	http.HandleFunc("/api/user/registered", business.Get_Test)

	//http.HandleFunc("/comment/v1/session_test", business.Session_test)
	//http.HandleFunc("/comment/v1/session_test_set", business.Session_test_set)

	http_srv = &http.Server{
		Addr:    comm.ConfigNode.ListenPort,
		Handler: http.DefaultServeMux,
		//sessionManager.Use(mux),
	}
	if err := http_srv.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}
}

func main() {
	log.Println("Waiting a man or women...!")
	go go_WebServer()
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	log.Println("Recv a kill signal and exit ...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := http_srv.Shutdown(ctx)
	log.Println("Server gracefully stopped:", err)

}
