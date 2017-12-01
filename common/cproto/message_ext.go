package gproto

import (
	"fmt"
)

// message SelfUser {
//     required string account = 1;
//     required string nickname = 2;
//     required uint8 gender = 3;
//     required string avatar = 4;
//     required string signature = 5;
//     required string extern = 6;
//     required string customVersion = 7;
//     required string data = 8;
//     required int64 last_online_timestamp = 9;
// }

type GProtoSelfUser struct {
	Account             string
	Nickname            string
	Gender              uint8
	Avatar              string
	Signature           string
	Extern              string
	CustomVersion       string
	Data                string
	LastOnlineTimestamp int64
}

func EmptyGProtoSelfUser(account string) *GProtoSelfUser {
	return &GProtoSelfUser{
		Account:             account,
		Nickname:            "",
		Gender:              0,
		Avatar:              "",
		Signature:           "",
		Extern:              "",
		CustomVersion:       "",
		Data:                "",
		LastOnlineTimestamp: 0,
	}
}

func (selfUser *GProtoSelfUser) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoSelfUser.Account
		if err = buffer.WriteString(selfUser.Account); err != nil {
			return nil, err
		}

		// GProtoSelfUser.Nickname
		if err = buffer.WriteString(selfUser.Nickname); err != nil {
			return nil, err
		}

		// GProtoSelfUser.Gender
		if err = buffer.WriteUInt8(selfUser.Gender); err != nil {
			return nil, err
		}

		// GProtoSelfUser.Avatar
		if err = buffer.WriteString(selfUser.Avatar); err != nil {
			return nil, err
		}

		// GProtoSelfUser.Signature
		if err = buffer.WriteString(selfUser.Signature); err != nil {
			return nil, err
		}

		// GProtoSelfUser.Extern
		if err = buffer.WriteString(selfUser.Extern); err != nil {
			return nil, err
		}

		// GProtoSelfUser.CustomVersion
		if err = buffer.WriteString(selfUser.CustomVersion); err != nil {
			return nil, err
		}

		// GProtoSelfUser.Data
		if err = buffer.WriteString(selfUser.Data); err != nil {
			return nil, err
		}

		// GProtoSelfUser.LastOnlineTimestamp
		if err = buffer.WriteInt64(selfUser.LastOnlineTimestamp); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (selfUser *GProtoSelfUser) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 10 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoSelfUser.Account
		if selfUser.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoSelfUser.Nickname
		if selfUser.Nickname, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoSelfUser.Gender
		if err = buffer.ReadUInt8(&selfUser.Gender); err != nil {
			return err
		}

		// GProtoSelfUser.Avatar
		if selfUser.Avatar, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoSelfUser.Signature
		if selfUser.Signature, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoSelfUser.Extern
		if selfUser.Extern, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoSelfUser.CustomVersion
		if selfUser.CustomVersion, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoSelfUser.Data
		if selfUser.Data, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoSelfUser.LastOnlineTimestamp
		if err = buffer.ReadInt64(&selfUser.LastOnlineTimestamp); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

func (selfUser *GProtoSelfUser) String() string {
	return fmt.Sprintf("%v_%v_%v_%v", selfUser.Account, selfUser.Nickname, selfUser.Gender, selfUser.CustomVersion)
}

// message OtherUser {
//     required string account = 1;
//     required string nickname = 2;
//     required uint8 gender = 3;
//     required string avatar = 4;
//     required string signature = 5;
//     required string extern = 6;
//     required string customVersion = 7;
//     required string data = 8;
//     required int64 last_online_timestamp = 9;
//     optional string markname = 10;
// }

type GProtoOtherUser struct {
	Account             string
	Nickname            string
	Gender              uint8
	Avatar              string
	Signature           string
	Extern              string
	CustomVersion       string
	Data                string
	LastOnlineTimestamp int64
	Markname            string
	MakrnameFlag        uint8
}

func EmptyGProtoOtherUser(account string) *GProtoOtherUser {
	return &GProtoOtherUser{
		Account:             account,
		Nickname:            "",
		Gender:              0,
		Avatar:              "",
		Signature:           "",
		Extern:              "",
		CustomVersion:       "",
		Data:                "",
		LastOnlineTimestamp: 0,
		Markname:            "",
		MakrnameFlag:        OPTIONAL_TRUE,
	}
}

func (otherUser *GProtoOtherUser) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoOtherUser.Account
		if err = buffer.WriteString(otherUser.Account); err != nil {
			return nil, err
		}

		// GProtoOtherUser.Nickname
		if err = buffer.WriteString(otherUser.Nickname); err != nil {
			return nil, err
		}

		// GProtoOtherUser.Gender
		if err = buffer.WriteUInt8(otherUser.Gender); err != nil {
			return nil, err
		}

		// GProtoOtherUser.Avatar
		if err = buffer.WriteString(otherUser.Avatar); err != nil {
			return nil, err
		}

		// GProtoOtherUser.Signature
		if err = buffer.WriteString(otherUser.Signature); err != nil {
			return nil, err
		}

		// GProtoOtherUser.Extern
		if err = buffer.WriteString(otherUser.Extern); err != nil {
			return nil, err
		}

		// GProtoOtherUser.CustomVersion
		if err = buffer.WriteString(otherUser.CustomVersion); err != nil {
			return nil, err
		}

		// GProtoOtherUser.Data
		if err = buffer.WriteString(otherUser.Data); err != nil {
			return nil, err
		}

		// GProtoOtherUser.LastOnlineTimestamp
		if err = buffer.WriteInt64(otherUser.LastOnlineTimestamp); err != nil {
			return nil, err
		}

		// GProtoOtherUser.Markname [optional]
		if otherUser.MakrnameFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(otherUser.Markname); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (otherUser *GProtoOtherUser) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 10 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoOtherUser.Account
		if otherUser.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoOtherUser.Nickname
		if otherUser.Nickname, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoOtherUser.Gender
		if err = buffer.ReadUInt8(&otherUser.Gender); err != nil {
			return err
		}

		// GProtoOtherUser.Avatar
		if otherUser.Avatar, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoOtherUser.Signature
		if otherUser.Signature, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoOtherUser.Extern
		if otherUser.Extern, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoOtherUser.CustomVersion
		if otherUser.CustomVersion, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoOtherUser.Data
		if otherUser.Data, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoOtherUser.LastOnlineTimestamp
		if err = buffer.ReadInt64(&otherUser.LastOnlineTimestamp); err != nil {
			return err
		}

		// GProtoOtherUser.Markname [optional]
		if err = buffer.ReadUInt8(&otherUser.MakrnameFlag); err != nil {
			return err
		}
		if otherUser.MakrnameFlag == OPTIONAL_TRUE {
			if otherUser.Markname, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

func (otherUser *GProtoOtherUser) String() string {
	return fmt.Sprintf("%v_%v_%v", otherUser.Account, otherUser.Nickname, otherUser.Gender)
}

///////////////////////////////////////
// message Friend {
//     required string account = 1;
//     required flag flag = 2;
//     optional string markname = 3;
// }

type GProtoFriend struct {
	Account      string
	Flag         uint8
	Markname     string
	MakrnameFlag uint8
}

func (friend *GProtoFriend) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoFriend.Account
		if err = buffer.WriteString(friend.Account); err != nil {
			return nil, err
		}

		// GProtoFriend.Flag
		if err = buffer.WriteUInt8(friend.Flag); err != nil {
			return nil, err
		}

		// GProtoFriend.Markname [optional]
		if friend.MakrnameFlag == OPTIONAL_FALSE {
			if err = buffer.WriteUInt8(OPTIONAL_FALSE); err != nil {
				return nil, err
			}
		} else {
			if err = buffer.WriteUInt8(OPTIONAL_TRUE); err != nil {
				return nil, err
			}
			if err = buffer.WriteString(friend.Markname); err != nil {
				return nil, err
			}
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (friend *GProtoFriend) Decode(version uint16, buf []byte) error {
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

		// GProtoFriend.Account
		if friend.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoFriend.Flag
		if err = buffer.ReadUInt8(&friend.Flag); err != nil {
			return err
		}

		// GProtoFriend.Markname [optional]
		if err = buffer.ReadUInt8(&friend.MakrnameFlag); err != nil {
			return err
		}
		if friend.MakrnameFlag == OPTIONAL_TRUE {
			if friend.Markname, err = buffer.ReadString(); err != nil {
				return err
			}
		}

		return nil
	}

	return InvalidVersionError
}

///////////////////////////////////////
// message Blacklist {
//     required string account = 1;
//     required flag flag = 2;
// }

type GProtoBlacklist struct {
	Account string
	Flag    uint8
}

func (blacklist *GProtoBlacklist) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoBlacklist.Account
		if err = buffer.WriteString(blacklist.Account); err != nil {
			return nil, err
		}

		// GProtoBlacklist.Flag
		if err = buffer.WriteUInt8(blacklist.Flag); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (blacklist *GProtoBlacklist) Decode(version uint16, buf []byte) error {
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

		// GProtoBlacklist.Account
		if blacklist.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoBlacklist.Flag
		if err = buffer.ReadUInt8(&blacklist.Flag); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

///////////////////////////////////////
// message ChatRoomInfo {
//     required string chatroom_id = 1;
//     requierd string name = 2;
//     requierd string description = 3;
//     required string owner_account = 4;
//     required uint16 max_user = 5;
//     required string data = 6;
//     required int64 created_timestamp = 7;
// }

type GProtoChatRoomInfo struct {
	ChatRoomId       string
	Name             string
	Description      string
	OwnerAccount     string
	MaxUser          uint16
	Data             string
	CreatedTimestamp int64
}

func (chatRoomInfo *GProtoChatRoomInfo) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoChatRoomInfo.ChatRoomId
		if err = buffer.WriteString(chatRoomInfo.ChatRoomId); err != nil {
			return nil, err
		}

		// GProtoChatRoomInfo.Name
		if err = buffer.WriteString(chatRoomInfo.Name); err != nil {
			return nil, err
		}

		// GProtoChatRoomInfo.Description
		if err = buffer.WriteString(chatRoomInfo.Description); err != nil {
			return nil, err
		}

		// GProtoChatRoomInfo.OwnerAccount
		if err = buffer.WriteString(chatRoomInfo.OwnerAccount); err != nil {
			return nil, err
		}

		// GProtoChatRoomInfo.MaxUser
		if err = buffer.WriteUInt16(chatRoomInfo.MaxUser); err != nil {
			return nil, err
		}

		// GProtoChatRoomInfo.Data
		if err = buffer.WriteString(chatRoomInfo.Data); err != nil {
			return nil, err
		}

		// GProtoChatRoomInfo.CreatedTimestamp
		if err = buffer.WriteInt64(chatRoomInfo.CreatedTimestamp); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (chatRoomInfo *GProtoChatRoomInfo) Decode(version uint16, buf []byte) error {
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

		// GProtoChatRoomInfo.ChatRoomId
		if chatRoomInfo.ChatRoomId, err = buffer.ReadString(); err != nil {
			return err
		}
		// GProtoChatRoomInfo.Name
		if chatRoomInfo.Name, err = buffer.ReadString(); err != nil {
			return err
		}
		// GProtoChatRoomInfo.Description
		if chatRoomInfo.Description, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoChatRoomInfo.OwnerAccount
		if chatRoomInfo.OwnerAccount, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoChatRoomInfo.MaxUser
		if err = buffer.ReadUInt16(&chatRoomInfo.MaxUser); err != nil {
			return err
		}

		// GProtoChatRoomInfo.Data
		if chatRoomInfo.Data, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoChatRoomInfo.CreatedTimestamp
		if err = buffer.ReadInt64(&chatRoomInfo.CreatedTimestamp); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}

///////////////////////////////////////
// message GroupInfo {
//     required string group_id = 1;
//     requierd string name = 2;
//     requierd string avatar = 3;
//     requierd string description = 4;
//     required string owner_account = 5;
//     required uint16 max_user = 6;
//     required string data = 7;
//     required uint8 flag = 8;
//     required int64 created_timestamp = 9;
//     repeated OtherUser users = 10;
//     repeated GroupMarkname marknames = 11;
// }

type GProtoGroupInfo struct {
	GroupId          string
	Name             string
	Avatar           string
	Description      string
	OwnerAccount     string
	MaxUser          uint16
	Data             string
	Flag             uint8
	CreatedTimestamp int64
	Users            []*GProtoOtherUser
	Marknames        []*GProtoGroupMarkname
}

func (groupInfo *GProtoGroupInfo) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGroupInfo.GroupId
		if err = buffer.WriteString(groupInfo.GroupId); err != nil {
			return nil, err
		}

		// GProtoGroupInfo.Name
		if err = buffer.WriteString(groupInfo.Name); err != nil {
			return nil, err
		}

		// GProtoGroupInfo.Avatar
		if err = buffer.WriteString(groupInfo.Avatar); err != nil {
			return nil, err
		}

		// GProtoGroupInfo.Description
		if err = buffer.WriteString(groupInfo.Description); err != nil {
			return nil, err
		}

		// GProtoGroupInfo.OwnerAccount
		if err = buffer.WriteString(groupInfo.OwnerAccount); err != nil {
			return nil, err
		}

		// GProtoGroupInfo.MaxUser
		if err = buffer.WriteUInt16(groupInfo.MaxUser); err != nil {
			return nil, err
		}

		// GProtoGroupInfo.Data
		if err = buffer.WriteString(groupInfo.Data); err != nil {
			return nil, err
		}

		// GProtoGroupInfo.Flag
		if err = buffer.WriteUInt8(groupInfo.Flag); err != nil {
			return nil, err
		}

		// GProtoGroupInfo.CreatedTimestamp
		if err = buffer.WriteInt64(groupInfo.CreatedTimestamp); err != nil {
			return nil, err
		}

		// GProtoGroupInfo.Users [array]
		if err = buffer.WriteArray(1, groupInfo.Users); err != nil {
			return nil, err
		}

		// GProtoGroupInfo.Marknames [array]
		if err = buffer.WriteArray(1, groupInfo.Marknames); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (groupInfo *GProtoGroupInfo) Decode(version uint16, buf []byte) error {
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

		// GProtoGroupInfo.GroupId
		if groupInfo.GroupId, err = buffer.ReadString(); err != nil {
			return err
		}
		// GProtoGroupInfo.Name
		if groupInfo.Name, err = buffer.ReadString(); err != nil {
			return err
		}
		// GProtoGroupInfo.Avatar
		if groupInfo.Avatar, err = buffer.ReadString(); err != nil {
			return err
		}
		// GProtoGroupInfo.Description
		if groupInfo.Description, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoGroupInfo.OwnerAccount
		if groupInfo.OwnerAccount, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoGroupInfo.MaxUser
		if err = buffer.ReadUInt16(&groupInfo.MaxUser); err != nil {
			return err
		}

		// GProtoGroupInfo.Data
		if groupInfo.Data, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoGroupInfo.Flag
		if err = buffer.ReadUInt8(&groupInfo.Flag); err != nil {
			return err
		}
		// GProtoGroupInfo.CreatedTimestamp
		if err = buffer.ReadInt64(&groupInfo.CreatedTimestamp); err != nil {
			return err
		}

		// GProtoGroupInfo.Users [array]
		var usersNum uint16
		if err = buffer.ReadUInt16(&usersNum); err != nil {
			return err
		} else {
			groupInfo.Users = make([]*GProtoOtherUser, usersNum, usersNum)
			for i := 0; i < int(usersNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						groupInfo.Users[i] = &GProtoOtherUser{}
						if err = groupInfo.Users[i].Decode(1, bufObj); err != nil {
							return err
						}
					}
				}
			}
		}

		// GProtoGroupInfo.Marknames [array]
		var marknamesNum uint16
		if err = buffer.ReadUInt16(&marknamesNum); err != nil {
			return err
		} else {
			groupInfo.Marknames = make([]*GProtoGroupMarkname, marknamesNum, marknamesNum)
			for i := 0; i < int(marknamesNum); i++ {
				var lenObj uint16
				if err = buffer.ReadUInt16(&lenObj); err != nil {
					return err
				} else {
					if bufObj, err := buffer.ReadRawBytes(lenObj); err != nil {
						return err
					} else {
						groupInfo.Marknames[i] = &GProtoGroupMarkname{}
						if err = groupInfo.Marknames[i].Decode(1, bufObj); err != nil {
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

//////////////////////////
// message GroupMarkname {
//     required string account = 1;
//     required string markname = 2;
// }

type GProtoGroupMarkname struct {
	Account  string
	Markname string
}

func (groupMarkname *GProtoGroupMarkname) Encode(version uint16) ([]byte, error) {
	if version == 1 {
		// New Buffer
		buffer := NewEmptyGBuffer()

		/////////////
		// Encode  //
		/////////////
		var err error

		// GProtoGroupMarkname.Account
		if err = buffer.WriteString(groupMarkname.Account); err != nil {
			return nil, err
		}

		// GProtoGroupMarkname.Markname
		if err = buffer.WriteString(groupMarkname.Markname); err != nil {
			return nil, err
		}

		return buffer.Bytes(), nil
	}

	return nil, InvalidVersionError
}

func (groupMarkname *GProtoGroupMarkname) Decode(version uint16, buf []byte) error {
	// 合法性判断
	if len(buf) < 10 {
		return InvalidDecodeBufferError
	}

	if version == 1 {
		// New Buffer
		buffer := NewGBuffer(buf)

		/////////////
		// Decode  //
		/////////////
		var err error

		// GProtoGroupMarkname.Account
		if groupMarkname.Account, err = buffer.ReadString(); err != nil {
			return err
		}

		// GProtoGroupMarkname.Markname
		if groupMarkname.Markname, err = buffer.ReadString(); err != nil {
			return err
		}

		return nil
	}

	return InvalidVersionError
}
