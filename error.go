package WeWorkFinanceSDK

import (
	"fmt"
)

const (
	SDKErrMsg = "sdk failed"
)

type Error struct {
	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}

func (this Error) Error() string {
	return fmt.Sprintf("%d:%s", this.ErrCode, this.ErrMsg)
}

func NewSDKErr(code int) Error {
	return Error{ErrCode: code, ErrMsg: SDKErrMsg}
}
