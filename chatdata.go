package wxworkfinancesdk

// ChatDataResponse 拉去消息返回值
type ChatDataResponse struct {
	Error
	ChatDataList []ChatData `json:"chatdata,omitempty"`
}

// IsError 判断是否错误
func (r ChatDataResponse) IsError() bool {
	return r.ErrCode != 0
}

// ChatData 加密消息数据
type ChatData struct {
	// Seq 消息的seq值，标识消息的序号。再次拉取需要带上上次回包中最大的seq。Uint64类型，范围0-pow(2,64)-1
	Seq uint64 `json:"seq,omitempty"`
	// MsgId 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。
	MsgId string `json:"msgid,omitempty"`
	// PublickeyVer 加密此条消息使用的公钥版本号。
	PublickeyVer uint32 `json:"publickey_ver,omitempty"`
	// EncryptRandomKey 使用publickey_ver指定版本的公钥进行非对称加密后base64加密的内容，需要业务方先base64 decode处理后，再使用指定版本的私钥进行解密，得出内容。
	EncryptRandomKey string `json:"encrypt_random_key,omitempty"`
	// EncryptChatMsg 消息密文。需要业务方使用将encrypt_random_key解密得到的内容，与encrypt_chat_msg，传入sdk接口DecryptData,得到消息明文。
	EncryptChatMsg string `json:"encrypt_chat_msg,omitempty"`
}
