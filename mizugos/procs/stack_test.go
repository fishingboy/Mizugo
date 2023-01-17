package procs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"

	"github.com/yinweli/Mizugo/mizugos/msgs"
	"github.com/yinweli/Mizugo/mizugos/utils"
	"github.com/yinweli/Mizugo/testdata"
)

func TestStack(t *testing.T) {
	suite.Run(t, new(SuiteStack))
}

type SuiteStack struct {
	suite.Suite
	testdata.TestEnv
	testdata.TestLeak
	key       []byte
	messageID MessageID
	message   *msgs.StackTest
}

func (this *SuiteStack) SetupSuite() {
	this.Change("test-procs-stack")
	this.key = utils.RandDesKey()
	this.messageID = MessageID(1)
	this.message = &msgs.StackTest{
		Data: "stack test",
	}
}

func (this *SuiteStack) TearDownSuite() {
	this.Restore()
}

func (this *SuiteStack) TearDownTest() {
	this.GoLeak(this.T(), true)
}

func (this *SuiteStack) TestNewStack() {
	assert.NotNil(this.T(), NewStack())
}

func (this *SuiteStack) TestEncode() {
	target := NewStack().Key(this.key)
	input := msgs.MarshalStackMsg([]msgs.TestMsg{
		{MessageID: this.messageID, Message: this.message},
	})

	encode, err := target.Encode(input)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), encode)

	_, err = target.Encode(nil)
	assert.NotNil(this.T(), err)

	_, err = target.Encode("!?")
	assert.NotNil(this.T(), err)

	decode, err := target.Decode(encode)
	assert.Nil(this.T(), err)
	assert.NotNil(this.T(), decode)
	assert.True(this.T(), proto.Equal(input, decode.(*msgs.StackMsg)))

	_, err = target.Decode(nil)
	assert.NotNil(this.T(), err)

	_, err = target.Decode([]byte("unknown encode"))
	assert.NotNil(this.T(), err)
}

func (this *SuiteStack) TestProcess() {
	target := NewStack().Key(this.key)
	input := msgs.MarshalStackMsg([]msgs.TestMsg{
		{MessageID: this.messageID, Message: this.message},
	})

	validSend := false
	target.Send(func(message any) {
		_, validSend = message.(*msgs.StackMsg)
	})
	validProcess := false
	target.Add(this.messageID, func(context any) {
		if stackContext, ok := context.(*StackContext); ok {
			if messageID, unit1, err := stackContext.Unmarshal(); err == nil {
				if unit2, ok := unit1.(*msgs.StackTest); ok {
					validProcess = this.messageID == messageID && proto.Equal(this.message, unit2)
				} // if
			} // if
		} // if
	})
	assert.Nil(this.T(), target.Process(input))
	assert.True(this.T(), validSend)
	assert.True(this.T(), validProcess)

	input = msgs.MarshalStackMsg([]msgs.TestMsg{
		{MessageID: 0, Message: this.message},
	})
	assert.NotNil(this.T(), target.Process(input))

	assert.NotNil(this.T(), target.Process(nil))
}

func (this *SuiteStack) TestStackContext() {
	target := &StackContext{}
	target.initialize(msgs.MarshalStackMsg([]msgs.TestMsg{
		{MessageID: 1, Message: this.message},
		{MessageID: 2, Message: this.message},
		{MessageID: 3, Message: this.message},
	}))

	assert.Equal(this.T(), MessageID(0), target.messageID())
	_, _, err := target.Unmarshal()
	assert.NotNil(this.T(), err)

	assert.True(this.T(), target.next())
	assert.Equal(this.T(), MessageID(1), target.messageID())
	messageID, output, err := target.Unmarshal()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), MessageID(1), messageID)
	assert.True(this.T(), proto.Equal(this.message, output))

	assert.True(this.T(), target.next())
	assert.Equal(this.T(), MessageID(2), target.messageID())
	messageID, output, err = target.Unmarshal()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), MessageID(2), messageID)
	assert.True(this.T(), proto.Equal(this.message, output))

	assert.True(this.T(), target.next())
	assert.Equal(this.T(), MessageID(3), target.messageID())
	messageID, output, err = target.Unmarshal()
	assert.Nil(this.T(), err)
	assert.Equal(this.T(), MessageID(3), messageID)
	assert.True(this.T(), proto.Equal(this.message, output))

	assert.False(this.T(), target.next())

	target = &StackContext{}
	assert.Nil(this.T(), target.AddRespond(1, this.message))
	assert.Nil(this.T(), target.AddRespond(2, this.message))
	assert.Nil(this.T(), target.AddRespond(3, this.message))
	result := target.result()
	assert.NotNil(this.T(), result)
	assert.True(this.T(), proto.Equal(msgs.MarshalStackMsg([]msgs.TestMsg{
		{MessageID: 1, Message: this.message},
		{MessageID: 2, Message: this.message},
		{MessageID: 3, Message: this.message},
	}), result))
}

func BenchmarkStackEncode(b *testing.B) {
	target := NewStack()
	input := msgs.MarshalStackMsg([]msgs.TestMsg{
		{MessageID: 1, Message: &msgs.StackTest{Data: "benchmark encode"}},
	})

	for i := 0; i < b.N; i++ {
		_, _ = target.Encode(input)
	} // for
}

func BenchmarkStackDecode(b *testing.B) {
	target := NewStack()
	input := msgs.MarshalStackMsg([]msgs.TestMsg{
		{MessageID: 1, Message: &msgs.StackTest{Data: "benchmark decode"}},
	})
	encode, _ := target.Encode(input)

	for i := 0; i < b.N; i++ {
		_, _ = target.Decode(encode)
	} // for
}
