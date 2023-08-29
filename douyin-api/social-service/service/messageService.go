package service

import (
	"douyin/douyin-api/social-service/dal/db"
	"douyin/douyin-api/social-service/proto"
	"log"
	"strconv"
)

type IChatService interface {
	PostMessage(SocialMessageChatRequest proto.SocialMessageChatRequest, userId int64) proto.SocialMessageChatResponse
	ListMessage(SocialMessageHistoryRequest proto.SocialMessageHistoryRequest, userId int64) proto.SocialMessageHistoryResponse
}

type ChatService struct {
	ChatRepository db.IChatRepository
}

func (c ChatService) ListMessage(SocialMessageHistoryRequest proto.SocialMessageHistoryRequest, userId int64) proto.SocialMessageHistoryResponse {

	//toUserID, _ := strconv.ParseInt(messageRequest.ToUserID, 10, 64)
	toUserID := SocialMessageHistoryRequest.ToUserId

	messages, err := c.ChatRepository.ListMessage(userId, toUserID, SocialMessageHistoryRequest.PreMsgTime)
	if err != nil {
		log.Printf("ListMessage|查询消息列表失败|%v", err)
		return proto.SocialMessageHistoryResponse{
			MessageList: nil,
			Response: proto.Response{
				StatusCode: 1,
				StatusMsg:  "get message list fail",
			}}
	}
	log.Printf("数据库消息列表为|%v", messages)

	var messageList []proto.Message
	messageList = make([]proto.Message, len(messages))
	for idx, message := range messages {
		messageList[idx].CreateTime = message.CreateTime
		messageList[idx].MessageId = message.MessageId
		messageList[idx].ToUserId = message.ToUserId
		messageList[idx].FromUserId = message.UserId
		messageList[idx].Content = message.Content
	}

	log.Printf("返回消息列表为|%v", messageList)
	return proto.SocialMessageHistoryResponse{
		MessageList: messageList,
		Response: proto.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		}}
}

func (c ChatService) PostMessage(SocialMessageChatRequest proto.SocialMessageChatRequest, userId int64) proto.SocialMessageChatResponse {

	content := SocialMessageChatRequest.Content
	toUserId := SocialMessageChatRequest.ToUserId
	actionType := SocialMessageChatRequest.ActionType

	err := c.ChatRepository.AddMessage(userId, strconv.FormatInt(toUserId, 10), content, string(actionType))

	if err != nil {
		log.Printf("PostMessage|增加消息失败|%v", err)
		return proto.SocialMessageChatResponse{
			Response: proto.Response{
				StatusCode: 1,
				StatusMsg:  "post message fail",
			},
		}
	}

	return proto.SocialMessageChatResponse{
		Response: proto.Response{
			StatusCode: 0,
			StatusMsg:  "post message success",
		},
	}

}

func NewChatService() IChatService {
	return ChatService{
		ChatRepository: db.NewChatRepository(),
	}
}