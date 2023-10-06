package msgpack

type Imessage interface {
	GetDataLen() uint32
	GetMsgId() uint32
	GetMsgData() []byte
}

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func NewMessage(id, datalen uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: datalen,
		Data:    data,
	}
}

func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}
func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}
func (msg *Message) GetMsgData() []byte {
	return msg.Data
}
