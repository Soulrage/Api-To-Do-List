package service

import (
	"To-Do/internal/models"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"time"
)
var ErrTaskNotFound = errors.New("task not found")



func GetCurrentTimeRFC3339() string {
	return time.Now().Format(time.RFC3339)
}

func GetLastID(db *gorm.DB) uint{
	var lastTask models.Tasks

	err := db.Order("id desc").First(&lastTask).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {

			return 0
		}
		return 0
	}
	return lastTask.ID
}


func CreateTask(db *gorm.DB, title string, description string, dueDate string) error {

	task := models.Tasks{
		ID:      GetLastID(db) + 1,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
		CreateAt:    GetCurrentTimeRFC3339(),
		UpdateAt: "",
	}
	fmt.Println(task)
	if err := db.Model(&models.Tasks{}).Create(&task).Error; err != nil {
		return err
	}
	return nil
}




func GetAllTasks(db *gorm.DB)([]models.Tasks, error){
	var tasks []models.Tasks
	if err := db.Model(&models.Tasks{}).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil

}





func GetTasksById(db *gorm.DB, CurrentID uint) (models.Tasks, error) {
	var tasks models.Tasks
	result := db.Where("id = ?", CurrentID).First(&tasks)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.Tasks{}, ErrTaskNotFound
		}
		return models.Tasks{}, result.Error
	}
	return tasks, nil
}



func UpdateTask(db *gorm.DB, ID uint, updatedTask models.Tasks) error {
	var task models.Tasks

	if err := db.Model(&models.Tasks{}).First(&task, ID).Error; err != nil {
		return err
	}

	if err := db.Model(&task).Updates(updatedTask).Error; err != nil {
		return err
	}

	return nil
}






func DeleteTaskById(db *gorm.DB, ID uint) error {
	result := db.Where("id = ?", ID).Delete(&models.Tasks{})
	if result.RowsAffected == 0 {
		return ErrTaskNotFound
	}
	if result.Error != nil {
		return result.Error
	}

	return nil
}




