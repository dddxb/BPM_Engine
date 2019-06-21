package model

import "time"

// token结构体
type Token struct {
	ID           string   `xorm:"varchar(32) pk notnull unique 'ID'"`
	PNID         string   `xorm:"varchar(64) 'PNID'"`        //此token所属PN
	Attribute    string   `xorm:"varchar(64) 'ATTRIBUTE'"`   //token属性，所有前置place必须拥有相同属性的token才可执行transition
	Desc         string   `xorm:"varchar(255) 'DESC'"`       //此token描述

	CreatedAt   time.Time `xorm:"created"`  //此token的初始创建时间
	UpdatedAt   time.Time `xorm:"updated"`  //此token的更新时间
	DeletedAt   time.Time `xorm:"deleted"`  //此token的删除时间
}

// TableName TableName
func (Token) TableName() string {
	return "token"
}
