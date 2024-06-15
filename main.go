package main

import (
	"fmt"
	"log"
	"mongo_benchmark_tool/handlers"
)

func main() {

    fmt.Println("Starting MongoDB Benchmark tool...")

    handler, connTime, err := handlers.NewMongoHandler("mongodb://localhost:27017/", "dev_db", "users")
    //handler, connTime, err := handlers.NewMongoHandler("mongodb://192.168.0.195/", "dev_db", "users")
    if err != nil {
        log.Fatalf("[DB_INIT] %v ", err)
    }
    fmt.Printf("Database Connection Success, took %v ms \n", connTime.Milliseconds())

    genTime, writeTime, err2 := handler.MakeNewUsers(10_000)
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