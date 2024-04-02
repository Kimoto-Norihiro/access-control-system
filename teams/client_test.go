package teams

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestSendEntryMessage(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		entryMember string
		entryAt     string
		members     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "send entry message",
			args: args{
				entryMember: "test1",
				entryAt:     "2021-01-01 00:00:00",
				members:     []string{"test1", "test2"},
			},
			wantErr: false,
		},
		{
			name: "send entry message with no members",
			args: args{
				entryMember: "test",
				entryAt:     "2021-01-01 00:00:00",
				members:     []string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient()
			err := client.SendEntryMessage(tt.args.entryMember, tt.args.entryAt, tt.args.members)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendEntryMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
