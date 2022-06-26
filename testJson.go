package main

import (
	"encoding/json"
	"fmt"
)

type TestData struct {
	ID   int
	Msg  string
	Data interface{}
}

// json.Marshal ： 是将结构体转为json格式的字节码，通过string函数把字节码转换为 json格式的字符串
// json.Unmarshal: 是将 json 格式的字节码转换为 结构体，并必须存储到结构体的引用中
func testJson() {
	data := TestData{
		ID:   0,
		Msg:  "OK",
		Data: nil,
	}
	m := make(map[string]interface{})
	v, _ := json.Marshal(&data) //v：[]byte类型
	var uv TestData
	err := json.Unmarshal(v, &uv)
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
	m["result"] = v
	fmt.Println("v的原始值：", v)
	fmt.Println("v的string值：", string(v))
	fmt.Println("uv值：", uv)
	bytes, _ := json.Marshal(&m)
	fmt.Println("m: ", string(bytes))
}
