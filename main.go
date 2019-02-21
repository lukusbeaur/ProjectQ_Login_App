package main

import (
	"fmt"
	"log"
	"net/http"

	helper "github.com/projects/ProjectQ_Login/ProjectQ_Login_App/helpers"
)

func main() {
	//better way to do this?
	//add username later
	email, emailConf, pswd, pswdConf := "", "", "", ""

	mux := http.NewServeMux()

	//Sign up new user
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) { 
		r.ParseForm()

		//call helper, checking if form is null
		//data from the form
		
		email = r.FormValue("email")
		pswd = r.FormValue("password")
		emailConf = r.FormValue("eConf")
		pswdConf = r.FormValue("passConf")

		//check for a return of true of IsEmpty
		if helper.IsEmpty(email) || helper.IsEmpty(pswd) || helper.IsEmpty(pswdConf) { //add username later
			fmt.Fprintf(w, "ErrorCode is -10: There is an error  \n")
			if len(email) == 0 {
				log.Printf("No username, Failed to signup\n")
				fmt.Fprintf(w, "There was no email entered, Try again\n") //change back to username check instead of email
			} else if len(pswd) == 0 {
				log.Printf("%v has failed to sign up, no password", email)
				fmt.Fprintf(w, "There was no password entered, Try again\n")
			} else {
				log.Printf(": %v failed signup", email)
				fmt.Fprint(w, "Potentian error: UserName already taken \n") //when checking for usernames
			}
			return
		}

		if helper.Comparable(pswd, pswdConf) && helper.Comparable(email, emailConf) {
			//This will be saved to database
			//will use mock data for now
			fmt.Fprintf(w, "Registration successfull\n")
		} else {
			fmt.Fprintf(w, "Passwords or Emails do not match\n")
		}
	})

	//login section
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { // concept proven
		r.ParseForm()

		email = r.FormValue("email")   //data from the form
		pswd = r.FormValue("password") //Data from the form

		if helper.IsEmpty(email) || helper.IsEmpty(pswd) {
			fmt.Fprintf(w, "Error Code is -10 : Missing Data")
			log.Printf(": %v failed to log in \n Empty Fields \n", email)
			return
			// find way to track number of attempts and send email potentially

		}

		//Mock data for testing
		//make this hidden eventually. make manditory password requirements
		dbPwb := "1234pass!" //temp
		dbEmail := "Test@email.com"

		if helper.Comparable(dbEmail, email) && helper.Comparable(dbPwb, pswd) {
			fmt.Fprintln(w, "Sucessfully login")
			log.Printf("%v has logged in", email) //usernames?
		} else {
			fmt.Fprintln(w, "Failed Login")
			log.Printf(": %v has failed to login", email)
		}
	})
	http.ListenAndServe(":8080", mux)
}
