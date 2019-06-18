package main

import (
	"encoding/json"
	"fmt"
)

// Person 其他库引用时, 必须是首字母大写的字段
// 字段被本包忽略
// Field int `json:"-"`
// // 字段在json里的键为"myName"
// Field int `json:"myName"`
// // 字段在json里的键为"myName"且如果字段为空值将在对象中省略掉
// Field int `json:"myName,omitempty"`
// // 字段在json里的键为"Field"（默认值），但如果字段为空值会跳过；注意前导的逗号
// Field int `json:",omitempty"`
// "string"选项标记一个字段在编码json时应编码为字符串。它只适用于字符串、浮点数、整数类型的字段。
// 这个额外水平的编码选项有时候会用于和javascript程序交互：
// Int64String int64 `json:",string"`
type Person struct {
	Name string `json:"name"`
	Age  uint
	like []string
	Sex  bool `json:",omitempty"`
}

func main06() {
	p := Person{"eiling", 17, []string{"music", "vodie"}, false}
	// {"name":"eiling","Age":17}
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json Marshal err: ", err)
	}
	fmt.Println(string(data))
	// {"name":"eiling","Age":17,"Sex":true} 其他库引用时, 必须是首字母大写的字段

	var p1 Person
	json.Unmarshal(data, &p1)
	fmt.Println(p1) // {eiling 17 [] false}

	d1 := make(map[string]interface{})
	json.Unmarshal(data, &d1)
	fmt.Println(d1) // map[Age:17 name:eiling]

	p3 := []Person{{"eiling", 17, []string{"music", "vodie"}, false},
		{"lin", 17, []string{"music", "vodie"}, true}}

	d2, _ := json.MarshalIndent(p3, "", "")
	fmt.Println(string(d2))
	// 	[
	// {
	// "name": "eiling",
	// "Age": 17
	// },
	// {
	// "name": "lin",
	// "Age": 17,
	// "Sex": true
	// }
	// ]
}
