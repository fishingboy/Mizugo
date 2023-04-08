package features

import (
	"fmt"

	"github.com/yinweli/Mizugo/mizugos"
	"github.com/yinweli/Mizugo/mizugos/redmos"
)

// NewRedmo 建立資料庫資料
func NewRedmo() *Redmo {
	return &Redmo{
		name: "redmo",
	}
}

// Redmo 資料庫資料
type Redmo struct {
	name   string      // 系統名稱
	config RedmoConfig // 配置資料
}

// RedmoConfig 配置資料
type RedmoConfig struct {
	MajorURI    redmos.RedisURI `yaml:"majorURI"`    // 主要資料庫連接字串
	MajorDebug  bool            `yaml:"majorDebug"`  // 主要資料庫除錯旗標
	MinorURI    redmos.MongoURI `yaml:"minorURI"`    // 次要資料庫連接字串
	MinorDBName string          `yaml:"minorDBName"` // 次要資料庫資料庫名稱
}

// Initialize 初始化處理
func (this *Redmo) Initialize() (err error) {
	if err = mizugos.Configmgr().Unmarshal(this.name, &this.config); err != nil {
		return fmt.Errorf("%v initialize: %w", this.name, err)
	} // if

	if DBMajor, err = mizugos.Redmomgr().AddMajor("dbmajor", this.config.MajorURI, this.config.MajorDebug); err != nil {
		return fmt.Errorf("%v initialize: %w", this.name, err)
	} // if

	if DBMinor, err = mizugos.Redmomgr().AddMinor("dbminor", this.config.MinorURI, this.config.MinorDBName); err != nil {
		return fmt.Errorf("%v initialize: %w", this.name, err)
	} // if

	if DBMixed, err = mizugos.Redmomgr().AddMixed("dbmixed", "dbmajor", "dbminor"); err != nil {
		return fmt.Errorf("%v initialize: %w", this.name, err)
	} // if

	LogSystem.Info(this.name).Caller(0).Message("initialize").KV("config", this.config).End()
	return nil
}

// Finalize 結束處理
func (this *Redmo) Finalize() {
	mizugos.Redmomgr().Finalize()
}

var DBMajor *redmos.Major // 主要資料庫
var DBMinor *redmos.Minor // 次要資料庫
var DBMixed *redmos.Mixed // 混合資料庫
