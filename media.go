package wxworkfinancesdk

// MediaData 媒体消息结构体
type MediaData struct {
	OutIndexBuf string `json:"outindexbuf,omitempty"` // 消息分片索引信息
	IsFinish    bool   `json:"is_finish,omitempty"`   //
	Data        []byte `json:"data,omitempty"`        // 媒体数据内通
}
