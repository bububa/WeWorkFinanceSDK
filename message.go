package wxworkfinancesdk

// ActionType  // 消息动作
type ActionType = string

const (
	// SEND_ACTION 发送消息
	SEND_ACTION ActionType = "send"
	// RECALL_ACTION 撤回消息
	RECALL_ACTION ActionType = "recall"
	// SWITCH_ACTION 切换企业日志
	SWITCH_ACTION ActionType = "switch"
)

// MessageType 消息类型
type MessageType = string

const (
	// UNKNOWN_MSG 未定义
	UNKNOWN_MSG MessageType = ""
	// TEXT_MSG 文本消息
	TEXT_MSG MessageType = "text"
	// IMG_MSG 图片消息
	IMG_MSG MessageType = "image"
	// REVOKE_MSG 撤回消息
	REVOKE_MSG MessageType = "revoke"
	// AGREE_MSG 同意消息
	AGREE_MSG MessageType = "agree"
	// DISAGREE_MSG 不同意消息
	DISAGREE_MSG MessageType = "disagree"
	// VOICE_MSG 语音消息
	VOICE_MSG MessageType = "voice"
	// VIDEO_MSG 视频消息
	VIDEO_MSG MessageType = "video"
	// CARD_MSG 名片消息
	CARD_MSG MessageType = "card"
	// LOC_MSG 位置消息
	LOC_MSG MessageType = "location"
	// EMOTION_MSG 表情消息
	EMOTION_MSG MessageType = "emotion"
	// FILE_MSG 文件消息
	FILE_MSG MessageType = "file"
	// LINK_MSG 链接消息
	LINK_MSG MessageType = "link"
	// WEAPP_MSG 小程序消息
	WEAPP_MSG MessageType = "weapp"
	// CHATRECORD_MSG 话记录消息
	CHATRECORD_MSG MessageType = "chatrecord"
	// TODO_MSG 待办消息
	TODO_MSG MessageType = "todo"
	// VOTE_MSG 投票消息
	VOTE_MSG MessageType = "vote"
	// COLLECT_MSG 填表消息
	COLLECT_MSG MessageType = "collect"
	// REDPACKET_MSG 红包消息
	REDPACKET_MSG MessageType = "redpacket"
	// MEETING_MSG 会议邀请消息
	MEETING_MSG MessageType = "meeting"
	// DOC_MSG 在线文档消息
	DOC_MSG MessageType = "docmsg"
	// MARKDOWN_MSG MarkDown格式消息
	MARKDOWN_MSG MessageType = "markdown"
	// NEWS_MSG 图文消息
	NEWS_MSG MessageType = "news"
	// CALENDAR_MSG 日程消息
	CALENDAR_MSG MessageType = "calendar"
	// MIXED_MSG 混合消息
	MIXED_MSG MessageType = "mixed"
	// MEETING_VOICE_CALL_MSG 音频存档消息
	MEETING_VOICE_CALL_MSG MessageType = "meeting_voice_call"
	// VOIP_DOC_SHARE_MSG 音频共享文档
	VOIP_DOC_SHARE_MSG MessageType = "voip_doc_share"
	// EXTERNAL_REDPACKET_MSG 红包消息
	EXTERNAL_REDPACKET_MSG MessageType = "external_redpacket"
	// SPHFEED_MSG 视频号消息
	SPHFEED_MSG MessageType = "sphfeed"
)

// EmotionType 表情类型
type EmotionType = uint32

const (
	// GIF gif表情
	GIF EmotionType = 1
	// PNG png表情
	PNG EmotionType = 2
)

// ChatRecordType 每条聊天记录的具体消息类型
type ChatRecordType = string

const (
	// TEXT_REC_MSG 文本消息
	TEXT_REC_MSG ChatRecordType = "ChatRecordText"
	// IMG_REC_MSG 图片消息
	IMG_REC_MSG ChatRecordType = "ChatRecordImage"
	// VOICE_REC_MSG 语音消息
	VOICE_REC_MSG ChatRecordType = "ChatRecordVoice"
	// VIDEO_REC_MSG 视频消息
	VIDEO_REC_MSG ChatRecordType = "ChatRecordVideo"
	// CARD_REC_MSG 名片消息
	CARD_REC_MSG ChatRecordType = "ChatRecordCard"
	// LOC_REC_MSG 位置消息
	LOC_REC_MSG ChatRecordType = "ChatRecordLocation"
	// EMOTION_REC_MSG 表情消息
	EMOTION_REC_MSG ChatRecordType = "ChatRecordEmotion"
	// FILE_REC_MSG 文件消息
	FILE_REC_MSG ChatRecordType = "ChatRecordFile"
	// LINK_REC_MSG 链接消息
	LINK_REC_MSG ChatRecordType = "ChatRecordLink"
	// WEAPP_REC_MSG 小程序消息
	WEAPP_REC_MSG ChatRecordType = "ChatRecordWeapp"
	// MIXED_REC_MSG 混合消息
	MIXED_REC_MSG ChatRecordType = "ChatRecordMixed"
)

// VoteType 投票类型
type VoteType = uint32

const (
	// LAUNCH_VOTE 发起投票
	LAUNCH_VOTE VoteType = 101
	// JOIN_VOTE 参与投票
	JOIN_VOTE VoteType = 102
)

// CollectDetailsType 表项类型
type CollectDetailsType = string

const (
	// TEXT 文本
	TEXT CollectDetailsType = "Text"
	// NUMBER 数字
	NUMBER CollectDetailsType = "Number"
	// DATE 日期
	DATE CollectDetailsType = "Date"
	// TIME 时间
	TIME CollectDetailsType = "Time"
)

// RedpacketType 红包类型
type RedpacketType = uint32

const (
	// NORMAL_REDPACKET 普通红包
	NORMAL_REDPACKET RedpacketType = 1
	// SPELL_REDPACKET 拼手气群红包
	SPELL_REDPACKET RedpacketType = 2
	// INCENTIVE_REDPACKET 励群红包
	INCENTIVE_REDPACKET RedpacketType = 3
)

// MeetingType 会议消息类型
type MeetingType = uint32

const (
	// LAUNCH_MEETING 发起会议邀请消息
	LAUNCH_MEETING MeetingType = 101
	// HANDLE_MEETING 处理会议邀请消息
	HANDLE_MEETING MeetingType = 102
)

// MeetingStatus 会议邀请处理状态
type MeetingStatus = uint32

const (
	// JOIN_MEETING 参加会议
	JOIN_MEETING MeetingStatus = 1
	// REJECT_MEETING 绝会议
	REJECT_MEETING MeetingStatus = 2
	// PENDING_MEETING 待定
	PENDING_MEETING MeetingStatus = 3
	// NOINVITE_MEETING 未被邀请
	NOINVITE_MEETING MeetingStatus = 4
	// CANCELED_MEETING 会议已取消
	CANCELED_MEETING MeetingStatus = 5
	// EXPIRED_MEETING 会议已过期
	EXPIRED_MEETING MeetingStatus = 6
	// ABSENT_MEETING 不在房间内
	ABSENT_MEETING MeetingStatus = 7
)

// SphfeedType 视频号消息类型
type SphfeedType = uint32

const (
	// SPHFEED_IMAGE 图片
	SPHFEED_IMAGE SphfeedType = 2
	// SPHFEED_VIDEO 视频
	SPHFEED_VIDEO SphfeedType = 4
	// SPHFEED_LIVE 直播
	SPHFEED_LIVE SphfeedType = 9
)

// Message 会话消息
type Message interface {
	// ID 消息ID
	ID() string
	// MessageType 消息类型
	MessageType() MessageType
	// ActionType 消息动作类型
	ActionType() ActionType
}

// BaseMessage 会话消息公共字段
type BaseMessage struct {
	// MsgId 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。
	MsgId string `json:"msgid,omitempty"`
	// Action 消息动作，目前有send(发送消息)/recall(撤回消息)/switch(切换企业日志)三种类型。
	Action ActionType `json:"action,omitempty"`
	// From 消息发送方id。同一企业内容为userid，非相同企业为external_userid。消息如果是机器人发出，也为external_userid。
	From string `json:"from,omitempty"`
	// ToList 消息接收方列表，可能是多个，同一个企业内容为userid，非相同企业为external_userid。
	ToList []string `json:"tolist,omitempty"`
	// RoomId 群聊消息的群id。如果是单聊则为空。
	RoomId string `json:"roomid,omitempty"`
	// MsgTime 消息发送时间戳，utc时间，ms单位。
	MsgTime int64 `json:"msgtime,omitempty"`
	// MsgType 文本消息为：text。
	MsgType MessageType `json:"msgtype,omitempty"`
}

// ID implement Message interface
func (m BaseMessage) ID() string {
	return m.MsgId
}

// MessageType implement Message interface
func (m BaseMessage) MessageType() MessageType {
	return m.MsgType
}

// ActionType Message interface
func (m BaseMessage) ActionType() ActionType {
	return m.Action
}

// TextMessage 文字消息
type TextMessage struct {
	BaseMessage
	// Content 消息内容。
	Content string `json:"content,omitempty"`
}

// ImageMessage 图片消息
type ImageMessage struct {
	BaseMessage
	// SdkFileId 媒体资源的id信息。
	SdkFileId string `json:"sdkfileid,omitempty"`
	// Md5Sum 图片资源的md5值，供进行校验。
	Md5Sum string `json:"md5sum,omitempty"`
	// FileSize 图片资源的文件大小。
	FileSize uint32 `json:"filesize,omitempty"`
}

// RevokeMessage 撤回消息
type RevokeMessage struct {
	BaseMessage
	// PreMsgId 标识撤回的原消息的msgid
	PreMsgId string `json:"pre_msgid,omitempty"`
}

// AggreeMessage 同意会话消息
type AggreeMessage struct {
	BaseMessage
	// UserId 同意/不同意协议者的userid，外部企业默认为external_userid。
	UserId string `json:"userid,omitempty"`
	// AgreeTime 同意/不同意协议的时间，utc时间，ms单位。
	AgreeTime int64 `json:"agree_time,omitempty"`
}

// VoiceMessage 语音消息
type VoiceMessage struct {
	BaseMessage
	// SdkFileId 媒体资源的id信息。
	SdkFileId string `json:"sdkfileid,omitempty"`
	// VoiceSize 语音消息大小。
	VoiceSize uint32 `json:"voice_size,omitempty"`
	// PlayLength 播放长度。
	PlayLength uint32 `json:"play_length,omitempty"`
	// Md5Sum 图片资源的md5值，供进行校验。
	Md5Sum string `json:"md5sum,omitempty"`
}

// VideoMessage 视频消息
type VideoMessage struct {
	BaseMessage
	// SdkFileId 媒体资源的id信息。
	SdkFileId string `json:"sdkfileid,omitempty"`
	// FileSize 图片资源的文件大小。
	FileSize uint32 `json:"filesize,omitempty"`
	// PlayLength 播放长度。
	PlayLength uint32 `json:"play_length,omitempty"`
	// Md5Sum 图片资源的md5值，供进行校验。
	Md5Sum string `json:"md5sum,omitempty"`
}

// CardMessage 名片消息
type CardMessage struct {
	BaseMessage
	// CorpName 名片所有者所在的公司名称。
	CorpName string `json:"corpname,omitempty"`
	// UserId 名片所有者的id，同一公司是userid，不同公司是external_userid
	UserId string `json:"userid,omitempty"`
}

// LocationMessage 位置消息
type LocationMessage struct {
	BaseMessage
	// Lng 经度，单位double
	Lng float64 `json:"longitude,omitempty"`
	// Lat 纬度，单位double
	Lat float64 `json:"latitude,omitempty"`
	// Address 地址信息
	Address string `json:"address,omitempty"`
	// Title 位置信息的title。
	Title string `json:"title,omitempty"`
	// Zoom 缩放比例。
	Zoom uint32 `json:"zoom,omitempty"`
}

// EmotionMessage 表情消息
type EmotionMessage struct {
	BaseMessage
	// Type 表情类型，png或者gif.1表示gif 2表示png。
	Type EmotionType `json:"type,omitempty"`
	// Width 表情图片宽度。
	Width uint32 `json:"width,omitempty"`
	// Height 表情图片高度。
	Height uint32 `json:"height,omitempty"`
	// ImageSize 资源的文件大小。
	ImageSize uint32 `json:"imagesize,omitempty"`
	// SdkFileId 媒体资源的id信息。
	SdkFileId string `json:"sdkfileid,omitempty"`
	// Md5Sum 图片资源的md5值，供进行校验。
	Md5Sum string `json:"md5sum,omitempty"`
}

// FileMessage 文件消息
type FileMessage struct {
	BaseMessage
	// FileName 文件名称。
	FileName string `json:"filename,omitempty"`
	// FileExt 文件类型后缀。
	FileExt string `json:"fileext,omitempty"`
	// SdkFileId 媒体资源的id信息。
	SdkFileId string `json:"sdkfileid,omitempty"`
	// FileSize 文件大小。
	FileSize uint32 `json:"filesize,omitempty"`
	// Md5Sum 资源的md5值，供进行校验。
	Md5Sum string `json:"md5sum,omitempty"`
}

// LinkMessage 链接消息
type LinkMessage struct {
	BaseMessage
	// Title 消息标题。
	Title string `json:"title,omitempty"`
	// Desc 消息描述。
	Desc string `json:"description,omitempty"`
	// LinkUrl 链接url地址
	LinkUrl string `json:"link_url,omitempty"`
	// ImageUrl 链接图片url。
	ImageUrl string `json:"image_url,omitempty"`
}

// WeappMessage 小程序消息
type WeappMessage struct {
	BaseMessage
	// Title  消息标题。
	Title string `json:"title,omitempty"`
	// Desc 消息描述。
	Desc string `json:"description,omitempty"`
	// Username 用户名称。
	Username string `json:"username,omitempty"`
	// DisplayName 小程序名称
	DisplayName string `json:"displayname,omitempty"`
}

// ChatRecordMessage 聊天记录消息
type ChatRecordMessage struct {
	BaseMessage
	// Title 聊天记录标题
	Title string `json:"title,omitempty"`
	// Item 消息记录内的消息内容，批量数据
	Item []ChatRecord `json:"item,omitempty"`
}

// TodoMessage 待办事项消息
type TodoMessage struct {
	BaseMessage
	// Title 待办的来源文本
	Title string `json:"title,omitempty"`
	// Content 待办的具体内容
	Content string `json:"content,omitempty"`
}

// VoteMessage 投票消息
type VoteMessage struct {
	BaseMessage
	// VoteTitle 投票主题。
	VoteTitle string `json:"votetitle,omitempty"`
	// VoteItem 投票选项，可能多个内容。
	VoteItem []string `json:"voteitem,omitempty"`
	// VoteType 投票类型.101发起投票、102参与投票。
	VoteType VoteType `json:"votetype,omitempty"`
	// VoteId 投票id，方便将参与投票消息与发起投票消息进行前后对照。
	VoteId string `json:"voteid,omitempty"`
}

// CollectMessage 填表消息
type CollectMessage struct {
	BaseMessage
	// RoomName 填表消息所在的群名称。
	RoomName string `json:"room_name,omitempty"`
	// Creator 创建者在群中的名字
	Creator string `json:"creator,omitempty"`
	// CreateTime 创建的时间
	CreateTime string `json:"create_time,omitempty"`
	// Details 表内容
	Details []CollectDetails `json:"details,omitempty"`
}

// RedpacketMessage 红包消息
type RedpacketMessage struct {
	BaseMessage
	// Type 红包消息类型。1 普通红包、2 拼手气群红包、3 激励群红包。
	Type RedpacketType `json:"type,omitempty"`
	// Wish 红包祝福语
	Wish string `json:"wish,omitempty"`
	// TotalCnt 红包总个数
	TotalCnt int `json:"totalcnt,omitempty"`
	// TotalAmount 红包总金额。单位为分。
	TotalAmount int64 `json:"totalamount,omitempty"`
}

// MeetingMessage 会议邀请消息
type MeetingMessage struct {
	BaseMessage
	// Topic 会议主题
	Topic string `json:"topic,omitempty"`
	// StartTime 会议开始时间。Utc时间
	StartTime int64 `json:"starttime,omitempty"`
	// EndTime 会议结束时间。Utc时间
	EndTime int64 `json:"endtime,omitempty"`
	// Address 会议地址
	Address string `json:"address,omitempty"`
	// Remarks 会议备注
	Remarks string `json:"remarks,omitempty"`
	// MeetingType 会议消息类型。101发起会议邀请消息、102处理会议邀请消息
	MeetingType MeetingType `json:"meetingtype,omitempty"`
	// MeetingId 会议id。方便将发起、处理消息进行对照
	MeetingId uint64 `json:"meetingid,omitempty"`
	// Status 会议邀请处理状态。1 参加会议、2 拒绝会议、3 待定、4 未被邀请、5 会议已取消、6 会议已过期、7 不在房间内。
	Status MeetingStatus `json:"status,omitempty"`
}

// DocMessage 在线文档消息
type DocMessage struct {
	BaseMessage
	// Title 在线文档名称
	Title string `json:"title,omitempty"`
	// LinkUrl 在线文档链接
	LinkUrl string `json:"link_url,omitempty"`
	// DocCreator 在线文档创建者。本企业成员创建为userid；外部企业成员创建为external_userid
	DocCreator string `json:"doc_creator,omitempty"`
}

// MarkdownMessage MarkDown格式消息
type MarkdownMessage struct {
	BaseMessage
	// Content markdown消息内容，目前为机器人发出的消息
	Content string `json:"content,omitempty"`
}

// NewsMessage 图文消息
type NewsMessage struct {
	BaseMessage
	// Info 图文消息的内容
	Info struct {
		// Item 图文消息数组
		Item []News `json:"item,omitempty"`
	} `json:"info,omitempty"`
}

// CalendarMessage 日程消息
type CalendarMessage struct {
	BaseMessage
	// Title 日程主题
	Title string `json:"title,omitempty"`
	// CreatorName 日程组织者
	CreatorName string `json:"creatorname,omitempty"`
	// AttendeeName 日程参与人。数组，内容为String类型
	AttendeeName []string `json:"attendeename,omitempty"`
	// StartTime 日程开始时间。Utc时间，单位秒
	StartTime int64 `json:"starttime,omitempty"`
	// EndTime 日程结束时间。Utc时间，单位秒
	EndTime int64 `json:"endtime,omitempty"`
	// Place 日程地点
	Place string `json:"place,omitempty"`
	// Remarks 日程备注
	Remarks string `json:"remarks,omitempty"`
}

// MixedMessage 混合消息
type MixedMessage struct {
	BaseMessage
	// Mixed 消息内容。可包含图片、文字、表情等多种消息。Object类型
	Mixed struct {
		Item []MixedMsg `json:"item,omitempty"`
	} `json:"mixed,omitempty"`
}

// MeetingVoiceCallMessage 音频存档消息
type MeetingVoiceCallMessage struct {
	BaseMessage
	// VoiceId 音频id
	VoiceId string `json:"voiceid,omitempty"`
	// MeetingVoiceCall 音频消息内容。包括结束时间、fileid，可能包括多个demofiledata、sharescreendata消息，demofiledata表示文档共享信息，sharescreendata表示屏幕共享信息。Object类型
	MeetingVoiceCall *MeetingVoiceCall `json:"meeting_voice_call,omitempty"`
}

// VoipDocShareMessage 音频共享文档消息
type VoipDocShareMessage struct {
	BaseMessage
	// VoipId 音频id
	VoipId string `json:"voipid,omitempty"`
	// VoipDocShare 共享文档消息内容。包括filename、md5sum、filesize、sdkfileid字段。Object类型
	VoipDocShare *VoipDocShare `json:"voip_doc_share,omitempty"`
}

// SphfeedMessage 视频号消息
type SphfeedMessage struct {
	BaseMessage
	// FeedType 视频号消息类型
	FeedType SphfeedType `json:"feed_type,omitempty"`
	// SphName 视频号账号名称
	SphName string `json:"sph_name,omitempty"`
	// FeedDesc 视频号消息描述
	FeedDesc string `json:"feed_desc,omitempty"`
}

// SwitchMessage 切换企业日志
type SwitchMessage struct {
	// MsgId 消息id，消息的唯一标识，企业可以使用此字段进行消息去重
	MsgId string `json:"msgid,omitempty"`
	// Action 消息动作，切换企业为switch
	Action string `json:"action,omitempty"`
	// Time 消息发送时间戳，utc时间，ms单位。
	Time int64 `json:"time,omitempty"`
	// User 具体为切换企业的成员的userid。
	User string `json:"user,omitempty"`
}

// ID implement Message interface
func (m SwitchMessage) ID() string {
	return m.MsgId
}

// MessageType implement Message interface
func (m SwitchMessage) MessageType() MessageType {
	return UNKNOWN_MSG
}

// ActionType implement ActionType interface
func (m SwitchMessage) ActionType() ActionType {
	return SWITCH_ACTION
}

// ChatRecord 消息记录内的消息内
type ChatRecord struct {
	// Type  每条聊天记录的具体消息类型：ChatRecordText/ ChatRecordFile/ ChatRecordImage/ ChatRecordVideo/ ChatRecordLink/ ChatRecordLocation/ ChatRecordMixed ….
	Type ChatRecordType `json:"type,omitempty"`
	// Content 消息内容。Json串，内容为对应类型的json
	Content string `json:"content,omitempty"`
	// MsgTime 消息时间，utc时间，ms单位。
	MsgTime int64 `json:"msgtime,omitempty"`
	// FromChatroom 是否来自群会话。
	FromChatroom bool `json:"from_chatroom,omitempty"`
}

// CollectDetails 填表消息表内容
type CollectDetails struct {
	// Id 表项id
	Id uint64 `json:"id,omitempty"`
	// Ques 表项名称
	Ques string `json:"ques,omitempty"`
	// Type 表项类型，有Text(文本),Number(数字),Date(日期),Time(时间)
	Type CollectDetailsType `json:"type,omitempty"`
}

// News 图文消息图文内容
type News struct {
	// Title 图文消息标题
	Title string `json:"title,omitempty"`
	// Desc 图文消息描述
	Desc string `json:"description,omitempty"`
	// Url 图文消息点击跳转地址
	Url string `json:"url,omitempty"`
	// PicUrl 图文消息配图的url
	PicUrl string `json:"picurl,omitempty"`
}

// MixedMsg 混合消息内容
type MixedMsg struct {
	Type MessageType `json:"type,omitempty"`
	// Content 消息内容。Json串，内容为对应类型的json
	Content string `json:"content,omitempty"`
}

// MeetingVoiceCall 音频消息内
type MeetingVoiceCall struct {
	// EndTime 音频结束时间
	EndTime int64 `json:"endtime,omitempty"`
	// SdkFileId 音频媒体下载的id
	SdkFileId string `json:"sdkfileid,omitempty"`
	// DemoFileData 文档分享对象，Object类型
	DemoFileData []DemoFileData `json:"demofiledata,omitempty"`
	// ShareScreenData 屏幕共享对象，Object类型
	ShareScreenData []ShareScreenData `json:"sharescreendata,omitempty"`
}

// DemoFileData 文档分享对
type DemoFileData struct {
	// FileName 文档共享名称
	FileName string `json:"filename,omitempty"`
	// DemoOperator 文档共享操作用户的id
	DemoOperator string `json:"demooperator,omitempty"`
	// StartTime 文档共享开始时间
	StartTime int64 `json:"starttime,omitempty"`
	// EndTime 文档共享结束时间
	EndTime int64 `json:"endtime,omitempty"`
}

// ShareScreenData 屏幕共享对象
type ShareScreenData struct {
	// Share 屏幕共享用户的id
	Share string `json:"share,omitempty"`
	// ShareTime 屏幕共享开始时间
	StartTime int64 `json:"starttime,omitempty"`
	// EndTime 屏幕共享结束时间
	EndTime int64 `json:"endtime,omitempty"`
}

// VoipDocShare 共享文档消息内容
type VoipDocShare struct {
	// FileName 文档共享文件名称
	FileName string `json:"filename,omitempty"`
	// Md5Sum 共享文件的md5值
	Md5Sum string `json:"md5sum,omitempty"`
	// FileSize 共享文件的大小
	FileSize int64 `json:"filesize,omitempty"`
	// SdkFileId 共享文件的sdkfile，通过此字段进行媒体数据下载
	SdkFileId string `json:"sdkfileid,omitempty"`
}
