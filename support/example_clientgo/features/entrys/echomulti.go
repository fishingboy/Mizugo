package entrys

import (
	"fmt"
	"time"

	"github.com/yinweli/Mizugo/mizugos"
	"github.com/yinweli/Mizugo/mizugos/msgs"
	"github.com/yinweli/Mizugo/mizugos/nets"
	"github.com/yinweli/Mizugo/support/example_clientgo/features/defines"
	"github.com/yinweli/Mizugo/support/example_clientgo/features/modules"
)

// NewEchoMulti 建立多次回音資料
func NewEchoMulti() *EchoMulti {
	return &EchoMulti{
		name: defines.EntryEchoMulti,
	}
}

// EchoMulti 多次回音資料
type EchoMulti struct {
	name   string          // 入口名稱
	config EchoMultiConfig // 設定資料
}

// EchoMultiConfig 設定資料
type EchoMultiConfig struct {
	IP         string        // 位址
	Port       string        // 埠號
	Timeout    time.Duration // 逾期時間(秒)
	EchoString string        // 回音字串
	EchoCount  int           // 回音次數
}

// Initialize 初始化處理
func (this *EchoMulti) Initialize() error {
	mizugos.Info(this.name).Message("entry initialize").End()

	if err := mizugos.Configmgr().ReadFile(this.name, defines.ConfigType); err != nil {
		return fmt.Errorf("%v initialize: %w", this.name, err)
	} // if

	if err := mizugos.Configmgr().Unmarshal(this.name, &this.config); err != nil {
		return fmt.Errorf("%v initialize: %w", this.name, err)
	} // if

	mizugos.Netmgr().AddConnect(nets.NewTCPConnect(this.config.IP, this.config.Port, this.config.Timeout), this)
	mizugos.Info(this.name).Message("entry start").KV("config", this.config).End()
	return nil
}

// Finalize 結束處理
func (this *EchoMulti) Finalize() {
	mizugos.Info(this.name).Message("entry finalize").End()
}

// Bind 綁定處理
func (this *EchoMulti) Bind(session nets.Sessioner) (content nets.Content, err error) {
	mizugos.Info(this.name).Message("session").KV("sessionID", session.SessionID()).End()
	entity := mizugos.Entitymgr().Add()

	if entity == nil {
		return content, fmt.Errorf("bind: entity nil")
	} // if

	if err := entity.SetSession(session); err != nil {
		return content, fmt.Errorf("bind: %w", err)
	} // if

	if err := entity.SetProcess(msgs.NewStringProc()); err != nil {
		return content, fmt.Errorf("bind: %w", err)
	} // if

	if err := entity.AddModule(modules.NewEchoMulti(this.config.EchoString, this.config.EchoCount)); err != nil {
		return content, fmt.Errorf("bind: %w", err)
	} // if

	if err := entity.Initialize(func() {
		mizugos.Entitymgr().Del(entity.EntityID())
		mizugos.Labelmgr().Erase(entity)
	}); err != nil {
		return content, fmt.Errorf("bind: %w", err)
	} // if

	mizugos.Labelmgr().Add(entity, defines.LabelEchoMulti)
	content.Unbind = entity.Finalize
	content.Encode = entity.GetProcess().Encode
	content.Decode = entity.GetProcess().Decode
	content.Receive = entity.GetProcess().Process
	return content, nil
}

// Error 錯誤處理
func (this *EchoMulti) Error(err error) {
	_ = mizugos.Error(this.name).EndError(err)
}
