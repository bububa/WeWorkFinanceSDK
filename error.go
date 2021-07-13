package WeWorkFinanceSDK

import (
	"fmt"
)

const (
	SDKErrMsg = "sdk failed" // SDK 错误消息
)

// SDK Error结构体
type Error struct {
	ErrCode int    `json:"errcode,omitempty"` // 错误代码
	ErrMsg  string `json:"errmsg,omitempty"`  // 错误说明
}

// Error implement error interface
func (this Error) Error() string {
	return fmt.Sprintf("%d:%s", this.ErrCode, this.ErrMsg)
}

// NewSDKError 新建错误
func NewSDKErr(code int) Error {
	return Error{ErrCode: code, ErrMsg: SDKErrMsg}
}
