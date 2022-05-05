package models

import (
	"github.com/BogdanT-1/calendar-backend/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var LoggedInUser string

type Task struct {
	gorm.Model
	Title        string `gorm:""json:"title"`
	Description  string `json:"description"`
	AssignedDate string `json:"assignedDate"`
	Importance   int64  `json:"importance"`
	User         string `json:"user_email"`
	Done         bool   `json:"done"`
}

type CompleteTask struct {
	IDs []int64 `json:"ids"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{}, &User{}, &Sessions{})
}

func (b *Task) CreateTask() *Task {
	b.User = LoggedInUser
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllTasks() []Task {
	var Tasks []Task
	db.Where("User=?", LoggedInUser).Find(&Tasks)
	return Tasks
}

func GetTaskById(Id int64) (*Task, *gorm.DB) {
	var getTask Task
	db := db.Where("ID=?", Id).Find(&getTask)
	return &getTask, db
}

func GetTasksByAssignedDate(assignedDate string) []Task {
	var getTasks []Task
	db.Where("User=?", LoggedInUser).Where("assigned_date=?", assignedDate).Find(&getTasks)
	return getTasks
}

func DeleteTask(ID int64) Task {
	var Task Task
	db.Where("ID=?", ID).Delete(Task)
	return Task
}

func CompleteTasksByIds(IDs []int64) {
	db.Table("tasks").Where("id IN (?)", IDs).Updates(Task{Done: true})
}
