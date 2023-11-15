package wxworkfinancesdk

import "sync"

// MediaData 媒体消息结构体
type MediaData struct {
	OutIndexBuf string `json:"outindexbuf,omitempty"` // 消息分片索引信息
	IsFinish    bool   `json:"is_finish,omitempty"`   //
	Data        []byte `json:"data,omitempty"`        // 媒体数据内通
}

var mediaDataPool = sync.Pool{
	New: func() any {
		return new(MediaData)
	},
}

func MediaDataPool() *MediaData {
	return mediaDataPool.Get().(*MediaData)
}

func MediaDataPoolRelease(m *MediaData) {
	m.OutIndexBuf = ""
	m.IsFinish = false
	m.Data = nil
	mediaDataPool.Put(m)
}
