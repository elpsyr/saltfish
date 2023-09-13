package win

import "testing"

func TestSendNotification(t *testing.T) {

	err := SendNotification(Message{
		AppID:       "helper",
		Title:       "haha",
		MessageText: "123",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
