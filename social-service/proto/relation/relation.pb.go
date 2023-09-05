// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: relation.proto

package rpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// first interface
type UserRelationActionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token      string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                              // 用户鉴权token
	ToUserId   int64  `protobuf:"varint,2,opt,name=to_user_id,json=toUserId,proto3" json:"to_user_id,omitempty"`     // 对方用户id
	ActionType int32  `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1-关注，2-取消关注
}

func (x *UserRelationActionReq) Reset() {
	*x = UserRelationActionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelationActionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelationActionReq) ProtoMessage() {}

func (x *UserRelationActionReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelationActionReq.ProtoReflect.Descriptor instead.
func (*UserRelationActionReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{0}
}

func (x *UserRelationActionReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UserRelationActionReq) GetToUserId() int64 {
	if x != nil {
		return x.ToUserId
	}
	return 0
}

func (x *UserRelationActionReq) GetActionType() int32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

type UserRelationActionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
}

func (x *UserRelationActionResp) Reset() {
	*x = UserRelationActionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelationActionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelationActionResp) ProtoMessage() {}

func (x *UserRelationActionResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelationActionResp.ProtoReflect.Descriptor instead.
func (*UserRelationActionResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{1}
}

func (x *UserRelationActionResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *UserRelationActionResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

// Second interface
type UserRelationFollowListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                  // 用户鉴权token
}

func (x *UserRelationFollowListReq) Reset() {
	*x = UserRelationFollowListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelationFollowListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelationFollowListReq) ProtoMessage() {}

func (x *UserRelationFollowListReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelationFollowListReq.ProtoReflect.Descriptor instead.
func (*UserRelationFollowListReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{2}
}

func (x *UserRelationFollowListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRelationFollowListReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserRelationFollowListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	UserList   []*User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`        // 用户信息列表
}

func (x *UserRelationFollowListResp) Reset() {
	*x = UserRelationFollowListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelationFollowListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelationFollowListResp) ProtoMessage() {}

func (x *UserRelationFollowListResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelationFollowListResp.ProtoReflect.Descriptor instead.
func (*UserRelationFollowListResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{3}
}

func (x *UserRelationFollowListResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *UserRelationFollowListResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *UserRelationFollowListResp) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                 // 用户id
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                              // 用户名称
	FollowCount     int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`            // 关注总数
	FollowerCount   int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`      // 粉丝总数
	IsFollow        bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`                     // true-已关注，false-未关注
	Avatar          string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`                                          //用户头像
	BackgroundImage string `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage,proto3" json:"background_image,omitempty"` //用户个人页顶部大图
	Signature       string `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`                                    //个人简介
	TotalFavorited  int64  `protobuf:"varint,9,opt,name=total_favorited,json=totalFavorited,proto3" json:"total_favorited,omitempty"`   //获赞数量
	WorkCount       int64  `protobuf:"varint,10,opt,name=work_count,json=workCount,proto3" json:"work_count,omitempty"`                 //作品数量
	FavoriteCount   int64  `protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`     //点赞数量
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *User) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *User) GetTotalFavorited() int64 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *User) GetWorkCount() int64 {
	if x != nil {
		return x.WorkCount
	}
	return 0
}

func (x *User) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

// Third interface
type UserRelationFollowerListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                  // 用户鉴权token
}

func (x *UserRelationFollowerListReq) Reset() {
	*x = UserRelationFollowerListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelationFollowerListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelationFollowerListReq) ProtoMessage() {}

func (x *UserRelationFollowerListReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelationFollowerListReq.ProtoReflect.Descriptor instead.
func (*UserRelationFollowerListReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{5}
}

func (x *UserRelationFollowerListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRelationFollowerListReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserRelationFollowerListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32   `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	UserList   []*User `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`        // 用户列表
}

func (x *UserRelationFollowerListResp) Reset() {
	*x = UserRelationFollowerListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelationFollowerListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelationFollowerListResp) ProtoMessage() {}

func (x *UserRelationFollowerListResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelationFollowerListResp.ProtoReflect.Descriptor instead.
func (*UserRelationFollowerListResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{6}
}

func (x *UserRelationFollowerListResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *UserRelationFollowerListResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *UserRelationFollowerListResp) GetUserList() []*User {
	if x != nil {
		return x.UserList
	}
	return nil
}

// Four interface
type UserRelationFriendListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 用户id
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`                  // 用户鉴权token
}

func (x *UserRelationFriendListReq) Reset() {
	*x = UserRelationFriendListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelationFriendListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelationFriendListReq) ProtoMessage() {}

func (x *UserRelationFriendListReq) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelationFriendListReq.ProtoReflect.Descriptor instead.
func (*UserRelationFriendListReq) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{7}
}

func (x *UserRelationFriendListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRelationFriendListReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserRelationFriendListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32         `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string        `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
	UserList   []*FriendUser `protobuf:"bytes,3,rep,name=user_list,json=userList,proto3" json:"user_list,omitempty"`        // 用户列表
}

func (x *UserRelationFriendListResp) Reset() {
	*x = UserRelationFriendListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRelationFriendListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRelationFriendListResp) ProtoMessage() {}

func (x *UserRelationFriendListResp) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRelationFriendListResp.ProtoReflect.Descriptor instead.
func (*UserRelationFriendListResp) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{8}
}

func (x *UserRelationFriendListResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *UserRelationFriendListResp) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *UserRelationFriendListResp) GetUserList() []*FriendUser {
	if x != nil {
		return x.UserList
	}
	return nil
}

type FriendUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                 // 用户id
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                              // 用户名称
	FollowCount     int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`            // 关注总数
	FollowerCount   int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`      // 粉丝总数
	IsFollow        bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`                     // true-已关注，false-未关注
	Avatar          string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`                                          //用户头像
	BackgroundImage string `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage,proto3" json:"background_image,omitempty"` //用户个人页顶部大图
	Signature       string `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`                                    //个人简介
	TotalFavorited  int64  `protobuf:"varint,9,opt,name=total_favorited,json=totalFavorited,proto3" json:"total_favorited,omitempty"`   //获赞数量
	WorkCount       int64  `protobuf:"varint,10,opt,name=work_count,json=workCount,proto3" json:"work_count,omitempty"`                 //作品数量
	FavoriteCount   int64  `protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`     //点赞数量
	Message         string `protobuf:"bytes,12,opt,name=message,proto3" json:"message,omitempty"`                                       // 和该好友的最新聊天消息
	MsgType         int64  `protobuf:"varint,13,opt,name=msgType,proto3" json:"msgType,omitempty"`                                      // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}

func (x *FriendUser) Reset() {
	*x = FriendUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relation_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendUser) ProtoMessage() {}

func (x *FriendUser) ProtoReflect() protoreflect.Message {
	mi := &file_relation_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendUser.ProtoReflect.Descriptor instead.
func (*FriendUser) Descriptor() ([]byte, []int) {
	return file_relation_proto_rawDescGZIP(), []int{9}
}

func (x *FriendUser) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FriendUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FriendUser) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *FriendUser) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *FriendUser) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *FriendUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *FriendUser) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *FriendUser) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *FriendUser) GetTotalFavorited() int64 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *FriendUser) GetWorkCount() int64 {
	if x != nil {
		return x.WorkCount
	}
	return 0
}

func (x *FriendUser) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *FriendUser) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *FriendUser) GetMsgType() int64 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

var File_relation_proto protoreflect.FileDescriptor

var file_relation_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x70, 0x62, 0x22, 0x6c, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x58, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x22, 0x4a, 0x0a,
	0x19, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x84, 0x01, 0x0a, 0x1a, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0xe1, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x29, 0x0a, 0x10,
	0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75,
	0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x86, 0x01, 0x0a, 0x1c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4d, 0x73, 0x67, 0x12, 0x26, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x4a, 0x0a, 0x19, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x1a, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70, 0x62, 0x2e,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x9b, 0x03, 0x0a, 0x0a, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72,
	0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x77, 0x6f,
	0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x32, 0xf3, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x72, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x59, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x72, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x72, 0x70,
	0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x5f, 0x0a, 0x18,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x72, 0x70, 0x62, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x72, 0x70, 0x62,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x59, 0x0a,
	0x16, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x72, 0x70, 0x62, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relation_proto_rawDescOnce sync.Once
	file_relation_proto_rawDescData = file_relation_proto_rawDesc
)

func file_relation_proto_rawDescGZIP() []byte {
	file_relation_proto_rawDescOnce.Do(func() {
		file_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_relation_proto_rawDescData)
	})
	return file_relation_proto_rawDescData
}

var file_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_relation_proto_goTypes = []interface{}{
	(*UserRelationActionReq)(nil),        // 0: rpb.UserRelationActionReq
	(*UserRelationActionResp)(nil),       // 1: rpb.UserRelationActionResp
	(*UserRelationFollowListReq)(nil),    // 2: rpb.UserRelationFollowListReq
	(*UserRelationFollowListResp)(nil),   // 3: rpb.UserRelationFollowListResp
	(*User)(nil),                         // 4: rpb.User
	(*UserRelationFollowerListReq)(nil),  // 5: rpb.UserRelationFollowerListReq
	(*UserRelationFollowerListResp)(nil), // 6: rpb.UserRelationFollowerListResp
	(*UserRelationFriendListReq)(nil),    // 7: rpb.UserRelationFriendListReq
	(*UserRelationFriendListResp)(nil),   // 8: rpb.UserRelationFriendListResp
	(*FriendUser)(nil),                   // 9: rpb.FriendUser
}
var file_relation_proto_depIdxs = []int32{
	4, // 0: rpb.UserRelationFollowListResp.user_list:type_name -> rpb.User
	4, // 1: rpb.UserRelationFollowerListResp.user_list:type_name -> rpb.User
	9, // 2: rpb.UserRelationFriendListResp.user_list:type_name -> rpb.FriendUser
	0, // 3: rpb.UserService.UserRelationAction:input_type -> rpb.UserRelationActionReq
	2, // 4: rpb.UserService.UserRelationFollowList:input_type -> rpb.UserRelationFollowListReq
	5, // 5: rpb.UserService.UserRelationFollowerList:input_type -> rpb.UserRelationFollowerListReq
	7, // 6: rpb.UserService.UserRelationFriendList:input_type -> rpb.UserRelationFriendListReq
	1, // 7: rpb.UserService.UserRelationAction:output_type -> rpb.UserRelationActionResp
	3, // 8: rpb.UserService.UserRelationFollowList:output_type -> rpb.UserRelationFollowListResp
	6, // 9: rpb.UserService.UserRelationFollowerList:output_type -> rpb.UserRelationFollowerListResp
	8, // 10: rpb.UserService.UserRelationFriendList:output_type -> rpb.UserRelationFriendListResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_relation_proto_init() }
func file_relation_proto_init() {
	if File_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelationActionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelationActionResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelationFollowListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelationFollowListResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelationFollowerListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelationFollowerListResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelationFriendListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRelationFriendListResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_relation_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendUser); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relation_proto_goTypes,
		DependencyIndexes: file_relation_proto_depIdxs,
		MessageInfos:      file_relation_proto_msgTypes,
	}.Build()
	File_relation_proto = out.File
	file_relation_proto_rawDesc = nil
	file_relation_proto_goTypes = nil
	file_relation_proto_depIdxs = nil
}