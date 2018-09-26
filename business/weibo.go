package business

//import (
//	"log"
//	"net/http"
//	"net/url"
//	"time"
//	"waiting/comm"
//	"waiting/db"

//	"github.com/codegangsta/negroni"
//	"github.com/dgrijalva/jwt-go"
//	"github.com/dgrijalva/jwt-go/request"*/
//)

//var endpotin = oauth2.Endpoint{
//	AuthURL:  "https://api.weibo.com/OAuth2/authorize",
//	TokenURL: "https://api.weibo.com/OAuth2/access_token",
//}

//var googleOauthConfig = &oauth2.Config{
//	ClientID:     "2302557195",
//	ClientSecret: "3e096effd83cd4d0553b5098159eced8",
//	RedirectURL:  "http://www.zhiaidian.com/WeiboCallback",
//	Scopes:       []string{"https://api.weibo.com/OAuth2/access_token"},
//	Endpoint:     endpotin,
//}

//const oauthStateString = "random"

//func WaitingLogin(w http.ResponseWriter, r *http.Request) {

//	log.Println("......WaitingLogin......")
//	t1 := time.Now()

//	var user UserCredentials

//	err := json.NewDecoder(r.Body).Decode(&user)

//	if err != nil {
//		w.WriteHeader(http.StatusForbidden)
//		fmt.Fprint(w, "Error in request", err.Error())
//		return
//	}
//	fmt.Println(user)

//	if strings.ToLower(user.Username) != "someone" {
//		if user.Password != "p@ssword" {
//			w.WriteHeader(http.StatusForbidden)
//			fmt.Println("Error logging in")
//			fmt.Fprint(w, "Invalid credentials")
//			return
//		}
//	}

//	token := jwt.New(jwt.SigningMethodHS256)
//	claims := make(jwt.MapClaims)
//	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
//	claims["iat"] = time.Now().Unix()
//	token.Claims = claims

//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		fmt.Fprintln(w, "Error extracting the key")
//		fatal(err)
//	}

//	tokenString, err := token.SignedString([]byte(SecretKey))
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		fmt.Fprintln(w, "Error while signing the token")
//		fatal(err)
//	}

//	response := Token{tokenString}
//	JsonResponse(response, w)

//	regUser := comm.Wait_User{User_IdNo: "22222222"}
//	db.Create_User(regUser)
//	log.Println("Create user elapsed time:", time.Since(t1))
//	return
//}

//func WeiboLogin(w http.ResponseWriter, r *http.Request) {
//	log.Println("......proc user reg......")
//	t1 := time.Now()
//	url := googleOauthConfig.AuthCodeURL(oauthStateString)
//	fmt.Println(url)
//	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
//	log.Println("Create user elapsed time:", time.Since(t1))
//	return
//}

//func WeiboCallback(w http.ResponseWriter, r *http.Request) {
//	log.Println("......proc user reg......")
//	t1 := time.Now()
//	state := r.FormValue("state")
//	if state != oauthStateString {
//		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
//		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
//		return
//	}
//	code := r.FormValue("code")
//	fmt.Println(code)
//	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
//	fmt.Println(token)
//	if err != nil {
//		fmt.Println("Code exchange failed with '%s'\n", err)
//		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
//		return
//	}
//	response, err := http.PostForm("https://api.weibo.com/Oauth2/get_token_info",url.Values{"access_token": {token.AccessToken}}
//	defer response.Body.Close()
//	contents, err := ioutil.ReadAll(response.Body)
//	fmt.Fprintf(w, "Content: %s\n", contents)
//	return
//}

//func WeiboCancelCallback(w http.ResponseWriter, r *http.Request) {
//	log.Println("......proc user reg......")
//	t1 := time.Now()
//	log.Println(t1)
//	return
//}
