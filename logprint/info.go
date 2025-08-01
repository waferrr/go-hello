package logprint

import (
	"fmt"
	"time"
)

func Info(msg interface{}) {
	now := time.Now()
	fmt.Printf("[info]%s:%s\n", now.Format("2006-01-02 15:04:05"), msg)
}
