package logprint

import (
	"fmt"
	"time"
)

func DebugInfo(mes interface{}) {
	now := time.Now()
	fmt.Printf("[debug]%s:%s\n", now.Format("2006-01-02 15:04:05"), mes)
}
