package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define the data
// They are not saved tho, so everytime you cancel the server
// and fire up again, it'll be restarted.
type Budget struct {
	ID          int
	Description string
	Price       int
}

// Define the collection
var budgets []Budget

func main() {
	fmt.Println("Fire up the server!")

	http.HandleFunc("/budget", // path
		// Handler
		// Response writer is the response you wrote to clients
		func(rw http.ResponseWriter, r *http.Request) {
			// Catch request from http
			// Different methods
			switch r.Method {
			case "GET":
				// Get the query
				query := r.URL.Query()

				if query.Get("id") == "" {
					json.NewEncoder(rw).Encode(budgets) // convert budgets and encode to json format
				}

				// convert string to int
				id, _ := strconv.Atoi(query.Get("id"))

				// loop budgets, match the id with the one we're looking for: (or maybe call the function)
				for _, budget := range budgets {
					if budget.ID == id {
						json.NewEncoder(rw).Encode(budget) // convert budgets and encode to json format
					}
				}

			case "POST":
				// create variable to save budget
				var budget Budget

				// decode `r.Body` from request to get data from request
				// decode the result to budget
				json.NewDecoder(r.Body).Decode(&budget)

				// append new budget in budgets:
				budgets = append(budgets, budget)

				// return to user:
				rw.Write([]byte("Budget added!"))

				// rw.Write([]byte("POST/ Hello world"))
			case "DELETE":
				// Get the query
				query := r.URL.Query()

				// convert string to int
				id, _ := strconv.Atoi(query.Get("id"))

				// loop budgets, match the id with the one we're looking for: (or maybe call the function)
				for i, budget := range budgets {
					if budget.ID == id {
						// Splice and replace
						budgets = append(budgets[:i], budgets[i+1:]...)
						rw.Write([]byte("Success to delete Budget!"))
						return
					}
				}
				rw.Write([]byte("ID not found!"))

			case "PUT":
				query := r.URL.Query()

				if query.Get("id") == "'" {
					rw.Write([]byte("Input ID is empty!"))
				}

				id, _ := strconv.Atoi(query.Get("id"))

				for i, budget := range budgets {
					json.NewDecoder(r.Body).Decode(&budget)

					if budget.ID == id {
						budgets[i].ID = budget.ID
						budgets[i].Description = budget.Description
						rw.Write([]byte("Success to update Budget!"))
						return
					}
				}

				rw.Write([]byte("ID not found!"))
			}
		})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
