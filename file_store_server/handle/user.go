package handle

import (
	mydb "WorkSpace/GoDevelopTest/file_store_server/db"
	"WorkSpace/GoDevelopTest/file_store_server/util"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	pwdSalt     = "*#890"
	tokenSalt   = "*#06#"
	expiredTime = 5
)

func SignUpHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "http://"+r.Host+"/static/view/signup.html", http.StatusFound)
		return
	}
	_ = r.ParseForm()
	userName := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(userName) < 3 || len(passwd) < 5 {
		_, _ = w.Write([]byte("invalid parameter"))
		return
	}
	encPasswd := util.Sha1([]byte(passwd + pwdSalt))
	status := mydb.UserSignUp(userName, encPasswd)
	if status {
		_, _ = w.Write([]byte("Succeed"))
	} else {
		_, _ = w.Write([]byte("failed"))
	}
}

func SignInHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "http://"+r.Host+"/static/view/signin.html", http.StatusFound)
		return
	}
	_ = r.ParseForm()
	userName := r.Form.Get("username")
	passwd := r.Form.Get("password")
	encPasswd := util.Sha1([]byte(passwd + pwdSalt))

	if !mydb.UserSignIn(userName, encPasswd) {
		_, _ = w.Write([]byte("sign in err: user_name or passwd error"))
		return
	}

	token := genToken(userName)
	if !mydb.UpdateToken(userName, token) {
		_, _ = w.Write([]byte("sign in err : gen token err"))
		return
	}
	//http.Redirect(w, r, "http://"+r.Host+"/static/view/home.html", http.StatusFound)
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location, Username, Token string
		}{
			Location: "http://" + r.Host + "/static/view/home.html",
			Username: userName, Token: token,
		},
	}
	_, _ = w.Write(resp.JSONBytes())
}

func genToken(userName string) string {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(userName + ts + tokenSalt))
	return tokenPrefix + ts[:8]
}

func isTokenValid(name, token string) bool {
	if "" == token {
		return false
	}
	if len(token) != 40 {
		return false
	}
	ts, err := strconv.ParseInt(token[32:], 16, 64)
	if err != nil {
		return false
	}
	if time.Now().Unix()-ts > expiredTime {
		return false
	}
	if token == mydb.GetToken(name) {
		return true
	}
	return false
}

func UserInfoHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	userName := r.Form.Get("username")

	user, err := mydb.GetUserInfo(userName)
	if err != nil {
		fmt.Println("user info handle err: ", err)
		return
	}
	resp := util.NewResp(0, "OK", user)
	_, _ = w.Write(resp.JSONBytes())
}
