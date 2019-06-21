package main

import (
	"BPMEngine/model"
	"BPMEngine/pnrunner"
)

func main() {
	//token
	token1 := model.Token{
		ID:"token01",
		PNID:"traffic_light",
		Attribute:"",
		Desc:"",
	}

	//red1,yellow1,green1三个place
	red1 := model.Place{
		ID:"red1",
		PNID:"traffic_light",
		Desc:"red_light",
		Token:[]model.Token{token1},
		ForwardTranID:[]string{"y2r"},
		BackwardTranID:[]string{"r2g"},
		Result:"",    //此处无后续条件
	}
	yellow1 := model.Place{
		ID:"yellow1",
		PNID:"traffic_light",
		Desc:"yellow_light",
		Token:[]model.Token{},
		ForwardTranID:[]string{"g2y"},
		BackwardTranID:[]string{"y2r"},
		Result:"",    //此处无后续条件
	}
	green1 := model.Place{
		ID:"green1",
		PNID:"traffic_light",
		Desc:"green_light",
		Token:[]model.Token{},
		ForwardTranID:[]string{"r2g"},
		BackwardTranID:[]string{"g2y"},
		Result:"",   //此处无后续条件
	}

	//y2r,r2g,g2y三个transition
	 y2r:= model.Transition{
		ID:"y2r",
		PNID:"traffic_light",
		Desc:"yellow2red",
		ForwardPlaceID:[]string{"yellow1"},
		BackwardPlaceID:[]string{"red1"},
		Condition:"",    //此处无前置条件
		Result:"",       //此处无后续条件
	}
	r2g:= model.Transition{
		ID:"r2g",
		PNID:"traffic_light",
		Desc:"red2green",
		ForwardPlaceID:[]string{"red1"},
		BackwardPlaceID:[]string{"green1"},
		Condition:"",    //此处无前置条件
		Result:"",       //此处无后续条件
	}
	g2y:= model.Transition{
		ID:"g2y",
		PNID:"traffic_light",
		Desc:"green2yellow",
		ForwardPlaceID:[]string{"green1"},
		BackwardPlaceID:[]string{"yellow1"},
		Condition:"",    //此处无前置条件
		Result:"",       //此处无后续条件
	}
	//获得库所和任务数组
	p := []model.Place{red1,yellow1,green1}
	t := []model.Transition{y2r,r2g,g2y}

	//fmt.Println(p)
	//fmt.Println(t)

	//调用函数遍历
	pnrnner.ContinueIfPossible(p,t)
}
