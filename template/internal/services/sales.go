package services

import (
    "github.com/joshua468/template/internal/models"
    "gorm.io/gorm"
)

func CreateSales(db *gorm.DB, sales *models.Sales) error {
    return db.Create(sales).Error
}

func GetSalesByID(db *gorm.DB, id uint) (*models.Sales, error) {
    var sales models.Sales
    if err := db.First(&sales, id).Error; err != nil {
        return nil, err
    }
    return &sales, nil
}

func UpdateSales(db *gorm.DB, id uint, updatedFields map[string]interface{}) (*models.Sales, error) {
    var sales models.Sales
    if err := db.First(&sales, id).Error; err != nil {
        return nil, err
    }
    if err := db.Model(&sales).Updates(updatedFields).Error; err != nil {
        return nil, err
    }
    return &sales, nil
}

func DeleteSales(db *gorm.DB, id uint) error {
    return db.Delete(&models.Sales{}, id).Error
}
