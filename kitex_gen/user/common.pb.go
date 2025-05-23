// Code generated by Kitex v0.13.1. DO NOT EDIT.

package user

import "github.com/cloudwego/prutal"

/* 类型定义 */
type Unit struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Phone      string `protobuf:"bytes,2,opt,name=phone" json:"phone,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Address    string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Contact    string `protobuf:"bytes,6,opt,name=contact" json:"contact,omitempty"`
	Level      int32  `protobuf:"varint,7,opt,name=level" json:"level,omitempty"`
	Status     string `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	CreateTime int64  `protobuf:"varint,9,opt,name=createTime" json:"createTime,omitempty"`
	UpdateTime int64  `protobuf:"varint,10,opt,name=updateTime" json:"updateTime,omitempty"`
	DeleteTime int64  `protobuf:"varint,11,opt,name=deleteTime" json:"deleteTime,omitempty"`
}

func (x *Unit) Reset() { *x = Unit{} }

func (x *Unit) Marshal(in []byte) ([]byte, error) { return prutal.MarshalAppend(in, x) }

func (x *Unit) Unmarshal(in []byte) error { return prutal.Unmarshal(in, x) }

func (x *Unit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Unit) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Unit) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Unit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Unit) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Unit) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *Unit) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Unit) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Unit) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Unit) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *Unit) GetDeleteTime() int64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

type User struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Phone      string `protobuf:"bytes,2,opt,name=phone" json:"phone,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Birth      string `protobuf:"bytes,5,opt,name=birth" json:"birth,omitempty"`
	Gender     string `protobuf:"bytes,6,opt,name=gender" json:"gender,omitempty"`
	Status     string `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	CreateTime int64  `protobuf:"varint,8,opt,name=createTime" json:"createTime,omitempty"`
	UpdateTime int64  `protobuf:"varint,9,opt,name=updateTime" json:"updateTime,omitempty"`
	DeleteTime int64  `protobuf:"varint,10,opt,name=deleteTime" json:"deleteTime,omitempty"`
}

func (x *User) Reset() { *x = User{} }

func (x *User) Marshal(in []byte) ([]byte, error) { return prutal.MarshalAppend(in, x) }

func (x *User) Unmarshal(in []byte) error { return prutal.Unmarshal(in, x) }

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetBirth() string {
	if x != nil {
		return x.Birth
	}
	return ""
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *User) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *User) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *User) GetDeleteTime() int64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

type UnitMembership struct {
	UnitId string `protobuf:"bytes,1,opt,name=unitId" json:"unitId,omitempty"`
	Level  int32  `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
}

func (x *UnitMembership) Reset() { *x = UnitMembership{} }

func (x *UnitMembership) Marshal(in []byte) ([]byte, error) { return prutal.MarshalAppend(in, x) }

func (x *UnitMembership) Unmarshal(in []byte) error { return prutal.Unmarshal(in, x) }

func (x *UnitMembership) GetUnitId() string {
	if x != nil {
		return x.UnitId
	}
	return ""
}

func (x *UnitMembership) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type View struct {
	Id              string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Phone           string            `protobuf:"bytes,2,opt,name=phone" json:"phone,omitempty"`
	Password        string            `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Name            string            `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Address         string            `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
	Contact         string            `protobuf:"bytes,6,opt,name=contact" json:"contact,omitempty"`
	UnitMemberships []*UnitMembership `protobuf:"bytes,7,rep,name=unitMemberships" json:"unitMemberships,omitempty"`
	Status          string            `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
	CreateTime      int64             `protobuf:"varint,9,opt,name=createTime" json:"createTime,omitempty"`
	UpdateTime      int64             `protobuf:"varint,10,opt,name=updateTime" json:"updateTime,omitempty"`
	DeleteTime      int64             `protobuf:"varint,11,opt,name=deleteTime" json:"deleteTime,omitempty"`
}

func (x *View) Reset() { *x = View{} }

func (x *View) Marshal(in []byte) ([]byte, error) { return prutal.MarshalAppend(in, x) }

func (x *View) Unmarshal(in []byte) error { return prutal.Unmarshal(in, x) }

func (x *View) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *View) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *View) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *View) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *View) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *View) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

func (x *View) GetUnitMemberships() []*UnitMembership {
	if x != nil {
		return x.UnitMemberships
	}
	return nil
}

func (x *View) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *View) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *View) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *View) GetDeleteTime() int64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

type UnitVerifyForm struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (x *UnitVerifyForm) Reset() { *x = UnitVerifyForm{} }

func (x *UnitVerifyForm) Marshal(in []byte) ([]byte, error) { return prutal.MarshalAppend(in, x) }

func (x *UnitVerifyForm) Unmarshal(in []byte) error { return prutal.Unmarshal(in, x) }

func (x *UnitVerifyForm) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *UnitVerifyForm) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type UnitVerify struct {
	Id         string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Verify     int32             `protobuf:"varint,2,opt,name=verify" json:"verify,omitempty"`
	Account    string            `protobuf:"bytes,3,opt,name=account" json:"account,omitempty"`
	Password   string            `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	Form       []*UnitVerifyForm `protobuf:"bytes,5,rep,name=form" json:"form,omitempty"`
	Status     string            `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
	CreateTime int64             `protobuf:"varint,7,opt,name=createTime" json:"createTime,omitempty"`
	UpdateTime int64             `protobuf:"varint,8,opt,name=updateTime" json:"updateTime,omitempty"`
	DeleteTime int64             `protobuf:"varint,9,opt,name=deleteTime" json:"deleteTime,omitempty"`
}

func (x *UnitVerify) Reset() { *x = UnitVerify{} }

func (x *UnitVerify) Marshal(in []byte) ([]byte, error) { return prutal.MarshalAppend(in, x) }

func (x *UnitVerify) Unmarshal(in []byte) error { return prutal.Unmarshal(in, x) }

func (x *UnitVerify) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UnitVerify) GetVerify() int32 {
	if x != nil {
		return x.Verify
	}
	return 0
}

func (x *UnitVerify) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *UnitVerify) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UnitVerify) GetForm() []*UnitVerifyForm {
	if x != nil {
		return x.Form
	}
	return nil
}

func (x *UnitVerify) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UnitVerify) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *UnitVerify) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *UnitVerify) GetDeleteTime() int64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}

type UnitModel struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Chat       string `protobuf:"bytes,2,opt,name=chat" json:"chat,omitempty"`
	Asr        string `protobuf:"bytes,3,opt,name=asr" json:"asr,omitempty"`
	Tts        string `protobuf:"bytes,4,opt,name=tts" json:"tts,omitempty"`
	Report     string `protobuf:"bytes,5,opt,name=report" json:"report,omitempty"`
	Option     string `protobuf:"bytes,6,opt,name=option" json:"option,omitempty"`
	Status     string `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	CreateTime int64  `protobuf:"varint,8,opt,name=createTime" json:"createTime,omitempty"`
	UpdateTime int64  `protobuf:"varint,9,opt,name=updateTime" json:"updateTime,omitempty"`
	DeleteTime int64  `protobuf:"varint,10,opt,name=deleteTime" json:"deleteTime,omitempty"`
}

func (x *UnitModel) Reset() { *x = UnitModel{} }

func (x *UnitModel) Marshal(in []byte) ([]byte, error) { return prutal.MarshalAppend(in, x) }

func (x *UnitModel) Unmarshal(in []byte) error { return prutal.Unmarshal(in, x) }

func (x *UnitModel) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UnitModel) GetChat() string {
	if x != nil {
		return x.Chat
	}
	return ""
}

func (x *UnitModel) GetAsr() string {
	if x != nil {
		return x.Asr
	}
	return ""
}

func (x *UnitModel) GetTts() string {
	if x != nil {
		return x.Tts
	}
	return ""
}

func (x *UnitModel) GetReport() string {
	if x != nil {
		return x.Report
	}
	return ""
}

func (x *UnitModel) GetOption() string {
	if x != nil {
		return x.Option
	}
	return ""
}

func (x *UnitModel) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UnitModel) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *UnitModel) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *UnitModel) GetDeleteTime() int64 {
	if x != nil {
		return x.DeleteTime
	}
	return 0
}
