package transition

import (
	"BPM_Engine/model"
	"fmt"
)

//使能(Enable)的算法，完全是按照 pn 的 token 算法。
//只要这个 Transition 的所有前置 place 中含有相同属性的 token，那么这个 Transition 就可以被激活
func Enable(tran model.Transition,p []model.Place) bool{
	//此tran的所有前向PlaceID，[]string
	forwardPlaceID := tran.ForwardPlaceID
	fmt.Println("遍历transition的前向库所的ID切片为：",forwardPlaceID)
	//获取对应的所有前向Place
	var forwardPlace []model.Place
	for i:=0;i<len(forwardPlaceID);i++ {
		for j:=0;j<len(p);j++ {
			if p[j].ID == forwardPlaceID[i] {
				forwardPlace = append(forwardPlace,p[j])
			}
		}
	}
	fmt.Println("遍历transition的前向库所的切片：",forwardPlace)
	//判断所有前向place的[]token中是否含有相同属性的token
	//enable transition只有一个前向库所
	if len(forwardPlace) == 1 {
		if len(forwardPlace[0].Token) > 0 {
			return true
		}else {
			return false
		}
	}else {//enable transition的前向库所个数大于1
		//TODO
	}
	return true
}