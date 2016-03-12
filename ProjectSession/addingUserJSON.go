//Joseph Vietti Project Part 4
//https://github.com/JEVietti/GOLang-CSCI130/blob/master/ProjectSession/addingUserJSON.go
package main

import (
	"encoding/json"
	"encoding/base64"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"net/http"
)
//define the structure for user
type User struct {
	Name string
	Age  string
}


func serveUser(res http.ResponseWriter, req *http.Request) {
	bakeCookie(res, req)
	temp, _ := template.ParseFiles("templates/template2.html")
	temp.Execute(res, nil)
}
//create the cookie
func createCookie(res *http.ResponseWriter, cookieName, cookieValue string) {
	cookie := &http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
		//Secure: true,
		HttpOnly: true,
	}

	http.SetCookie(*res, cookie)
}
//"bake" it = get form values and encode
func bakeCookie(res http.ResponseWriter, req *http.Request) {
	id, _ := uuid.NewV4()
	createCookie(&res, "session-fino", id.String())
	//set user values
	user := User{
		Name: req.FormValue("name"),
		Age:  req.FormValue("age"),
	}
	if req.Method == "POST" {
		userBytes, _ := json.Marshal(user)
		createCookie(&res, "userData", base64.StdEncoding.EncodeToString(userBytes))
	}
}


func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", serveUser)
	http.ListenAndServe(":8080", nil)
}