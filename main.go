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
			log.Printf(": %v tried to log in and failed\n", userName)
		}

		if pswd == pswdConf {
			//This will be saved to database
			//will use mock data for now
			fmt.Fprintln(w, "Registration successfull")
		} else {
			fmt.Fprint(w, "Passwords do not match.")
		}
	})

	//login section
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email = r.FormValue("email")   //data from the form
		pswd = r.FormValue("password") //Data from the form

		//IsEmpty Check
		emailCheck := IsEmpty(email)
		pswdCheck := IsEmpty(pswd)

		//make generic less code typed.
		if emailCheck || pswdCheck {
			fmt.Fprintf(w, "Error Code is -10 : Missing Data")
			return
			// find way to track number of attempts and send email potentially
			log.Printf(": %v failed to log in \n Empty Fields \n")
		}

		//Mock data for testing
		//make this hidden eventually. make manditory password requirements
		dbPwb := "1234pass!" //temp
		dbEmail := "Test@email.com"

		if email == dbEmail && pswd == dbPwb {
			fmt.Fprintln(w, "Sucessfully login")
			log.Printf("%v has logged in", dbEmail) //usernames?
		} else {
			fmt.Fprintln(w, "Failed Login")
			log.Printf(": %v has failed to login", dbEmail)
		}
	})
	http.ListenAndServe(":8080", mux)
}

//IsEmpty returns if field is empty
func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}
