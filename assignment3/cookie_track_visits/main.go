//Joe Vietti
//Tracks # of visits
package main
import(
	"io"
	"net/http"
	"strconv"
)

func serve (res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("count")
	if err == http.ErrNoCookie{
		cookie = &http.Cookie{
			Name: "count",
			Value: "0",
		}
	}
	count,_ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(res, cookie)
	io.WriteString(res, cookie.Value)
}
func fav_icon_null(res http.ResponseWriter, req *http.Request){ /*nothing*/}

func main() {
	http.HandleFunc("/",serve)
	http.HandleFunc("/favicon.io",fav_icon_null)
	http.ListenAndServe(":8080",nil)
}
