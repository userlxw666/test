package pack

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Ipack interface {
	Pack(msg IMessage) ([]byte, error)
	UnPack(datapack []byte) (IMessage, error)
	GetHeadLen() uint32
}

type Pack struct {
}

func NewPack() *Pack {
	return &Pack{}
}

func (p *Pack) GetHeadLen() uint32 {
	return 8
}

func (p *Pack) Pack(msg IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetId()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

func (p *Pack) UnPack(datapack []byte) (IMessage, error) {
	dataBuff := bytes.NewBuffer(datapack)
	msg := &Message{}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	if msg.DataLen > 1024 {
		fmt.Println("Too Big ")
		return nil, nil
	}
	return msg, nil
}
