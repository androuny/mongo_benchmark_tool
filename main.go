package main

import (
	"fmt"
	"log"
	"mongo_benchmark_tool/handlers"
)

func main() {
    var uri string
    var db_name string
    var collection_name string
    var insert_users int

    fmt.Println("Starting MongoDB Benchmark tool... ")
    fmt.Println("")

    fmt.Println("Please enter your mongodb connection uri: ")
    fmt.Scanf("%s \n", &uri)

    fmt.Println("Please enter your mongodb db name: ")
    fmt.Scanf("%s \n", &db_name)

    fmt.Println("Please enter your mongodb db collection name: ")
    fmt.Scanf("%s \n", &collection_name)

    fmt.Println("Please enter the amount of users you want to insert for the test: ")
    fmt.Scanf("%d \n", &insert_users)


    handler, connTime, err := handlers.NewMongoHandler(uri, db_name, collection_name)
    //handler, connTime, err := handlers.NewMongoHandler("mongodb://192.168.0.195/", "dev_db", "users")
    if err != nil {
        log.Fatalf("[DB_INIT] %v ", err)
    }
    fmt.Printf("Database Connection Success, took %v ms \n", connTime.Milliseconds())

    genTime, writeTime, err2 := handler.MakeNewUsers(insert_users)
    if err2 != nil {
        log.Fatalf("[DB_WRITE] %v", err2)
    }
    fmt.Printf("Database WRITE Success, took %v ms (UserGen took %v ms) \n", writeTime.Milliseconds(), genTime.Milliseconds())

    readTime, err1 := handler.ReadAllUsers()
    if err1 != nil {
        log.Fatalf("[DB_READ_ALL] %v", err1)
    }
    fmt.Printf("Database READ Success, took %v ms \n", readTime.Milliseconds())
}