package pack

type IMessage interface {
	GetDataLen() uint32
	GetId() uint32
	GetData() []byte
	SetId(id uint32)
	SetDataLen(data []byte)
	SetData(data []byte)
}

type Message struct {
	ID      uint32
	DataLen uint32
	Data    []byte
}

func NewMessage(ID uint32, data []byte) *Message {
	return &Message{
		ID:      ID,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetId() uint32 {
	return m.ID
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetId(id uint32) {
	m.ID = id
}

func (m *Message) SetDataLen(data []byte) {
	m.DataLen = uint32(len(data))
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
