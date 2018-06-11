package socket_server

import (
	"sync"
)

type ConnMgr struct {
	connections map[[11]byte]*Connection
	mutex       *sync.RWMutex
}

func NewConnMgr() *ConnMgr {
	return &ConnMgr{
		connections: make(map[[11]byte]*Connection),
		mutex:       new(sync.RWMutex),
	}
}

func (cm *ConnMgr) Put(id [11]byte, c *Connection) {
	cm.mutex.Lock()
	cm.connections[id] = c
	cm.mutex.Unlock()
}

func (cm *ConnMgr) Get(id [11]byte) *Connection {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if c, ok := cm.connections[id]; ok {
		return c
	}

	return nil
}

func (cm *ConnMgr) Del(c *Connection) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if cc, ok := cm.connections[c.ID]; ok && cc.Equal(c) {
		delete(cm.connections, c.ID)
	}
}
