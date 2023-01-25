package task

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Task struct {
	ID        int
	Uid       uint
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted soft_delete.DeletedAt `gorm:"softDelete:flag"`
}

func BuildTaskResponse(task *Task) *TaskResponse {
	return &TaskResponse{
		TaskDetail: &TaskModel{
			ID:        int64(task.ID),
			UID:       int64(task.Uid),
			Title:     task.Title,
			Content:   task.Content,
			CreatedAt: task.CreatedAt.Unix(),
			UpdatedAt: task.UpdatedAt.Unix(),
		},
		Code: 20000,
	}
}

func BuildUpdateTask(tasku *TaskRequest, tasko *Task) *Task {
	tasko.Content = tasku.Content
	tasko.Title = tasku.Title
	return tasko
}

func BuildTaskModel(task *Task) *TaskModel {
	return &TaskModel{
		ID:        int64(task.ID),
		UID:       int64(task.Uid),
		Title:     task.Title,
		Content:   task.Content,
		UpdatedAt: task.UpdatedAt.Unix(),
		CreatedAt: task.CreatedAt.Unix(),
	}
}

func BuildTaskListResponse(list []*TaskModel, count int64) *TaskListResponse {
	return &TaskListResponse{
		TaskModel: list,
		Count:     count,
	}
}
