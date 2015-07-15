package sessionstate

import (
	"github.com/astaxie/beego/session"
)

var (
	GlobalSessions *session.Manager
)

func init() {
	GlobalSessions, _ = session.NewManager("memory", `{"cookieName":"gosessionid","gclifetime":3600}`)
	go GlobalSessions.GC()
}
