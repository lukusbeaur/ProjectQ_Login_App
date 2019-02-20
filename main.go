package main

import (
	"fmt"
	"log"
	"net/http"
	//helper "./helpers" fix this later fucking local package garbage
	//adding a func in main.go calle isEmpty
)

func main() {
	userName, email, pswd, pswdConf := "", "", "", ""

	mux := http.NewServeMux()

	//Sign up new user
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		userName = r.FormValue("username")
		email = r.FormValue("email")
		pswd = r.FormValue("password")
		pswdConf = r.FormValue("confirm")

		//call helper, checking if form is null
		//make generic eventually and call each func as a list ? instead of hard coding
		//temp change - IsEmpty is now in main.go until local packages work

		/*uNameCheck := helper.IsEmpty(userName)
		emailCheck := helper.IsEmpty(email)
		pswdCheck := helper.IsEmpty(pswd)
		pswdConfCheck := helper.IsEmpty(pswdConf)*/
		uNameCheck := IsEmpty(userName)
		emailCheck := IsEmpty(email)
		pswdCheck := IsEmpty(pswd)
		pswdConfCheck := IsEmpty(pswdConf)

		//check for a return of true of IsEmpty
		if uNameCheck || emailCheck || pswdCheck || pswdConfCheck {
			fmt.Fprintf(w, "ErrorCode is -10: There is an error")
			log.Printf(": %v tried to log in and failed", userName)
		}
	})
}

//IsEmpty returns if field is empty
func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}
