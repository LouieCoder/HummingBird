package service

import "sync"


type GID struct {
	ID int64
	mutex sync.Mutex
}

var gid GID

//加锁的全局id生成器
func GenerateID() int64  {
	gid.mutex.Lock()
	gid.ID++
	gid.mutex.Unlock()
	return gid.ID
}