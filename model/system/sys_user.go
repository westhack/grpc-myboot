package system

import (
	"bytes"
	"encoding/gob"
	"grpc-myboot/global"
	"log"
)

type SysUser struct {
	global.Model
	Username string `json:"username" gorm:"comment:用户登录名"` // 用户登录名
	Password string `json:"-" gorm:"comment:用户登录密码"`       // 用户登录密码
	Mobile   string `json:"mobile" gorm:"comment:手机号"`     // 用户登录密码
	Nickname string `json:"nickname" gorm:"comment:用户昵称"`
}

// 编码，把结构体数据编码成字节流数据
func (u *SysUser) Encode() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf) // 构造编码器，并把数据写进buf中
	if err := encoder.Encode(u); err != nil {
		log.Printf("encode error: %v\n", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

// 解码，把字节流数据解析成结构体数据
func (u *SysUser) Decode(b []byte) (SysUser, error) {
	//var buf bytes.Buffer
	bufPtr := bytes.NewBuffer(b)      // 返回的类型是 *Buffer，而不是 Buffer。注意一下
	decoder := gob.NewDecoder(bufPtr) // 从 bufPtr 中获取数据
	var p SysUser
	if err := decoder.Decode(&p); err != nil { // 将数据写进变量 p 中
		return p, err
	}
	return p, nil
}
