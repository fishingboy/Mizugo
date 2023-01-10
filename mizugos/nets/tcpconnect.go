package nets

import (
	"fmt"
	"net"
	"time"
)

// TODO: 研究一下怎麼改用context

// NewTCPConnect 建立tcp連接器
func NewTCPConnect(ip, port string, timeout time.Duration) *TCPConnect {
	return &TCPConnect{
		address: net.JoinHostPort(ip, port),
		timeout: timeout,
	}
}

// TCPConnect tcp連接器
type TCPConnect struct {
	address string        // 位址字串
	timeout time.Duration // 逾時時間
}

// Connect 啟動連接
func (this *TCPConnect) Connect(bind Bind, unbind Unbind, wrong Wrong) {
	go func() {
		conn, err := net.DialTimeout("tcp", this.address, this.timeout)

		if err != nil {
			wrong.Do(fmt.Errorf("tcp connect: %v: %w", this.address, err))
			return
		} // if

		session := NewTCPSession(conn)
		session.Start(bind, unbind, wrong)
	}()
}

// Address 取得位址
func (this *TCPConnect) Address() string {
	return this.address
}
