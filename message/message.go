package message

import (
	"encoding/binary"
	"io"
)

// https://wiki.theory.org/BitTorrentSpecification#Messages
type messageID uint8

const (
	MsgChoke         messageID = 0
	MsgUnchoke       messageID = 1
	MsgInterested    messageID = 2
	MsgNotInterested messageID = 3
	MsgHave          messageID = 4
	MsgBitfield      messageID = 5
	MsgRequest       messageID = 6
	MsgPiece         messageID = 7
	MsgCancel        messageID = 8
)

type Message struct {
	ID      messageID
	Payload []byte
}

// <length prefix><message ID><payload>
// the `length prefix` is a four byte big-endian value. The `message ID` is a single decimal byte. The `payload` is message dependent.
func (m *Message) Serialize() []byte {

	if m == nil {
		return make([]byte, 4) // Return a 4-byte slice if nil
	}

	length := uint32(len(m.Payload) + 1)         // Calculate the total length (1 byte for ID + length of Payload)
	buf := make([]byte, 4+length)                // Create a buffer of the appropriate size
	binary.BigEndian.PutUint32(buf[0:4], length) // Write the length in big-endian format to the first 4 bytes
	buf[4] = byte(m.ID)                          // Set the ID at the 5th byte
	copy(buf[5:], m.Payload)                     // Copy the payload into the buffer starting from the 6th byte
	return buf                                   // Return the serialized byte slice
}

// <length prefix><message ID><payload>
func Read(r io.Reader) (*Message, error) {
	lengthBuf := make([]byte, 4)
	_, err := io.ReadFull(r, lengthBuf)
	if err != nil {
		return nil, err
	}
	length := binary.BigEndian.Uint32(lengthBuf)

	if length == 0 {
		return nil, nil
	}

	messageBuf := make([]byte, length)
	_, err = io.ReadFull(r, messageBuf)
	if err != nil {
		return nil, err
	}

	m := Message{
		ID:      messageID(messageBuf[0]),
		Payload: messageBuf[1:],
	}

	return &m, nil
}
