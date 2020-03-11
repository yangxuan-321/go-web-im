package util

import (
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var groupIdCrease int32 = 0
var mutex sync.Mutex

func IntId() int64 {
	nowStr := time.Now().Format("19991212010101")
	mutex.Lock()
	int32 := atomic.AddInt32(&groupIdCrease, 1)
	nowInt64, _ := strconv.ParseInt(nowStr+string(int32), 10, 64)
	mutex.Unlock()

	return nowInt64
}
