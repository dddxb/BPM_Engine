package place

import (
	"BPMEngine/model"
	"fmt"
)

func Tokenmove(enabletran []model.Transition,p []model.Place) []model.Place {
	//取随机下标
	//rand.Seed(time.Now().Unix())
	//m := rand.Intn(len(enabletran))
	//placesID := enabletran[m].ForwardPlaceID
	//for i:=0;i<len(placesID);i++ {
	//}

	//只有一个transition满足条件，此时所有前向库所具有相同属性的token减一，所有后向库所增加一个具有相同属性的token
	if len(enabletran) == 1 {
		//此transition只有一个前向库所，可能含有多个不同属性的token
		if len(enabletran[0].ForwardPlaceID) == 1 {
			var (
				forwardPlace model.Place
				moveToken model.Token
				backwardPlace []model.Place
			)
			forwardPlaceID := enabletran[0].ForwardPlaceID[0]
			fmt.Println("前向库所ID：",forwardPlaceID)
			for i:=0;i<len(p);i++ {
				if forwardPlaceID == p[i].ID {
					forwardPlace = p[i]
				}
			}
			fmt.Println("前向place：",forwardPlace)
			//如果此前向库所只有一个token
			if len(forwardPlace.Token) == 1 {
				//取出需要转移的token
				//此前向place置空
				moveToken = forwardPlace.Token[0]
				var a []model.Token
				forwardPlace.Token = a
				//找出所有的后向place
				for i:=0;i<len(enabletran[0].BackwardPlaceID);i++ {
					for j:=0;j<len(p);j++ {
						if enabletran[0].BackwardPlaceID[i] == p[j].ID {
							backwardPlace = append(backwardPlace,p[j])
						}
					}
				}
				fmt.Println("后向place:",backwardPlace)
				//将token转移进后向place的[]Token中
				for i:=0;i<len(backwardPlace);i++ {
					backwardPlace[i].Token = append(backwardPlace[i].Token,moveToken)
				}
				//所有发生改变的前向库所，后向库所切片
				//暂不考虑循环结构，即前后place不同
				var changedPlace []model.Place
				changedPlace = append(changedPlace,forwardPlace)
				changedPlace = append(changedPlace,backwardPlace...)
				return changedPlace
				fmt.Println(changedPlace)

			}else {
				//TODO
			}
		}else {
			//TODO
		}
	}
	return []model.Place{}
}
