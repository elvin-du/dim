package goproto

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

// 全局错误码
const (
	CODE_SUCCESS                  = 0  // Success
	CODE_IGNORE                   = 1  // Ignore
	CODE_ERROR_NO_USER            = 2  // User Not Found
	CODE_ERROR_PASSWORD_ERROR     = 3  // Password Error
	CODE_ERROR_UNKNOWN            = 4  // Unknown Error
	CODE_ERROR_REPLACED           = 5  // Session Replaced
	CODE_ERROR_CLIENT             = 6  // Client Logci Error
	CODE_ERROR_SERVER             = 7  // Server Internal Error
	CODE_ERROR_DB                 = 8  // DB Ops Error
	CODE_ERROR_NOT_FOUND          = 9  // Not Found Error
	CODE_ERROR_AUTH_ERROR         = 10 // AUTH Error
	CODE_ERROR_UNSUPPORT_API      = 11 // UnSupport API
	CODE_ERROR_INVALID_JSON       = 12 // Invalid JSON Format
	CODE_ERROR_ALREADY_REGISTERED = 13 // Already Registered Error
	CODE_ERROR_NO_APP             = 14 // No App Error
	CODE_ERROR_ALREADY_EXIST      = 15 // Already Exist
)

func CodeToMessage(code int) string {
	switch code {
	case CODE_SUCCESS:
		return "Success"
	case CODE_IGNORE:
		return "Ignore"
	case CODE_ERROR_NO_USER:
		return "User Not Found"
	case CODE_ERROR_PASSWORD_ERROR:
		return "Password Error"
	case CODE_ERROR_UNKNOWN:
		return "Unknown Error"
	case CODE_ERROR_REPLACED:
		return "Session Replaced"
	case CODE_ERROR_CLIENT:
		return "Client Logci Error"
	case CODE_ERROR_SERVER:
		return "Server Internal Error"
	case CODE_ERROR_DB:
		return "DB Ops Error"
	case CODE_ERROR_NOT_FOUND:
		return "Not Found Error"
	case CODE_ERROR_AUTH_ERROR:
		return "Auth Error"
	case CODE_ERROR_UNSUPPORT_API:
		return "UnSupport API"
	case CODE_ERROR_INVALID_JSON:
		return "Invalid JSON Format"
	case CODE_ERROR_ALREADY_REGISTERED:
		return "Already Registered Error"
	case CODE_ERROR_NO_APP:
		return "No App Error"
	default:
		return fmt.Sprintf("Unknown Code: %v", code)
	}
}

const (
	MSG_PING_REQUEST                 = 0x1001 // C -> S
	MSG_PING_RESPONSE                = 0x1002 // S -> C
	MSG_ECHO_REQUEST                 = 0x1003 // C -> S
	MSG_ECHO_RESPONSE                = 0x1004 // S -> C
	MSG_KICKOFF_NOTIFY               = 0x1005 // S -> C
	MSG_AUTH_REQUEST                 = 0x1006 // C -> S
	MSG_AUTH_RESPONSE                = 0x1007 // S -> C
	MSG_SET_PROFILE_REQUEST          = 0x1008 // C -> S
	MSG_SET_PROFILE_RESPONSE         = 0x1009 // S -> C
	MSG_SET_FRIEND_MARKNAME_REQUEST  = 0x100a // C -> S
	MSG_SET_FRIEND_MARKNAME_RESPONSE = 0x100b // S -> C
	MSG_GET_PROFILES_REQUEST         = 0x100c // C -> S
	MSG_GET_PROFILES_RESPONSE        = 0x100d // S -> C
	MSG_MESSAGE_REQUEST              = 0x100e // C -> S
	MSG_MESSAGE_RESPONSE             = 0x100f // S -> C
	MSG_MESSAGE_NOTIFY               = 0x1010 // S -> C
	MSG_MESSAGE_ACK                  = 0x1011 // C -> S
	MSG_ADD_FRIEND_REQUEST           = 0x1012 // C -> S
	MSG_ADD_FRIEND_RESPONSE          = 0x1013 // S -> C
	MSG_ADD_FRIEND_NOTIFY            = 0x1014 // S -> C
	MSG_ANSWER_FRIEND_REQUEST        = 0x1015 // C -> S
	MSG_ANSWER_FRIEND_RESPONSE       = 0x1016 // S -> C
	MSG_ANSWER_FRIEND_NOTIFY         = 0x1017 // S -> C
	MSG_DEL_FRIEND_REQUEST           = 0x1018 // C -> S
	MSG_DEL_FRIEND_RESPONSE          = 0x1019 // S -> C
	MSG_DEL_FRIEND_NOTIFY            = 0x101a // S -> C
	MSG_GET_FRIENDS_REQUEST          = 0x101b // C -> S
	MSG_GET_FRIENDS_RESPONSE         = 0x101c // S -> C
	MSG_GET_CHATROOM_INFO_REQUEST    = 0x101d // C -> S
	MSG_GET_CHATROOM_INFO_RESPONSE   = 0x101e // S -> S
	MSG_ENTER_CHATROOM_REQUEST       = 0x101f // C -> S
	MSG_ENTER_CHATROOM_RESPONSE      = 0x1020 // S -> C
	MSG_ENTER_CHATROOM_NOTIFY        = 0x1021 // S -> C
	MSG_LEAVE_CHATROOM_REQUEST       = 0x1022 // C -> S
	MSG_LEAVE_CHATROOM_RESPONSE      = 0x1023 // S -> C
	MSG_LEAVE_CHATROOM_NOTIFY        = 0x1024 // S -> C
	MSG_CHATROOM_KICKOFF_NOTIFY      = 0x1025 // S -> C
	MSG_CREATE_GROUP_REQUEST         = 0x1026 // C -> S
	MSG_CREATE_GROUP_RESPONSE        = 0x1027 // S -> C
	MSG_GET_GROUP_INFO_REQUEST       = 0x1028 // C -> S
	MSG_GET_GROUP_INFO_RESPONSE      = 0x1029 // S -> S
	MSG_SET_GROUP_INFO_REQUEST       = 0x102a // C -> S
	MSG_SET_GROUP_INFO_RESPONSE      = 0x102b // S -> S
	MSG_GET_GROUP_LIST_REQUEST       = 0x102c // C -> S
	MSG_GET_GROUP_LIST_RESPONSE      = 0x102d // S -> S
	MSG_ENTER_GROUP_REQUEST          = 0x102e // C -> S
	MSG_ENTER_GROUP_RESPONSE         = 0x1030 // S -> C
	MSG_ENTER_GROUP_NOTIFY           = 0x1031 // S -> C
	MSG_LEAVE_GROUP_REQUEST          = 0x1032 // C -> S
	MSG_LEAVE_GROUP_RESPONSE         = 0x1033 // S -> C
	MSG_LEAVE_GROUP_NOTIFY           = 0x1034 // S -> C
	MSG_GROUP_KICKOFF_NOTIFY         = 0x1035 // S -> C
	MSG_ADD_BLACKLIST_REQUEST        = 0x1035 // C -> S
	MSG_ADD_BLACKLIST_RESPONSE       = 0x1036 // S -> C
	MSG_ADD_BLACKLIST_NOTIFY         = 0x1037 // S -> C
	MSG_DEL_BLACKLIST_REQUEST        = 0x1038 // C -> S
	MSG_DEL_BLACKLIST_RESPONSE       = 0x1039 // S -> C
	MSG_DEL_BLACKLIST_NOTIFY         = 0x103a // S -> C
	MSG_GET_BLACKLISTS_REQUEST       = 0x103b // C -> S
	MSG_GET_BLACKLISTS_RESPONSE      = 0x103c // S -> C
	MSG_SET_GROUP_MARKNAME_REQUEST   = 0x103d // C -> S
	MSG_SET_GROUP_MARKNAME_RESPONSE  = 0x103e // S -> C
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
	//TODO
	messageCreators[MSG_PING_REQUEST] = func() IMessage { return new(PingRequest) }
	messageCreators[MSG_PING_RESPONSE] = func() IMessage { return new(PingResponse) }
	messageCreators[MSG_ECHO_REQUEST] = func() IMessage { return new(EchoRequest) }
	messageCreators[MSG_ECHO_RESPONSE] = func() IMessage { return new(EchoResponse) }
	messageCreators[MSG_KICKOFF_NOTIFY] = func() IMessage { return new(KickOffNotify) }
	messageCreators[MSG_AUTH_REQUEST] = func() IMessage { return new(AuthRequest) }
	messageCreators[MSG_AUTH_RESPONSE] = func() IMessage { return new(AuthResponse) }
	messageCreators[MSG_SET_PROFILE_REQUEST] = func() IMessage { return new(SetProfileRequest) }
	messageCreators[MSG_SET_PROFILE_RESPONSE] = func() IMessage { return new(SetProfileResponse) }
	messageCreators[MSG_GET_PROFILES_REQUEST] = func() IMessage { return new(GetProfilesRequest) }
	messageCreators[MSG_GET_PROFILES_RESPONSE] = func() IMessage { return new(GetProfilesResponse) }
	messageCreators[MSG_SET_FRIEND_MARKNAME_REQUEST] = func() IMessage { return new(SetFriendMarknameRequest) }
	messageCreators[MSG_SET_FRIEND_MARKNAME_RESPONSE] = func() IMessage { return new(SetFriendMarknameResponse) }
	messageCreators[MSG_MESSAGE_REQUEST] = func() IMessage { return new(MessageRequest) }
	messageCreators[MSG_MESSAGE_RESPONSE] = func() IMessage { return new(MessageResponse) }
	messageCreators[MSG_MESSAGE_NOTIFY] = func() IMessage { return new(MessageNotify) }
	messageCreators[MSG_MESSAGE_ACK] = func() IMessage { return new(MessageAck) }
	messageCreators[MSG_ADD_FRIEND_REQUEST] = func() IMessage { return new(AddFriendRequest) }
	messageCreators[MSG_ADD_FRIEND_RESPONSE] = func() IMessage { return new(AddFriendResponse) }
	messageCreators[MSG_ADD_FRIEND_NOTIFY] = func() IMessage { return new(AddFriendNotify) }
	messageCreators[MSG_ANSWER_FRIEND_REQUEST] = func() IMessage { return new(AnswerFriendRequest) }
	messageCreators[MSG_ANSWER_FRIEND_RESPONSE] = func() IMessage { return new(AnswerFriendResponse) }
	messageCreators[MSG_ANSWER_FRIEND_NOTIFY] = func() IMessage { return new(AnswerFriendNotify) }
	messageCreators[MSG_DEL_FRIEND_REQUEST] = func() IMessage { return new(DelFriendRequest) }
	messageCreators[MSG_DEL_FRIEND_RESPONSE] = func() IMessage { return new(DelFriendResponse) }
	messageCreators[MSG_DEL_FRIEND_NOTIFY] = func() IMessage { return new(DelFriendNotify) }
	messageCreators[MSG_GET_FRIENDS_REQUEST] = func() IMessage { return new(GetFriendsRequest) }
	messageCreators[MSG_GET_FRIENDS_RESPONSE] = func() IMessage { return new(GetFriendsResponse) }
	messageCreators[MSG_CREATE_GROUP_REQUEST] = func() IMessage { return new(CreateGroupRequest) }
	messageCreators[MSG_CREATE_GROUP_RESPONSE] = func() IMessage { return new(CreateGroupResponse) }
	messageCreators[MSG_GET_GROUP_INFO_REQUEST] = func() IMessage { return new(GetGroupInfoRequest) }
	messageCreators[MSG_GET_GROUP_INFO_RESPONSE] = func() IMessage { return new(GetGroupInfoResponse) }
	messageCreators[MSG_SET_GROUP_INFO_REQUEST] = func() IMessage { return new(SetGroupInfoRequest) }
	messageCreators[MSG_SET_GROUP_INFO_RESPONSE] = func() IMessage { return new(SetGroupInfoResponse) }
	messageCreators[MSG_GET_GROUP_LIST_REQUEST] = func() IMessage { return new(GetGroupListRequest) }
	messageCreators[MSG_GET_GROUP_LIST_RESPONSE] = func() IMessage { return new(GetGroupListResponse) }
	messageCreators[MSG_ENTER_GROUP_REQUEST] = func() IMessage { return new(EnterGroupRequest) }
	messageCreators[MSG_ENTER_GROUP_RESPONSE] = func() IMessage { return new(EnterGroupResponse) }
	messageCreators[MSG_ENTER_GROUP_NOTIFY] = func() IMessage { return new(EnterGroupNotify) }
	messageCreators[MSG_LEAVE_GROUP_REQUEST] = func() IMessage { return new(LeaveGroupRequest) }
	messageCreators[MSG_LEAVE_GROUP_RESPONSE] = func() IMessage { return new(LeaveGroupResponse) }
	messageCreators[MSG_LEAVE_GROUP_NOTIFY] = func() IMessage { return new(LeaveGroupNotify) }
	messageCreators[MSG_GROUP_KICKOFF_NOTIFY] = func() IMessage { return new(GroupKickOffNotify) }
	messageCreators[MSG_ADD_BLACKLIST_REQUEST] = func() IMessage { return new(AddBlacklistRequest) }
	messageCreators[MSG_ADD_BLACKLIST_RESPONSE] = func() IMessage { return new(AddBlacklistResponse) }
	messageCreators[MSG_ADD_BLACKLIST_NOTIFY] = func() IMessage { return new(AddBlacklistNotify) }
	messageCreators[MSG_DEL_BLACKLIST_REQUEST] = func() IMessage { return new(DelBlacklistRequest) }
	messageCreators[MSG_DEL_BLACKLIST_RESPONSE] = func() IMessage { return new(DelBlacklistResponse) }
	messageCreators[MSG_DEL_BLACKLIST_NOTIFY] = func() IMessage { return new(DelBlacklistNotify) }
	messageCreators[MSG_GET_BLACKLISTS_REQUEST] = func() IMessage { return new(GetBlacklistsRequest) }
	messageCreators[MSG_GET_BLACKLISTS_RESPONSE] = func() IMessage { return new(GetBlacklistsResponse) }
	messageCreators[MSG_SET_GROUP_MARKNAME_REQUEST] = func() IMessage { return new(SetGroupMarknameRequest) }
	messageCreators[MSG_SET_GROUP_MARKNAME_RESPONSE] = func() IMessage { return new(SetGroupMarknameResponse) }
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
	//	Encode(version uint16) ([]byte, error)
	Marshal() (dAtA []byte, err error)

	// 解码IMessage, 如果error=nil, 则解码成功
	//	Decode(version uint16, buf []byte) error
	Unmarshal(dAtA []byte) error
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
			//			return m.Encode(message.Version)
			return m.Marshal()
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
		//		err := iMessage.Decode(message.Version, buf)
		err := iMessage.Unmarshal(buf)
		message.Body = iMessage // 解码的结果
		return err
	}

	panic(fmt.Errorf("Message.Command Invalid %v", message.Command))
}
