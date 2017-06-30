package apptest

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Eddy_User struct {
	Timestamp time.Time
	User_Name   string
	Email string
}


func NewEddy_User() *Eddy_User {
	return &Eddy_User{
		Timestamp: time.Now(),
	}
}


func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/userinfo", userinfo)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, userinfoForm)
}

// http://shadynasty.biz/blog/2012/07/30/quick-and-clean-in-go/

func userinfo(w http.ResponseWriter, req *http.Request) {

	user := NewEddy_User( )
	user.User_Name = req.FormValue("User_Name")
	user.Email = req.FormValue("Email")

	if user.User_Name == "" {
		user.User_Name = "Some dummy who forgot a name"
	}
	if user.Email  == "" {
		user.Email  = "Some dummy who forgot a message."
	}


	err := userinfoTemplate.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


const userinfoForm = `
<html>
  <body>
    <form action="/userinfo" method="post">
      <div><input  name="User_Name"   ></input ></div>
       <div><input  name="Email"  ></input ></div>
      <div><input type="submit" value="Enter User Info"></div>
    </form>
  </body>
</html>
`


var userinfoTemplate = template.Must(template.New("userinfo").Parse(tmpl))

const tmpl = `User name: {{.User_Name}},  User email id is {{.Email}} `