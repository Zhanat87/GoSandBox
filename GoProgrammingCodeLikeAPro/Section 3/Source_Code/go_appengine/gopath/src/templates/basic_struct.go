package main

//import (
//	"os"
//	"text/template"
//)

type Basic_Struct struct {
	//exported field since it begins
	//with a capital letter
	Basic_field_one string
}


//func main() {
//	//define an instance
//	s := Basic_Struct{"value for a basic  field \n"}
//
//	//create a new template with some name
//	tmpl := template.New("test")
//
//	//parse some content and generate a template
//	tmpl, err := tmpl.Parse("Template value is  {{.Basic_field_one}}!")
//
//	//A common use of panic is to abort if a function
//	//returns an error value that we don't know how to
//	//(or want to) handle.
//	if err != nil {
//		panic(err)
//	}
//
//	//merge template 'tmpl' with content of 's'
//	err1 := tmpl.Execute(os.Stdout, s)
//
//	if err1 != nil {
//		panic(err1)
//	}
//}