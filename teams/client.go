package teams

import (
	"fmt"
	"os"
	"strings"

	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/messagecard"
)

type TeamsNotify struct {
	client *goteamsnotify.TeamsClient
}

func NewClient() *TeamsNotify {
	client := goteamsnotify.NewTeamsClient()

	return &TeamsNotify{
		client: client,
	}
}

func (t *TeamsNotify) SendEntryMessage(entryMember string, entryAt string, members []string) error {
	msgCard := messagecard.NewMessageCard()
	msgCard.ThemeColor = "005B94"
	msgCard.Title = "入室通知 （テスト）"
	msgCard.Text = "研究室に新規入室がありました"

	msgSection := messagecard.NewSection()
	msgSection.ActivitySubtitle = "入室者情報"
	msgSection.ActivityText = entryMember
	msgSection.ActivitySubtitle = "入室時刻: "+entryAt
	msgSection.ActivityImage = "https://icon-library.com/images/enter-icon-png/enter-icon-png-5.jpg"

	if len(members) == 0 {
		msgSection.Text = "現在在室者はいません"
	} else {
		msgSection.Text = fmt.Sprintf("現在の在室者一覧: %s", strings.Join(members, ", "))
	}

	msgCard.Sections = append(msgCard.Sections, msgSection)

	return t.client.Send(os.Getenv("PRO_URL"), msgCard)
}

func (t *TeamsNotify) SendExitMessage(exitMember string, exitAt string, members []string) error {
	msgCard := messagecard.NewMessageCard()
	msgCard.ThemeColor = "005B94"
	msgCard.Title = "退室通知 （テスト）"
	msgCard.Text = "研究室に新規退室がありました"

	msgSection := messagecard.NewSection()
	msgSection.ActivitySubtitle = "退室者情報"
	msgSection.ActivityText = exitMember
	msgSection.ActivitySubtitle = "退室時刻: " + exitAt
	msgSection.ActivityImage = "https://upload.wikimedia.org/wikipedia/commons/thumb/7/74/PublicInformationSymbol_EmergencyExit.svg/1200px-PublicInformationSymbol_EmergencyExit.svg.png"

	if len(members) == 0 {
		msgSection.Text = "現在在室者はいません"
	} else {
		msgSection.Text = fmt.Sprintf("現在の在室者一覧: %s", strings.Join(members, ", "))
	}

	msgCard.Sections = append(msgCard.Sections, msgSection)

	return t.client.Send(os.Getenv("PRO_URL"), msgCard)
}
