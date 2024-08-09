package services

import (
    "github.com/joshua468/template/internal/models"
    "gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
    return db.Create(user).Error
}

func GetUserByID(db *gorm.DB, id uint) (*models.User, error) {
    var user models.User
    if err := db.First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func UpdateUser(db *gorm.DB, id uint, updatedFields map[string]interface{}) (*models.User, error) {
    var user models.User
    if err := db.First(&user, id).Error; err != nil {
        return nil, err
    }
    if err := db.Model(&user).Updates(updatedFields).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func DeleteUser(db *gorm.DB, id uint) error {
    return db.Delete(&models.User{}, id).Error
}
