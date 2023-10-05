package msgpack

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"msgpack/message"
	"net"
)

type IMsgPack interface {
	GetHeadLen() uint32
	Pack(message.Imessage) ([]byte, error)
	Unpack() (message.Imessage, error)
}

type MsgPack struct {
	HeadLen uint32
	conn    net.Conn
}

func NewMsgPack(headlen uint32, conn net.Conn) *MsgPack {
	return &MsgPack{
		HeadLen: headlen,
		conn:    conn,
	}
}

func (mp *MsgPack) GetHeadLen() uint32 {
	return mp.HeadLen
}
func (mp *MsgPack) Pack(msg message.Imessage) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
func (mp *MsgPack) Unpack() (message.Imessage, error) {
	headDate := make([]byte, mp.GetHeadLen())
	_, err := io.ReadFull(mp.conn, headDate)
	if err != nil {
		log.Println("read headData failed:", err)
		return nil, err
	}
	buffer := bytes.NewReader(headDate)
	msg := message.NewMessage(0, 0, nil)
	if err := binary.Read(buffer, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	fmt.Println("msg.GetDataLen():", msg.GetDataLen())
	if msg.GetDataLen() > 0 {
		msg.Data = make([]byte, msg.GetDataLen())
		_, err := io.ReadFull(mp.conn, msg.Data)
		if err != nil {
			log.Println("read msgData failed:", err)
			return nil, err
		}
	}
	return msg, nil
}
