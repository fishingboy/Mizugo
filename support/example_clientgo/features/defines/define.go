package defines

import (
	"github.com/yinweli/Mizugo/mizugos/msgs"
)

const ConfigPath = "config" // 配置路徑
const ConfigType = "yaml"   // 配置類型

const ( // 入口名稱
	EntryEchoOnce  = "echoonce"
	EntryEchoMulti = "echomulti"
)

const ( // 標籤名稱
	LabelEchoOnce  = "echoonce"
	LabelEchoMulti = "echomulti"
)

const ( // 訊息編號
	MessageIDEcho = msgs.MessageID(1)
)
