package result

import (
	"encoding/json"
	"errors"
)

type GlobalError struct {
	Id     string `json:"id"`
	Code   int    `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

type IError struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type CommonResult struct {
	Code   int         `json:"code"`
	Detail string      `json:"detail"`
	Data   interface{} `json:"data"`
}

func Ok(detail string) CommonResult {
	res := CommonResult{
		Code:   200,
		Detail: detail,
		Data:   nil,
	}
	return res
}

func NewResult(code int, detail string) CommonResult {
	res := CommonResult{
		Code:   code,
		Detail: detail,
	}
	return res
}

func NewIError(code int, detail string) error {
	jsonBytes, _ := json.Marshal(map[string]interface{}{
		"Code":   code,
		"Detail": detail,
	})
	return errors.New(string(jsonBytes))
}

func NewError(code int, detail string) GlobalError {
	err := GlobalError{
		Code:   code,
		Detail: detail,
	}
	return err
}

func Data(data interface{}) CommonResult {
	res := CommonResult{
		Code: 200,
		Data: data,
	}
	return res
}
