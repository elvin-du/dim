package gproto

import (
	"fmt"
)

///////////////////
// 消息相关的模块
///////////////////

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

// optional字段前缀
const (
	OPTIONAL_TRUE  = 1 // 字段写入
	OPTIONAL_FALSE = 0 // 字段不写(忽略), 默认值
)

////////////////////////////////////
// message PingRequest {
// }

type GProtoPingRequest struct {
}

////////////////////////////////////

func (pingRequest *GProtoPingRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		return nil, nil
	}

	return nil, InvalidVersionError
}

func (pingRequest *GProtoPingRequest) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message PingResponse {
//     required uint8 code = 1;
// }

type GProtoPingResponse struct {
	Code uint8
}

////////////////////////////////////

func (pingResponse *GProtoPingResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoPingResponse.Code
		if err = buffer.WriteUInt8(pingResponse.Code); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (pingResponse *GProtoPingResponse) Decode(version uint16, buf []byte) error {
	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoPingResponse.Code
		if err = buffer.ReadUInt8(&pingResponse.Code); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message EchoRequest {
//     required string data = 1;
// }

type GProtoEchoRequest struct {
	Data string
}

////////////////////////////////////

func (echoRequest *GProtoEchoRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoEchoRequest.Data
		if err = buffer.WriteString(echoRequest.Data); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (echoRequest *GProtoEchoRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 2 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoEchoRequest.Data
		if echoRequest.Data, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message EchoResponse {
//     required uint8 code = 1;
//     required string data = 2;
// }

type GProtoEchoResponse struct {
	Code uint8
	Data string
}

////////////////////////////////////

func (echoResponse *GProtoEchoResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoEchoResponse.Code
		if err = buffer.WriteUInt8(echoResponse.Code); err != nil {
			return nil, err
		}

		// GProtoEchoResponse.Data
		if err = buffer.WriteString(echoResponse.Data); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (echoResponse *GProtoEchoResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoEchoResponse.Code
		if err = buffer.ReadUInt8(&echoResponse.Code); err != nil {
			return err
		}

		// GProtoEchoResponse.Data
		if echoResponse.Data, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message KickoffNotify {
//     required uint8 code = 1;
// }

type GProtoKickOffNotify struct {
	Code uint8
}

////////////////////////////////////

func (kickOffNotify *GProtoKickOffNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoKickOffNotify.Code
		if err = buffer.WriteUInt8(kickOffNotify.Code); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (kickOffNotify *GProtoKickOffNotify) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoKickOffNotify.Code
		if err = buffer.ReadUInt8(&kickOffNotify.Code); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message AuthRequest {
//     required string appId = 1;
//     required string appSecret = 2;
//     required string account = 3;
//     required string token = 4;
//     required string customVersion = 5;
//     optional uint8 groupFlag = 6;
// }

type GProtoAuthRequest struct {
	AppId         string
	AppSecret     string
	Account       string
	Token         string
	CustomVersion string
	GroupFlag     uint8 // optional 1表示不下发group信息
	GroupFlagFlag uint8
}

////////////////////////////////////

func (authRequest *GProtoAuthRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoAuthRequest.AppId
		if err = buffer.WriteString(authRequest.AppId); err != nil {
			return nil, err
		}

		// GProtoAuthRequest.AppSecret
		if err = buffer.WriteString(authRequest.AppSecret); err != nil {
			return nil, err
		}

		// GProtoAuthRequest.Account
		if err = buffer.WriteString(authRequest.Account); err != nil {
			return nil, err
		}

		// GProtoAuthRequest.Token
		if err = buffer.WriteString(authRequest.Token); err != nil {
			return nil, err
		}

		// GProtoAuthRequest.CustomVersion
		if err = buffer.WriteString(authRequest.CustomVersion); err != nil {
			return nil, err
		}

		// GProtoAuthRequest.GroupFlag [optional]
		if authRequest.GroupFlagFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteUInt8(authRequest.GroupFlag); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (authRequest *GProtoAuthRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 6 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoAuthRequest.AppId
		if authRequest.AppId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAuthRequest.AppSecret
		if authRequest.AppSecret, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAuthRequest.Account
		if authRequest.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAuthRequest.Token
		if authRequest.Token, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAuthRequest.CustomVersion
		if authRequest.CustomVersion, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAuthRequest.Gender [optional]
		if err = buffer.ReadUInt8(&authRequest.GroupFlagFlag); err != nil {
			// 注意:
			// 兼容老版本客户端(读不到这个字段直接返回成功)
			authRequest.GroupFlag = 0
			authRequest.GroupFlagFlag = OPTIONAL_FALSE
			return nil
		}
		if authRequest.GroupFlagFlag == OPTIONAL_TRUE {
			if err = buffer.ReadUInt8(&authRequest.GroupFlag); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message AuthResponse {
//     required uint8 code = 1;
//     required SelfUser user = 2;
//     repeated Friend friends = 3;
//	   repeated GroupInfo groups = 4;
//     repeated Blacklist blacklists = 5;
// }

type GProtoAuthResponse struct {
	Code       uint8
	User       *GProtoSelfUser
	Friends    []*GProtoFriend
	Groups     []*GProtoGroupInfo
	Blacklists []*GProtoBlacklist
}

////////////////////////////////////

func (authResponse *GProtoAuthResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoAuthResponse.Code
		if err = buffer.WriteUInt8(authResponse.Code); err != nil {
			return nil, err
		}

		// GProtoAuthResponse.User
		if err = buffer.WriteStruct(1, authResponse.User); err != nil {
			return nil, err
		}

		// GProtoAuthResponse.Friends [array]
		if err = buffer.WriteArray(1, authResponse.Friends); err != nil {
			return nil, err
		}

		// GProtoAuthResponse.Groups [array]
		if err = buffer.WriteArray(1, authResponse.Groups); err != nil {
			return nil, err
		}

		// GProtoAuthResponse.Blacklists [array]
		if err = buffer.WriteArray(1, authResponse.Blacklists); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (authResponse *GProtoAuthResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 8 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoAuthResponse.Code
		if err = buffer.ReadUInt8(&authResponse.Code); err != nil {
			return err
		}

		// GProtoAuthResponse.User
		authResponse.User = &GProtoSelfUser{}
		if err = buffer.ReadStruct(1, authResponse.User); err != nil {
			return err
		}

		// GProtoAuthResponse.Friends [array]
		var friendsNum uint16
		if err = buffer.ReadUInt16(&friendsNum); err != nil {
			return err
		} else {
			authResponse.Friends = make([]*GProtoFriend, friendsNum, friendsNum)
			for i := 0; i < int(friendsNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						authResponse.Friends[i] = &GProtoFriend{}
						if err = authResponse.Friends[i].Decode(1, bufObj); err != nil {
							return err
						}
					}
				}
			}
		}

		// GProtoAuthResponse.Groups [array]
		var groupsNum uint16
		if err = buffer.ReadUInt16(&groupsNum); err != nil {
			return err
		} else {
			authResponse.Groups = make([]*GProtoGroupInfo, groupsNum, groupsNum)
			for i := 0; i < int(groupsNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						authResponse.Groups[i] = &GProtoGroupInfo{}
						if err = authResponse.Groups[i].Decode(1, bufObj); err != nil {
							return err
						}
					}
				}
			}
		}

		// GProtoAuthResponse.Blacklists [array]
		var blacklistsNum uint16
		if err = buffer.ReadUInt16(&blacklistsNum); err != nil {
			return err
		} else {
			authResponse.Blacklists = make([]*GProtoBlacklist, blacklistsNum, blacklistsNum)
			for i := 0; i < int(blacklistsNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						authResponse.Blacklists[i] = &GProtoBlacklist{}
						if err = authResponse.Blacklists[i].Decode(1, bufObj); err != nil {
							return err
						}
					}
				}
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message SetProfileRequest {
//     optional string nickname = 1;
//     optional uint8 gender = 2;
//     optional string avatar = 3;
//     optional string signature = 4;
//     optional string extern = 5;
// }

type GProtoSetProfileRequest struct {
	Nickname      string // optional
	NicknameFlag  uint8
	Gender        uint8 // optional
	GenderFlag    uint8
	Avatar        string // optional
	AvatarFlag    uint8
	Signature     string // optional
	SignatureFlag uint8
	Extern        string // optional
	ExternFlag    uint8
}

func (setProfileRequest *GProtoSetProfileRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoSetProfileRequest.Nickname [optional]
		if setProfileRequest.NicknameFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(setProfileRequest.Nickname); err != nil {
				return nil, err
			}
		}

		// GProtoSetProfileRequest.Gender [optional]
		if setProfileRequest.GenderFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteUInt8(setProfileRequest.Gender); err != nil {
				return nil, err
			}
		}

		// GProtoSetProfileRequest.Avatar [optional]
		if setProfileRequest.AvatarFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(setProfileRequest.Avatar); err != nil {
				return nil, err
			}
		}

		// GProtoSetProfileRequest.Signature [optional]
		if setProfileRequest.SignatureFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(setProfileRequest.Signature); err != nil {
				return nil, err
			}
		}

		// GProtoSetProfileRequest.Extern [optional]
		if setProfileRequest.ExternFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(setProfileRequest.Extern); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (setProfileRequest *GProtoSetProfileRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoSetProfileRequest.Nickname [optional]
		if err = buffer.ReadUInt8(&setProfileRequest.NicknameFlag); err != nil {
			return err
		}
		if setProfileRequest.NicknameFlag == OPTIONAL_TRUE {
			if setProfileRequest.Nickname, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoSetProfileRequest.Gender [optional]
		if err = buffer.ReadUInt8(&setProfileRequest.GenderFlag); err != nil {
			return err
		}
		if setProfileRequest.GenderFlag == OPTIONAL_TRUE {
			if err = buffer.ReadUInt8(&setProfileRequest.Gender); err != nil {
				return err
			}
		}

		// GProtoSetProfileRequest.Avatar [optional]
		if err = buffer.ReadUInt8(&setProfileRequest.AvatarFlag); err != nil {
			return err
		}
		if setProfileRequest.AvatarFlag == OPTIONAL_TRUE {
			if setProfileRequest.Avatar, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoSetProfileRequest.Signature [optional]
		if err = buffer.ReadUInt8(&setProfileRequest.SignatureFlag); err != nil {
			return err
		}
		if setProfileRequest.SignatureFlag == OPTIONAL_TRUE {
			if setProfileRequest.Signature, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoSetProfileRequest.Extern [optional]
		if err = buffer.ReadUInt8(&setProfileRequest.ExternFlag); err != nil {
			return err
		}
		if setProfileRequest.ExternFlag == OPTIONAL_TRUE {
			if setProfileRequest.Extern, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message SetProfileResponse {
//     required uint8 code = 1;
//     optional SelfUser user = 2;
// }

type GProtoSetProfileResponse struct {
	Code     uint8
	User     *GProtoSelfUser // optional
	UserFlag uint8
}

func (setProfileResponse *GProtoSetProfileResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoSetProfileResponse.Code
		if err = buffer.WriteUInt8(setProfileResponse.Code); err != nil {
			return nil, err
		}

		// GProtoSetProfileResponse.User [optional]
		if setProfileResponse.UserFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteStruct(1, setProfileResponse.User); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (setProfileResponse *GProtoSetProfileResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoSetProfileResponse.Code
		if err = buffer.ReadUInt8(&setProfileResponse.Code); err != nil {
			return err
		}

		// GProtoSetProfileResponse.User [optional]
		if err = buffer.ReadUInt8(&setProfileResponse.UserFlag); err != nil {
			return err
		}
		if setProfileResponse.UserFlag == OPTIONAL_TRUE {
			setProfileResponse.User = &GProtoSelfUser{}
			if err = buffer.ReadStruct(1, setProfileResponse.User); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message SetFriendMarknameRequest {
//     required string account = 1;
//     required string markname = 2;
// }

type GProtoSetFriendMarknameRequest struct {
	Account  string
	Markname string
}

func (setFriendMarknameRequest *GProtoSetFriendMarknameRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoSetFriendMarknameRequest.Account
		if err = buffer.WriteString(setFriendMarknameRequest.Account); err != nil {
			return nil, err
		}

		// GProtoSetFriendMarknameRequest.Markname
		if err = buffer.WriteString(setFriendMarknameRequest.Markname); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (setFriendMarknameRequest *GProtoSetFriendMarknameRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoSetFriendMarknameRequest.Account
		if setFriendMarknameRequest.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoSetFriendMarknameRequest.Markname
		if setFriendMarknameRequest.Markname, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

///////////////////////////////////////
// message SetFriendMarknameResponse {
//     required uint8 code = 1;
//     optional OtherUser user = 2;
// }
type GProtoSetFriendMarknameResponse struct {
	Code     uint8
	User     *GProtoOtherUser // optional
	UserFlag uint8
}

func (setFriendMarknameResponse *GProtoSetFriendMarknameResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoSetFriendMarknameResponse.Code
		if err = buffer.WriteUInt8(setFriendMarknameResponse.Code); err != nil {
			return nil, err
		}

		// GProtoSetFriendMarknameResponse.User [optional]
		if setFriendMarknameResponse.UserFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteStruct(1, setFriendMarknameResponse.User); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (setFriendMarknameResponse *GProtoSetFriendMarknameResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoSetFriendMarknameResponse.Code
		if err = buffer.ReadUInt8(&setFriendMarknameResponse.Code); err != nil {
			return err
		}

		// GProtoSetFriendMarknameResponse.User [optional]
		if err = buffer.ReadUInt8(&setFriendMarknameResponse.UserFlag); err != nil {
			return err
		}
		if setFriendMarknameResponse.UserFlag == OPTIONAL_TRUE {
			setFriendMarknameResponse.User = &GProtoOtherUser{}
			if err = buffer.ReadStruct(1, setFriendMarknameResponse.User); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetProfilesRequest {
//     repeated string accounts = 1;
// }

type GProtoGetProfilesRequest struct {
	Accounts []string
}

func (getProfilesRequest *GProtoGetProfilesRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGetProfilesRequest.Accounts [array]
		accountNum := len(getProfilesRequest.Accounts)
		if err := buffer.WriteUInt16(uint16(accountNum)); err != nil {
			return nil, err
		}
		for i := 0; i < accountNum; i++ {
			if err = buffer.WriteString(getProfilesRequest.Accounts[i]); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (getProfilesRequest *GProtoGetProfilesRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoGetProfilesRequest.Accounts [array]
		var accountsNum uint16
		if err = buffer.ReadUInt16(&accountsNum); err != nil {
			return err
		} else {
			getProfilesRequest.Accounts = make([]string, accountsNum, accountsNum)
			for i := 0; i < int(accountsNum); i++ {
				if getProfilesRequest.Accounts[i], err = buffer.ReadString(); err != nil {
					return err
				}
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetProfilesResponse {
//     required uint8 code = 1;
//     repeated OtherUser users = 2;
// }

type GProtoGetProfilesResponse struct {
	Code  uint8
	Users []*GProtoOtherUser
}

func (getProfilesResponse *GProtoGetProfilesResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGetProfilesResponse.Code
		if err = buffer.WriteUInt8(getProfilesResponse.Code); err != nil {
			return nil, err
		}

		// GProtoGetProfilesResponse.Users [array]
		if err = buffer.WriteArray(1, getProfilesResponse.Users); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (getProfilesResponse *GProtoGetProfilesResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoGetProfilesResponse.Code
		if err = buffer.ReadUInt8(&getProfilesResponse.Code); err != nil {
			return err
		}

		// GProtoGetProfilesResponse.Users [array]
		var usersNum uint16
		if err = buffer.ReadUInt16(&usersNum); err != nil {
			return err
		} else {
			getProfilesResponse.Users = make([]*GProtoOtherUser, usersNum, usersNum)
			for i := 0; i < int(usersNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						getProfilesResponse.Users[i] = &GProtoOtherUser{}
						if err = getProfilesResponse.Users[i].Decode(1, bufObj); err != nil {
							return err
						}
					}
				}
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message MessageRequest {
//     required string msgId = 1;
//     required uint8 msgType = 2;

//     required string content = 3;
//     required string from = 4;
//     required string to = 5;

//     optional uint32 file_size = 6;
//     optional uint32 record_time = 7;
//     optional string extern = 8;
// }

const (
	// 0 - 9
	MESSAGE_TYPE_PLAYER_TEXT              = 0
	MESSAGE_TYPE_PLAYER_IMAGE             = 1
	MESSAGE_TYPE_PLAYER_SHORT_VOICE       = 2
	MESSAGE_TYPE_PLAYER_SHORT_VEDEO       = 3
	MESSAGE_TYPE_PLAYER_CUSTOMIZE_ONLINE  = 4 // 自定义格式消息(不会离线存储)
	MESSAGE_TYPE_PLAYER_CUSTOMIZE_OFFLINE = 5 // 自定义格式消息(会离线存储)

	// 10 - 19
	MESSAGE_TYPE_GROUP_TEXT              = 10
	MESSAGE_TYPE_GROUP_IMAGE             = 11
	MESSAGE_TYPE_GROUP_SHORT_VOICE       = 12
	MESSAGE_TYPE_GROUP_SHORT_VEDEO       = 13
	MESSAGE_TYPE_GROUP_CUSTOMIZE_ONLINE  = 14 // 自定义格式消息(不会离线存储)
	MESSAGE_TYPE_GROUP_CUSTOMIZE_OFFLINE = 15 // 自定义格式消息(会离线存储)

	// 20 - 29
	MESSAGE_TYPE_CHATROOM_TEXT             = 20
	MESSAGE_TYPE_CHATROOM_IMAGE            = 21
	MESSAGE_TYPE_CHATROOM_SHORT_VOICE      = 22
	MESSAGE_TYPE_CHATROOM_SHORT_VEDEO      = 23
	MESSAGE_TYPE_CHATROOM_CUSTOMIZE_ONLINE = 24 // 自定义格式消息(不会离线存储)

	// 30 - 39
	MESSAGE_TYPE_BROADCAST_TEXT             = 30 // (不会离线存储)
	MESSAGE_TYPE_BROADCAST_IMAGE            = 31 // (不会离线存储)
	MESSAGE_TYPE_BROADCAST_SHORT_VOICE      = 32 // (不会离线存储)
	MESSAGE_TYPE_BROADCAST_SHORT_VIDEO      = 33 // (不会离线存储)
	MESSAGE_TYPE_BROADCAST_CUSTOMIZE_ONLINE = 34 // 自定义格式消息(不会离线存储)

)

type GProtoMessageRequest struct {
	MsgId                  string
	MsgType                uint8
	Content                string
	From                   string
	To                     string
	FileSize               uint32 // optional
	FileSizeOptionalFlag   uint8
	RecordTime             uint32 // optional
	RecordTimeOptionalFlag uint8
	Extern                 string // optional
	ExternOptionalFlag     uint8
}

func (messageRequest *GProtoMessageRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoMessageRequest.MsgId
		if err = buffer.WriteString(messageRequest.MsgId); err != nil {
			return nil, err
		}

		// GProtoMessageRequest.MsgType
		if err = buffer.WriteUInt8(messageRequest.MsgType); err != nil {
			return nil, err
		}

		// GProtoMessageRequest.Content
		if err = buffer.WriteString(messageRequest.Content); err != nil {
			return nil, err
		}

		// GProtoMessageRequest.From
		if err = buffer.WriteString(messageRequest.From); err != nil {
			return nil, err
		}

		// GProtoMessageRequest.To
		if err = buffer.WriteString(messageRequest.To); err != nil {
			return nil, err
		}

		// GProtoMessageRequest.FileSize [optional]
		if messageRequest.FileSizeOptionalFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteUInt32(messageRequest.FileSize); err != nil {
				return nil, err
			}
		}

		// GProtoMessageRequest.RecordTime [optional]
		if messageRequest.RecordTimeOptionalFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteUInt32(messageRequest.RecordTime); err != nil {
				return nil, err
			}
		}

		// GProtoMessageRequest.Extern [optional]
		if messageRequest.ExternOptionalFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(messageRequest.Extern); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (messageRequest *GProtoMessageRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoMessageRequest.MsgId
		if messageRequest.MsgId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoMessageRequest.MsgType
		if err = buffer.ReadUInt8(&messageRequest.MsgType); err != nil {
			return err
		}

		// GProtoMessageRequest.Content
		if messageRequest.Content, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoMessageRequest.From
		if messageRequest.From, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoMessageRequest.To
		if messageRequest.To, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoMessageRequest.FileSize [optional]
		if err = buffer.ReadUInt8(&messageRequest.FileSizeOptionalFlag); err != nil {
			return err
		}
		if messageRequest.FileSizeOptionalFlag == OPTIONAL_TRUE {
			if err = buffer.ReadUInt32(&messageRequest.FileSize); err != nil {
				return err
			}
		}

		// GProtoMessageRequest.RecordTime [optional]
		if err = buffer.ReadUInt8(&messageRequest.RecordTimeOptionalFlag); err != nil {
			return err
		}
		if messageRequest.RecordTimeOptionalFlag == OPTIONAL_TRUE {
			if err = buffer.ReadUInt32(&messageRequest.RecordTime); err != nil {
				return err
			}
		}

		// GProtoMessageRequest.Extern [optional]
		if err = buffer.ReadUInt8(&messageRequest.ExternOptionalFlag); err != nil {
			return err
		}
		if messageRequest.ExternOptionalFlag == OPTIONAL_TRUE {
			if messageRequest.Extern, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message MessageResponse {
//     required uint8 code = 1;
//     required string msgId = 2;
//     optional uint8 msgType = 3;
//     optional string content = 4;
// }

type GProtoMessageResponse struct {
	Code        uint8
	MsgId       string
	MsgType     uint8 // optional
	MsgTypeFlag uint8
	Content     string // optional
	ContentFlag uint8
}

func (messageResponse *GProtoMessageResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoMessageResponse.Code
		if err = buffer.WriteUInt8(messageResponse.Code); err != nil {
			return nil, err
		}

		// GProtoMessageResponse.MsgId
		if err = buffer.WriteString(messageResponse.MsgId); err != nil {
			return nil, err
		}

		// GProtoMessageResponse.MsgType [optional]
		if messageResponse.MsgTypeFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteUInt8(messageResponse.MsgType); err != nil {
				return nil, err
			}
		}

		// GProtoMessageResponse.ContentFlag [optional]
		if messageResponse.ContentFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(messageResponse.Content); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (messageResponse *GProtoMessageResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoMessageResponse.Code
		if err = buffer.ReadUInt8(&messageResponse.Code); err != nil {
			return err
		}

		// GProtoMessageResponse.MsgId
		if messageResponse.MsgId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoMessageResponse.MsgType [optional]
		if err = buffer.ReadUInt8(&messageResponse.MsgTypeFlag); err != nil {
			return err
		}
		if messageResponse.MsgTypeFlag == OPTIONAL_TRUE {
			if err = buffer.ReadUInt8(&messageResponse.MsgType); err != nil {
				return err
			}
		}

		// GProtoMessageResponse.ContentFlag [optional]
		if err = buffer.ReadUInt8(&messageResponse.ContentFlag); err != nil {
			return err
		}
		if messageResponse.ContentFlag == OPTIONAL_TRUE {
			if messageResponse.Content, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message MessageNotify {
//     required string msgId = 1;
//     required uint8 msgType = 2;

//     required string content = 3;
//     required string from = 4;
//     required string to = 5;

//     required uint64 send_time = 6;
//	   required uint64 send_time_ack = 7;

//     optional uint32 file_size = 8;
//     optional uint32 record_time = 9;
//     optional string extern = 10;
// }

type GProtoMessageNotify struct {
	MsgId                  string
	MsgType                uint8
	Content                string
	From                   string
	To                     string
	SendTime               uint64
	SendTimeAck            uint64
	FileSize               uint32 // optional
	FileSizeOptionalFlag   uint8
	RecordTime             uint32 // optional
	RecordTimeOptionalFlag uint8
	Extern                 string // optional
	ExternOptionalFlag     uint8

	FixTo string // 这个字段不发送, 只是内部使用(群组的时候To逻辑需要做修改)
}

func (messageNotify *GProtoMessageNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoMessageNotify.MsgId
		if err = buffer.WriteString(messageNotify.MsgId); err != nil {
			return nil, err
		}

		// GProtoMessageNotify.MsgType
		if err = buffer.WriteUInt8(messageNotify.MsgType); err != nil {
			return nil, err
		}

		// GProtoMessageNotify.Content
		if err = buffer.WriteString(messageNotify.Content); err != nil {
			return nil, err
		}

		// GProtoMessageNotify.From
		if err = buffer.WriteString(messageNotify.From); err != nil {
			return nil, err
		}

		// GProtoMessageNotify.To
		if err = buffer.WriteString(messageNotify.To); err != nil {
			return nil, err
		}

		// GProtoMessageNotify.SendTime
		if err = buffer.WriteUInt64(messageNotify.SendTime); err != nil {
			return nil, err
		}

		// GProtoMessageNotify.SendTimeAck
		if err = buffer.WriteUInt64(messageNotify.SendTimeAck); err != nil {
			return nil, err
		}

		// GProtoMessageNotify.FileSize [optional]
		if messageNotify.FileSizeOptionalFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteUInt32(messageNotify.FileSize); err != nil {
				return nil, err
			}
		}

		// GProtoMessageNotify.RecordTime [optional]
		if messageNotify.RecordTimeOptionalFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteUInt32(messageNotify.RecordTime); err != nil {
				return nil, err
			}
		}

		// GProtoMessageNotify.Extern [optional]
		if messageNotify.ExternOptionalFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(messageNotify.Extern); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (messageNotify *GProtoMessageNotify) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoMessageNotify.MsgId
		if messageNotify.MsgId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoMessageNotify.MsgType
		if err = buffer.ReadUInt8(&messageNotify.MsgType); err != nil {
			return err
		}

		// GProtoMessageNotify.Content
		if messageNotify.Content, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoMessageNotify.From
		if messageNotify.From, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoMessageNotify.To
		if messageNotify.To, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoMessageNotify.SendTime
		if err = buffer.ReadUInt64(&messageNotify.SendTime); err != nil {
			return err
		}

		// GProtoMessageNotify.SendTimeAck
		if err = buffer.ReadUInt64(&messageNotify.SendTimeAck); err != nil {
			return err
		}

		// GProtoMessageNotify.FileSize [optional]
		if err = buffer.ReadUInt8(&messageNotify.FileSizeOptionalFlag); err != nil {
			return err
		}
		if messageNotify.FileSizeOptionalFlag == OPTIONAL_TRUE {
			if err = buffer.ReadUInt32(&messageNotify.FileSize); err != nil {
				return err
			}
		}

		// GProtoMessageNotify.RecordTime [optional]
		if err = buffer.ReadUInt8(&messageNotify.RecordTimeOptionalFlag); err != nil {
			return err
		}
		if messageNotify.RecordTimeOptionalFlag == OPTIONAL_TRUE {
			if err = buffer.ReadUInt32(&messageNotify.RecordTime); err != nil {
				return err
			}
		}

		// GProtoMessageNotify.Extern [optional]
		if err = buffer.ReadUInt8(&messageNotify.ExternOptionalFlag); err != nil {
			return err
		}
		if messageNotify.ExternOptionalFlag == OPTIONAL_TRUE {
			if messageNotify.Extern, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message MessageAck {
//     required string msgId = 1;
// }

type GProtoMessageAck struct {
	MsgId string
}

func (messageAck *GProtoMessageAck) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoMessageAck.MsgId
		if err = buffer.WriteString(messageAck.MsgId); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (messageAck *GProtoMessageAck) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoMessageAck.MsgId
		if messageAck.MsgId, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message AddFriendRequest {
//     required string account = 1;
//     required string desc = 2;
// }

type GProtoAddFriendRequest struct {
	Account string
	Desc    string
}

func (addFriendRequest *GProtoAddFriendRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoAddFriendRequest.Account
		if err = buffer.WriteString(addFriendRequest.Account); err != nil {
			return nil, err
		}

		// GProtoAddFriendRequest.Desc
		if err = buffer.WriteString(addFriendRequest.Desc); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (addFriendRequest *GProtoAddFriendRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoAddFriendRequest.Account
		if addFriendRequest.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAddFriendRequest.Desc
		if addFriendRequest.Desc, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message AddFriendResponse {
//     required uint8 code = 1;
//     required string account = 2;
//     required uint8 flag = 3;
// }

type GProtoAddFriendResponse struct {
	Code    uint8
	Account string
	Flag    uint8
}

func (addFriendResponse *GProtoAddFriendResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoAddFriendResponse.Code
		if err = buffer.WriteUInt8(addFriendResponse.Code); err != nil {
			return nil, err
		}

		// GProtoAddFriendResponse.Account
		if err = buffer.WriteString(addFriendResponse.Account); err != nil {
			return nil, err
		}

		// GProtoAddFriendResponse.Flag
		if err = buffer.WriteUInt8(addFriendResponse.Flag); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (addFriendResponse *GProtoAddFriendResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoAddFriendResponse.Code
		if err = buffer.ReadUInt8(&addFriendResponse.Code); err != nil {
			return err
		}

		// GProtoAddFriendResponse.Account
		if addFriendResponse.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAddFriendResponse.Flag
		if err = buffer.ReadUInt8(&addFriendResponse.Flag); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message AddFriendNotify {
//     required OtherUser user = 1;
//     required string desc = 2;
//     required uint8 flag = 3;
//     required uint64 send_time = 4;
//     required uint8 is_offline = 5;
// }

type GProtoAddFriendNotify struct {
	User      *GProtoOtherUser
	Desc      string
	Flag      uint8
	SendTime  uint64
	IsOffline uint8
}

func (addFriendNotify *GProtoAddFriendNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoAddFriendNotify.User
		if err = buffer.WriteStruct(1, addFriendNotify.User); err != nil {
			return nil, err
		}

		// GProtoAddFriendNotify.Desc
		if err = buffer.WriteString(addFriendNotify.Desc); err != nil {
			return nil, err
		}

		// GProtoAddFriendNotify.Flag
		if err = buffer.WriteUInt8(addFriendNotify.Flag); err != nil {
			return nil, err
		}

		// GProtoAddFriendNotify.SendTime
		if err = buffer.WriteUInt64(addFriendNotify.SendTime); err != nil {
			return nil, err
		}

		// GProtoAddFriendNotify.IsOffline
		if err = buffer.WriteUInt8(addFriendNotify.IsOffline); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (addFriendNotify *GProtoAddFriendNotify) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoAddFriendNotify.User
		addFriendNotify.User = &GProtoOtherUser{}
		if err = buffer.ReadStruct(1, addFriendNotify.User); err != nil {
			return err
		}

		// GProtoAddFriendNotify.Desc
		if addFriendNotify.Desc, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAddFriendNotify.Flag
		if err = buffer.ReadUInt8(&addFriendNotify.Flag); err != nil {
			return err
		}

		// GProtoAddFriendNotify.SendTime
		if err = buffer.ReadUInt64(&addFriendNotify.SendTime); err != nil {
			return err
		}

		// GProtoAddFriendNotify.IsOffline
		if err = buffer.ReadUInt8(&addFriendNotify.IsOffline); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message AnswerFriendRequest {
//     required string account = 1;
//     required string desc = 2;
//     required uint8 ops = 3;
// }

const (
	ANSWER_FRIEND_OPS_DENY          = 1
	ANSWER_FRIEND_OPS_AGREE         = 2
	ANSWER_FRIEND_OPS_AGREE_AND_ADD = 3
)

type GProtoAnswerFriendRequest struct {
	Account string
	Desc    string
	Ops     uint8 // 1拒绝, 2同意, 3同意并添加
}

func (answerFriendRequest *GProtoAnswerFriendRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoAnswerFriendRequest.Account
		if err = buffer.WriteString(answerFriendRequest.Account); err != nil {
			return nil, err
		}

		// GProtoAnswerFriendRequest.Desc
		if err = buffer.WriteString(answerFriendRequest.Desc); err != nil {
			return nil, err
		}

		// GProtoAnswerFriendRequest.Ops
		if err = buffer.WriteUInt8(answerFriendRequest.Ops); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (answerFriendRequest *GProtoAnswerFriendRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoAnswerFriendRequest.Account
		if answerFriendRequest.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAnswerFriendRequest.Desc
		if answerFriendRequest.Desc, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAnswerFriendRequest.Ops
		if err = buffer.ReadUInt8(&answerFriendRequest.Ops); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message AnswerFriendResponse {
//     required uint8 code = 1;
//     required string account = 2;
//     required uint8 flag = 3;
//     required uint8 ops = 4;
// }

type GProtoAnswerFriendResponse struct {
	Code    uint8
	Account string
	Flag    uint8
	Ops     uint8
}

func (answerFriendResponse *GProtoAnswerFriendResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoAnswerFriendResponse.Code
		if err = buffer.WriteUInt8(answerFriendResponse.Code); err != nil {
			return nil, err
		}

		// GProtoAnswerFriendResponse.Account
		if err = buffer.WriteString(answerFriendResponse.Account); err != nil {
			return nil, err
		}

		// GProtoAnswerFriendResponse.Flag
		if err = buffer.WriteUInt8(answerFriendResponse.Flag); err != nil {
			return nil, err
		}

		// GProtoAnswerFriendResponse.Ops
		if err = buffer.WriteUInt8(answerFriendResponse.Ops); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (answerFriendResponse *GProtoAnswerFriendResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoAnswerFriendResponse.Code
		if err = buffer.ReadUInt8(&answerFriendResponse.Code); err != nil {
			return err
		}

		// GProtoAnswerFriendResponse.Account
		if answerFriendResponse.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAnswerFriendResponse.Flag
		if err = buffer.ReadUInt8(&answerFriendResponse.Flag); err != nil {
			return err
		}

		// GProtoAnswerFriendResponse.Ops
		if err = buffer.ReadUInt8(&answerFriendResponse.Ops); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message AnswerFriendNotify {
//     required OtherUser user = 1;
//     required string desc = 2;
//     required uint8 flag = 3;
//     required uint8 ops = 4;
//     required uint64 send_time = 5;
//     required uint8 is_offline = 6;
// }

type GProtoAnswerFriendNotify struct {
	User      *GProtoOtherUser
	Desc      string
	Flag      uint8
	Ops       uint8
	SendTime  uint64
	IsOffline uint8
}

func (answerFriendNotify *GProtoAnswerFriendNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoAnswerFriendNotify.User
		if err = buffer.WriteStruct(1, answerFriendNotify.User); err != nil {
			return nil, err
		}

		// GProtoAnswerFriendNotify.Desc
		if err = buffer.WriteString(answerFriendNotify.Desc); err != nil {
			return nil, err
		}

		// GProtoAnswerFriendNotify.Flag
		if err = buffer.WriteUInt8(answerFriendNotify.Flag); err != nil {
			return nil, err
		}

		// GProtoAnswerFriendNotify.Ops
		if err = buffer.WriteUInt8(answerFriendNotify.Ops); err != nil {
			return nil, err
		}

		// GProtoAnswerFriendNotify.SendTime
		if err = buffer.WriteUInt64(answerFriendNotify.SendTime); err != nil {
			return nil, err
		}

		// GProtoAnswerFriendNotify.IsOffline
		if err = buffer.WriteUInt8(answerFriendNotify.IsOffline); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (answerFriendNotify *GProtoAnswerFriendNotify) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoAnswerFriendNotify.User
		answerFriendNotify.User = &GProtoOtherUser{}
		if err = buffer.ReadStruct(1, answerFriendNotify.User); err != nil {
			return err
		}

		// GProtoAnswerFriendNotify.Desc
		if answerFriendNotify.Desc, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAnswerFriendNotify.Flag
		if err = buffer.ReadUInt8(&answerFriendNotify.Flag); err != nil {
			return err
		}

		// GProtoAnswerFriendNotify.Ops
		if err = buffer.ReadUInt8(&answerFriendNotify.Ops); err != nil {
			return err
		}

		// GProtoAnswerFriendNotify.SendTime
		if err = buffer.ReadUInt64(&answerFriendNotify.SendTime); err != nil {
			return err
		}

		// GProtoAnswerFriendNotify.IsOffline
		if err = buffer.ReadUInt8(&answerFriendNotify.IsOffline); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message DelFriendRequest {
//     required string account = 1;
// 	   required uint8 ops = 2;
// }

const (
	DEL_FRIEND_OPS_NORMAL = 1
	DEL_FRIEND_OPS_BOTH   = 2
)

type GProtoDelFriendRequest struct {
	Account string
	Ops     uint8 // 1删除单向好友; 2把我也从对方列表中删除
}

func (delFriendRequest *GProtoDelFriendRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoDelFriendRequest.Account
		if err = buffer.WriteString(delFriendRequest.Account); err != nil {
			return nil, err
		}

		// GProtoDelFriendRequest.Ops
		if err = buffer.WriteUInt8(delFriendRequest.Ops); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (delFriendRequest *GProtoDelFriendRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoDelFriendRequest.Account
		if delFriendRequest.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoDelFriendRequest.Ops
		if err = buffer.ReadUInt8(&delFriendRequest.Ops); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message DelFriendResponse {
//     required uint8 code = 1;
//     required string account = 2;
//     required uint8 flag = 3;
// }

type GProtoDelFriendResponse struct {
	Code    uint8
	Account string
	Flag    uint8
}

func (delFriendResponse *GProtoDelFriendResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoDelFriendResponse.Code
		if err = buffer.WriteUInt8(delFriendResponse.Code); err != nil {
			return nil, err
		}

		// GProtoDelFriendResponse.Account
		if err = buffer.WriteString(delFriendResponse.Account); err != nil {
			return nil, err
		}

		// GProtoDelFriendResponse.Flag
		if err = buffer.WriteUInt8(delFriendResponse.Flag); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (delFriendResponse *GProtoDelFriendResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoDelFriendResponse.Code
		if err = buffer.ReadUInt8(&delFriendResponse.Code); err != nil {
			return err
		}

		// GProtoDelFriendResponse.Account
		if delFriendResponse.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoDelFriendResponse.Flag
		if err = buffer.ReadUInt8(&delFriendResponse.Flag); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message DelFriendNotify {
//     required OtherUser user = 1;
//     required uint8 flag = 2;
//     required uint64 send_time = 3;
//     required uint8 is_offline = 4;
// }

type GProtoDelFriendNotify struct {
	User      *GProtoOtherUser
	Flag      uint8
	SendTime  uint64
	IsOffline uint8
}

func (elFriendNotify *GProtoDelFriendNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoDelFriendNotify.User
		if err = buffer.WriteStruct(1, elFriendNotify.User); err != nil {
			return nil, err
		}

		// GProtoDelFriendNotify.Flag
		if err = buffer.WriteUInt8(elFriendNotify.Flag); err != nil {
			return nil, err
		}

		// GProtoDelFriendNotify.SendTime
		if err = buffer.WriteUInt64(elFriendNotify.SendTime); err != nil {
			return nil, err
		}

		// GProtoDelFriendNotify.IsOffline
		if err = buffer.WriteUInt8(elFriendNotify.IsOffline); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (elFriendNotify *GProtoDelFriendNotify) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoDelFriendNotify.User
		elFriendNotify.User = &GProtoOtherUser{}
		if err = buffer.ReadStruct(1, elFriendNotify.User); err != nil {
			return err
		}

		// GProtoDelFriendNotify.Flag
		if err = buffer.ReadUInt8(&elFriendNotify.Flag); err != nil {
			return err
		}

		// GProtoDelFriendNotify.SendTime
		if err = buffer.ReadUInt64(&elFriendNotify.SendTime); err != nil {
			return err
		}

		// GProtoDelFriendNotify.IsOffline
		if err = buffer.ReadUInt8(&elFriendNotify.IsOffline); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetFriendsRequest {
//     required string customVersion = 1;
// }

type GProtoGetFriendsRequest struct {
	CustomVersion string
}

func (getFriendsRequest *GProtoGetFriendsRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGetFriendsRequest.CustomVersion
		if err = buffer.WriteString(getFriendsRequest.CustomVersion); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (getFriendsRequest *GProtoGetFriendsRequest) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoGetFriendsRequest.CustomVersion
		if getFriendsRequest.CustomVersion, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetFriendsResponse {
//     required uint8 code = 1;
//     repeated Friend friends = 2;
//     optional string CustomVersion = 3;
// }

type GProtoGetFriendsResponse struct {
	Code              uint8
	Friends           []*GProtoFriend
	CustomVersion     string // optional
	CustomVersionFlag uint8
}

func (getFriendsResponse *GProtoGetFriendsResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGetFriendsResponse.Code
		if err = buffer.WriteUInt8(getFriendsResponse.Code); err != nil {
			return nil, err
		}

		// GProtoGetFriendsResponse.Friends [array]
		if err = buffer.WriteArray(1, getFriendsResponse.Friends); err != nil {
			return nil, err
		}

		// GProtoGetFriendsResponse.CustomVersion [optional]
		if getFriendsResponse.CustomVersionFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(getFriendsResponse.CustomVersion); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (getFriendsResponse *GProtoGetFriendsResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoGetFriendsResponse.Code
		if err = buffer.ReadUInt8(&getFriendsResponse.Code); err != nil {
			return err
		}

		// GProtoGetFriendsResponse.Friends [array]
		var friendsNum uint16
		if err = buffer.ReadUInt16(&friendsNum); err != nil {
			return err
		} else {
			getFriendsResponse.Friends = make([]*GProtoFriend, friendsNum, friendsNum)
			for i := 0; i < int(friendsNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						getFriendsResponse.Friends[i] = &GProtoFriend{}
						if err = getFriendsResponse.Friends[i].Decode(1, bufObj); err != nil {
							return err
						}
					}
				}
			}
		}

		// GProtoGetFriendsResponse.CustomVersion [optional]
		if err = buffer.ReadUInt8(&getFriendsResponse.CustomVersionFlag); err != nil {
			return err
		}
		if getFriendsResponse.CustomVersionFlag == OPTIONAL_TRUE {
			if getFriendsResponse.CustomVersion, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetChatRoomInfoRequest {
//     required string chatroom_id = 1;
// }

type GProtoGetChatRoomInfoRequest struct {
	ChatRoomId string
}

func (getChatRoomInfoRequest *GProtoGetChatRoomInfoRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGetChatRoomInfoRequest.ChatRoomId
		if err = buffer.WriteString(getChatRoomInfoRequest.ChatRoomId); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (getChatRoomInfoRequest *GProtoGetChatRoomInfoRequest) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoGetChatRoomInfoRequest.ChatRoomId
		if getChatRoomInfoRequest.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetChatRoomInfoResponse {
//     required uint8 code = 1;
//     optional ChatRoomInfo info = 2;
// }

type GProtoGetChatRoomInfoResponse struct {
	Code     uint8
	Info     *GProtoChatRoomInfo // optional
	InfoFlag uint8
}

func (getChatRoomInfoResponse *GProtoGetChatRoomInfoResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGetChatRoomInfoResponse.Code
		if err = buffer.WriteUInt8(getChatRoomInfoResponse.Code); err != nil {
			return nil, err
		}

		// GProtoGetChatRoomInfoResponse.Info [optional]
		if getChatRoomInfoResponse.InfoFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteStruct(1, getChatRoomInfoResponse.Info); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (getChatRoomInfoResponse *GProtoGetChatRoomInfoResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoGetChatRoomInfoResponse.Code
		if err = buffer.ReadUInt8(&getChatRoomInfoResponse.Code); err != nil {
			return err
		}

		// GProtoGetChatRoomInfoResponse.Info [optional]
		if err = buffer.ReadUInt8(&getChatRoomInfoResponse.InfoFlag); err != nil {
			return err
		}
		if getChatRoomInfoResponse.InfoFlag == OPTIONAL_TRUE {
			getChatRoomInfoResponse.Info = &GProtoChatRoomInfo{}
			if err = buffer.ReadStruct(1, getChatRoomInfoResponse.Info); err != nil {
				return err
			}

		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message EnterChatRoomRequest {
//     required string chatroom_id = 1;
// }

type GProtoEnterChatRoomRequest struct {
	ChatRoomId string
}

func (enterChatRoomRequest *GProtoEnterChatRoomRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoEnterChatRoomRequest.ChatRoomId
		if err = buffer.WriteString(enterChatRoomRequest.ChatRoomId); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (enterChatRoomRequest *GProtoEnterChatRoomRequest) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoEnterChatRoomRequest.ChatRoomId
		if enterChatRoomRequest.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message EnterChatRoomResponse {
//     required uint8 code = 1;
//     optional ChatRoomInfo info = 2;
//     repeated OtherUser users = 3;
// }

type GProtoEnterChatRoomResponse struct {
	Code     uint8
	Info     *GProtoChatRoomInfo // optional
	InfoFlag uint8
	Users    []*GProtoOtherUser
}

func (enterChatRoomResponse *GProtoEnterChatRoomResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoEnterChatRoomResponse.Code
		if err = buffer.WriteUInt8(enterChatRoomResponse.Code); err != nil {
			return nil, err
		}

		// GProtoEnterChatRoomResponse.Info [optional]
		if enterChatRoomResponse.InfoFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteStruct(1, enterChatRoomResponse.Info); err != nil {
				return nil, err
			}
		}

		// GProtoEnterChatRoomResponse.Users [array]
		if err = buffer.WriteArray(1, enterChatRoomResponse.Users); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (enterChatRoomResponse *GProtoEnterChatRoomResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoEnterChatRoomResponse.Code
		if err = buffer.ReadUInt8(&enterChatRoomResponse.Code); err != nil {
			return err
		}

		// GProtoEnterChatRoomResponse.Info [optional]
		if err = buffer.ReadUInt8(&enterChatRoomResponse.InfoFlag); err != nil {
			return err
		}
		if enterChatRoomResponse.InfoFlag == OPTIONAL_TRUE {
			enterChatRoomResponse.Info = &GProtoChatRoomInfo{}
			if err = buffer.ReadStruct(1, enterChatRoomResponse.Info); err != nil {
				return err
			}

		}

		// GProtoEnterChatRoomResponse.Users [array]
		var usersNum uint16
		if err = buffer.ReadUInt16(&usersNum); err != nil {
			return err
		} else {
			enterChatRoomResponse.Users = make([]*GProtoOtherUser, usersNum, usersNum)
			for i := 0; i < int(usersNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						enterChatRoomResponse.Users[i] = &GProtoOtherUser{}
						if err = enterChatRoomResponse.Users[i].Decode(1, bufObj); err != nil {
							return err
						}
					}
				}
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message EnterChatRoomNotify {
//     required string chatroom_id = 1;
//     required OtherUser user = 2;
// }

type GProtoEnterChatRoomNotify struct {
	ChatRoomId string
	User       *GProtoOtherUser
}

func (enterChatRoomNotify *GProtoEnterChatRoomNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoEnterChatRoomNotify.ChatRoomId
		if err = buffer.WriteString(enterChatRoomNotify.ChatRoomId); err != nil {
			return nil, err
		}

		// GProtoEnterChatRoomNotify.User
		if err = buffer.WriteStruct(1, enterChatRoomNotify.User); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (enterChatRoomNotify *GProtoEnterChatRoomNotify) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoEnterChatRoomNotify.ChatRoomId
		if enterChatRoomNotify.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoEnterChatRoomNotify.User
		enterChatRoomNotify.User = &GProtoOtherUser{}
		if err = buffer.ReadStruct(1, enterChatRoomNotify.User); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message LeaveChatRoomRequest {
//     required string chatroom_id = 1;
// }

type GProtoLeaveChatRoomRequest struct {
	ChatRoomId string
}

func (leaveChatRoomRequest *GProtoLeaveChatRoomRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoLeaveChatRoomRequest.ChatRoomId
		if err = buffer.WriteString(leaveChatRoomRequest.ChatRoomId); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (leaveChatRoomRequest *GProtoLeaveChatRoomRequest) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoLeaveChatRoomRequest.ChatRoomId
		if leaveChatRoomRequest.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message LeaveChatRoomResponse {
//     required uint8 code = 1;
//     optional string chatroom_id = 2;
// }

type GProtoLeaveChatRoomResponse struct {
	Code           uint8
	ChatRoomId     string // optional
	ChatRoomIdFlag uint8
}

func (leaveChatRoomResponse *GProtoLeaveChatRoomResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoLeaveChatRoomResponse.Code
		if err = buffer.WriteUInt8(leaveChatRoomResponse.Code); err != nil {
			return nil, err
		}

		// GProtoLeaveChatRoomResponse.ChatRoomId [optional]
		if leaveChatRoomResponse.ChatRoomIdFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(leaveChatRoomResponse.ChatRoomId); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (leaveChatRoomResponse *GProtoLeaveChatRoomResponse) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoLeaveChatRoomResponse.Code
		if err = buffer.ReadUInt8(&leaveChatRoomResponse.Code); err != nil {
			return err
		}

		// GProtoLeaveChatRoomResponse.ChatRoomId [optional]
		if err = buffer.ReadUInt8(&leaveChatRoomResponse.ChatRoomIdFlag); err != nil {
			return err
		}
		if leaveChatRoomResponse.ChatRoomIdFlag == OPTIONAL_TRUE {
			if leaveChatRoomResponse.ChatRoomId, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message LeaveChatRoomNotify {
//     required string chatroom_id = 1;
//     required string account = 2;
// }

type GProtoLeaveChatRoomNotify struct {
	ChatRoomId string
	Account    string
}

func (leaveChatRoomNotify *GProtoLeaveChatRoomNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoLeaveChatRoomNotify.ChatRoomId
		if err = buffer.WriteString(leaveChatRoomNotify.ChatRoomId); err != nil {
			return nil, err
		}

		// GProtoLeaveChatRoomNotify.Account
		if err = buffer.WriteString(leaveChatRoomNotify.Account); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (leaveChatRoomNotify *GProtoLeaveChatRoomNotify) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoLeaveChatRoomNotify.ChatRoomId
		if leaveChatRoomNotify.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoLeaveChatRoomNotify.Account
		if leaveChatRoomNotify.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message ChatRoomKickoffNotify {
//     required string chatroom_id = 1;
//     required string account = 2;
// }
type GProtoChatRoomKickoffNotify struct {
	ChatRoomId string
	Account    string
}

func (chatRoomKickoffNotify *GProtoChatRoomKickoffNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoChatRoomKickoffNotify.ChatRoomId
		if err = buffer.WriteString(chatRoomKickoffNotify.ChatRoomId); err != nil {
			return nil, err
		}

		// GProtoChatRoomKickoffNotify.Account
		if err = buffer.WriteString(chatRoomKickoffNotify.Account); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (chatRoomKickoffNotify *GProtoChatRoomKickoffNotify) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoChatRoomKickoffNotify.ChatRoomId
		if chatRoomKickoffNotify.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoChatRoomKickoffNotify.Account
		if chatRoomKickoffNotify.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message CreateGroupRequest {
//     optional string group_id = 1;
//     optional string name = 2;
//     optional string avatar = 3;
//     optional string description = 4;
//     optional uint16 max_user = 5;
//     optional string data = 6;
//     optional uint8 flag = 7;
//     repeated string invite_tos = 8;
// }

type GProtoCreateGroupRequest struct {
	GroupId         string // optional
	GroupIdFlag     uint8
	Name            string // optional
	NameFlag        uint8
	Avatar          string // optional
	AvatarFlag      uint8
	Description     string // optional
	DescriptionFlag uint8
	MaxUser         uint16 // optional
	MaxUserFlag     uint8
	Data            string // optional
	DataFlag        uint8
	Flag            uint8 // optional
	FlagFlag        uint8
	InviteTos       []string
}

func (createGroupRequest *GProtoCreateGroupRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoCreateGroupRequest.GroupId [optional]
		if createGroupRequest.GroupIdFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(createGroupRequest.GroupId); err != nil {
				return nil, err
			}
		}

		// GProtoCreateGroupRequest.Name [optional]
		if createGroupRequest.NameFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(createGroupRequest.Name); err != nil {
				return nil, err
			}
		}

		// GProtoCreateGroupRequest.Avatar [optional]
		if createGroupRequest.AvatarFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(createGroupRequest.Avatar); err != nil {
				return nil, err
			}
		}

		// GProtoCreateGroupRequest.Description [optional]
		if createGroupRequest.DescriptionFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(createGroupRequest.Description); err != nil {
				return nil, err
			}
		}

		// GProtoCreateGroupRequest.MaxUser [optional]
		if createGroupRequest.MaxUserFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteUInt16(createGroupRequest.MaxUser); err != nil {
				return nil, err
			}
		}

		// GProtoCreateGroupRequest.Data [optional]
		if createGroupRequest.DataFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(createGroupRequest.Data); err != nil {
				return nil, err
			}
		}

		// GProtoCreateGroupRequest.Flag [optional]
		if createGroupRequest.FlagFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteUInt8(createGroupRequest.Flag); err != nil {
				return nil, err
			}
		}

		// GProtoCreateGroupRequest.InviteTos [array]
		inviteTosNum := len(createGroupRequest.InviteTos)
		if err := buffer.WriteUInt16(uint16(inviteTosNum)); err != nil {
			return nil, err
		}
		for i := 0; i < inviteTosNum; i++ {
			if err = buffer.WriteString(createGroupRequest.InviteTos[i]); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (createGroupRequest *GProtoCreateGroupRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoCreateGroupRequest.GroupId [optional]
		if err = buffer.ReadUInt8(&createGroupRequest.GroupIdFlag); err != nil {
			return err
		}
		if createGroupRequest.GroupIdFlag == OPTIONAL_TRUE {
			if createGroupRequest.GroupId, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoCreateGroupRequest.Name [optional]
		if err = buffer.ReadUInt8(&createGroupRequest.NameFlag); err != nil {
			return err
		}
		if createGroupRequest.NameFlag == OPTIONAL_TRUE {
			if createGroupRequest.Name, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoCreateGroupRequest.Avatar [optional]
		if err = buffer.ReadUInt8(&createGroupRequest.AvatarFlag); err != nil {
			return err
		}
		if createGroupRequest.AvatarFlag == OPTIONAL_TRUE {
			if createGroupRequest.Avatar, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoCreateGroupRequest.Description [optional]
		if err = buffer.ReadUInt8(&createGroupRequest.DescriptionFlag); err != nil {
			return err
		}
		if createGroupRequest.DescriptionFlag == OPTIONAL_TRUE {
			if createGroupRequest.Description, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoCreateGroupRequest.MaxUser [optional]
		if err = buffer.ReadUInt8(&createGroupRequest.MaxUserFlag); err != nil {
			return err
		}
		if createGroupRequest.MaxUserFlag == OPTIONAL_TRUE {
			if err = buffer.ReadUInt16(&createGroupRequest.MaxUser); err != nil {
				return err
			}
		}

		// GProtoCreateGroupRequest.Data [optional]
		if err = buffer.ReadUInt8(&createGroupRequest.DataFlag); err != nil {
			return err
		}
		if createGroupRequest.DataFlag == OPTIONAL_TRUE {
			if createGroupRequest.Data, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoCreateGroupRequest.Flag [optional]
		if err = buffer.ReadUInt8(&createGroupRequest.FlagFlag); err != nil {
			return err
		}
		if createGroupRequest.FlagFlag == OPTIONAL_TRUE {
			if err = buffer.ReadUInt8(&createGroupRequest.Flag); err != nil {
				return err
			}
		}

		// GProtoCreateGroupRequest.InviteTos [array]
		var inviteTosNum uint16
		if err = buffer.ReadUInt16(&inviteTosNum); err != nil {
			return err
		} else {
			createGroupRequest.InviteTos = make([]string, inviteTosNum, inviteTosNum)
			for i := 0; i < int(inviteTosNum); i++ {
				if createGroupRequest.InviteTos[i], err = buffer.ReadString(); err != nil {
					return err
				}
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message CreateGroupResponse {
//     required uint8 code = 1;
//     optional GroupInfo info = 2;
// }

type GProtoCreateGroupResponse struct {
	Code     uint8
	Info     *GProtoGroupInfo // optional
	InfoFlag uint8
}

func (createGroupResponse *GProtoCreateGroupResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoCreateGroupResponse.Code
		if err = buffer.WriteUInt8(createGroupResponse.Code); err != nil {
			return nil, err
		}

		// GProtoCreateGroupResponse.Info [optional]
		if createGroupResponse.InfoFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteStruct(1, createGroupResponse.Info); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (createGroupResponse *GProtoCreateGroupResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoCreateGroupResponse.Code
		if err = buffer.ReadUInt8(&createGroupResponse.Code); err != nil {
			return err
		}

		// GProtoCreateGroupResponse.Info [optional]
		if err = buffer.ReadUInt8(&createGroupResponse.InfoFlag); err != nil {
			return err
		}
		if createGroupResponse.InfoFlag == OPTIONAL_TRUE {
			createGroupResponse.Info = &GProtoGroupInfo{}
			if err = buffer.ReadStruct(1, createGroupResponse.Info); err != nil {
				return err
			}

		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetGroupInfoRequest {
//     required string group_id = 1;
// }

type GProtoGetGroupInfoRequest struct {
	GroupId string
}

func (getGroupInfoRequest *GProtoGetGroupInfoRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGetGroupInfoRequest.GroupId
		if err = buffer.WriteString(getGroupInfoRequest.GroupId); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (getGroupInfoRequest *GProtoGetGroupInfoRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		/////////////
		// Decode  //
		/////////////

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoCreateGroupResponse.Code
		if getGroupInfoRequest.GroupId, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetGroupInfoResponse {
//     required uint8 code = 1;
//     optional GroupInfo info = 2;
// }

type GProtoGetGroupInfoResponse struct {
	Code     uint8
	Info     *GProtoGroupInfo // optional
	InfoFlag uint8
}

func (getGroupInfoResponse *GProtoGetGroupInfoResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGetGroupInfoResponse.Code
		if err = buffer.WriteUInt8(getGroupInfoResponse.Code); err != nil {
			return nil, err
		}

		// GProtoGetGroupInfoResponse.Info [optional]
		if getGroupInfoResponse.InfoFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteStruct(1, getGroupInfoResponse.Info); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (getGroupInfoResponse *GProtoGetGroupInfoResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoGetGroupInfoResponse.Code
		if err = buffer.ReadUInt8(&getGroupInfoResponse.Code); err != nil {
			return err
		}

		// GProtoGetGroupInfoResponse.Info [optional]
		if err = buffer.ReadUInt8(&getGroupInfoResponse.InfoFlag); err != nil {
			return err
		}
		if getGroupInfoResponse.InfoFlag == OPTIONAL_TRUE {
			getGroupInfoResponse.Info = &GProtoGroupInfo{}
			if err = buffer.ReadStruct(1, getGroupInfoResponse.Info); err != nil {
				return err
			}

		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message SetGroupInfoRequest {
//     required string group_id = 1;
//     optional string name = 2;
//     optional string avatar = 3;
//     optional string description = 4;
//     optional string data = 5;
// }

type GProtoSetGroupInfoRequest struct {
	GroupId         string
	Name            string // optional
	NameFlag        uint8
	Avatar          string // optional
	AvatarFlag      uint8
	Description     string // optional
	DescriptionFlag uint8
	Data            string // optional
	DataFlag        uint8
}

func (setGroupInfoRequest *GProtoSetGroupInfoRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoSetGroupInfoRequest.GroupId
		if err = buffer.WriteString(setGroupInfoRequest.GroupId); err != nil {
			return nil, err
		}

		// GProtoSetGroupInfoRequest.Name [optional]
		if setGroupInfoRequest.NameFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(setGroupInfoRequest.Name); err != nil {
				return nil, err
			}
		}

		// GProtoSetGroupInfoRequest.Avatar [optional]
		if setGroupInfoRequest.AvatarFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(setGroupInfoRequest.Avatar); err != nil {
				return nil, err
			}
		}

		// GProtoSetGroupInfoRequest.Description [optional]
		if setGroupInfoRequest.DescriptionFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(setGroupInfoRequest.Description); err != nil {
				return nil, err
			}
		}

		// GProtoSetGroupInfoRequest.Data [optional]
		if setGroupInfoRequest.DataFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(setGroupInfoRequest.Data); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (setGroupInfoRequest *GProtoSetGroupInfoRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoSetGroupInfoRequest.GroupId
		if setGroupInfoRequest.GroupId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoSetGroupInfoRequest.Name [optional]
		if err = buffer.ReadUInt8(&setGroupInfoRequest.NameFlag); err != nil {
			return err
		}
		if setGroupInfoRequest.NameFlag == OPTIONAL_TRUE {
			if setGroupInfoRequest.Name, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoSetGroupInfoRequest.Avatar [optional]
		if err = buffer.ReadUInt8(&setGroupInfoRequest.AvatarFlag); err != nil {
			return err
		}
		if setGroupInfoRequest.AvatarFlag == OPTIONAL_TRUE {
			if setGroupInfoRequest.Avatar, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoSetGroupInfoRequest.Description [optional]
		if err = buffer.ReadUInt8(&setGroupInfoRequest.DescriptionFlag); err != nil {
			return err
		}
		if setGroupInfoRequest.DescriptionFlag == OPTIONAL_TRUE {
			if setGroupInfoRequest.Description, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoSetGroupInfoRequest.Data [optional]
		if err = buffer.ReadUInt8(&setGroupInfoRequest.DataFlag); err != nil {
			return err
		}
		if setGroupInfoRequest.DataFlag == OPTIONAL_TRUE {
			if setGroupInfoRequest.Data, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message SetGroupInfoResponse {
//     required uint8 code = 1;
//     optional GroupInfo info = 2;
// }

type GProtoSetGroupInfoResponse struct {
	Code     uint8
	Info     *GProtoGroupInfo // optional
	InfoFlag uint8
}

func (setGroupInfoResponse *GProtoSetGroupInfoResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoSetGroupInfoResponse.Code
		if err = buffer.WriteUInt8(setGroupInfoResponse.Code); err != nil {
			return nil, err
		}

		// GProtoSetGroupInfoResponse.Info [optional]
		if setGroupInfoResponse.InfoFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteStruct(1, setGroupInfoResponse.Info); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (setGroupInfoResponse *GProtoSetGroupInfoResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoSetGroupInfoResponse.Code
		if err = buffer.ReadUInt8(&setGroupInfoResponse.Code); err != nil {
			return err
		}

		// GProtoSetGroupInfoResponse.Info [optional]
		if err = buffer.ReadUInt8(&setGroupInfoResponse.InfoFlag); err != nil {
			return err
		}
		if setGroupInfoResponse.InfoFlag == OPTIONAL_TRUE {
			setGroupInfoResponse.Info = &GProtoGroupInfo{}
			if err = buffer.ReadStruct(1, setGroupInfoResponse.Info); err != nil {
				return err
			}

		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetGroupListRequest {
// }

type GProtoGetGroupListRequest struct {
}

func (getGroupListRequest *GProtoGetGroupListRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		/////////////
		// Encode  //
		/////////////

		return nil, nil
	}

	return nil, InvalidVersionError
}

func (getGroupListRequest *GProtoGetGroupListRequest) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetGroupListResponse {
//     required uint8 code = 1;
//     repeated string group_ids = 2;
// }

type GProtoGetGroupListResponse struct {
	Code     uint8
	GroupIds []string
}

func (getGroupListResponse *GProtoGetGroupListResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGetGroupListResponse.Code
		if err = buffer.WriteUInt8(getGroupListResponse.Code); err != nil {
			return nil, err
		}

		// GProtoGetGroupListResponse.GroupIds [array]
		groupIdsNum := len(getGroupListResponse.GroupIds)
		if err := buffer.WriteUInt16(uint16(groupIdsNum)); err != nil {
			return nil, err
		}
		for i := 0; i < groupIdsNum; i++ {
			if err = buffer.WriteString(getGroupListResponse.GroupIds[i]); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (getGroupListResponse *GProtoGetGroupListResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoGetGroupListResponse.Code
		if err = buffer.ReadUInt8(&getGroupListResponse.Code); err != nil {
			return err
		}

		// GProtoGetGroupListResponse.GroupIds [array]
		var groupIdsNum uint16
		if err = buffer.ReadUInt16(&groupIdsNum); err != nil {
			return err
		} else {
			getGroupListResponse.GroupIds = make([]string, groupIdsNum, groupIdsNum)
			for i := 0; i < int(groupIdsNum); i++ {
				if getGroupListResponse.GroupIds[i], err = buffer.ReadString(); err != nil {
					return err
				}
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message EnterGroupRequest {
//     required string group_id = 1;
//     required string invite_from = 2;
//     repeated string invite_tos = 3;
// }

type GProtoEnterGroupRequest struct {
	GroupId    string
	InviteFrom string
	InviteTos  []string
}

func (enterGroupRequest *GProtoEnterGroupRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoEnterGroupRequest.GroupId
		if err = buffer.WriteString(enterGroupRequest.GroupId); err != nil {
			return nil, err
		}

		// GProtoEnterGroupRequest.InviteFrom
		if err = buffer.WriteString(enterGroupRequest.InviteFrom); err != nil {
			return nil, err
		}

		// GProtoEnterGroupRequest.InviteTos [array]
		inviteTosNum := len(enterGroupRequest.InviteTos)
		if err := buffer.WriteUInt16(uint16(inviteTosNum)); err != nil {
			return nil, err
		}
		for i := 0; i < inviteTosNum; i++ {
			if err = buffer.WriteString(enterGroupRequest.InviteTos[i]); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (enterGroupRequest *GProtoEnterGroupRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoEnterGroupRequest.GroupId
		if enterGroupRequest.GroupId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoEnterGroupRequest.InviteFrom
		if enterGroupRequest.InviteFrom, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoEnterGroupRequest.InviteTos [array]
		var inviteTosNum uint16
		if err = buffer.ReadUInt16(&inviteTosNum); err != nil {
			return err
		} else {
			enterGroupRequest.InviteTos = make([]string, inviteTosNum, inviteTosNum)
			for i := 0; i < int(inviteTosNum); i++ {
				if enterGroupRequest.InviteTos[i], err = buffer.ReadString(); err != nil {
					return err
				}
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message EnterGroupResponse {
//     required uint8 code = 1;
//     optional string invite_from = 2;
//     repeated string invite_tos = 3;
//     optional GroupInfo info = 4;
// }

type GProtoEnterGroupResponse struct {
	Code           uint8
	InviteFrom     string
	InviteFromFlag uint8
	InviteTos      []string
	Info           *GProtoGroupInfo
	InfoFlag       uint8
}

func (enterGroupResponse *GProtoEnterGroupResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoEnterGroupResponse.Code
		if err = buffer.WriteUInt8(enterGroupResponse.Code); err != nil {
			return nil, err
		}

		// GProtoEnterGroupResponse.InviteFrom [optional]
		if enterGroupResponse.InviteFromFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(enterGroupResponse.InviteFrom); err != nil {
				return nil, err
			}
		}

		// GProtoEnterGroupResponse.InviteTos [array]
		inviteTosNum := len(enterGroupResponse.InviteTos)
		if err := buffer.WriteUInt16(uint16(inviteTosNum)); err != nil {
			return nil, err
		}
		for i := 0; i < inviteTosNum; i++ {
			if err = buffer.WriteString(enterGroupResponse.InviteTos[i]); err != nil {
				return nil, err
			}
		}

		// GProtoEnterGroupResponse.Info [optional]
		if enterGroupResponse.InfoFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteStruct(1, enterGroupResponse.Info); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (enterGroupResponse *GProtoEnterGroupResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoEnterGroupResponse.Code
		if err = buffer.ReadUInt8(&enterGroupResponse.Code); err != nil {
			return err
		}

		// GProtoEnterGroupResponse.InviteFrom [optional]
		if err = buffer.ReadUInt8(&enterGroupResponse.InviteFromFlag); err != nil {
			return err
		}
		if enterGroupResponse.InviteFromFlag == OPTIONAL_TRUE {
			if enterGroupResponse.InviteFrom, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoEnterGroupResponse.InviteTos [array]
		var inviteTosNum uint16
		if err = buffer.ReadUInt16(&inviteTosNum); err != nil {
			return err
		} else {
			enterGroupResponse.InviteTos = make([]string, inviteTosNum, inviteTosNum)
			for i := 0; i < int(inviteTosNum); i++ {
				if enterGroupResponse.InviteTos[i], err = buffer.ReadString(); err != nil {
					return err
				}
			}
		}

		// GProtoEnterGroupResponse.Info [optional]
		if err = buffer.ReadUInt8(&enterGroupResponse.InfoFlag); err != nil {
			return err
		}
		if enterGroupResponse.InfoFlag == OPTIONAL_TRUE {
			enterGroupResponse.Info = &GProtoGroupInfo{}
			if err = buffer.ReadStruct(1, enterGroupResponse.Info); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message EnterGroupNotify {
//     required string group_id = 1;
//     required OtherUser invite_from = 2;
//     repeated OtherUser invite_tos = 3;
//     optional GroupInfo info = 4;
//     required uint64 send_time = 5;
//     required uint8 is_offline = 6;
// }

type GProtoEnterGroupNotify struct {
	GroupId    string
	InviteFrom *GProtoOtherUser
	InviteTos  []*GProtoOtherUser
	Info       *GProtoGroupInfo
	InfoFlag   uint8
	SendTime   uint64
	IsOffline  uint8
}

func (enterGroupNotify *GProtoEnterGroupNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoEnterGroupNotify.GroupId
		if err = buffer.WriteString(enterGroupNotify.GroupId); err != nil {
			return nil, err
		}

		// GProtoEnterGroupNotify.InviteFrom
		if err = buffer.WriteStruct(1, enterGroupNotify.InviteFrom); err != nil {
			return nil, err
		}

		// GProtoEnterGroupNotify.InviteTos [array]
		inviteTosNum := len(enterGroupNotify.InviteTos)
		if err := buffer.WriteUInt16(uint16(inviteTosNum)); err != nil {
			return nil, err
		}
		for i := 0; i < inviteTosNum; i++ {
			if err = buffer.WriteStruct(1, enterGroupNotify.InviteTos[i]); err != nil {
				return nil, err
			}
		}

		// GProtoEnterGroupNotify.Info [optional]
		if enterGroupNotify.InfoFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteStruct(1, enterGroupNotify.Info); err != nil {
				return nil, err
			}
		}

		// GProtoEnterGroupNotify.SendTime
		if err = buffer.WriteUInt64(enterGroupNotify.SendTime); err != nil {
			return nil, err
		}

		// GProtoEnterGroupNotify.IsOffline
		if err = buffer.WriteUInt8(enterGroupNotify.IsOffline); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (enterGroupNotify *GProtoEnterGroupNotify) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoEnterGroupNotify.GroupId
		if enterGroupNotify.GroupId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoEnterGroupNotify.InviteFrom
		enterGroupNotify.InviteFrom = &GProtoOtherUser{}
		if err = buffer.ReadStruct(1, enterGroupNotify.InviteFrom); err != nil {
			return err
		}

		// GProtoEnterGroupNotify.InviteTos [array]
		var inviteTosNum uint16
		if err = buffer.ReadUInt16(&inviteTosNum); err != nil {
			return err
		} else {
			enterGroupNotify.InviteTos = make([]*GProtoOtherUser, inviteTosNum, inviteTosNum)
			for i := 0; i < int(inviteTosNum); i++ {
				enterGroupNotify.InviteTos[i] = &GProtoOtherUser{}
				if err = buffer.ReadStruct(1, enterGroupNotify.InviteTos[i]); err != nil {
					return err
				}
			}
		}

		// GProtoEnterGroupNotify.Info [optional]
		if err = buffer.ReadUInt8(&enterGroupNotify.InfoFlag); err != nil {
			return err
		}
		if enterGroupNotify.InfoFlag == OPTIONAL_TRUE {
			enterGroupNotify.Info = &GProtoGroupInfo{}
			if err = buffer.ReadStruct(1, enterGroupNotify.Info); err != nil {
				return err
			}
		}

		// GProtoEnterGroupNotify.SendTime
		if err = buffer.ReadUInt64(&enterGroupNotify.SendTime); err != nil {
			return err
		}

		// GProtoEnterGroupNotify.IsOffline
		if err = buffer.ReadUInt8(&enterGroupNotify.IsOffline); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message LeaveGroupRequest {
//     required string group_id = 1;
//     required string leave_from = 2;
//     repeated string leave_tos = 3;
// }

type GProtoLeaveGroupRequest struct {
	GroupId   string
	LeaveFrom string
	LeaveTos  []string
}

func (leaveGroupRequest *GProtoLeaveGroupRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoLeaveGroupRequest.GroupId
		if err = buffer.WriteString(leaveGroupRequest.GroupId); err != nil {
			return nil, err
		}

		// GProtoLeaveGroupRequest.LeaveFrom
		if err = buffer.WriteString(leaveGroupRequest.LeaveFrom); err != nil {
			return nil, err
		}

		// GProtoLeaveGroupRequest.LeaveTos [array]
		leaveTosNum := len(leaveGroupRequest.LeaveTos)
		if err := buffer.WriteUInt16(uint16(leaveTosNum)); err != nil {
			return nil, err
		}
		for i := 0; i < leaveTosNum; i++ {
			if err = buffer.WriteString(leaveGroupRequest.LeaveTos[i]); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (leaveGroupRequest *GProtoLeaveGroupRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoLeaveGroupRequest.GroupId
		if leaveGroupRequest.GroupId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoLeaveGroupRequest.LeaveFrom
		if leaveGroupRequest.LeaveFrom, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoLeaveGroupRequest.LeaveTos [array]
		var leaveTosNum uint16
		if err = buffer.ReadUInt16(&leaveTosNum); err != nil {
			return err
		} else {
			leaveGroupRequest.LeaveTos = make([]string, leaveTosNum, leaveTosNum)
			for i := 0; i < int(leaveTosNum); i++ {
				if leaveGroupRequest.LeaveTos[i], err = buffer.ReadString(); err != nil {
					return err
				}
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message LeaveGroupResponse {
//     required uint8 code = 1;
//     optional string leave_from = 2;
//     repeated string leave_tos = 3;
//     optional string group_id = 4;
// }

type GProtoLeaveGroupResponse struct {
	Code          uint8
	LeaveFrom     string
	LeaveFromFlag uint8
	LeaveTos      []string
	GroupId       string
	GroupIdFlag   uint8
}

func (leaveGroupResponse *GProtoLeaveGroupResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoLeaveGroupResponse.Code
		if err = buffer.WriteUInt8(leaveGroupResponse.Code); err != nil {
			return nil, err
		}

		// GProtoLeaveGroupResponse.LeaveFrom [optional]
		if leaveGroupResponse.LeaveFromFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(leaveGroupResponse.LeaveFrom); err != nil {
				return nil, err
			}
		}

		// GProtoLeaveGroupResponse.LeaveTos [array]
		leaveTosNum := len(leaveGroupResponse.LeaveTos)
		if err := buffer.WriteUInt16(uint16(leaveTosNum)); err != nil {
			return nil, err
		}
		for i := 0; i < leaveTosNum; i++ {
			if err = buffer.WriteString(leaveGroupResponse.LeaveTos[i]); err != nil {
				return nil, err
			}
		}

		// GProtoLeaveGroupResponse.GroupId [optional]
		if leaveGroupResponse.GroupIdFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(leaveGroupResponse.GroupId); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (leaveGroupResponse *GProtoLeaveGroupResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoLeaveGroupResponse.Code
		if err = buffer.ReadUInt8(&leaveGroupResponse.Code); err != nil {
			return err
		}

		// GProtoLeaveGroupResponse.LeaveFrom [optional]
		if err = buffer.ReadUInt8(&leaveGroupResponse.LeaveFromFlag); err != nil {
			return err
		}
		if leaveGroupResponse.LeaveFromFlag == OPTIONAL_TRUE {
			if leaveGroupResponse.LeaveFrom, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		// GProtoLeaveGroupResponse.LeaveTos [array]
		var leaveTosNum uint16
		if err = buffer.ReadUInt16(&leaveTosNum); err != nil {
			return err
		} else {
			leaveGroupResponse.LeaveTos = make([]string, leaveTosNum, leaveTosNum)
			for i := 0; i < int(leaveTosNum); i++ {
				if leaveGroupResponse.LeaveTos[i], err = buffer.ReadString(); err != nil {
					return err
				}
			}
		}

		// GProtoLeaveGroupResponse.GroupId [optional]
		if err = buffer.ReadUInt8(&leaveGroupResponse.GroupIdFlag); err != nil {
			return err
		}
		if leaveGroupResponse.GroupIdFlag == OPTIONAL_TRUE {
			if leaveGroupResponse.GroupId, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message LeaveGroupNotify {
//     required string group_id = 1;
//     required string leave_from = 2;
//     repeated string leave_tos = 3;
//     required uint64 send_time = 4;
//     required uint8 is_offline = 5;
// }

type GProtoLeaveGroupNotify struct {
	GroupId   string
	LeaveFrom string
	LeaveTos  []string
	SendTime  uint64
	IsOffline uint8
}

func (leaveGroupNotify *GProtoLeaveGroupNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoLeaveGroupNotify.GroupId
		if err = buffer.WriteString(leaveGroupNotify.GroupId); err != nil {
			return nil, err
		}

		// GProtoLeaveGroupNotify.LeaveFrom
		if err = buffer.WriteString(leaveGroupNotify.LeaveFrom); err != nil {
			return nil, err
		}

		// GProtoLeaveGroupNotify.LeaveTos [array]
		leaveTosNum := len(leaveGroupNotify.LeaveTos)
		if err := buffer.WriteUInt16(uint16(leaveTosNum)); err != nil {
			return nil, err
		}
		for i := 0; i < leaveTosNum; i++ {
			if err = buffer.WriteString(leaveGroupNotify.LeaveTos[i]); err != nil {
				return nil, err
			}
		}

		// GProtoLeaveGroupNotify.SendTime
		if err = buffer.WriteUInt64(leaveGroupNotify.SendTime); err != nil {
			return nil, err
		}

		// GProtoLeaveGroupNotify.IsOffline
		if err = buffer.WriteUInt8(leaveGroupNotify.IsOffline); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (leaveGroupNotify *GProtoLeaveGroupNotify) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoLeaveGroupNotify.GroupId
		if leaveGroupNotify.GroupId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoLeaveGroupNotify.LeaveFrom
		if leaveGroupNotify.LeaveFrom, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoLeaveGroupNotify.LeaveTos [array]
		var leaveTosNum uint16
		if err = buffer.ReadUInt16(&leaveTosNum); err != nil {
			return err
		} else {
			leaveGroupNotify.LeaveTos = make([]string, leaveTosNum, leaveTosNum)
			for i := 0; i < int(leaveTosNum); i++ {
				if leaveGroupNotify.LeaveTos[i], err = buffer.ReadString(); err != nil {
					return err
				}
			}
		}

		// GProtoLeaveGroupNotify.SendTime
		if err = buffer.ReadUInt64(&leaveGroupNotify.SendTime); err != nil {
			return err
		}

		// GProtoLeaveGroupNotify.IsOffline
		if err = buffer.ReadUInt8(&leaveGroupNotify.IsOffline); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GroupKickoffNotify {
//     required string group_id = 1;
//     required string account = 2;
//     required uint64 send_time = 3;
//     required uint8 is_offline = 4;
// }

type GProtoGroupKickoffNotify struct {
	GroupId   string
	Account   string
	SendTime  uint64
	IsOffline uint8
}

func (groupKickoffNotify *GProtoGroupKickoffNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGroupKickoffNotify.GroupId
		if err = buffer.WriteString(groupKickoffNotify.GroupId); err != nil {
			return nil, err
		}

		// GProtoGroupKickoffNotify.Account
		if err = buffer.WriteString(groupKickoffNotify.Account); err != nil {
			return nil, err
		}

		// GProtoGroupKickoffNotify.SendTime
		if err = buffer.WriteUInt64(groupKickoffNotify.SendTime); err != nil {
			return nil, err
		}

		// GProtoGroupKickoffNotify.IsOffline
		if err = buffer.WriteUInt8(groupKickoffNotify.IsOffline); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (groupKickoffNotify *GProtoGroupKickoffNotify) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {

		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoGroupKickoffNotify.GroupId
		if groupKickoffNotify.GroupId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoGroupKickoffNotify.Account
		if groupKickoffNotify.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoGroupKickoffNotify.SendTime
		if err = buffer.ReadUInt64(&groupKickoffNotify.SendTime); err != nil {
			return err
		}

		// GProtoGroupKickoffNotify.IsOffline
		if err = buffer.ReadUInt8(&groupKickoffNotify.IsOffline); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message SetGroupMarknameRequest {
//     required string group_id = 1;
//     required string markname = 2;
// }

type GProtoSetGroupMarknameRequest struct {
	GroupId  string
	Markname string
}

func (setGroupMarknameRequest *GProtoSetGroupMarknameRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoSetGroupMarknameRequest.GroupId
		if err = buffer.WriteString(setGroupMarknameRequest.GroupId); err != nil {
			return nil, err
		}

		// GProtoSetGroupMarknameRequest.Markname
		if err = buffer.WriteString(setGroupMarknameRequest.Markname); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (setGroupMarknameRequest *GProtoSetGroupMarknameRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoSetGroupMarknameRequest.GroupId
		if setGroupMarknameRequest.GroupId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoSetGroupMarknameRequest.Markname
		if setGroupMarknameRequest.Markname, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

type GProtoSetGroupMarknameResponse struct {
	Code     uint8
	GroupId  string
	Markname string
}

func (setGroupMarknameResponse *GProtoSetGroupMarknameResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoSetGroupMarknameResponse.Code
		if err = buffer.WriteUInt8(setGroupMarknameResponse.Code); err != nil {
			return nil, err
		}

		// GProtoSetGroupMarknameResponse.GroupId
		if err = buffer.WriteString(setGroupMarknameResponse.GroupId); err != nil {
			return nil, err
		}

		// GProtoSetGroupMarknameResponse.Markname
		if err = buffer.WriteString(setGroupMarknameResponse.Markname); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (setGroupMarknameResponse *GProtoSetGroupMarknameResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoSetGroupMarknameResponse.Code
		if err = buffer.ReadUInt8(&setGroupMarknameResponse.Code); err != nil {
			return err
		}

		// GProtoSetGroupMarknameResponse.GroupId
		if setGroupMarknameResponse.GroupId, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoSetGroupMarknameResponse.Markname
		if setGroupMarknameResponse.Markname, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message SetGroupMarknameResponse {
//     required uint8 code = 1;
//     required string group_id = 2;
//     required string markname = 3;
// }

////////////////////////////////////
// message AddBlacklistRequest {
//     required string account = 1;
// }

type GProtoAddBlacklistRequest struct {
	Account string
}

func (addBlacklistRequest *GProtoAddBlacklistRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoAddBlacklistRequest.Account
		if err = buffer.WriteString(addBlacklistRequest.Account); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (addBlacklistRequest *GProtoAddBlacklistRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoAddBlacklistRequest.Account
		if addBlacklistRequest.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message AddBlacklistResponse {
//     required uint8 code = 1;
//     required string account = 2;
//     required uint8 flag = 3;
// }

type GProtoAddBlacklistResponse struct {
	Code    uint8
	Account string
	Flag    uint8
}

func (addBlacklistResponse *GProtoAddBlacklistResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoAddBlacklistResponse.Code
		if err = buffer.WriteUInt8(addBlacklistResponse.Code); err != nil {
			return nil, err
		}

		// GProtoAddBlacklistResponse.Account
		if err = buffer.WriteString(addBlacklistResponse.Account); err != nil {
			return nil, err
		}

		// GProtoAddBlacklistResponse.Flag
		if err = buffer.WriteUInt8(addBlacklistResponse.Flag); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (addBlacklistResponse *GProtoAddBlacklistResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoAddBlacklistResponse.Code
		if err = buffer.ReadUInt8(&addBlacklistResponse.Code); err != nil {
			return err
		}

		// GProtoAddBlacklistResponse.Account
		if addBlacklistResponse.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoAddBlacklistResponse.Flag
		if err = buffer.ReadUInt8(&addBlacklistResponse.Flag); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message AddBlacklistNotify {
//     required OtherUser user = 1;
//     required uint8 flag = 2;
//     required uint64 send_time = 3;
//     required uint8 is_offline = 4;
// }

type GProtoAddBlacklistNotify struct {
	User      *GProtoOtherUser
	Flag      uint8
	SendTime  uint64
	IsOffline uint8
}

func (addBlacklistNotify *GProtoAddBlacklistNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoAddBlacklistNotify.User
		if err = buffer.WriteStruct(1, addBlacklistNotify.User); err != nil {
			return nil, err
		}

		// GProtoAddBlacklistNotify.Flag
		if err = buffer.WriteUInt8(addBlacklistNotify.Flag); err != nil {
			return nil, err
		}

		// GProtoAddBlacklistNotify.SendTime
		if err = buffer.WriteUInt64(addBlacklistNotify.SendTime); err != nil {
			return nil, err
		}

		// GProtoAddBlacklistNotify.IsOffline
		if err = buffer.WriteUInt8(addBlacklistNotify.IsOffline); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (addBlacklistNotify *GProtoAddBlacklistNotify) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoAddBlacklistNotify.User
		addBlacklistNotify.User = &GProtoOtherUser{}
		if err = buffer.ReadStruct(1, addBlacklistNotify.User); err != nil {
			return err
		}

		// GProtoAddBlacklistNotify.Flag
		if err = buffer.ReadUInt8(&addBlacklistNotify.Flag); err != nil {
			return err
		}

		// GProtoAddBlacklistNotify.SendTime
		if err = buffer.ReadUInt64(&addBlacklistNotify.SendTime); err != nil {
			return err
		}

		// GProtoAddBlacklistNotify.IsOffline
		if err = buffer.ReadUInt8(&addBlacklistNotify.IsOffline); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message DelBlacklistRequest {
//     required string account = 1;
// 	   required uint8 ops = 2;
// }

type GProtoDelBlacklistRequest struct {
	Account string
}

func (delBlacklistRequest *GProtoDelBlacklistRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoDelBlacklistRequest.Account
		if err = buffer.WriteString(delBlacklistRequest.Account); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (delBlacklistRequest *GProtoDelBlacklistRequest) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoDelBlacklistRequest.Account
		if delBlacklistRequest.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message DelBlacklistResponse {
//     required uint8 code = 1;
//     required string account = 2;
//     required uint8 flag = 3;
// }

type GProtoDelBlacklistResponse struct {
	Code    uint8
	Account string
	Flag    uint8
}

func (delBlacklistResponse *GProtoDelBlacklistResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoDelBlacklistResponse.Code
		if err = buffer.WriteUInt8(delBlacklistResponse.Code); err != nil {
			return nil, err
		}

		// GProtoDelBlacklistResponse.Account
		if err = buffer.WriteString(delBlacklistResponse.Account); err != nil {
			return nil, err
		}

		// GProtoDelBlacklistResponse.Flag
		if err = buffer.WriteUInt8(delBlacklistResponse.Flag); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (delBlacklistResponse *GProtoDelBlacklistResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoDelBlacklistResponse.Code
		if err = buffer.ReadUInt8(&delBlacklistResponse.Code); err != nil {
			return err
		}

		// GProtoDelBlacklistResponse.Account
		if delBlacklistResponse.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoDelBlacklistResponse.Flag
		if err = buffer.ReadUInt8(&delBlacklistResponse.Flag); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message DelBlacklistNotify {
//     required OtherUser user = 1;
//     required uint8 flag = 2;
//     required uint64 send_time = 3;
//     required uint8 is_offline = 4;
// }

type GProtoDelBlacklistNotify struct {
	User      *GProtoOtherUser
	Flag      uint8
	SendTime  uint64
	IsOffline uint8
}

func (elBlacklistNotify *GProtoDelBlacklistNotify) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoDelBlacklistNotify.User
		if err = buffer.WriteStruct(1, elBlacklistNotify.User); err != nil {
			return nil, err
		}

		// GProtoDelBlacklistNotify.Flag
		if err = buffer.WriteUInt8(elBlacklistNotify.Flag); err != nil {
			return nil, err
		}

		// GProtoDelBlacklistNotify.SendTime
		if err = buffer.WriteUInt64(elBlacklistNotify.SendTime); err != nil {
			return nil, err
		}

		// GProtoDelBlacklistNotify.IsOffline
		if err = buffer.WriteUInt8(elBlacklistNotify.IsOffline); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (elBlacklistNotify *GProtoDelBlacklistNotify) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 3 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoDelBlacklistNotify.User
		elBlacklistNotify.User = &GProtoOtherUser{}
		if err = buffer.ReadStruct(1, elBlacklistNotify.User); err != nil {
			return err
		}

		// GProtoDelBlacklistNotify.Flag
		if err = buffer.ReadUInt8(&elBlacklistNotify.Flag); err != nil {
			return err
		}

		// GProtoDelBlacklistNotify.SendTime
		if err = buffer.ReadUInt64(&elBlacklistNotify.SendTime); err != nil {
			return err
		}

		// GProtoDelBlacklistNotify.IsOffline
		if err = buffer.ReadUInt8(&elBlacklistNotify.IsOffline); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetBlacklistsRequest {
// }

type GProtoGetBlacklistsRequest struct {
}

func (getBlacklistsRequest *GProtoGetBlacklistsRequest) Encode(version uint16) ([]byte, error) {
	if version == 1 {

		/////////////
		// Encode  //
		/////////////

		return nil, nil
	}

	return nil, InvalidVersionError
}

func (getBlacklistsRequest *GProtoGetBlacklistsRequest) Decode(version uint16, buf []byte) error {
	if version == 1 {

		/////////////
		// Decode  //
		/////////////

		return nil
	}

	return InvalidVersionError
}

////////////////////////////////////
// message GetBlacklistsResponse {
//     required uint8 code = 1;
//     repeated Blacklist blacklists = 2;
// }

type GProtoGetBlacklistsResponse struct {
	Code       uint8
	Blacklists []*GProtoBlacklist
}

func (getBlacklistsResponse *GProtoGetBlacklistsResponse) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGetBlacklistsResponse.Code
		if err = buffer.WriteUInt8(getBlacklistsResponse.Code); err != nil {
			return nil, err
		}

		// GProtoGetBlacklistsResponse.Blacklists [array]
		if err = buffer.WriteArray(1, getBlacklistsResponse.Blacklists); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (getBlacklistsResponse *GProtoGetBlacklistsResponse) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 1 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoGetBlacklistsResponse.Code
		if err = buffer.ReadUInt8(&getBlacklistsResponse.Code); err != nil {
			return err
		}

		// GProtoGetBlacklistsResponse.Blacklists [array]
		var blacklistsNum uint16
		if err = buffer.ReadUInt16(&blacklistsNum); err != nil {
			return err
		} else {
			getBlacklistsResponse.Blacklists = make([]*GProtoBlacklist, blacklistsNum, blacklistsNum)
			for i := 0; i < int(blacklistsNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						getBlacklistsResponse.Blacklists[i] = &GProtoBlacklist{}
						if err = getBlacklistsResponse.Blacklists[i].Decode(1, bufObj); err != nil {
							return err
						}
					}
				}
			}
		}

		return nil
	}

	return InvalidVersionError
}
