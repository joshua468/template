package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/joshua468/template/internal/models"
    "github.com/joshua468/template/internal/services"
    "gorm.io/gorm"
)

func CreateSalesHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var sales models.Sales
        if err := json.NewDecoder(r.Body).Decode(&sales); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if err := services.CreateSales(db, &sales); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(sales)
    }
}

func GetSalesHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.Atoi(r.URL.Query().Get("id"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        sales, err := services.GetSalesByID(db, uint(id))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(sales)
    }
}
