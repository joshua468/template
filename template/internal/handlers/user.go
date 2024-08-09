package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/joshua468/template/internal/models"
    "github.com/joshua468/template/internal/services"
    "gorm.io/gorm"
)

func CreateUserHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var user models.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        if err := services.CreateUser(db, &user); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user)
    }
}

func GetUserHandler(db *gorm.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id, err := strconv.Atoi(r.URL.Query().Get("id"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        user, err := services.GetUserByID(db, uint(id))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(user)
    }
}
