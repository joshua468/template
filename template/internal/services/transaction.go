package services

import (
    "github.com/joshua468/template/internal/models"
    "gorm.io/gorm"
)

func CreateTransaction(db *gorm.DB, transaction *models.Transaction) error {
    return db.Create(transaction).Error
}

func GetTransactionByID(db *gorm.DB, id uint) (*models.Transaction, error) {
    var transaction models.Transaction
    if err := db.First(&transaction, id).Error; err != nil {
        return nil, err
    }
    return &transaction, nil
}

func UpdateTransaction(db *gorm.DB, id uint, updatedFields map[string]interface{}) (*models.Transaction, error) {
    var transaction models.Transaction
    if err := db.First(&transaction, id).Error; err != nil {
        return nil, err
    }
    if err := db.Model(&transaction).Updates(updatedFields).Error; err != nil {
        return nil, err
    }
    return &transaction, nil
}

func DeleteTransaction(db *gorm.DB, id uint) error {
    return db.Delete(&models.Transaction{}, id).Error
}
