package wxworkfinancesdk

import (
	"fmt"
)

const (
	// SDKErrMsg 错误消息
	// SDKInitErrMsg sdk 初始化失败
	SDKInitErrMsg = "sdk init failed"
	// GetChatDataErrMsg 调用 GetChatData API 失败
	GetChatDataErrMsg = "call GetChatData failed"
	// DecryptErrMsg 调用 DecryptMessage API 失败
	DecryptErrMsg = "call DecryptMessage API failed"
	// GetMediaDataErrMsg 调用 GetMediaData API 失败
	GetMediaDataErrMsg = "call GetMediaData API failed"
)

// Error SDK Error结构体
type Error struct {
	// ErrCode 错误代码
	ErrCode int `json:"errcode,omitempty"`
	// ErrMsg 错误说明
	ErrMsg string `json:"errmsg,omitempty"`
	Hint   string `json:"hint,omitempty"`
}

// Error implement error interface
func (e Error) Error() string {
	if e.Hint != "" {
		return fmt.Sprintf("errcode:%d, errmsg:%s, hint:%s", e.ErrCode, e.ErrMsg, e.Hint)
	}
	return fmt.Sprintf("errcode:%d, errmsg:%s", e.ErrCode, e.ErrMsg)
}

// NewSDKErr 新建错误
func NewSDKErr(code int, msg string) Error {
	return Error{ErrCode: code, ErrMsg: msg}
}
