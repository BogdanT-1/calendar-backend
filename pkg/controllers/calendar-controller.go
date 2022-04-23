package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BogdanT-1/calendar-backend/pkg/models"
	"github.com/BogdanT-1/calendar-backend/pkg/utils"
	"github.com/gorilla/mux"
)

var NewTask models.Task

func GetTasks(w http.ResponseWriter, r *http.Request){
	tasks:=models.GetAllTasks()
	res, _ :=json.Marshal(tasks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTaskById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	ID, err:= strconv.ParseInt(taskId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	taskDetails, _:= models.GetTaskById(ID)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request){
	CreateTask := &models.Task{}
	utils.ParseBody(r, CreateTask)
	b:= CreateTask.CreateTask()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTask(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	ID, err := strconv.ParseInt(taskId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	task := models.DeleteTask(ID)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTask(w http.ResponseWriter, r *http.Request){
	var updateTask = &models.Task{}
	utils.ParseBody(r, updateTask)
	vars := mux.Vars(r)
	taskId := vars["taskId"]
	ID, err := strconv.ParseInt(taskId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	taskDetails, db:=models.GetTaskById(ID)
	taskDetails = updateTask;
	db.Save(&taskDetails)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CompleteTasks(w http.ResponseWriter, r *http.Request){
	var toCompleteTasks = &models.CompleteTask{}
	utils.ParseBody(r, toCompleteTasks)
	models.CompleteTasksByIds(toCompleteTasks.IDs)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
}