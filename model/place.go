package model

import "time"

// 库所结构体
type Place struct {
	ID           string   `xorm:"varchar(64) pk notnull unique 'ID'"`
	PNID         string   `xorm:"varchar(64) 'PNID'"`          //此库所所属PN
	Desc         string   `xorm:"varchar(255) 'DESC'"`         //此库所表示的状态描述
	Token        []Token     `xorm:"varchar(255) notnull 'TOKEN'"`    //此place所包含token切片
	ForwardTranID    []string	`xorm:"varchar(255) 'FORWARDTASKID'"`   //前向transition的ID，-1表示此为起始库所，例["2","3"]
	BackwardTranID   []string   `xorm:"varchar(255) 'BACKWARDTASKID'"`  //后向transition的ID，-2表示此为结束库所，例["3","4"]
	Result           string     `xorm:"varchar(255) 'RESULT'"`          //此place执行后的结果，作为选取后续transition的条件

	CreatedAt   time.Time `xorm:"created"`  //此库所的初始创建时间
	UpdatedAt   time.Time `xorm:"updated"`  //此库所的更新时间
	DeletedAt   time.Time `xorm:"deleted"`  //此库所的删除时间
}

// TableName TableName
func (Place) TableName() string {
	return "place"
}