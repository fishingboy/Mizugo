package entrys //nolint:dupl

import (
	"fmt"
	"sync/atomic"

	"github.com/yinweli/Mizugo/mizugos"
	"github.com/yinweli/Mizugo/mizugos/entitys"
	"github.com/yinweli/Mizugo/mizugos/events"
	"github.com/yinweli/Mizugo/mizugos/nets"
	"github.com/yinweli/Mizugo/mizugos/procs"
	"github.com/yinweli/Mizugo/support/test-server/internal/defines"
	"github.com/yinweli/Mizugo/support/test-server/internal/modules"
)

// NewJson 建立Json入口
func NewJson() *Json {
	return &Json{
		name: "json",
	}
}

// Json Json入口
type Json struct {
	name     string        // 入口名稱
	config   JsonConfig    // 配置資料
	count    atomic.Int64  // 計數器
	listenID nets.ListenID // 接聽編號
}

// JsonConfig 配置資料
type JsonConfig struct {
	IP   string `yaml:"ip"`   // 位址
	Port string `yaml:"port"` // 埠號
}

// Initialize 初始化處理
func (this *Json) Initialize() error {
	mizugos.Info(this.name).Caller(0).Message("entry initialize").End()

	if err := mizugos.Configmgr().Unmarshal(this.name, &this.config); err != nil {
		return fmt.Errorf("%v initialize: %w", this.name, err)
	} // if

	this.listenID = mizugos.Netmgr().AddListenTCP(this.config.IP, this.config.Port, this.bind, this.unbind, this.wrong)
	mizugos.Info(this.name).Caller(0).Message("entry start").KV("config", this.config).End()
	return nil
}

// Finalize 結束處理
func (this *Json) Finalize() {
	mizugos.Netmgr().DelListen(this.listenID)
	mizugos.Info(this.name).Caller(0).Message("entry finalize").End()
}

// bind 綁定處理
func (this *Json) bind(session nets.Sessioner) *nets.Bundle {
	mizugos.Info(this.name).Caller(0).Message("bind").End()
	entity := mizugos.Entitymgr().Add()

	var wrong error

	if entity == nil {
		wrong = fmt.Errorf("bind: entity nil")
		goto Error
	} // if

	if err := entity.SetModulemgr(entitys.NewModulemgr()); err != nil {
		wrong = fmt.Errorf("bind: %w", err)
		goto Error
	} // if

	if err := entity.SetEventmgr(events.NewEventmgr(defines.EventCapacity)); err != nil {
		wrong = fmt.Errorf("bind: %w", err)
		goto Error
	} // if

	if err := entity.SetProcess(procs.NewJson()); err != nil {
		wrong = fmt.Errorf("bind: %w", err)
		goto Error
	} // if

	if err := entity.SetSession(session); err != nil {
		wrong = fmt.Errorf("bind: %w", err)
		goto Error
	} // if

	if err := entity.AddModule(modules.NewJson(this.incr)); err != nil {
		wrong = fmt.Errorf("bind: %w", err)
		goto Error
	} // if

	if err := entity.Initialize(this.wrong); err != nil {
		wrong = fmt.Errorf("bind: %w", err)
		goto Error
	} // if

	mizugos.Labelmgr().Add(entity, "json")
	session.SetOwner(entity)
	return entity.Bundle()

Error:
	if entity != nil {
		entity.Finalize()
		mizugos.Entitymgr().Del(entity.EntityID())
		mizugos.Labelmgr().Erase(entity)
	} // if

	session.Stop()
	mizugos.Error(this.name).Caller(0).EndError(wrong)
	return nil
}

// unbind 解綁處理
func (this *Json) unbind(session nets.Sessioner) {
	if entity, ok := session.GetOwner().(*entitys.Entity); ok {
		entity.Finalize()
		mizugos.Entitymgr().Del(entity.EntityID())
		mizugos.Labelmgr().Erase(entity)
	} // if
}

// wrong 錯誤處理
func (this *Json) wrong(fail bool, err error) {
	if fail {
		mizugos.Error(this.name).Caller(1).EndError(err)
	} else {
		mizugos.Warn(this.name).Caller(1).EndError(err)
	} // if
}

// incr 增加計數
func (this *Json) incr() int64 {
	return this.count.Add(1)
}
