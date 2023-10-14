package tcpack

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"net"
)

// IMsgPack is an interface that defined a packager,
// which carries HeadLen, and provides Pack() and Unpack() method.
type IMsgPack interface {
	// Get the head length of the message package.
	GetHeadLen() uint32
	// Pack returns bytes stream packed from a message.
	Pack(Imessage) (uint32, error)
	// Unpack returns a message from bytes stream.
	Unpack() (Imessage, error)
}

// MsgPack implements the interface IMsgPack,
// carrying HeadLen and conn for Pack() and Unpack().
type MsgPack struct {
	headLen uint32
	conn    net.Conn
}

// NewMsgPack returns a packager *MsgPack.
func NewMsgPack(headlen uint32, conn net.Conn) *MsgPack {
	return &MsgPack{
		headLen: headlen,
		conn:    conn,
	}
}

// GetHeadLen return HeadLen of the message.
func (mp *MsgPack) GetHeadLen() uint32 {
	return mp.headLen
}

// Pack packs a message to bytes stream and sends it.
func (mp *MsgPack) Pack(msg Imessage) (uint32, error) {
	buffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return 0, err
	}
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return 0, err
	}
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return 0, err
	}
	num, err := mp.conn.Write(buffer.Bytes())
	if err != nil {
		return 0, err
	}
	return uint32(num), nil
}

// Unpack unpacks a certain length bytes stream to a message.
func (mp *MsgPack) Unpack() (Imessage, error) {
	headDate := make([]byte, mp.GetHeadLen())
	_, err := io.ReadFull(mp.conn, headDate)
	if err != nil {
		log.Println("read headData failed:", err)
		return nil, err
	}
	buffer := bytes.NewReader(headDate)
	msg := NewMessage(0, 0, nil)
	if err := binary.Read(buffer, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
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
