package models

import (
	"time"

	"github.com/yinweli/Mizugo/mizugos/depots"
)

// Auth 認證資料
type Auth struct {
	Account string    // [主索引] 帳號
	Token   string    // token
	Time    time.Time // 更新時間
}

// NewAuthGet 建立取得認證資料行為
func NewAuthGet(key string, data *Auth) *depots.Getter[Auth] {
	return &depots.Getter[Auth]{
		Key:  key,
		Data: data,
	}
}

// NewAuthSet 建立設定認證資料行為
func NewAuthSet(key string, data *Auth) *depots.Setter[Auth] {
	return &depots.Setter[Auth]{
		Field: "Account",
		Key:   key,
		Data:  data,
	}
}
