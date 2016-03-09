//Joseph Vietti
//Project 1 Part 2 Uploaded at 3/9
package main

import (
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"net/http"
)

func serveCookie(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("templates/template2.html")
	if err != nil {
		log.Fatalln(err)
	}
	//create cookie "session-fino"
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String(),
			HttpOnly: true,
			//Secure: true,
		}
		http.SetCookie(res, cookie)
	}
	tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", serveCookie)
	http.ListenAndServe(":8080", nil)

}
