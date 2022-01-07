package message_body

type MsgBody struct {
	ObjectType string        `json:"docType"`   // 消息类型
	MsgId      string        `json:"msg_id"`    // 信息 id
	Sender     string        `json:"sender"`    // 发送者
	Recipient  string        `json:"recipient"` // 接受者
	Body       string        `json:"body"`      // 消息体
	DataHash   string        `json:"hash"`      // 整体数据 hash
	TimeDate   string        `json:"time_date"` // 日期
	History    []HistoryItem // 当前 msg 的历史记录
}

type HistoryItem struct {
	TxId    string // 交易 id
	MsgBody MsgBody
}
