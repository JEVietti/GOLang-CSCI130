//Joseph Vietti CSCI 130 HMAC in cookie
package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"io/ioutil"
	"net/http"
)

type User struct {
	Uuid, Name, Age, Hmac string
}

var loginFile string

func init() {
	temp, _ := ioutil.ReadFile("user_templates/template2.html")
	loginFile = string(temp)
}

func getHMAC(data string) string {
	hmac_code := hmac.New(sha256.New, []byte(data+"plebKey"))
	return string(hmac_code.Sum(nil))
}

func bakeUserCookie(cookie *http.Cookie, req *http.Request) string {
	jsonVal, _ := undoJSON(cookie)
	jsonVal.Name = req.FormValue("name")
	jsonVal.Age = req.FormValue("age")
	jsonVal.Hmac = req.FormValue("HMAC")
	return redoJSON(jsonVal)
}

func redoJSON(jsonVal User) string {
	encode, _ := json.Marshal(jsonVal)
	return base64.StdEncoding.EncodeToString(encode)
}

func undoJSON(cookie *http.Cookie) (User, bool) {
	decode, _ := base64.StdEncoding.DecodeString(cookie.Value)
	var jsonVal User
	json.Unmarshal(decode, &jsonVal)
	if hmac.Equal([]byte(jsonVal.Hmac), []byte(getHMAC(jsonVal.Uuid+jsonVal.Name+jsonVal.Age))) {
		return jsonVal, true
	}
	return jsonVal, false
}

func userCookie() *http.Cookie {
	id, _ := uuid.NewV4()
	temp := User{Uuid: id.String(), Hmac: getHMAC(id.String())}
	b, _ := json.Marshal(temp)
	encode := base64.StdEncoding.EncodeToString(b)
	return &http.Cookie{
		Name:     "session-fino",
		Value:    encode,
		HttpOnly: true,
		//Secure : true,
	}
}

func serveLogin(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		cookie = userCookie()
		http.SetCookie(res, cookie)
	}
	if req.Method == "POST" {
		cookie.Value = bakeUserCookie(cookie, req)
	}
	obj, status := undoJSON(cookie)
	if status {
		t, _ := template.New("name").Parse(loginFile)
		t.Execute(res, obj)
	}
}

func main() {
	http.HandleFunc("/", serveLogin)
	http.ListenAndServe(":8080", nil)
}
