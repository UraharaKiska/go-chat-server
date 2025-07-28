package converter

import (
	"github.com/UraharaKiska/go-chat-server/internal/model"
	"github.com/UraharaKiska/go-chat-server/internal/utils"
	desc "github.com/UraharaKiska/go-chat-server/pkg/chat_v1"
)

func ToChatFromDesc(chat *desc.CreateRequest) (*model.Chat) {
	return &model.Chat{
		Info: ToChatInfoFromDesc(chat.GetChatInfo()),
		Users: chat.GetUsernames(),
	}
}

func ToChatInfoFromDesc(info *desc.ChatInfo) (model.ChatInfo) {
	return model.ChatInfo{
		Name: info.GetName(),
	}
}

func ToMessageInfoFromDesc(info *desc.MessageInfo) (*model.MessageInfo) {
	t, _ := utils.ParseDateTime(info.GetDatetime())
	return &model.MessageInfo{
		From: info.GetFrom(),
		Text: info.GetText(),
		Timestamp: t,
	}
}