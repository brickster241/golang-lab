package learnlab

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func textTemplates() {
	tmpl, err := template.New("intro").Parse("Welcome, {{.name}}! How are you doing :) ??")
	if err != nil {
		panic(err)
	}

	/* Another Way without err
	tmpl = template.Must(template.New("intro").Parse("Welcome, {{.name}}! How are you doing :) ??"))
	*/

	// Define data for intro message template
	data := map[string]interface{}{
		"name": "AMV",
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

	// Use Template with custom Input as well
	reader := bufio.NewReader(os.Stdin)
	st, _ := reader.ReadString('\n')
	st = strings.TrimSpace(st)

	// Define Named Templates for different types of templates
	templates := map[string]string {
		"welcome": "Welcome, {{.name}}! We are glad you joined.",
		"notification": "{{.nm}}, you have a new notification : {{.ntf}}",
		"error": "Oops! An error occured: {{.errMsg}}",
	}

	// Parse and Store Templates
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// Show Menu
		fmt.Println("\nMenu :")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Printf("Choose an Option : ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string] interface{}
		var tmp *template.Template

		switch choice {
		case "1":
			tmp = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": st}
		
		case "2":
			tmp = parsedTemplates["notification"]
			fmt.Print("Enter Notification Message : ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			data = map[string]interface{}{"nm": st, "ntf": notification}
		
		case "3":
			tmp = parsedTemplates["error"]
			fmt.Print("Enter Error Message : ")
			errorMsg, _ := reader.ReadString('\n')
			errorMsg = strings.TrimSpace(errorMsg)
			data = map[string]interface{}{"errMsg" : errorMsg}
		
		case "4":
			fmt.Println("Exiting....")
			return
		
		default:
			fmt.Println("Invalid Choice. Please select a valid option !")
			continue
		}

		err := tmp.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error Executing Template :(")
			break
		}

	}
}
