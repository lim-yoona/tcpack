package safetcpack

import (
	"bytes"
	"encoding/binary"
	"github.com/lim-yoona/tcpack"
	"io"
	"log"
	"net"
	"sync"
)

// connMap is a map ensures that only one SafeMsgPack instance for one conn.
var connMap map[net.Conn]*SafeMsgPack

// mu is a mutex lock for Concurrently reading and writing connMap.
var mu sync.Mutex

// SafeMsgPack implements the interface IMsgPack,
// carrying HeadLen and conn for Pack() and Unpack(),
// and mutex for concurrent Pack() and Unpack().
type SafeMsgPack struct {
	headLen uint32
	conn    net.Conn
	mutex   sync.Mutex
}

// NewSafeMsgPack returns a thread-safe packager *SafeMsgPack.
// NewSafeMsgPack returns the same packager for the same TCP connection,
// so the value of the headLen is consistent with the first time you new a SafeMsgPack.
func NewSafeMsgPack(headleng uint32, conn net.Conn) *SafeMsgPack {
	defer mu.Unlock()
	mu.Lock()
	smp, ok := connMap[conn]
	if !ok {
		smp = &SafeMsgPack{
			headLen: headleng,
			conn:    conn,
		}
		connMap[conn] = smp
	}
	return smp
}

// GetHeadLen return headLen of the message.
func (smp *SafeMsgPack) GetHeadLen() uint32 {
	return smp.headLen
}

// Pack packs a message to bytes stream and sends it.
func (smp *SafeMsgPack) Pack(msg tcpack.Imessage) (uint32, error) {
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
	smp.mutex.Lock()
	num, err := smp.conn.Write(buffer.Bytes())
	smp.mutex.Unlock()
	if err != nil {
		return 0, err
	}
	return uint32(num), nil
}

// Unpack unpacks a certain length bytes stream to a message.
func (smp *SafeMsgPack) Unpack() (tcpack.Imessage, error) {
	defer smp.mutex.Unlock()
	smp.mutex.Lock()
	headDate := make([]byte, smp.GetHeadLen())
	_, err := io.ReadFull(smp.conn, headDate)
	if err != nil {
		log.Println("read headData failed:", err)
		return nil, err
	}
	buffer := bytes.NewReader(headDate)
	msg := tcpack.NewMessage(0, 0, nil)
	if err := binary.Read(buffer, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	if msg.GetDataLen() > 0 {
		msg.Data = make([]byte, msg.GetDataLen())
		_, err := io.ReadFull(smp.conn, msg.Data)
		if err != nil {
			log.Println("read msgData failed:", err)
			return nil, err
		}
	}
	return msg, nil
}

func init() {
	connMap = make(map[net.Conn]*SafeMsgPack, 5)
}
