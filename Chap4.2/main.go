package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World!") // write data to response
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		//Verification
		//username
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("username")); !m {
			fmt.Println("Username is invalid!")
		} else {
			fmt.Println("Username is correct.")
		}

		//password
		if pwlen := len(r.Form["password"][0]); pwlen < 4 || pwlen > 16 {
			fmt.Println("Password is too short or long!")
		} else {
			fmt.Println("Password is correct.")
		}

		//email
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("Email is invalid!")
		} else {
			fmt.Println("Email is correct.")
		}

		//chinese name
		if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("chinese")); !m {
			fmt.Println("Chinese name is invalid!")
		} else {
			fmt.Println("Chinese name is correct.")
		}

		//fruit
		slice := []string{"apple", "pear", "banana"}
		found := false
		for _, v := range slice {
			if v == r.Form.Get("fruit") {
				fmt.Println("Fruit is correct.")
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Fruit is invalid!")
		}

		//interest
		slice = []string{"football", "basketball", "tennis"}
		found = false
		for _, v := range slice {
			if v == r.Form.Get("interest") {
				fmt.Println("Interest is correct.")
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Interest is invalid!")
		}

		//age
		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Println("Age is invalid!")
		} else if getint < 0 {
			fmt.Println("Age must be a positve number!")
		} else {
			fmt.Println("Age is correct.")
		}

		//gender
		slice = []string{"1", "2", "3"}
		found = false
		for _, v := range slice {
			if v == r.Form.Get("gender") {
				fmt.Println("Gender is correct.")
				found = true
			}
		}
		if !found {
			fmt.Println("Gender is invalid!")
		}

		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("chinese:", r.Form["chinese"])
		fmt.Println("fruit:", r.Form["fruit"])
		fmt.Println("interest:", r.Form["interest"])
		fmt.Println("age:", r.Form["age"])
		fmt.Println("gender:", r.Form["gender"])
	}
}

func main() {
	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
