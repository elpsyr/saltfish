package job

import (
	"fmt"
	"github.com/elpsyr/saltfish/pkg/win"
	"sync"
)

const (
	WindowsTitle        = "咸鱼之王"
	AppID               = "咸鱼助手"
	NotificationTitle   = "请登录微信"
	NotificationMessage = "然后再打开小程序"
)

type Single struct {
	mu sync.Mutex
}

type ManagerCallback func()

type Manager struct {
	singleInstance *Single
	lock           *sync.Mutex
	countReward    int
	countFish      int
	callback       ManagerCallback
}

func NewManager() *Manager {
	return &Manager{
		lock: &sync.Mutex{},
	}
}

// GetInstance 双重校验获取单例
func (m *Manager) GetInstance() *Single {
	if m.singleInstance == nil {
		m.lock.Lock()
		defer m.lock.Unlock()
		if m.singleInstance == nil {
			m.singleInstance = &Single{}
		} else {
		}
	} else {
	}
	return m.singleInstance
}

func (*Manager) GetHwnd() uintptr {

	hwndByTitle := win.GetHwndByTitle(WindowsTitle)
	if hwndByTitle == 0 {
		err := win.SendNotification(win.Message{
			AppID:       AppID,
			Title:       NotificationTitle,
			MessageText: NotificationMessage,
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		// 通知
	}
	return hwndByTitle
}

func (m *Manager) GetCountReward() int {
	return m.countReward
}

func (m *Manager) GetCountFish() int {
	return m.countFish
}

func (m *Manager) SetCallBack(fn func()) *Manager {
	m.callback = fn
	return m
}
