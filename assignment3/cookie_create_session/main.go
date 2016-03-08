//Joseph Vietti
/*
Create a webpage which writes a cookie to the client's machine. This cookie should be designed to create a session and should use a UUID, HttpOnly, and Secure (though you'll need to comment secure out).
*/

package main

import (
	"github.com/nu7hatch/gouuid"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("session-id")
		if err == http.ErrNoCookie {
			//set the uuid for the cookie so that each user has their own session
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:     "session-id",
				Value:    id.String(),
				HttpOnly: true,
				//Secure:true,
			}
		}
		//set the form for entering their name
		if req.FormValue("name") != "" && !strings.Contains(cookie.Value, "name") {
			cookie.Value = cookie.Value + `name= ` + req.FormValue("name")
		}

		http.SetCookie(res, cookie)
		io.WriteString(res, `<!DOCTYPE html>
		<html>
		  <body>
		   <form method="POST">
		 <label for="Name">What is your name? -> </label>
		 <input type="text" name="name">
		 <input type="submit" value="Enter!">
		</form>
		  </body>
		</html>`)
		//check if the name and session id are being stored properly
		io.WriteString(res,cookie.Value+", that's a good name!");
	})
	http.ListenAndServe(":8080", nil)
}
