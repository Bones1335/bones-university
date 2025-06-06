package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func handlerGetIndex(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/layout.html", "templates/index.html")
	if err != nil {
		fmt.Printf("error parsing html templates: %v", err)
		return
	}

	err = temp.Execute(w, temp)
	if err != nil {
		fmt.Printf("problem executing template data: %v", err)
		return
	}
}

func handlerCreateEnrollment(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/layout.html", "templates/enrollment.html")
	if err != nil {
		fmt.Printf("error parsing html templates: %v", err)
		return
	}

	err = temp.Execute(w, temp)
	if err != nil {
		fmt.Printf("problem executing template data: %v", err)
		return
	}
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/layout.html", "templates/login.html")
	if err != nil {
		fmt.Printf("error parsing html templates: %v", err)
		return
	}

	err = temp.Execute(w, temp)
	if err != nil {
		fmt.Printf("problem executing template data: %v", err)
		return
	}
}

func handlerStudentDashboard(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/layout.html", "templates/student_dashboard.html")
	if err != nil {
		fmt.Printf("error parsing html templates: %v", err)
		return
	}

	err = temp.Execute(w, temp)
	if err != nil {
		fmt.Printf("problem executing template data: %v", err)
		return
	}
}

func handlerAdminDashboard(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/layout.html", "templates/admin_dashboard.html")
	if err != nil {
		fmt.Printf("error parsing html templates: %v", err)
		return
	}

	err = temp.Execute(w, temp)
	if err != nil {
		fmt.Printf("problem executing template data: %v", err)
		return
	}
}
