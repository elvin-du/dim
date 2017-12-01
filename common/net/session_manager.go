package libnet

import (
	"sync"
	"sync/atomic"
)

// session manager
type SessionManager struct {
	maxSessionId uint64
	sessions     map[uint64]*Session
	sessionMutex sync.Mutex
}

// create a SessionManager Instance
func NewSessionManager() *SessionManager {
	return &SessionManager{
		maxSessionId: 1,
		sessions:     make(map[uint64]*Session),
	}
}

// 生成一个SessionId
func (sm *SessionManager) GenSessionId() uint64 {
	return atomic.AddUint64(&sm.maxSessionId, 1)
}

// 将Session加入到SessionManager
func (sm *SessionManager) SetSession(session *Session) {
	sm.sessionMutex.Lock()
	defer sm.sessionMutex.Unlock()
	sm.sessions[session.id] = session
}

// 把Session从SessionManager中删除
func (sm *SessionManager) DelSession(session *Session) {
	sm.sessionMutex.Lock()
	defer sm.sessionMutex.Unlock()
	delete(sm.sessions, session.id)
}

// 返回SessionManager中当前Session的数量
func (sm *SessionManager) Count() int {
	sm.sessionMutex.Lock()
	defer sm.sessionMutex.Unlock()
	return len(sm.sessions)
}

// 遍历SessionManager当前所有的Session
//
// 注意:
// 这个操作会锁住SessionManager直到遍历结束, 所以不适合费时的操作
func (sm *SessionManager) Foreach(callback func(*Session)) {
	sm.sessionMutex.Lock()
	defer sm.sessionMutex.Unlock()

	for _, session := range sm.sessions {
		callback(session)
	}
}

// 关闭SessionManager中所有的Session
//
// 注意:
// 通过遍历调用Session.Close(err error)的方式来关闭
func (sm *SessionManager) CloseSessions() {
	sm.Foreach(func(session *Session) { session.Close(nil) })
}
