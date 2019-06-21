package pnrnner

import (
	"BPMEngine/model"
	"BPMEngine/place"
	"BPMEngine/transition"
	"fmt"
)

//遍历transition判断token的移动
func ContinueIfPossible (p []model.Place,t []model.Transition) {
	fmt.Println(p)
	fmt.Println(t)
	//遍历transition
	var enableTran []model.Transition
	for i:=0;i<len(t);i++ {
		//如果允许使能，则分情况执行，这里存在 PetriNet 的经典算法:
		//任何一次 token 的位置变迁（从 place 中移出 token），都会引起整个遍历的重新执行
		if  transition.Enable(t[i],p) {
			//前置place可能有多个token，可能有多个处于使能(Enable)状态的transition，先把所有处于使能状态的transition置于切片中，
			//再考虑执行逻辑,后续解决冲突、竞争
			enableTran = append(enableTran,t[i])
		}
	}
	fmt.Println("可以使能的transition切片enableTran：",enableTran)
	//执行token转移逻辑，随机选取使能transition并执行
	changPlace := place.Tokenmove(enableTran,p)
	fmt.Println("发生改变的前后向place切片(前后向place不同)changePlace：",changPlace)
	//token发生移动，此PN需要重新遍历
	//所有的transition不会改变，始终为 t []model.Transition
	//部分place发生改变，p为新的切片
	for i:=0;i<len(p);i++ {
		for j:=0;j<len(changPlace);j++ {
			if changPlace[j].ID == p[i].ID {
				p = append(p[:i],p[i+1:]...)
				break
			}
		}
	}
	fmt.Println("没有改变的place：",p)
	for i:=0;i<len(changPlace);i++ {
		p = append(p,changPlace[i])
	}
	fmt.Println("遍历后的新的place切片p：",p)
	fmt.Println("所有的transition变迁切片t：",t)
	fmt.Println("")
	ContinueIfPossible(p,t)
}

















//map转slice
//enabletran := make([]model.Transition,len(enableTran))
//i := 0
//for _ ,v := range enableTran {
//	enabletran[i] = v
//	i++
//}