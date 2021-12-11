//curl https://api.securitytrails.com/v1/prototype/dslv2?page=3 -X POST --header "Content-Type: application/json" --header "apikey: 35ZoQkxGLPHIYYgepbnrihU0Km6oS8Uh"  --data "{\"query\": \"QUERYHERE\"}"

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Query struct {
	Query string `json:"query"`
}

// the main dslv2 function
func dsl(query string) {
	queryStruct := Query{query}
	queryJSON, err := json.Marshal(queryStruct)
	if err != nil {
		log.Println("Could not encode query to JSON", err)
		os.Exit(1)
	}
	queryString := string(queryJSON)
	response := getResponse("POST", "search/list", queryString)
	var results map[string]interface{}
	json.Unmarshal([]byte(response), &results)
	metaInterface, ok := results["meta"].(map[string]interface{})
	if !ok { // no results
		if output == "list" {
			return
		} else {
			fmt.Println(response)
		}
		return
	}

	totalPages := metaInterface["total_pages"].(float64) // total number of pages
	fmt.Println(response)                                // print the first page
	// print all the other pages
	for i := 2; i <= int(totalPages); i++ {
		response = getResponse("POST", "search/list?page="+strconv.Itoa(i), queryString)
		fmt.Println(response)
	}
}
