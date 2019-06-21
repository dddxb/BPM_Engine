package model

import "time"

//变迁结构体
type Transition struct {
	ID               string   `xorm:"varchar(64) pk notnull unique 'ID'"`
	PNID             string   `xorm:"varchar(64) 'PNID'"`               //此变迁所属PN
	Desc             string   `xorm:"varchar(255) 'DESC'"`              //transition描述
	ForwardPlaceID    []string  `xorm:"varchar(255) 'FORWARDPLACEID'"`  //前向place的ID，例["001","002"]
	BackwardPlaceID   []string  `xorm:"varchar(255) 'BACKWARDPLACEID'"` //后向place的ID，例["003","004"]
	Condition          string   `xorm:"varchar(255) 'CONDITION'"`       //此transition执行的条件，判断是否执行此transition,空表示无条件执行(前置条件)
	Result             string   `xorm:"varchar(255) 'RESULT'"`          //此transition执行后的结果，作为选取后续place的条件(后续条件)

	CreatedAt   time.Time `xorm:"created"`  //此transition的初始创建时间
	UpdatedAt   time.Time `xorm:"updated"`  //此transition的更新时间
	DeletedAt   time.Time `xorm:"deleted"`  //此transition的删除时间
}

// TableName TableName
func (Transition) TableName() string {
	return "transition"
}