//Joseph Vietti Project Part 3 serve template form
//https://github.com/JEVietti/GOLang-CSCI130/blob/master/ProjectSession/templateForm.go
package main

import (
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
)

func serveForm(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("templates/template2.html")
	if err != nil {
		log.Fatalln(err)
	}

	//create a cookie session-fino
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
	}
	//Add the cookie values
	cookie.Value = cookie.Value +
		`Name=` + req.FormValue("name") +
		`Age=` + req.FormValue("age")
	http.SetCookie(res, cookie)

	tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", serveForm)
	http.ListenAndServe(":8080", nil)
}
