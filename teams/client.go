package teams

import (
	"fmt"
	"strings"

	goteamsnotify "github.com/dasrick/go-teams-notify/v2"
)

type TeamsNotify struct {
	client goteamsnotify.API
}

func NewClient() *TeamsNotify {
	client := goteamsnotify.NewClient()

	return &TeamsNotify{
		client: client,
	}
}

func (t *TeamsNotify) SendEntryMessage(entryMember string, members []string) error {
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = "入室通知 （テスト）"
	msgCard.Text = "研究室に新規入室がありました"

	msgSection := goteamsnotify.NewMessageCardSection()
	msgSection.ActivitySubtitle = "入室者情報"
	msgSection.ActivityText = entryMember
	msgSection.ActivityImage = "https://icon-library.com/images/enter-icon-png/enter-icon-png-5.jpg"

	if len(members) == 0 {
		msgSection.Text = "現在在室者はいません"
	} else {
		msgSection.Text = fmt.Sprintf("現在の在室者一覧: %s", strings.Join(members, ", "))
	}

	msgCard.Sections = append(msgCard.Sections, msgSection)
	
	return t.client.Send("test webhocks", msgCard)
}

func (t *TeamsNotify) SendExitMessage(exitMember string, members []string) error {
	msgCard := goteamsnotify.NewMessageCard()
	msgCard.Title = "退室通知 （テスト）"
	msgCard.Text = "研究室に新規退室がありました"

	msgSection := goteamsnotify.NewMessageCardSection()
	msgSection.ActivitySubtitle = "退室者情報"

	if len(members) == 0 {
		msgSection.Text = "現在在室者はいません"
	} else {
		msgSection.Text = fmt.Sprintf("現在の在室者一覧: %s", strings.Join(members, ", "))
	}

	msgCard.Sections = append(msgCard.Sections, msgSection)

	return t.client.Send("test webhocks", msgCard)
}
