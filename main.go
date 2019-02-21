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
	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) { // still has issues. Automatic login
		r.ParseForm()

		//call helper, checking if form is null
		//make generic eventually and call each func as a list ? instead of hard coding

		//data from the form
		email = r.FormValue("email")
		pswd = r.FormValue("password")
		emailConf = r.FormValue("eConf")
		pswdConf = r.FormValue("passConf")

		//uNameBool := helper.IsEmpty(userName)
		emailBool := helper.IsEmpty(email)
		pswdBool := helper.IsEmpty(pswd)
		pswdConfBool := helper.IsEmpty(pswdConf)

		//check for a return of true of IsEmpty
		if emailBool || pswdBool || pswdConfBool { //add username later
			fmt.Fprintf(w, "ErrorCode is -10: There is an error  \n")
			if len(email) == 0 {
				log.Printf("No username, Failed to signup")
				fmt.Fprintf(w, "There was no email entered, Try again") //change back to username check instead of email
			} else {
				log.Printf(": %v failed signup", email)
				fmt.Fprint(w, "Potentian error: UserName already taken") //when checking for usernames
			}
		}

		if pswd == pswdConf && email == emailConf {
			//This will be saved to database
			//will use mock data for now
			fmt.Fprintln(w, "Registration successfull")
		} else {
			fmt.Fprint(w, "Passwords or Emails do not match")
		}
	})

	//login section
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { // concept proven
		r.ParseForm()

		email = r.FormValue("email")   //data from the form
		pswd = r.FormValue("password") //Data from the form

		//IsEmpty Check
		//put this stupid line into one line of code in the if statement. and eliminate the variable all together. when concept is done
		emailBool := helper.IsEmpty(email)
		pswdBool := helper.IsEmpty(pswd)

		//make generic less code typed.
		if emailBool || pswdBool {
			fmt.Fprintf(w, "Error Code is -10 : Missing Data")
			log.Printf(": %v failed to log in \n Empty Fields \n", email)
			return
			// find way to track number of attempts and send email potentially

		}

		//Mock data for testing
		//make this hidden eventually. make manditory password requirements
		dbPwb := "1234pass!" //temp
		dbEmail := "Test@email.com"

		if email == dbEmail && pswd == dbPwb {
			fmt.Fprintln(w, "Sucessfully login")
			log.Printf("%v has logged in", email) //usernames?
		} else {
			fmt.Fprintln(w, "Failed Login")
			log.Printf(": %v has failed to login", email)
		}
	})
	http.ListenAndServe(":8080", mux)
}
