package wxworkfinancesdk

// #cgo LDFLAGS: -L${SRCDIR}/lib -lWeWorkFinanceSdk_C
// #cgo CFLAGS: -Wall
// #cgo CFLAGS: -I ${SRCDIR}/lib/
// #include <stdlib.h>
// #include "WeWorkFinanceSdk_C.h"
import "C"
import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"unsafe"
)

// Client Client对象
type Client struct {
	ptr *C.WeWorkFinanceSdk_t
}

// NewClient 初始化函数
// @param [in]  corpid      调用企业的企业id，例如：wwd08c8exxxx5ab44d，可以在企业微信管理端--我的企业--企业信息查看
// @param [in]  secret      聊天内容存档的Secret，可以在企业微信管理端--管理工具--聊天内容存档查看
func NewClient(corpId string, corpSecret string) (*Client, error) {
	ptr := C.NewSdk()
	corpIdC := C.CString(corpId)
	corpSecretC := C.CString(corpSecret)
	defer func() {
		C.free(unsafe.Pointer(corpIdC))
		C.free(unsafe.Pointer(corpSecretC))
	}()
	retC := C.Init(ptr, corpIdC, corpSecretC)
	ret := int(retC)
	if ret != 0 {
		return nil, NewSDKErr(ret, SDKInitErrMsg)
	}
	return &Client{
		ptr: ptr,
	}, nil
}

// Destroy  释放sdk，和NewClient成对使用
func (c *Client) Destroy() {
	C.DestroySdk(c.ptr)
}

// GetChatData 拉取聊天记录函数
//
// @param [in]  seq             从指定的seq开始拉取消息，注意的是返回的消息从seq+1开始返回，seq为之前接口返回的最大seq值。首次使用请使用seq:0
// @param [in]  limit           一次拉取的消息条数，最大值1000条，超过1000条会返回错误
// @param [in]  proxy           使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
// @param [in]  passwd          代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
// @param [in]  timeout         超时时间，单位秒
// @return chatDatas       返回本次拉取消息的数据，slice结构体.内容包括errcode/errmsg，以及每条消息内容。示例如下：
//
// {"errcode":0,"errmsg":"ok","chatdata":[{"seq":196,"msgid":"CAQQ2fbb4QUY0On2rYSAgAMgip/yzgs=","publickey_ver":3,"encrypt_random_key":"ftJ+uz3n/z1DsxlkwxNgE+mL38H42/KCvN8T60gbbtPD+Rta1hKTuQPzUzO6Hzne97MgKs7FfdDxDck/v8cDT6gUVjA2tZ/M7euSD0L66opJ/IUeBtpAtvgVSD5qhlaQjvfKJc/zPMGNK2xCLFYqwmQBZXbNT7uA69Fflm512nZKW/piK2RKdYJhRyvQnA1ISxK097sp9WlEgDg250fM5tgwMjujdzr7ehK6gtVBUFldNSJS7ndtIf6aSBfaLktZgwHZ57ONewWq8GJe7WwQf1hwcDbCh7YMG8nsweEwhDfUz+u8rz9an+0lgrYMZFRHnmzjgmLwrR7B/32Qxqd79A==","encrypt_chat_msg":"898WSfGMnIeytTsea7Rc0WsOocs0bIAerF6de0v2cFwqo9uOxrW9wYe5rCjCHHH5bDrNvLxBE/xOoFfcwOTYX0HQxTJaH0ES9OHDZ61p8gcbfGdJKnq2UU4tAEgGb8H+Q9n8syRXIjaI3KuVCqGIi4QGHFmxWenPFfjF/vRuPd0EpzUNwmqfUxLBWLpGhv+dLnqiEOBW41Zdc0OO0St6E+JeIeHlRZAR+E13Isv9eS09xNbF0qQXWIyNUi+ucLr5VuZnPGXBrSfvwX8f0QebTwpy1tT2zvQiMM2MBugKH6NuMzzuvEsXeD+6+3VRqL"}]}
func (c *Client) GetChatData(seq uint64, limit uint64, proxy string, passwd string, timeout int) ([]ChatData, error) {
	proxyC := C.CString(proxy)
	passwdC := C.CString(passwd)
	chatSlice := C.NewSlice()
	defer func() {
		C.free(unsafe.Pointer(proxyC))
		C.free(unsafe.Pointer(passwdC))
		C.FreeSlice(chatSlice)
	}()

	retC := C.GetChatData(c.ptr, C.ulonglong(seq), C.uint(limit), proxyC, passwdC, C.int(timeout), chatSlice)
	if ret := int(retC); ret != 0 {
		return nil, NewSDKErr(ret, GetChatDataErrMsg)
	}
	buf := BufferPool()
	defer BufferPoolRelease(buf)
	c.GetContentFromSlice(chatSlice, buf)
	var data ChatDataResponse
	if err := json.NewDecoder(buf).Decode(&data); err != nil {
		return nil, errors.Join(err, NewSDKErr(-1, GetChatDataErrMsg))
	}
	if data.IsError() {
		return nil, data.Error
	}
	return data.ChatDataList, nil
}

// DecryptData 解析密文.企业微信自有解密内容
// @param [in]  encrypt_key, getchatdata返回的encrypt_random_key,使用企业自持对应版本秘钥RSA解密后的内容
// @param [in]  encrypt_msg, getchatdata返回的encrypt_chat_msg
// @return msg, 解密的消息明文
func (c *Client) DecryptData(encryptKey string, encryptMsg string) (Message, error) {
	encryptKeyC := C.CString(encryptKey)
	encryptMsgC := C.CString(encryptMsg)
	msgSlice := C.NewSlice()
	defer func() {
		C.free(unsafe.Pointer(encryptKeyC))
		C.free(unsafe.Pointer(encryptMsgC))
		C.FreeSlice(msgSlice)
	}()

	retC := C.DecryptData(encryptKeyC, encryptMsgC, msgSlice)
	if ret := int(retC); ret != 0 {
		return nil, NewSDKErr(ret, DecryptErrMsg)
	}
	buf := BufferPool()
	defer BufferPoolRelease(buf)
	c.GetContentFromSlice(msgSlice, buf)
	bs := buf.Bytes()
	// handle illegal escape character in text
	for i := 0; i < len(bs); {
		if bs[i] < 0x20 {
			bs = append(bs[:i], bs[i+1:]...)
			continue
		}
		i++
	}
	r := ReaderPool(bs)
	defer ReaderPoolRelease(r)
	var baseMessage BaseMessage
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&baseMessage); err != nil {
		return nil, errors.Join(err, NewSDKErr(-1, DecryptErrMsg))
	}
	var msg Message
	if baseMessage.ActionType() == SWITCH_ACTION {
		msg = SwitchMessage{}
	} else {
		switch baseMessage.MessageType() {
		case TEXT_MSG:
			msg = TextMessage{}
		case IMG_MSG:
			msg = ImageMessage{}
		case REVOKE_MSG:
			msg = RevokeMessage{}
		case AGREE_MSG, DISAGREE_MSG:
			msg = AggreeMessage{}
		case VOICE_MSG:
			msg = VoiceMessage{}
		case VIDEO_MSG:
			msg = VideoMessage{}
		case CARD_MSG:
			msg = CardMessage{}
		case LOC_MSG:
			msg = LocationMessage{}
		case EMOTION_MSG:
			msg = EmotionMessage{}
		case FILE_MSG:
			msg = FileMessage{}
		case LINK_MSG:
			msg = LinkMessage{}
		case WEAPP_MSG:
			msg = WeappMessage{}
		case CHATRECORD_MSG:
			msg = ChatRecordMessage{}
		case TODO_MSG:
			msg = TodoMessage{}
		case VOTE_MSG:
			msg = VoteMessage{}
		case COLLECT_MSG:
			msg = CollectMessage{}
		case REDPACKET_MSG, EXTERNAL_REDPACKET_MSG:
			msg = RedpacketMessage{}
		case MEETING_MSG:
			msg = MeetingMessage{}
		case DOC_MSG:
			msg = DocMessage{}
		case MARKDOWN_MSG:
			msg = MarkdownMessage{}
		case NEWS_MSG:
			msg = NewsMessage{}
		case CALENDAR_MSG:
			msg = CalendarMessage{}
		case MIXED_MSG:
			msg = MixedMessage{}
		case MEETING_VOICE_CALL_MSG:
			msg = MeetingVoiceCallMessage{}
		case VOIP_DOC_SHARE_MSG:
			msg = VoipDocShareMessage{}
		case SPHFEED_MSG:
			msg = SphfeedMessage{}
		}
	}
	r.Seek(0, 0)
	if err := decoder.Decode(&msg); err != nil {
		return nil, errors.Join(err, NewSDKErr(-1, DecryptErrMsg))
	}
	return msg, nil
}

// GetMediaData 拉取媒体消息函数
// Return值=0表示该API调用成功
//
// @param [in]  sdkFileid       从GetChatData返回的聊天消息中，媒体消息包括的sdkfileid
// @param [in]  proxy           使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
// @param [in]  passwd          代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
// @param [in]  indexbuf        媒体消息分片拉取，需要填入每次拉取的索引信息。首次不需要填写，默认拉取512k，后续每次调用只需要将上次调用返回的outindexbuf填入即可。
// @param [in]  timeout         超时时间，单位秒
// @return media_data      返回本次拉取的媒体数据.MediaData结构体.内容包括data(数据内容)/outindexbuf(下次索引)/is_finish(拉取完成标记)
func (c *Client) GetMediaData(indexBuf string, sdkFileId string, proxy string, passwd string, timeout int, data *MediaData) error {
	indexBufC := C.CString(indexBuf)
	sdkFileIdC := C.CString(sdkFileId)
	proxyC := C.CString(proxy)
	passwdC := C.CString(passwd)
	mediaDataC := C.NewMediaData()
	defer func() {
		C.free(unsafe.Pointer(indexBufC))
		C.free(unsafe.Pointer(sdkFileIdC))
		C.free(unsafe.Pointer(proxyC))
		C.free(unsafe.Pointer(passwdC))
		C.FreeMediaData(mediaDataC)
	}()

	retC := C.GetMediaData(c.ptr, indexBufC, sdkFileIdC, proxyC, passwdC, C.int(timeout), mediaDataC)
	if ret := int(retC); ret != 0 {
		return NewSDKErr(ret, GetMediaDataErrMsg)
	}
	data.OutIndexBuf = C.GoString(C.GetOutIndexBuf(mediaDataC))
	data.Data = C.GoBytes(unsafe.Pointer(C.GetData(mediaDataC)), C.GetDataLen(mediaDataC))
	data.IsFinish = int(C.IsMediaDataFinish(mediaDataC)) == 1
	return nil
}

// DownloadMedia 下载MediaData
func (c *Client) DownloadMedia(w io.Writer, sdkFileId string, proxy string, passwd string, timeout int) error {
	var indexBuf string
	mediaData := MediaDataPool()
	defer MediaDataPoolRelease(mediaData)
	for {
		if err := c.GetMediaData(indexBuf, sdkFileId, proxy, passwd, timeout, mediaData); err != nil {
			return err
		}
		w.Write(mediaData.Data)
		if mediaData.IsFinish {
			break
		}
		indexBuf = mediaData.OutIndexBuf
	}
	return nil
}

// DownloadMediaZeroAlloc 下载MediaData zero mem alloc version
func (c *Client) DownloadMediaZeroAlloc(w io.Writer, sdkFileId string, proxy string, passwd string, timeout int) error {
	var indexBuf string
	mediaData := MediaDataPool()
	defer MediaDataPoolRelease(mediaData)
	indexBufC := C.CString(indexBuf)
	sdkFileIdC := C.CString(sdkFileId)
	proxyC := C.CString(proxy)
	passwdC := C.CString(passwd)
	mediaDataC := C.NewMediaData()
	defer func() {
		C.free(unsafe.Pointer(indexBufC))
		C.free(unsafe.Pointer(sdkFileIdC))
		C.free(unsafe.Pointer(proxyC))
		C.free(unsafe.Pointer(passwdC))
		C.FreeMediaData(mediaDataC)
	}()
	for {
		retC := C.GetMediaData(c.ptr, indexBufC, sdkFileIdC, proxyC, passwdC, C.int(timeout), mediaDataC)
		if ret := int(retC); ret != 0 {
			return NewSDKErr(ret, GetMediaDataErrMsg)
		}
		indexBufC = C.GetOutIndexBuf(mediaDataC)
		l := C.GetDataLen(mediaDataC)
		bs := (*[1 << 30]byte)(unsafe.Pointer(C.GetData(mediaDataC)))[:l:l]
		// bs := C.GoBytes(unsafe.Pointer(C.GetData(mediaDataC)), C.GetDataLen(mediaDataC))
		w.Write(bs)
		if int(C.IsMediaDataFinish(mediaDataC)) == 1 {
			break
		}
	}
	return nil
}

// GetContentFromSlice 转换C.struct_Slice_t为go bytes
func (c Client) GetContentFromSlice(slice *C.struct_Slice_t, buf *bytes.Buffer) {
	l := C.GetSliceLen(slice)
	bs := (*[1 << 30]byte)(unsafe.Pointer(C.GetContentFromSlice(slice)))[:l:l]
	// buf.Write(C.GoBytes(unsafe.Pointer(C.GetContentFromSlice(slice)), C.GetSliceLen(slice)))
	buf.Write(bs)
}
