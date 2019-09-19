# idCardParser-go
身份证号码校验、生日、性别，支持15和18位

身份证编码规则如下：根据〖中华人民共和国国家标准GB11643-1999〗中有关公民身份号码的规定，公民身份号码是特征组合码，由十七位数字本体码和一位数字校验码组成。

用法

go get github.com/lieney/idCardParser-go

//用法

package main

import (
	"fmt"
	"github.com/lieney/idCardParser-go"
)

func main() {

	//设置身份证号
	idCardParser.SetIdCard("111111111111111")

	//获取生日
	birthday := idCardParser.GetBirthday()

	//验证是否合法
	isTrue := idCardParser.IsValidate("111111111111111")

	//获取性别
	gender := idCardParser.GetGender()
	fmt.Println(birthday, isTrue, gender)

}


