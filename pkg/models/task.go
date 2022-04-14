package models

import (
	"github.com/BogdanT-1/calendar-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Task struct{
	gorm.Model
	Title string `gorm:""json:"Title"`
	Description string `json:"description"`
	AssignedDate string `json:"assignedDate"`
	Importance int64 `json:"importance"`
	Done bool `json:"done"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{})
}

func (b *Task) CreateTask() *Task{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllTasks() []Task{
	var Tasks []Task
	db.Find(&Tasks)
	return Tasks
}

func GetTaskById(Id int64) (*Task, *gorm.DB){
	var getTask Task
	db:=db.Where("ID=?", Id).Find(&getTask)
	return &getTask, db
}

func DeleteTask(ID int64) Task{
	var Task Task
	db.Where("ID=?", ID).Delete(Task)
	return Task
}