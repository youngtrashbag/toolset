package user

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
)

// RenderPage : Renders the page of a user
func (u *User) RenderPage() []byte {

	log.Printf("Rendering User ID=\"%d\"", u.ID)

	hTmpl, err := ioutil.ReadFile("templates/head.html")
	if err != nil {
		log.Panicln(err.Error())
	}
	uTmpl, err := ioutil.ReadFile("templates/user.html")
	if err != nil {
		log.Panicln(err.Error())
	}

	tmpl, err := template.New("user").Parse(string(append(hTmpl[:], uTmpl[:]...)))
	if err != nil {
		log.Panicln(err.Error())
	}

	var templateBytes bytes.Buffer

	err = tmpl.Execute(&templateBytes, u)
	if err != nil {
		log.Println(err.Error())
	}

	return templateBytes.Bytes()
}
