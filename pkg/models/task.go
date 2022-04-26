package models

import (
	"github.com/BogdanT-1/calendar-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Task struct {
	gorm.Model
	Title        string `gorm:""json:"title"`
	Description  string `json:"description"`
	AssignedDate string `json:"assignedDate"`
	Importance   int64  `json:"importance"`
	Done         bool   `json:"done"`
}

type CompleteTask struct {
	IDs []int64 `json:"ids"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{}, &User{})
}

func (b *Task) CreateTask() *Task {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllTasks() []Task {
	var Tasks []Task
	db.Find(&Tasks)
	return Tasks
}

func GetTaskById(Id int64) (*Task, *gorm.DB) {
	var getTask Task
	db := db.Where("ID=?", Id).Find(&getTask)
	return &getTask, db
}

func DeleteTask(ID int64) Task {
	var Task Task
	db.Where("ID=?", ID).Delete(Task)
	return Task
}

func CompleteTasksByIds(IDs []int64) {
	db.Table("tasks").Where("id IN (?)", IDs).Updates(Task{Done: true})
}
