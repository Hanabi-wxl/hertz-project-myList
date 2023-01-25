// Code generated by hertz generator.

package user

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"hertz-mylist/base/result"
	user "hertz-mylist/biz/model/user"
)

// UserRegister .
// @router /user/register [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var req user.UserRequest
	if err := c.BindAndValidate(&req); err != nil {
		panic(result.NewError(10000, "数据绑定异常"))
	}
	resp, err := user.Register(&req)

	if err != nil {
		panic(err)
	}

	c.JSON(consts.StatusOK, resp)
}