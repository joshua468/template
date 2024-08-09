package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/joshua468/template/internal/models"
    "github.com/joshua468/template/internal/services"
    "gorm.io/gorm"
)

func CreateTransactionHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var transaction models.Transaction
        if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if err := services.CreateTransaction(db, &transaction); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(transaction)
    }
}

func GetTransactionHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.Atoi(r.URL.Query().Get("id"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        transaction, err := services.GetTransactionByID(db, uint(id))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(transaction)
    }
}
