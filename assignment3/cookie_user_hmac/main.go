//Joseph Vietti
/*
Create a webpage which writes a cookie to the client's machine. Though this is NOT A BEST PRACTICE, you will store some session data in the cookie. Make sure you use HMAC to ensure that session data is not changed by a user.
*/

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("session-id")
		if err == http.ErrNoCookie {
			cookie = &http.Cookie{
				Name:     "session-id",
				Value:    "",
				HttpOnly: true,
				//Secure:true,
			}
		}
		//set the form for entering their name
		if req.FormValue("name") != "" {
			cookie.Value = cookie.Value + `name= ` + getSHA("name")
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
		<a href="/auth">Valid HMAC</a>
		  </body>
		</html>`)
		//check if the name and session id are being stored properly
		io.WriteString(res, cookie.Value+", that's a good name!")
	})

	http.HandleFunc("/auth", checkAuth)
	http.ListenAndServe(":8080", nil)
}

func getSHA(data string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func checkAuth(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("cookie")
	if err != nil {
		http.Redirect(res, req, "/", 303)
		return
	}
	if cookie.Value == "" {
		http.Redirect(res, req, "/", 303)
		return
	}

	xs := strings.Split(cookie.Value, "|")
	name := xs[0]
	shaRCVD := xs[1]
	shaCheck := getSHA(name)

	if shaRCVD != shaCheck {
		log.Println("HMAC codes don't match")
		log.Println(shaRCVD)
		log.Println(shaCheck)
		http.Redirect(res, req, "/", 303)
		return
	}
}
