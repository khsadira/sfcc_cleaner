package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"os"
)

func GetLocalKey(method string, idEnv string, pwdEnv string) string {
	id := os.Getenv(idEnv)
	pwd := os.Getenv(pwdEnv)
	ids := id + ":" + pwd
	return method + " " + base64.StdEncoding.EncodeToString([]byte(ids))
}

func SendPostAuth(path string, id string, pwd string, buffer *bytes.Buffer) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", path, buffer)
	key := GetLocalKey("Basic", id, pwd)
	req.Header.Set("Authorization", key)
	return client.Do(req)
}

func SendGetAuth(path string, id string, pwd string) (*http.Response, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", path, nil)
	key := GetLocalKey("Basic", id, pwd)
	req.Header.Set("Authorization", key)
	return client.Do(req)
}

func BasicAuth(w http.ResponseWriter, r *http.Request, realm string) bool {

	user, pass, ok := r.BasicAuth()

	baseURL := "https://store-dev.ubi.com/on/demandware.store/Sites-Site"
	path := "/default/ViewApplication-ProcessLogin"
	auth := scrapeAuth(user, pass, baseURL, path)
	if !auth || !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
		w.WriteHeader(401)
		w.Write([]byte("Unauthorised.\n"))
		return false
	}
	code := user + ":" + pass
	value := hex.EncodeToString([]byte(code))
	cookie := http.Cookie{Name: "cleaner_auth_token", Value: value, Path: "/", MaxAge: 60 * 15}
	http.SetCookie(w, &cookie)
	tmp := Auth{Username: user, Password: pass, Token: value}
	auths = append(auths, tmp)
	return true
}

func TryConnection(w http.ResponseWriter, r *http.Request) bool {
	key := GetLocalKey("Basic", "CLIENT_ID_SFCC", "CLIENT_PW_SFCC")
	if r.Header.Get("Authorization") == key {
		println("author right")
		return true
	}
	cookie, err := r.Cookie("cleaner_auth_token")
	if err != nil {
		println("no cookie found")
		return BasicAuth(w, r, "Provide Business Manager's username and password")
	} else {
		for _, auth := range auths {
			println("auth")
			if auth.Token == cookie.Value {
				println("we got a true")
				return true
			}
		}
		return BasicAuth(w, r, "Provide username and password")
	}
	return false
}

func AuthStart(muxH *http.ServeMux, path string, fn func(http.ResponseWriter, *http.Request)) {
	muxH.Handle(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if TryConnection(w, r) {
			fn(w, r)
		}
	}))
}
