package main

// Example from "How to Use Templates in Go"
// https://www.digitalocean.com/community/tutorials/how-to-use-templates-in-go#step-2-creating-template-data

import (
	"os"
	"strings"

	// "text/template"
	"html/template"
)

type Pet struct {
	Name   string
	Sex    string
	Intact bool
	Age    string
	Breed  string
	// Breed  []string
}

func main() {
	dogs := []Pet{
		{
			Name:   "<script>alert(\"Gotcha!\");</script>Jujube",
			Sex:    "Female",
			Intact: false,
			Age:    "10 months",
			Breed:  "German Shepherd/Pitbull",
			// Breed: []string{"German Shepherd", "Pitbull"},
		},
		{
			Name:   "Zephyr",
			Sex:    "Male",
			Intact: true,
			Age:    "13 years, 3 months",
			Breed:  "German Shepherd/Border Collie",
			// Breed: []string{"German Shepherd", "Border Collie"},
		},
		{
			Name:   "Bruce Wayne",
			Sex:    "Male",
			Intact: false,
			Age:    "3 years, 8 months",
			Breed:  "Chihuahua",
			// Breed: []string{"Chihuahua"},
		},
	}

	funcMap := template.FuncMap{
		"dec":     func(i int) int { return i - 1 },
		"replace": strings.ReplaceAll,
		"join":    strings.Join,
	}

	// tmplFile := "pets.tmpl"
	// tmplFile := "lastPet.tmpl"
	tmplFile := "petsHtml.tmpl"
	// tmplFilePath := "cmd/pets/" + tmplFile

	tmpl, err := template.New(tmplFile).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("pets.html")
	if err != nil {
		panic(err)
	}

	// err = tmpl.Execute(os.Stdout, dogs)
	err = tmpl.Execute(f, dogs)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
