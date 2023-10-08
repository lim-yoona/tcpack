package msgpack

// Imessage is an interface that definds a message,
// which carries DataLen, MsgId and MsgData.
type Imessage interface {
	GetDataLen() uint32
	GetMsgId() uint32
	GetMsgData() []byte
}

// Message implements Imessage.
type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

// NewMessage returns a message typed *Message.
func NewMessage(id, datalen uint32, data []byte) *Message {
	if data == nil {
		data = []uint8{}
	}
	return &Message{
		Id:      id,
		DataLen: datalen,
		Data:    data,
	}
}

// GetDataLen returns DataLen typed uint32.
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// GetMsgId returns Id typed uint32.
func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

// GetMsgData returns Data typed []byte.
func (msg *Message) GetMsgData() []byte {
	return msg.Data
}
