package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	// get configuration
	address := flag.String("server", "http://localhost:8080", "HTTP gateway url, e.g. http://localhost:8080")
	flag.Parse()

	var body string

	resp, err := http.Post(*address+"/v1/multiply", "application/json", strings.NewReader(
		fmt.Sprintf(
			`
				{
					"api":"v1",
					"number1":%f,
					"number2":%f
				}
			`, 5, 3)))

	resp, err := http.Post(*address+"/v1/todo", "application/json", strings.NewReader(fmt.Sprintf(`
		{
			"api":"v1",
			"toDo": {
				"title":"title (%s)",
				"description":"description (%s)",
				"reminder":"%s"
			}
		}
	`, pfx, pfx, pfx)))

	if err != nil {
		log.Fatalf("failed to call multiply method: %v", err)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		body = fmt.Sprintf("failed read multiply response body: %v", err)
	} else {
		body = string(bodyBytes)
	}

	log.Printf("multiply response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

	var numResponse struct {
		API     string  `json:"api"`
		Number1 float64 `json:"number_1"`
	}

	err = json.Unmarshal(bodyBytes, &numResponse)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON response of multiply method: %v", err)
		fmt.Println("error:", err)
	}

}
