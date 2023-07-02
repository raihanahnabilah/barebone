# Development Basics of HTTP and REST API (Barebone)

A few days ago, I read a <a href="https://twitter.com/lynxluna/status/1674789858571870212?s=20">tweet</a> from an influencer saying that most developers these days don't know how to create a REST API without a database connection. Feeling challenged, I tried to do this in Python and Golang.

## Python implementation

Have yet to be implemented. I am planning to try two things: (1) using a framework, (2) without any framework.
## Golang

Since I am not very familiar with Golang's http syntax without using framework, I followed a tutorial from Adiatma Kamaarudin <a href="https://dev.to/adiatma/create-crud-apps-with-http-net-golang-without-database-18ip">here</a> but with some tweaks and additional queries. 

Run the following command to run the program:
```
go run main.go
```

I set the port to :8080. Feel free to change. 

### APIs

The APIs are as follows:
1. GET /budgets. To get a specific budget data by ID, use query param: "id".
2. POST /budgets. Put the data in the input body.
3. DELETE /budgets. Use query param: "id".
4. PUT /budgets. Use query param: "id" and put the data change in the input body. 

### Learning points

I learned how the HTTP request-response work if we're not using any Go framework (i.e. Gin, Gorilla/Mux). When I first learned Go, I immediately used these frameworks, but now I know how we connect to the HTTP server and how these requests and responses are handled without these frameworks. Moreover, I also learned that without a database connection, these data are just stored in the data structure that we specified (i.e. array, hashmap, etc), and when we restart the server, no data is saved. Though learning these things might seem "minor", they are quite fundamental in my learning progress.

Note that the queries are very slow because the data structure to store the data is an array, and to query, we are doing a for loop. A better data structure would be a hash map or a dictionary. This is for future development.