package myHandler

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"hertz-mylist/base/logging"
	"hertz-mylist/base/result"
)

func Handler(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
	logging.Info(c, "[Recovery] myHandler =", err, string(stack), string(ctx.Request.Header.UserAgent()))

	errStr := make(map[string]interface{}, 3)

	switch err.(type) {
	case result.GlobalError:
		globalError := err.(result.GlobalError)
		errStr = map[string]interface{}{
			"code":   globalError.Code,
			"detail": globalError.Detail,
			"data":   nil,
		}
	case error:
		var ie result.IError
		iError := err.(error)
		_ = json.Unmarshal([]byte(iError.Error()), &ie)
		errStr = map[string]interface{}{
			"code":   ie.Code,
			"detail": ie.Detail,
			"data":   nil,
		}
	}

	ctx.JSON(500, errStr)
	ctx.AbortWithStatus(consts.StatusInternalServerError)
}
