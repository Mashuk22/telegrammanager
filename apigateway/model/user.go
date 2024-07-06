package model

type User struct {
	Id           int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ChatId       int64  `protobuf:"varint,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Username     string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	FirstName    string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	RoleId       int64  `protobuf:"varint,6,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	IsSubscribed bool   `protobuf:"varint,7,opt,name=is_subscribed,json=isSubscribed,proto3" json:"is_subscribed,omitempty"`
	CreatedAt    string `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

type CreateUser struct {
	ChatId       int64  `protobuf:"varint,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	Username     string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	FirstName    string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName     string `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	RoleId       int64  `protobuf:"varint,6,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	IsSubscribed bool   `protobuf:"varint,7,opt,name=is_subscribed,json=isSubscribed,proto3" json:"is_subscribed,omitempty"`
}
