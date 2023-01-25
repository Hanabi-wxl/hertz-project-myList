// Code generated by hertz generator.

package task

import (
	"context"
	"hertz-mylist/biz/router/middleware"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	task "hertz-mylist/biz/model/task"
)

// CreateTask .
// @router /task/create [POST]
func CreateTask(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.TaskRequest
	err = c.BindAndValidate(&req)
	claim, exists := c.Get(middleware.IdentityKey)
	if exists {
		req.Uid = uint64(claim.(*middleware.Claim).ID)
	}
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	taskResponse, err := task.CreateTask(&req)
	if err != nil {
		panic(err)
	}
	c.JSON(consts.StatusOK, taskResponse)
}

// UpdateTask .
// @router /task/update [PUT]
func UpdateTask(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.TaskRequest
	err = c.BindAndValidate(&req)
	claim, exists := c.Get(middleware.IdentityKey)
	if exists {
		req.Uid = uint64(claim.(*middleware.Claim).ID)
	}
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	taskResponse, err := task.UpdateTask(&req)
	if err != nil {
		panic(err)
	}
	c.JSON(consts.StatusOK, taskResponse)
}

// DeleteTask .
// @router /task/delete/:id [DELETE]
func DeleteTask(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.TaskRequest
	err = c.BindAndValidate(&req)
	claim, exists := c.Get(middleware.IdentityKey)
	if exists {
		req.Uid = uint64(claim.(*middleware.Claim).ID)
	}
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	taskResponse, err := task.DeleteTask(&req)
	if err != nil {
		panic(err)
	}
	c.JSON(consts.StatusOK, taskResponse)
}

// GetTaskDetail .
// @router /task/getDetail/:id [GET]
func GetTaskDetail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.TaskRequest
	err = c.BindAndValidate(&req)
	claim, exists := c.Get(middleware.IdentityKey)
	if exists {
		req.Uid = uint64(claim.(*middleware.Claim).ID)
	}
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	taskResponse, err := task.GetTaskDetail(&req)
	if err != nil {
		panic(err)
	}
	c.JSON(consts.StatusOK, taskResponse)
}

// GetTaskList .
// @router /task/getList/:pageSize/:pageNum [GET]
func GetTaskList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req task.TaskRequest
	err = c.BindAndValidate(&req)
	claim, exists := c.Get(middleware.IdentityKey)
	if exists {
		req.Uid = uint64(claim.(*middleware.Claim).ID)
	}
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	taskResponse, err := task.GetTaskList(&req)
	if err != nil {
		panic(err)
	}
	c.JSON(consts.StatusOK, taskResponse)
}
