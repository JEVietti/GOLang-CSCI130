//Joe Vietti
/*Create a webpage that serves a form and allows the user to upload a txt file. You do not need to check if the file is a txt; bad programming but just trust the user to follow the instructions. Once a user has uploaded a txt file, copy the text from the file and display it on the webpage. Use req.FormFile and io.Copy to do this
*/
package main

import(
	"io"
	"net/http"
	"os"
	"path/filepath"

	"io/ioutil"
)

func serve(res http.ResponseWriter, req *http.Request){
	if req.Method == "POST" {
		file, _, err := req.FormFile("input")
		if err !=nil{
			http.Error(res, err.Error(),500)
			return
		}
		defer file.Close()
		src := io.LimitReader(file,404)
		dst, err := os.Create(filepath.Join(".","input.txt"))
		if err != nil{
			http.Error(res,err.Error(),500)
			return
		}
		defer dst.Close()
		io.Copy(dst, src)
		dat, err := ioutil.ReadFile("input.txt")
		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res, string(dat))
	}
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res, `
      <form method="POST" enctype="multipart/form-data">
        <input type="file" name="input">
        <input type="submit">
      </form>`)
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(serve))
}
