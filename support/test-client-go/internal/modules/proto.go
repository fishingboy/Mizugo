package modules

import (
	"fmt"
	"time"

	"github.com/yinweli/Mizugo/mizugos/entitys"
	"github.com/yinweli/Mizugo/mizugos/procs"
	"github.com/yinweli/Mizugo/support/test-client-go/internal/defines"
	"github.com/yinweli/Mizugo/support/test-client-go/internal/features"
	"github.com/yinweli/Mizugo/support/test-client-go/msgs"
)

// NewProto 建立Proto模組
func NewProto(delay time.Duration, disconnect bool) *Proto {
	return &Proto{
		Module:     entitys.NewModule(defines.ModuleIDProto),
		name:       "proto",
		delay:      delay,
		disconnect: disconnect,
	}
}

// Proto Proto模組
type Proto struct {
	*entitys.Module               // 模組資料
	name            string        // 系統名稱
	delay           time.Duration // 延遲時間
	disconnect      bool          // 斷線旗標
}

// Awake 喚醒處理
func (this *Proto) Awake() error {
	this.Entity().Subscribe(defines.EventBegin, this.eventBegin)
	this.Entity().AddMessage(procs.MessageID(msgs.MsgID_ProtoA), this.procMProtoA)
	return nil
}

// Start 啟動處理
func (this *Proto) Start() error {
	this.Entity().PublishDelay(defines.EventBegin, nil, this.delay)
	return nil
}

// event 開始事件
func (this *Proto) eventBegin(_ any) {
	this.sendMProtoQ()
}

// procMProtoA 處理回應Proto
func (this *Proto) procMProtoA(message any) {
	_, msg, err := procs.ProtoUnmarshal[msgs.MProtoA](message)

	if err != nil {
		features.LogSystem.Warn(this.name).Caller(0).EndError(fmt.Errorf("proto procMProtoA: %w", err))
		return
	} // if

	if msg.ErrID != msgs.ErrID_Success {
		features.LogSystem.Warn(this.name).Caller(0).EndError(fmt.Errorf("proto procMProtoA: %v", msg.ErrID))
		return
	} // if

	duration := time.Duration(time.Now().UnixNano() - msg.From.Time)
	features.MeterProto.Add(duration)

	if this.disconnect {
		this.Entity().GetSession().Stop()
	} else {
		this.sendMProtoQ()
	} // if

	features.LogSystem.Info(this.name).Caller(0).KV("duration", duration).KV("count", msg.Count).End()
}

// sendMProtoQ 傳送要求Proto
func (this *Proto) sendMProtoQ() {
	msg, err := procs.ProtoMarshal(procs.MessageID(msgs.MsgID_ProtoQ), &msgs.MProtoQ{
		Time: time.Now().UnixNano(),
	})

	if err != nil {
		features.LogSystem.Warn(this.name).Caller(0).EndError(err)
		return
	} // if

	this.Entity().Send(msg)
}
