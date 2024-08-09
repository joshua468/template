package main

import (
    "net/http"
    "os"
    "github.com/joshua468/template/config"
    "github.com/joshua468/template/internal/handlers"
    "gorm.io/gorm"
    "log"
)

func main() {
    db := config.InitDB()
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    http.HandleFunc("/create_sales", handlers.CreateSalesHandler(db))
    http.HandleFunc("/get_sales", handlers.GetSalesHandler(db))

    http.HandleFunc("/create_transaction", handlers.CreateTransactionHandler(db))
    http.HandleFunc("/get_transaction", handlers.GetTransactionHandler(db))

    http.HandleFunc("/create_user", handlers.CreateUserHandler(db))
    http.HandleFunc("/get_user", handlers.GetUserHandler(db))

    log.Println("Server is running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
