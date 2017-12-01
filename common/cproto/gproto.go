package gproto

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// 协议定义
// GProto协议

// 1. 基本类型
// <1> Int(Big-Endian编码)
// int8, int16, int32, int64
// uint8, uint16, uint32, uint64

// <2> Bytes/String
// uint16长度 + 具体字节流

// <3> Struct虚拟的约束
// uint16长度 + 具体字节流

// <4> List
// uint16具体元素个数 + 具体元素

// 2. 协议
// <1> 协议格式
// | Header | Payload | Header | Payload | ....

// <2> Header
// | MagicWord | Version | Command | PayloadLength |

// type GProtoMessageHeader struct {
// 	MagicWord `gproto:uint32`
// 	Version `gproto:uint16`
// 	Command `gproto:uint16`
// 	PayloadLength `gproto:uint32`
// }

// 协议实现

const (
	MAGIC_WORLD = 0x88888888 // 4 bytes
	MAX_VERSION = 10
	MIN_VERSION = 1
)

var (
	InvalidDecodeBufferError = errors.New("Invalid Decode Buffer")
	InvalidVersionError      = errors.New("Invalid Version")
)

func init() {
	// 初始化Message Descriptions
	messageDescriptions[MSG_PING_REQUEST] = "MSG_PING_REQUEST"
	messageDescriptions[MSG_PING_RESPONSE] = "MSG_PING_RESPONSE"
	messageDescriptions[MSG_ECHO_REQUEST] = "MSG_ECHO_REQUEST"
	messageDescriptions[MSG_ECHO_RESPONSE] = "MSG_ECHO_RESPONSE"
	messageDescriptions[MSG_KICKOFF_NOTIFY] = "MSG_KICKOFF_NOTIFY"
	messageDescriptions[MSG_AUTH_REQUEST] = "MSG_AUTH_REQUEST"
	messageDescriptions[MSG_AUTH_RESPONSE] = "MSG_AUTH_RESPONSE"
	messageDescriptions[MSG_SET_PROFILE_REQUEST] = "MSG_SET_PROFILE_REQUEST"
	messageDescriptions[MSG_SET_PROFILE_RESPONSE] = "MSG_SET_PROFILE_RESPONSE"
	messageDescriptions[MSG_GET_PROFILES_REQUEST] = "MSG_GET_PROFILES_REQUEST"
	messageDescriptions[MSG_GET_PROFILES_RESPONSE] = "MSG_GET_PROFILES_RESPONSE"
	messageDescriptions[MSG_SET_FRIEND_MARKNAME_REQUEST] = "MSG_SET_FRIEND_MARKNAME_REQUEST"
	messageDescriptions[MSG_SET_FRIEND_MARKNAME_RESPONSE] = "MSG_SET_FRIEND_MARKNAME_RESPONSE"
	messageDescriptions[MSG_MESSAGE_REQUEST] = "MSG_MESSAGE_REQUEST"
	messageDescriptions[MSG_MESSAGE_RESPONSE] = "MSG_MESSAGE_RESPONSE"
	messageDescriptions[MSG_MESSAGE_NOTIFY] = "MSG_MESSAGE_NOTIFY"
	messageDescriptions[MSG_MESSAGE_ACK] = "MSG_MESSAGE_ACK"
	messageDescriptions[MSG_ADD_FRIEND_REQUEST] = "MSG_ADD_FRIEND_REQUEST"
	messageDescriptions[MSG_ADD_FRIEND_RESPONSE] = "MSG_ADD_FRIEND_RESPONSE"
	messageDescriptions[MSG_ADD_FRIEND_NOTIFY] = "MSG_ADD_FRIEND_NOTIFY"
	messageDescriptions[MSG_ANSWER_FRIEND_REQUEST] = "MSG_ANSWER_FRIEND_REQUEST"
	messageDescriptions[MSG_ANSWER_FRIEND_RESPONSE] = "MSG_ANSWER_FRIEND_RESPONSE"
	messageDescriptions[MSG_ANSWER_FRIEND_NOTIFY] = "MSG_ANSWER_FRIEND_NOTIFY"
	messageDescriptions[MSG_DEL_FRIEND_REQUEST] = "MSG_DEL_FRIEND_REQUEST"
	messageDescriptions[MSG_DEL_FRIEND_RESPONSE] = "MSG_DEL_FRIEND_RESPONSE"
	messageDescriptions[MSG_DEL_FRIEND_NOTIFY] = "MSG_DEL_FRIEND_NOTIFY"
	messageDescriptions[MSG_GET_FRIENDS_REQUEST] = "MSG_GET_FRIENDS_REQUEST"
	messageDescriptions[MSG_GET_FRIENDS_RESPONSE] = "MSG_GET_FRIENDS_RESPONSE"
	messageDescriptions[MSG_GET_CHATROOM_INFO_REQUEST] = "MSG_GET_CHATROOM_INFO_REQUEST"
	messageDescriptions[MSG_GET_CHATROOM_INFO_RESPONSE] = "MSG_GET_CHATROOM_INFO_RESPONSE"
	messageDescriptions[MSG_ENTER_CHATROOM_REQUEST] = "MSG_ENTER_CHATROOM_REQUEST"
	messageDescriptions[MSG_ENTER_CHATROOM_RESPONSE] = "MSG_ENTER_CHATROOM_RESPONSE"
	messageDescriptions[MSG_ENTER_CHATROOM_NOTIFY] = "MSG_ENTER_CHATROOM_NOTIFY"
	messageDescriptions[MSG_LEAVE_CHATROOM_REQUEST] = "MSG_LEAVE_CHATROOM_REQUEST"
	messageDescriptions[MSG_LEAVE_CHATROOM_RESPONSE] = "MSG_LEAVE_CHATROOM_RESPONSE"
	messageDescriptions[MSG_LEAVE_CHATROOM_NOTIFY] = "MSG_LEAVE_CHATROOM_NOTIFY"
	messageDescriptions[MSG_CHATROOM_KICKOFF_NOTIFY] = "MSG_CHATROOM_KICKOFF_NOTIFY"
	messageDescriptions[MSG_CREATE_GROUP_REQUEST] = "MSG_CREATE_GROUP_REQUEST"
	messageDescriptions[MSG_CREATE_GROUP_RESPONSE] = "MSG_CREATE_GROUP_RESPONSE"
	messageDescriptions[MSG_GET_GROUP_INFO_REQUEST] = "MSG_GET_GROUP_INFO_REQUEST"
	messageDescriptions[MSG_GET_GROUP_INFO_RESPONSE] = "MSG_GET_GROUP_INFO_RESPONSE"
	messageDescriptions[MSG_SET_GROUP_INFO_REQUEST] = "MSG_SET_GROUP_INFO_REQUEST"
	messageDescriptions[MSG_SET_GROUP_INFO_RESPONSE] = "MSG_SET_GROUP_INFO_RESPONSE"
	messageDescriptions[MSG_GET_GROUP_LIST_REQUEST] = "MSG_GET_GROUP_LIST_REQUEST"
	messageDescriptions[MSG_GET_GROUP_LIST_RESPONSE] = "MSG_GET_GROUP_LIST_RESPONSE"
	messageDescriptions[MSG_ENTER_GROUP_REQUEST] = "MSG_ENTER_GROUP_REQUEST"
	messageDescriptions[MSG_ENTER_GROUP_RESPONSE] = "MSG_ENTER_GROUP_RESPONSE"
	messageDescriptions[MSG_ENTER_GROUP_NOTIFY] = "MSG_ENTER_GROUP_NOTIFY"
	messageDescriptions[MSG_LEAVE_GROUP_REQUEST] = "MSG_LEAVE_GROUP_REQUEST"
	messageDescriptions[MSG_LEAVE_GROUP_RESPONSE] = "MSG_LEAVE_GROUP_RESPONSE"
	messageDescriptions[MSG_LEAVE_GROUP_NOTIFY] = "MSG_LEAVE_GROUP_NOTIFY"
	messageDescriptions[MSG_GROUP_KICKOFF_NOTIFY] = "MSG_GROUP_KICKOFF_NOTIFY"
	messageDescriptions[MSG_ADD_BLACKLIST_REQUEST] = "MSG_ADD_BLACKLIST_REQUEST"
	messageDescriptions[MSG_ADD_BLACKLIST_RESPONSE] = "MSG_ADD_BLACKLIST_RESPONSE"
	messageDescriptions[MSG_ADD_BLACKLIST_NOTIFY] = "MSG_ADD_BLACKLIST_NOTIFY"
	messageDescriptions[MSG_DEL_BLACKLIST_REQUEST] = "MSG_DEL_BLACKLIST_REQUEST"
	messageDescriptions[MSG_DEL_BLACKLIST_RESPONSE] = "MSG_DEL_BLACKLIST_RESPONSE"
	messageDescriptions[MSG_DEL_BLACKLIST_NOTIFY] = "MSG_DEL_BLACKLIST_NOTIFY"
	messageDescriptions[MSG_GET_BLACKLISTS_REQUEST] = "MSG_GET_BLACKLISTS_REQUEST"
	messageDescriptions[MSG_GET_BLACKLISTS_RESPONSE] = "MSG_GET_BLACKLISTS_RESPONSE"
	messageDescriptions[MSG_SET_GROUP_MARKNAME_REQUEST] = "MSG_SET_GROUP_MARKNAME_REQUEST"
	messageDescriptions[MSG_SET_GROUP_MARKNAME_RESPONSE] = "MSG_SET_GROUP_MARKNAME_RESPONSE"

	// 初始化Message Creators
	messageCreators[MSG_PING_REQUEST] = func() IMessage { return new(GProtoPingRequest) }
	messageCreators[MSG_PING_RESPONSE] = func() IMessage { return new(GProtoPingResponse) }
	messageCreators[MSG_ECHO_REQUEST] = func() IMessage { return new(GProtoEchoRequest) }
	messageCreators[MSG_ECHO_RESPONSE] = func() IMessage { return new(GProtoEchoResponse) }
	messageCreators[MSG_KICKOFF_NOTIFY] = func() IMessage { return new(GProtoKickOffNotify) }
	messageCreators[MSG_AUTH_REQUEST] = func() IMessage { return new(GProtoAuthRequest) }
	messageCreators[MSG_AUTH_RESPONSE] = func() IMessage { return new(GProtoAuthResponse) }
	messageCreators[MSG_SET_PROFILE_REQUEST] = func() IMessage { return new(GProtoSetProfileRequest) }
	messageCreators[MSG_SET_PROFILE_RESPONSE] = func() IMessage { return new(GProtoSetProfileResponse) }
	messageCreators[MSG_GET_PROFILES_REQUEST] = func() IMessage { return new(GProtoGetProfilesRequest) }
	messageCreators[MSG_GET_PROFILES_RESPONSE] = func() IMessage { return new(GProtoGetProfilesResponse) }
	messageCreators[MSG_SET_FRIEND_MARKNAME_REQUEST] = func() IMessage { return new(GProtoSetFriendMarknameRequest) }
	messageCreators[MSG_SET_FRIEND_MARKNAME_RESPONSE] = func() IMessage { return new(GProtoSetFriendMarknameResponse) }
	messageCreators[MSG_MESSAGE_REQUEST] = func() IMessage { return new(GProtoMessageRequest) }
	messageCreators[MSG_MESSAGE_RESPONSE] = func() IMessage { return new(GProtoMessageResponse) }
	messageCreators[MSG_MESSAGE_NOTIFY] = func() IMessage { return new(GProtoMessageNotify) }
	messageCreators[MSG_MESSAGE_ACK] = func() IMessage { return new(GProtoMessageAck) }
	messageCreators[MSG_ADD_FRIEND_REQUEST] = func() IMessage { return new(GProtoAddFriendRequest) }
	messageCreators[MSG_ADD_FRIEND_RESPONSE] = func() IMessage { return new(GProtoAddFriendResponse) }
	messageCreators[MSG_ADD_FRIEND_NOTIFY] = func() IMessage { return new(GProtoAddFriendNotify) }
	messageCreators[MSG_ANSWER_FRIEND_REQUEST] = func() IMessage { return new(GProtoAnswerFriendRequest) }
	messageCreators[MSG_ANSWER_FRIEND_RESPONSE] = func() IMessage { return new(GProtoAnswerFriendResponse) }
	messageCreators[MSG_ANSWER_FRIEND_NOTIFY] = func() IMessage { return new(GProtoAnswerFriendNotify) }
	messageCreators[MSG_DEL_FRIEND_REQUEST] = func() IMessage { return new(GProtoDelFriendRequest) }
	messageCreators[MSG_DEL_FRIEND_RESPONSE] = func() IMessage { return new(GProtoDelFriendResponse) }
	messageCreators[MSG_DEL_FRIEND_NOTIFY] = func() IMessage { return new(GProtoDelFriendNotify) }
	messageCreators[MSG_GET_FRIENDS_REQUEST] = func() IMessage { return new(GProtoGetFriendsRequest) }
	messageCreators[MSG_GET_FRIENDS_RESPONSE] = func() IMessage { return new(GProtoGetFriendsResponse) }
	messageCreators[MSG_GET_CHATROOM_INFO_REQUEST] = func() IMessage { return new(GProtoGetChatRoomInfoRequest) }
	messageCreators[MSG_GET_CHATROOM_INFO_RESPONSE] = func() IMessage { return new(GProtoGetChatRoomInfoResponse) }
	messageCreators[MSG_ENTER_CHATROOM_RESPONSE] = func() IMessage { return new(GProtoGetChatRoomInfoResponse) }
	messageCreators[MSG_ENTER_CHATROOM_REQUEST] = func() IMessage { return new(GProtoEnterChatRoomRequest) }
	messageCreators[MSG_ENTER_CHATROOM_RESPONSE] = func() IMessage { return new(GProtoEnterChatRoomResponse) }
	messageCreators[MSG_ENTER_CHATROOM_NOTIFY] = func() IMessage { return new(GProtoEnterChatRoomNotify) }
	messageCreators[MSG_LEAVE_CHATROOM_REQUEST] = func() IMessage { return new(GProtoLeaveChatRoomRequest) }
	messageCreators[MSG_LEAVE_CHATROOM_RESPONSE] = func() IMessage { return new(GProtoLeaveChatRoomResponse) }
	messageCreators[MSG_LEAVE_CHATROOM_NOTIFY] = func() IMessage { return new(GProtoLeaveChatRoomNotify) }
	messageCreators[MSG_CHATROOM_KICKOFF_NOTIFY] = func() IMessage { return new(GProtoKickOffNotify) }
	messageCreators[MSG_CREATE_GROUP_REQUEST] = func() IMessage { return new(GProtoCreateGroupRequest) }
	messageCreators[MSG_CREATE_GROUP_RESPONSE] = func() IMessage { return new(GProtoCreateGroupResponse) }
	messageCreators[MSG_GET_GROUP_INFO_REQUEST] = func() IMessage { return new(GProtoGetGroupInfoRequest) }
	messageCreators[MSG_GET_GROUP_INFO_RESPONSE] = func() IMessage { return new(GProtoGetGroupInfoResponse) }
	messageCreators[MSG_SET_GROUP_INFO_REQUEST] = func() IMessage { return new(GProtoSetGroupInfoRequest) }
	messageCreators[MSG_SET_GROUP_INFO_RESPONSE] = func() IMessage { return new(GProtoSetGroupInfoResponse) }
	messageCreators[MSG_GET_GROUP_LIST_REQUEST] = func() IMessage { return new(GProtoGetGroupListRequest) }
	messageCreators[MSG_GET_GROUP_LIST_RESPONSE] = func() IMessage { return new(GProtoGetGroupListResponse) }
	messageCreators[MSG_ENTER_GROUP_REQUEST] = func() IMessage { return new(GProtoEnterGroupRequest) }
	messageCreators[MSG_ENTER_GROUP_RESPONSE] = func() IMessage { return new(GProtoEnterGroupResponse) }
	messageCreators[MSG_ENTER_GROUP_NOTIFY] = func() IMessage { return new(GProtoEnterGroupNotify) }
	messageCreators[MSG_LEAVE_GROUP_REQUEST] = func() IMessage { return new(GProtoLeaveGroupRequest) }
	messageCreators[MSG_LEAVE_GROUP_RESPONSE] = func() IMessage { return new(GProtoLeaveGroupResponse) }
	messageCreators[MSG_LEAVE_GROUP_NOTIFY] = func() IMessage { return new(GProtoLeaveGroupNotify) }
	messageCreators[MSG_GROUP_KICKOFF_NOTIFY] = func() IMessage { return new(GProtoGroupKickoffNotify) }
	messageCreators[MSG_ADD_BLACKLIST_REQUEST] = func() IMessage { return new(GProtoAddBlacklistRequest) }
	messageCreators[MSG_ADD_BLACKLIST_RESPONSE] = func() IMessage { return new(GProtoAddBlacklistResponse) }
	messageCreators[MSG_ADD_BLACKLIST_NOTIFY] = func() IMessage { return new(GProtoAddBlacklistNotify) }
	messageCreators[MSG_DEL_BLACKLIST_REQUEST] = func() IMessage { return new(GProtoDelBlacklistRequest) }
	messageCreators[MSG_DEL_BLACKLIST_RESPONSE] = func() IMessage { return new(GProtoDelBlacklistResponse) }
	messageCreators[MSG_DEL_BLACKLIST_NOTIFY] = func() IMessage { return new(GProtoDelBlacklistNotify) }
	messageCreators[MSG_GET_BLACKLISTS_REQUEST] = func() IMessage { return new(GProtoGetBlacklistsRequest) }
	messageCreators[MSG_GET_BLACKLISTS_RESPONSE] = func() IMessage { return new(GProtoGetBlacklistsResponse) }
	messageCreators[MSG_SET_GROUP_MARKNAME_REQUEST] = func() IMessage { return new(GProtoSetGroupMarknameRequest) }
	messageCreators[MSG_SET_GROUP_MARKNAME_RESPONSE] = func() IMessage { return new(GProtoSetGroupMarknameResponse) }
}

// 全局的Message Descriptions
var messageDescriptions map[uint16]string = make(map[uint16]string)

// 全局的Message Creators
type MessageCreator func() IMessage

var messageCreators map[uint16]MessageCreator = make(map[uint16]MessageCreator)

// Command
type Command uint16

func (cmd Command) String() string {
	c := uint16(cmd)
	if desc, ok := messageDescriptions[c]; ok {
		return desc
	} else {
		return fmt.Sprintf("Unknown Command=%v", c)
	}
}

type IMessage interface {
	// 编码IMessage, 如果error=nil, 则编码成功
	Encode(version uint16) ([]byte, error)

	// 解码IMessage, 如果error=nil, 则解码成功
	Decode(version uint16, buf []byte) error
}

type Message struct {
	MagicWord     uint32
	Version       uint16
	Command       uint16
	PayloadLength uint32
	Body          interface{} // IMessage
}

func NewEmptyMessage() *Message {
	return &Message{}
}

func NewMessage(version, command uint16, body interface{}) *Message {
	return &Message{
		MagicWord:     MAGIC_WORLD,
		Version:       version,
		Command:       command,
		PayloadLength: 0, // 发送的时候会计算 & 长度
		Body:          body,
	}
}

func (message *Message) EncodeHead() ([]byte, error) {
	buf := make([]byte, 12, 12)
	binary.BigEndian.PutUint32(buf[0:], message.MagicWord)
	binary.BigEndian.PutUint16(buf[4:], message.Version)
	binary.BigEndian.PutUint16(buf[6:], message.Command)
	binary.BigEndian.PutUint32(buf[8:], message.PayloadLength)
	return buf[0:], nil
}

func (message *Message) DecodeHead(buf []byte) error {
	if len(buf) != 12 {
		return InvalidDecodeBufferError
	}
	message.MagicWord = binary.BigEndian.Uint32(buf[0:4])
	message.Version = binary.BigEndian.Uint16(buf[4:6])
	message.Command = binary.BigEndian.Uint16(buf[6:8])
	message.PayloadLength = binary.BigEndian.Uint32(buf[8:])
	return nil
}

// 编码Message.Body, 如果error=nil, 则编码成功
func (message *Message) Encode() ([]byte, error) {
	if message.Body != nil {
		if m, ok := message.Body.(IMessage); ok {
			return m.Encode(message.Version)
		}
		panic(fmt.Errorf("Message.Body is NOT IMessage"))
	} else {
		return nil, nil
	}
}

// 解码Message.Body, 如果error=nil, 则解码成功
func (message *Message) Decode(buf []byte) error {
	if creator, ok := messageCreators[message.Command]; ok {
		iMessage := creator()
		err := iMessage.Decode(message.Version, buf)
		message.Body = iMessage // 解码的结果
		return err
	}

	panic(fmt.Errorf("Message.Command Invalid %v", message.Command))
}
