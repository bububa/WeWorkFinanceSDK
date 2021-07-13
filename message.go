package WeWorkFinanceSDK

// ActionType  // 消息动作
type ActionType = string

const (
	SEND_ACTION   ActionType = "send"   // 发送消息
	RECALL_ACTION ActionType = "recall" // 撤回消息
	SWITCH_ACTION ActionType = "switch" // 切换企业日志
)

// MessageType 消息类型
type MessageType = string

const (
	UNKNOWN_MSG            MessageType = ""                   // 未定义
	TEXT_MSG               MessageType = "text"               // 	文本消息
	IMG_MSG                MessageType = "image"              // 	图片消息
	REVOKE_MSG             MessageType = "revoke"             // 撤回消息
	AGREE_MSG              MessageType = "agree"              // 同意消息
	DISAGREE_MSG           MessageType = "disagree"           // 不同意消息
	VOICE_MSG              MessageType = "voice"              // 语音消息
	VIDEO_MSG              MessageType = "video"              // 视频消息
	CARD_MSG               MessageType = "card"               // 名片消息
	LOC_MSG                MessageType = "location"           // 位置消息
	EMOTION_MSG            MessageType = "emotion"            // 表情消息
	FILE_MSG               MessageType = "file"               // 文件消息
	LINK_MSG               MessageType = "link"               // 链接消息
	WEAPP_MSG              MessageType = "weapp"              // 小程序消息
	CHATRECORD_MSG         MessageType = "chatrecord"         // 会话记录消息
	TODO_MSG               MessageType = "todo"               // 待办消息
	VOTE_MSG               MessageType = "vote"               // 投票消息
	COLLECT_MSG            MessageType = "collect"            // 填表消息
	REDPACKET_MSG          MessageType = "redpacket"          // 红包消息
	MEETING_MSG            MessageType = "meeting"            // 会议邀请消息
	DOC_MSG                MessageType = "docmsg"             // 在线文档消息
	MARKDOWN_MSG           MessageType = "markdown"           // MarkDown格式消息
	NEWS_MSG               MessageType = "news"               // 图文消息
	CALENDAR_MSG           MessageType = "calendar"           // 日程消息
	MIXED_MSG              MessageType = "mixed"              // 混合消息
	MEETING_VOICE_CALL_MSG MessageType = "meeting_voice_call" // 音频存档消息
	VOIP_DOC_SHARE_MSG     MessageType = "voip_doc_share"     // 音频共享文档
	EXTERNAL_REDPACKET_MSG MessageType = "external_redpacket" // 互通红包消息
	SPHFEED_MSG            MessageType = "sphfeed"            // 视频号消息
)

// EmotionType 表情类型
type EmotionType = uint32

const (
	GIF EmotionType = 1 // gif表情
	PNG EmotionType = 2 // png表情
)

// ChatRecordType 每条聊天记录的具体消息类型
type ChatRecordType = string

const (
	TEXT_REC_MSG    ChatRecordType = "ChatRecordText"     // 文本消息
	IMG_REC_MSG     ChatRecordType = "ChatRecordImage"    // 图片消息
	VOICE_REC_MSG   ChatRecordType = "ChatRecordVoice"    // 语音消息
	VIDEO_REC_MSG   ChatRecordType = "ChatRecordVideo"    // 视频消息
	CARD_REC_MSG    ChatRecordType = "ChatRecordCard"     // 名片消息
	LOC_REC_MSG     ChatRecordType = "ChatRecordLocation" // 位置消息
	EMOTION_REC_MSG ChatRecordType = "ChatRecordEmotion"  // 表情消息
	FILE_REC_MSG    ChatRecordType = "ChatRecordFile"     // 文件消息
	LINK_REC_MSG    ChatRecordType = "ChatRecordLink"     // 链接消息
	WEAPP_REC_MSG   ChatRecordType = "ChatRecordWeapp"    // 小程序消息
	MIXED_REC_MSG   ChatRecordType = "ChatRecordMixed"    // 混合消息
)

// VoteType 投票类型
type VoteType = uint32

const (
	LAUNCH_VOTE VoteType = 101 // 发起投票
	JOIN_VOTE   VoteType = 102 // 参与投票
)

// CollectDetailsType 表项类型
type CollectDetailsType = string

const (
	TEXT   CollectDetailsType = "Text"   // 文本
	NUMBER CollectDetailsType = "Number" // 数字
	DATE   CollectDetailsType = "Date"   // 日期
	TIME   CollectDetailsType = "Time"   // 时间
)

// RedpacketType 红包类型
type RedpacketType = uint32

const (
	NORMAL_REDPACKET    RedpacketType = 1 // 普通红包
	SPELL_REDPACKET     RedpacketType = 2 // 拼手气群红包
	INCENTIVE_REDPACKET RedpacketType = 3 // 激励群红包
)

// MeetingType 会议消息类型
type MeetingType = uint32

const (
	LAUNCH_MEETING MeetingType = 101 // 发起会议邀请消息
	HANDLE_MEETING MeetingType = 102 // 处理会议邀请消息
)

// MeetingStatus 会议邀请处理状态
type MeetingStatus = uint32

const (
	JOIN_MEETING     MeetingStatus = 1 // 参加会议
	REJECT_MEETING   MeetingStatus = 2 // 拒绝会议
	PENDING_MEETING  MeetingStatus = 3 // 待定
	NOINVITE_MEETING MeetingStatus = 4 // 未被邀请
	CANCELED_MEETING MeetingStatus = 5 // 会议已取消
	EXPIRED_MEETING  MeetingStatus = 6 // 会议已过期
	ABSENT_MEETING   MeetingStatus = 7 // 不在房间内
)

// SphfeedType 视频号消息类型
type SphfeedType = uint32

const (
	SPHFEED_IMAGE SphfeedType = 2
	SPHFEED_VIDEO SphfeedType = 4
	SPHFEED_LIVE  SphfeedType = 9
)

// Message 会话消息
type Message interface {
	ID() string                // 消息ID
	GetMsgType() MessageType   // 消息类型
	GetActionType() ActionType // 消息动作类型
}

// BaseMessage 会话消息公共字段
type BaseMessage struct {
	MsgId   string      `json:"msgid,omitempty"`   // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重。
	Action  ActionType  `json:"action,omitempty"`  // 消息动作，目前有send(发送消息)/recall(撤回消息)/switch(切换企业日志)三种类型。
	From    string      `json:"from,omitempty"`    // 消息发送方id。同一企业内容为userid，非相同企业为external_userid。消息如果是机器人发出，也为external_userid。
	ToList  []string    `json:"tolist,omitempty"`  // 消息接收方列表，可能是多个，同一个企业内容为userid，非相同企业为external_userid。
	RoomId  string      `json:"roomid,omitempty"`  // 群聊消息的群id。如果是单聊则为空。
	MsgTime int64       `json:"msgtime,omitempty"` // 消息发送时间戳，utc时间，ms单位。
	MsgType MessageType `json:"msgtype,omitempty"` // 文本消息为：text。
}

// ID implement Message interface
func (this BaseMessage) ID() string {
	return this.MsgId
}

// MessageType implement Message interface
func (this BaseMessage) MessageType() MessageType {
	return this.MsgType
}

// ActionType Message interface
func (this BaseMessage) ActionType() ActionType {
	return this.Action
}

// TextMessage 文字消息
type TextMessage struct {
	BaseMessage
	Content string `json:"content,omitempty"` // 消息内容。
}

// ImageMessage 图片消息
type ImageMessage struct {
	BaseMessage
	SdkFileId string `json:"sdkfileid,omitempty"` // 媒体资源的id信息。
	Md5Sum    string `json:"md5sum,omitempty"`    // 图片资源的md5值，供进行校验。
	FileSize  uint32 `json:"filesize,omitempty"`  // 图片资源的文件大小。
}

// RevokeMessage 撤回消息
type RevokeMessage struct {
	BaseMessage
	PreMsgId string `json:"pre_msgid,omitempty"` // 标识撤回的原消息的msgid
}

// AggreeMessage 同意会话消息
type AggreeMessage struct {
	BaseMessage
	UserId    string `json:"userid,omitempty"`     // 同意/不同意协议者的userid，外部企业默认为external_userid。
	AgreeTime int64  `json:"agree_time,omitempty"` // 同意/不同意协议的时间，utc时间，ms单位。
}

// VoiceMessage 语音消息
type VoiceMessage struct {
	BaseMessage
	SdkFileId  string `json:"sdkfileid,omitempty"`   // 媒体资源的id信息。
	VoiceSize  uint32 `json:"voice_size,omitempty"`  // 语音消息大小。
	PlayLength uint32 `json:"play_length,omitempty"` // 播放长度。
	Md5Sum     string `json:"md5sum,omitempty"`      // 图片资源的md5值，供进行校验。
}

// VideoMessage 视频消息
type VideoMessage struct {
	BaseMessage
	SdkFileId  string `json:"sdkfileid,omitempty"`   // 媒体资源的id信息。
	FileSize   uint32 `json:"filesize,omitempty"`    // 图片资源的文件大小。
	PlayLength uint32 `json:"play_length,omitempty"` // 播放长度。
	Md5Sum     string `json:"md5sum,omitempty"`      // 图片资源的md5值，供进行校验。
}

// CardMessage 名片消息
type CardMessage struct {
	BaseMessage
	CorpName string `json:"corpname,omitempty"` // 名片所有者所在的公司名称。
	UserId   string `json:"userid,omitempty"`   // 名片所有者的id，同一公司是userid，不同公司是external_userid
}

// LocationMessage 位置消息
type LocationMessage struct {
	BaseMessage
	Lng     float64 `json:"longitude,omitempty"` // 经度，单位double
	Lat     float64 `json:"latitude,omitempty"`  // 纬度，单位double
	Address string  `json:"address,omitempty"`   // 地址信息
	Title   string  `json:"title,omitempty"`     // 位置信息的title。
	Zoom    uint32  `json:"zoom,omitempty"`      // 缩放比例。
}

// EmotionMessage 表情消息
type EmotionMessage struct {
	BaseMessage
	Type      EmotionType `json:"type,omitempty"`      // 表情类型，png或者gif.1表示gif 2表示png。
	Width     uint32      `json:"width,omitempty"`     // 表情图片宽度。
	Height    uint32      `json:"height,omitempty"`    // 表情图片高度。
	ImageSize uint32      `json:"imagesize,omitempty"` // 资源的文件大小。
	SdkFileId string      `json:"sdkfileid,omitempty"` // 媒体资源的id信息。
	Md5Sum    string      `json:"md5sum,omitempty"`    // 图片资源的md5值，供进行校验。
}

// FileMessage 文件消息
type FileMessage struct {
	BaseMessage
	FileName  string `json:"filename,omitempty"`  // 文件名称。
	FileExt   string `json:"fileext,omitempty"`   // 文件类型后缀。
	SdkFileId string `json:"sdkfileid,omitempty"` // 媒体资源的id信息。
	FileSize  uint32 `json:"filesize,omitempty"`  // 文件大小。
	Md5Sum    string `json:"md5sum,omitempty"`    // 资源的md5值，供进行校验。
}

// LinkMessage 链接消息
type LinkMessage struct {
	BaseMessage
	Title    string `json:"title,omitempty"`       // 消息标题。
	Desc     string `json:"description,omitempty"` // 消息描述。
	LinkUrl  string `json:"link_url,omitempty"`    // 链接url地址
	ImageUrl string `json:"image_url,omitempty"`   // 链接图片url。
}

// WeappMessage 小程序消息
type WeappMessage struct {
	BaseMessage
	Title       string `json:"title,omitempty"`       // 消息标题。
	Desc        string `json:"description,omitempty"` // 消息描述。
	Username    string `json:"username,omitempty"`    // 用户名称。
	DisplayName string `json:"displayname,omitempty"` // 小程序名称
}

// ChatRecordMessage 聊天记录消息
type ChatRecordMessage struct {
	BaseMessage
	Title string       `json:"title,omitempty"` // 聊天记录标题
	Item  []ChatRecord `json:"item,omitempty"`  // 消息记录内的消息内容，批量数据
}

// TodoMessage 待办事项消息
type TodoMessage struct {
	BaseMessage
	Title   string `json:"title,omitempty"`   // 待办的来源文本
	Content string `json:"content,omitempty"` // 待办的具体内容
}

// VoteMessage 投票消息
type VoteMessage struct {
	BaseMessage
	VoteTitle string   `json:"votetitle,omitempty"` // 投票主题。
	VoteItem  []string `json:"voteitem,omitempty"`  // 投票选项，可能多个内容。
	VoteType  VoteType `json:"votetype,omitempty"`  // 投票类型.101发起投票、102参与投票。
	VoteId    string   `json:"voteid,omitempty"`    // 投票id，方便将参与投票消息与发起投票消息进行前后对照。
}

// CollectMessage 填表消息
type CollectMessage struct {
	BaseMessage
	RoomName   string           `json:"room_name,omitempty"`   // 填表消息所在的群名称。
	Creator    string           `json:"creator,omitempty"`     // 创建者在群中的名字
	CreateTime string           `json:"create_time,omitempty"` // 创建的时间
	Details    []CollectDetails `json:"details,omitempty"`     // 表内容
}

// RedpacketMessage 红包消息
type RedpacketMessage struct {
	BaseMessage
	Type        RedpacketType `json:"type,omitempty"`        // 红包消息类型。1 普通红包、2 拼手气群红包、3 激励群红包。
	Wish        string        `json:"wish,omitempty"`        // 红包祝福语
	TotalCnt    uint32        `json:"totalcnt,omitempty"`    // 红包总个数
	TotalAmount uint32        `json:"totalamount,omitempty"` // 红包总金额。单位为分。
}

// MeetingMessage 会议邀请消息
type MeetingMessage struct {
	BaseMessage
	Topic       string        `json:"topic,omitempty"`       // 会议主题
	StartTime   int64         `json:"starttime,omitempty"`   // 会议开始时间。Utc时间
	EndTime     int64         `json:"endtime,omitempty"`     // 会议结束时间。Utc时间
	Address     string        `json:"address,omitempty"`     // 会议地址
	Remarks     string        `json:"remarks,omitempty"`     // 会议备注
	MeetingType MeetingType   `json:"meetingtype,omitempty"` // 会议消息类型。101发起会议邀请消息、102处理会议邀请消息
	MeetingId   uint64        `json:"meetingid,omitempty"`   // 会议id。方便将发起、处理消息进行对照
	Status      MeetingStatus `json:"status,omitempty"`      // 会议邀请处理状态。1 参加会议、2 拒绝会议、3 待定、4 未被邀请、5 会议已取消、6 会议已过期、7 不在房间内。
}

// DocMessage 在线文档消息
type DocMessage struct {
	BaseMessage
	Title      string `json:"title,omitempty"`       // 在线文档名称
	LinkUrl    string `json:"link_url,omitempty"`    // 在线文档链接
	DocCreator string `json:"doc_creator,omitempty"` // 在线文档创建者。本企业成员创建为userid；外部企业成员创建为external_userid
}

// MarkdownMessage MarkDown格式消息
type MarkdownMessage struct {
	BaseMessage
	Content string `json:"content,omitempty"` // markdown消息内容，目前为机器人发出的消息
}

// NewsMessage 图文消息
type NewsMessage struct {
	BaseMessage
	Info struct {
		Item []News `json:"item,omitempty"` // 图文消息数组
	} `json:"info,omitempty"` // 图文消息的内容
}

// CalendarMessage 日程消息
type CalendarMessage struct {
	BaseMessage
	Title        string   `json:"title,omitempty"`        // 日程主题
	CreatorName  string   `json:"creatorname,omitempty"`  // 日程组织者
	AttendeeName []string `json:"attendeename,omitempty"` // 日程参与人。数组，内容为String类型
	StartTime    int64    `json:"starttime,omitempty"`    // 日程开始时间。Utc时间，单位秒
	EndTime      int64    `json:"endtime,omitempty"`      // 日程结束时间。Utc时间，单位秒
	Place        string   `json:"place,omitempty"`        // 日程地点
	Remarks      string   `json:"remarks,omitempty"`      // 日程备注
}

// MixedMessage 混合消息
type MixedMessage struct {
	BaseMessage
	Mixed struct {
		Item []MixedMsg `json:"item,omitempty"`
	} `json:"mixed,omitempty"` // 消息内容。可包含图片、文字、表情等多种消息。Object类型
}

// MeetingVoiceCallMessage 音频存档消息
type MeetingVoiceCallMessage struct {
	BaseMessage
	VoiceId          string            `json:"voiceid,omitempty"`            // 音频id
	MeetingVoiceCall *MeetingVoiceCall `json:"meeting_voice_call,omitempty"` // 音频消息内容。包括结束时间、fileid，可能包括多个demofiledata、sharescreendata消息，demofiledata表示文档共享信息，sharescreendata表示屏幕共享信息。Object类型
}

// VoipDocShareMessage 音频共享文档消息
type VoipDocShareMessage struct {
	BaseMessage
	VoipId       string        `json:"voipid,omitempty"`         // 音频id
	VoipDocShare *VoipDocShare `json:"voip_doc_share,omitempty"` // 共享文档消息内容。包括filename、md5sum、filesize、sdkfileid字段。Object类型
}

// SphfeedMessage 视频号消息
type SphfeedMessage struct {
	BaseMessage
	FeedType SphfeedType `json:"feed_type,omitempty"` // 视频号消息类型
	SphName  string      `json:"sph_name,omitempty"`  // 视频号账号名称
	FeedDesc string      `json:"feed_desc,omitempty"` // 视频号消息描述
}

// SwitchMessage 切换企业日志
type SwitchMessage struct {
	MsgId  string `json:"msgid,omitempty"`  // 消息id，消息的唯一标识，企业可以使用此字段进行消息去重
	Action string `json:"action,omitempty"` // 消息动作，切换企业为switch
	Time   int64  `json:"time,omitempty"`   // 消息发送时间戳，utc时间，ms单位。
	User   string `json:"user,omitempty"`   // 具体为切换企业的成员的userid。
}

// ID implement Message interface
func (this SwitchMessage) ID() string {
	return this.MsgId
}

// MessageType implement Message interface
func (this SwitchMessage) MessageType() MessageType {
	return UNKNOWN_MSG
}

// ActionType implement ActionType interface
func (this SwitchMessage) ActionType() ActionType {
	return SWITCH_ACTION
}

// ChatRecord 消息记录内的消息内
type ChatRecord struct {
	Type         ChatRecordType `json:"type,omitempty"`          // 每条聊天记录的具体消息类型：ChatRecordText/ ChatRecordFile/ ChatRecordImage/ ChatRecordVideo/ ChatRecordLink/ ChatRecordLocation/ ChatRecordMixed ….
	Content      string         `json:"content,omitempty"`       // 消息内容。Json串，内容为对应类型的json
	MsgTime      int64          `json:"msgtime,omitempty"`       // 消息时间，utc时间，ms单位。
	FromChatroom bool           `json:"from_chatroom,omitempty"` // 是否来自群会话。
}

// CollectDetails 填表消息表内容
type CollectDetails struct {
	Id   uint64             `json:"id,omitempty"`   // 表项id
	Ques string             `json:"ques,omitempty"` // 表项名称
	Type CollectDetailsType `json:"type,omitempty"` // 表项类型，有Text(文本),Number(数字),Date(日期),Time(时间)
}

// News 图文消息图文内容
type News struct {
	Title  string `json:"title,omitempty"`       // 图文消息标题
	Desc   string `json:"description,omitempty"` // 图文消息描述
	Url    string `json:"url,omitempty"`         // 图文消息点击跳转地址
	PicUrl string `json:"picurl,omitempty"`      // 图文消息配图的url
}

// MixedMsg 混合消息内容
type MixedMsg struct {
	Type    MessageType `json:"type,omitempty"`
	Content string      `json:"content,omitempty"`
}

// MeetingVoiceCall 音频消息内
type MeetingVoiceCall struct {
	EndTime         int64             `json:"endtime,omitempty"`         // 音频结束时间
	SdkFileId       string            `json:"sdkfileid,omitempty"`       // 音频媒体下载的id
	DemoFileData    []DemoFileData    `json:"demofiledata,omitempty"`    // 文档分享对象，Object类型
	ShareScreenData []ShareScreenData `json:"sharescreendata,omitempty"` // 屏幕共享对象，Object类型
}

// DemoFileData 文档分享对
type DemoFileData struct {
	FileName     string `json:"filename,omitempty"`     // 文档共享名称
	DemoOperator string `json:"demooperator,omitempty"` // 文档共享操作用户的id
	StartTime    int64  `json:"starttime,omitempty"`    // 文档共享开始时间
	EndTime      int64  `json:"endtime,omitempty"`      // 文档共享结束时间
}

// ShareScreenData 屏幕共享对象
type ShareScreenData struct {
	Share     string `json:"share,omitempty"`     // 屏幕共享用户的id
	StartTime int64  `json:"starttime,omitempty"` // 屏幕共享开始时间
	EndTime   int64  `json:"endtime,omitempty"`   // 屏幕共享结束时间
}

// VoipDocShare 共享文档消息内容
type VoipDocShare struct {
	FileName  string `json:"filename,omitempty"`  // 文档共享文件名称
	Md5Sum    string `json:"md5sum,omitempty"`    // 共享文件的md5值
	FileSize  uint64 `json:"filesize,omitempty"`  // 共享文件的大小
	SdkFileId string `json:"sdkfileid,omitempty"` // 共享文件的sdkfile，通过此字段进行媒体数据下载
}
