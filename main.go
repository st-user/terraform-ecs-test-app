package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("activity: This is activity %s\n", r.URL.Path)
		fmt.Printf("log: This is normal log %s\n", r.URL.Path)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("activity: This is activity %s\n", r.URL.Path)
		fmt.Printf("log: This is normal log %s\n", r.URL.Path)
		fmt.Fprintf(w, "Hi, Hi")
	})

	http.HandleFunc("/checkconfig", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("activity: This is activity %s\n", r.URL.Path)
		fmt.Printf("log: This is normal log %s\n", r.URL.Path)

		// send http request to 'http://localhost:2772/'
		// sample URL: $ curl "http://localhost:2772/applications/application_name/environments/environment_name/configurations/configuration_name"
		applicationName := "my-appconfig-ecs-application"
		environmentName := "Beta"
		profileName := "my-appconfig-configuration-profile"
		resp, err := http.Get("http://localhost:2772/applications/" + applicationName + "/environments/" + environmentName + "/configurations/" + profileName)
		if err != nil {
			fmt.Fprintf(w, "Error: %s", err)
			return
		}
		defer resp.Body.Close()
		respBody := make([]byte, 1024)
		// Read all bytes from response body and convert to string
		resp.Body.Read(respBody)
		fmt.Fprintf(w, "Response from appconfig: %s", string(respBody))

	})

	fmt.Println("Start.")

	log.Fatal(http.ListenAndServe(":8081", nil))
}
