package gproto

//import (
//	"gimcloud/gim_data/db"
//	"gimcloud/skynet/misc"
//)

//func MessageRequestToMessageNotify(gProtoMessageRequest *GProtoMessageRequest) *GProtoMessageNotify {
//	timestamp := misc.UnixTimestamp()
//	return &GProtoMessageNotify{
//		MsgId:                  gProtoMessageRequest.MsgId,
//		MsgType:                gProtoMessageRequest.MsgType,
//		Content:                gProtoMessageRequest.Content,
//		From:                   gProtoMessageRequest.From,
//		To:                     gProtoMessageRequest.To,
//		SendTime:               uint64(timestamp),
//		SendTimeAck:            uint64(timestamp),
//		FileSize:               gProtoMessageRequest.FileSize,
//		FileSizeOptionalFlag:   gProtoMessageRequest.FileSizeOptionalFlag,
//		RecordTime:             gProtoMessageRequest.RecordTime,
//		RecordTimeOptionalFlag: gProtoMessageRequest.RecordTimeOptionalFlag,
//		Extern:                 gProtoMessageRequest.Extern,
//		ExternOptionalFlag:     gProtoMessageRequest.ExternOptionalFlag,
//	}
//}

//func DBUserToSelfUser(dbUser *db.DBUser) *GProtoSelfUser {
//	return &GProtoSelfUser{
//		Account:             dbUser.Account,
//		Nickname:            dbUser.Nickname,
//		Gender:              uint8(dbUser.Gender),
//		Avatar:              dbUser.Avatar,
//		Signature:           dbUser.Signature,
//		Extern:              dbUser.Extern,
//		CustomVersion:       dbUser.CustomVersion,
//		Data:                dbUser.Data,
//		LastOnlineTimestamp: dbUser.LastOnlineTimestamp,
//	}
//}

//func DBUserToOtherUser(dbUser *db.DBUser) *GProtoOtherUser {
//	return &GProtoOtherUser{
//		Account:             dbUser.Account,
//		Nickname:            dbUser.Nickname,
//		Gender:              uint8(dbUser.Gender),
//		Avatar:              dbUser.Avatar,
//		Signature:           dbUser.Signature,
//		Extern:              dbUser.Extern,
//		CustomVersion:       dbUser.CustomVersion,
//		Data:                dbUser.Data,
//		LastOnlineTimestamp: dbUser.LastOnlineTimestamp,
//		Markname:            "",
//		MakrnameFlag:        OPTIONAL_TRUE,
//	}
//}

//func DBUserToOtherUserWithMarkname(dbUser *db.DBUser, markname string) *GProtoOtherUser {
//	return &GProtoOtherUser{
//		Account:             dbUser.Account,
//		Nickname:            dbUser.Nickname,
//		Gender:              uint8(dbUser.Gender),
//		Avatar:              dbUser.Avatar,
//		Signature:           dbUser.Signature,
//		Extern:              dbUser.Extern,
//		CustomVersion:       dbUser.CustomVersion,
//		Data:                dbUser.Data,
//		LastOnlineTimestamp: dbUser.LastOnlineTimestamp,
//		Markname:            markname,
//		MakrnameFlag:        OPTIONAL_TRUE,
//	}
//}

//func DBFriendsToFriends(dbFriends []*db.DBFriend) []*GProtoFriend {
//	friends := make([]*GProtoFriend, len(dbFriends), len(dbFriends))
//	for i, dbFriend := range dbFriends {
//		friends[i] = DBFriendToFriend(dbFriend)
//	}
//	return friends
//}

//func DBFriendToFriend(dbFriend *db.DBFriend) *GProtoFriend {
//	var markNameFlag uint8 = OPTIONAL_FALSE
//	if dbFriend.Markname == "" {
//		markNameFlag = OPTIONAL_TRUE
//	}

//	return &GProtoFriend{
//		Account:      dbFriend.AccountF,
//		Flag:         uint8(dbFriend.Flag),
//		Markname:     dbFriend.Markname,
//		MakrnameFlag: markNameFlag,
//	}
//}

//func DBBlacklistsToBlacklists(dbBlacklists []*db.DBBlacklist) []*GProtoBlacklist {
//	blacklists := make([]*GProtoBlacklist, len(dbBlacklists), len(dbBlacklists))
//	for i, dbBlacklist := range dbBlacklists {
//		blacklists[i] = DBBlacklistToBlacklist(dbBlacklist)
//	}
//	return blacklists
//}

//func DBBlacklistToBlacklist(dbBlacklist *db.DBBlacklist) *GProtoBlacklist {
//	return &GProtoBlacklist{
//		Account: dbBlacklist.AccountF,
//		Flag:    uint8(dbBlacklist.Flag),
//	}
//}

//func DBChatRoomToChatRoomInfo(dbChatRoom *db.DBChatRoom) *GProtoChatRoomInfo {
//	return &GProtoChatRoomInfo{
//		ChatRoomId:       dbChatRoom.ChatRoomId,
//		Name:             dbChatRoom.Name,
//		Description:      dbChatRoom.Description,
//		OwnerAccount:     dbChatRoom.OwnerAccount,
//		MaxUser:          uint16(dbChatRoom.MaxUser),
//		Data:             dbChatRoom.Data,
//		CreatedTimestamp: dbChatRoom.Timestamp,
//	}
//}

//// 备注:
//// DBGroup.Users和DBGroup.Marknames需要额外赋值
//func DBGroupToGroupInfo(dbGroup *db.DBGroup) *GProtoGroupInfo {
//	return &GProtoGroupInfo{
//		GroupId:          dbGroup.GroupId,
//		Name:             dbGroup.Name,
//		Avatar:           dbGroup.Avatar,
//		Description:      dbGroup.Description,
//		OwnerAccount:     dbGroup.OwnerAccount,
//		MaxUser:          uint16(dbGroup.MaxUser),
//		Data:             dbGroup.Data,
//		Flag:             uint8(dbGroup.Flag),
//		CreatedTimestamp: dbGroup.Timestamp,
//	}
//}
