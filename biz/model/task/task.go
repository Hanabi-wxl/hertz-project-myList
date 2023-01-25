package task

import (
	"hertz-mylist/base/result"
	"hertz-mylist/biz/model/gorm"
)

func CreateTask(req *TaskRequest) (*TaskResponse, error) {
	task := Task{
		Uid:     uint(req.Uid),
		Title:   req.Title,
		Content: req.Content,
	}
	if err := gorm.DB.Create(&task).Error; err != nil {
		return nil, result.NewIError(10006, "创建备忘录失败")
	}
	return BuildTaskResponse(&task), nil
}

func UpdateTask(req *TaskRequest) (*TaskResponse, error) {
	var task Task
	var count int64
	if err := gorm.DB.Where("id = ? and uid = ?", req.Id, req.Uid).Find(&task).Count(&count).Error; err != nil {
		return nil, result.NewIError(10006, "查询数据库失败")
	}
	if count == 0 {
		return nil, result.NewIError(10007, "请勿修改他人笔记")
	}
	updateTask := BuildUpdateTask(req, &task)

	if err := gorm.DB.Save(updateTask).Error; err != nil {
		return nil, result.NewIError(10006, "修改失败")
	}
	return BuildTaskResponse(updateTask), nil
}

func DeleteTask(req *TaskRequest) (*TaskResponse, error) {
	var task Task
	var count int64
	if err := gorm.DB.Model(&task).Where("id = ? and uid = ?", req.Id, req.Uid).Count(&count).Error; err != nil {
		return nil, result.NewIError(10006, "查询数据库失败")
	}
	if count == 0 {
		return nil, result.NewIError(10007, "请勿删除他人笔记")
	}
	task.ID = int(req.Id)
	if err := gorm.DB.Delete(&task).Error; err != nil {
		return nil, result.NewIError(10006, "查询数据库失败")
	}
	return BuildTaskResponse(&task), nil
}

func GetTaskList(req *TaskRequest) (*TaskListResponse, error) {
	var tasks []Task
	size := int(req.PageSize)
	pageNum := int(req.PageNum)
	offset := (pageNum - 1) * size
	uid := req.Uid
	var count int64
	if err := gorm.DB.Where("uid = ?", uid).Offset(offset).Limit(size).Find(&tasks).Error; err != nil {
		return nil, result.NewIError(10006, "查询数据库失败")
	}
	if err := gorm.DB.Model(&Task{}).Where("uid = ?", uid).Count(&count).Error; err != nil {
		return nil, result.NewIError(10006, "查询数据库失败")
	}
	var tasksDetail []*TaskModel
	for _, task := range tasks {
		tasksDetail = append(tasksDetail, BuildTaskModel(&task))
	}
	return BuildTaskListResponse(tasksDetail, count), nil
}

func GetTaskDetail(req *TaskRequest) (*TaskResponse, error) {
	task := Task{
		ID:  int(req.Id),
		Uid: uint(req.Uid),
	}
	if err := gorm.DB.First(&task).Error; err != nil {
		return nil, result.NewIError(10006, "获取备忘录失败")
	}
	return BuildTaskResponse(&task), nil
}
