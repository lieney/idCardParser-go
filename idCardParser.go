package idCardParser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	idCardNum   string
	idCardLenth int
)
var salt = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
var checksum = [...]string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}
var genderMale int = 1
var genderFemale int = 0

func SetIdCard(idCard string) error {
	if idCard == "" {
		return errors.New("param idCard must not be empty")
	}
	idCardLenth = len(idCard)

	if idCardLenth != 18 && idCardLenth != 15 {
		return errors.New("param idCard error")
	}

	idCardNum = idCard
	return nil
}

func IsValidate(idCard string) bool {
	if idCardNum == "" {
		SetIdCard(idCard)
	}
	if checkFormat() && checkBirthday() && checkLastCode() {
		return true
	}
	return false
}

func GetGender() (gender int) {

	if idCardLenth == 18 {
		temp := fmt.Sprintf("%c", idCardNum[16])
		gender, _ = strconv.Atoi(temp)
	} else {
		temp := fmt.Sprintf("%c", idCardNum[14])
		gender, _ = strconv.Atoi(temp)
	}
	if gender%2 == 0 {
		gender = genderFemale
	} else {
		gender = genderMale
	}
	return
}

func checkFormat() bool {
	res, err := regexp.MatchString(`^([\d]{17}[xX\d]|[\d]{15})$`, idCardNum)
	if err != nil {
		return false
	}
	return res
}

func checkBirthday() bool {
	birthday := GetBirthday()
	pattern := `(([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})(((0[13578]|1[02])(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)(0[1-9]|[12][0-9]|30))|(02(0[1-9]|[1][0-9]|2[0-8]))))|((([0-9]{2})(0[48]|[2468][048]|[13579][26])|((0[48]|[2468][048]|[3579][26])00))0229)`
	res, err := regexp.MatchString(pattern, birthday)
	if err != nil {
		return false
	}
	return res
}

func checkLastCode() bool {
	if idCardLenth == 15 {
		return true
	}
	sum := 0
	for i := 0; i < 17; i++ {
		temp := fmt.Sprintf("%c", idCardNum[i])
		num, _ := strconv.Atoi(temp)
		sum += num * salt[i]
	}
	seek := sum % 11
	if checksum[seek] != strings.ToUpper(fmt.Sprintf("%c", idCardNum[17])) {
		return false
	}
	return true
}

func GetBirthday() (birthday string) {
	if idCardLenth == 18 {
		birthday = idCardNum[6:14]
	} else {
		for i := 0; i < 8; i++ {
			birthday = strings.Join([]string{"19", idCardNum[6:12]}, "")
		}
	}
	return
}
