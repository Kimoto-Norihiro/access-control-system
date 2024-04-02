package teams

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestSendEnterMessage(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		enterMember string
		enterAt     string
		members     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "send enter message",
			args: args{
				enterMember: "test1",
				enterAt:     "2021-01-01 00:00:00",
				members:     []string{"test1", "test2"},
			},
			wantErr: false,
		},
		{
			name: "send enter message with no members",
			args: args{
				enterMember: "test",
				enterAt:     "2021-01-01 00:00:00",
				members:     []string{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient()
			err := client.SendEnterMessage(tt.args.enterMember, tt.args.enterAt, tt.args.members)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendEnterMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
