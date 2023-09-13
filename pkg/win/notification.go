package win

import (
	"github.com/go-toast/toast"
)

type Message struct {
	AppID       string
	Title       string
	MessageText string
}

func SendNotification(m Message) error {
	notification := toast.Notification{
		AppID:   m.AppID,       // 应用程序ID
		Title:   m.Title,       // 通知标题
		Message: m.MessageText, // 通知内容
		//Icon:    "icon.ico",    // 可选的通知图标文件
		Audio: toast.Default, // 可选的通知声音（默认为Default）
	}

	return notification.Push()
}
