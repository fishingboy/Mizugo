package redmos

import (
	"time"
)

const ( // redis定義
	Timeout  = time.Second * 30 // redis超時時間
	RedisNil = ""               // redis回應空字串, 通常在GET命令找不到索引時, 會以此字串回報給使用者
	RedisOk  = "OK"             // redis回應完成, 通常在SET命令順利完成後, 會以此字串回報給使用者
)

// Metaer 元資料介面, 提供主要/次要資料庫操作時所需的必要資訊
type Metaer interface {
	// MajorKey 取得主要資料庫索引值
	// 當使用到主要資料庫, 必須提供索引值
	MajorKey(key any) string

	// MinorTable 取得次要資料庫表格名稱
	// 當使用到次要資料庫, 必須提供表格名稱
	MinorTable() string

	// MinorField 取得次要資料庫欄位名稱
	// 當使用到次要資料庫時, 必須提供欄位名稱
	MinorField() string

	// MinorKey 取得次要資料庫索引值
	// 當使用到次要資料庫時, 必須提供索引值
	MinorKey(key any) string
}
