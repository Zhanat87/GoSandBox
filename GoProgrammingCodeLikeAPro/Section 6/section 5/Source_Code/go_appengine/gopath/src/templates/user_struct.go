package main

import (
	"os"
	"text/template"
)

type User struct {
	User_Name   string
	Emails []string
}

const tmpl = `{{$name := .User_Name}}
{{range .Emails}}
  User name: {{$name}},  User email id is {{.}}
{{end}}
`

func main() {
	user := User{
		User_Name:   "Eddy",
		Emails: []string{"eddy@eduonix.com", "eddy@eddysmail.com"},
	}

	t := template.New("User template")

	t, err := t.Parse(tmpl)

	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}
}