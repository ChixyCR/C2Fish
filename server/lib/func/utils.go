package lib

import (
	"encoding/base32"
	"encoding/base64"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// Time
func GetTimeUnix(Nano bool) int64{
	if Nano{
		return time.Now().UnixNano()
	}
	return time.Now().Unix()

}
func GetTimeDateTime()string{
	return time.Now().Format("2006-01-02 15:04:05")
}
func GetTimeDateTimeFromUnix(ts int64)string{
	return time.Unix(ts, 0).Format("2006-01-02 15:04:05")
}
func GetTimeUnixFromTimeDate(s string,Nano bool)int64{
	format,_ := time.Parse("2006-01-02 15:04:05",s)
	if Nano{
		return format.UnixNano()
	}
	return format.Unix()
}


func IsValueList(value string, list []string) bool {

	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func IsAnyNil(i ...interface{}) bool{
	for _,val := range i{
		switch val.(type) {
		case nil:
			return true
		default:
			continue
		}
	}
	return false
}

func MergeSliceStringValue(a ...[]string) (reValue []interface{}){
	reValue = make([]interface{},0)
	for _,each := range a{
		for _,val := range each{
			reValue = append(reValue,val)
		}
	}
	return
}

// get random str
func GetRandStr(n int,chars string) ( randStr string){
	if chars == ""{
		chars = "ABCDEFGHIJKMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz23456789"
	}
	charsLen := len(chars)
	rand.Seed(GetTimeUnix(true))
	for i := 0; i < n;i ++{
		randIndex := rand.Intn(charsLen)
		randStr += chars[randIndex:randIndex+1]
	}
	return
}

// base
func BaseEn(s []byte,b int) string{
	switch b {
	case 32:
		return base32.StdEncoding.EncodeToString(s)
	default:
		return base64.StdEncoding.EncodeToString(s)
	}
}

func BaseDe(s string,b int)([] byte,error){
	switch b {
	case 32:
		return base32.StdEncoding.DecodeString(s)
	default:
		return base64.StdEncoding.DecodeString(s)
	}
}

//module verify
func VerifyModuleType(moduleString string) string {

	if ok, _ := regexp.MatchString("^userModule", strings.TrimSpace(moduleString)); ok {
		return "user"
	}
	if ok, _ := regexp.MatchString("^publicModule", strings.TrimSpace(moduleString)); ok {
		return "public"
	}

	return "error"

}

// bool verify
func VerifyBooleanType(b string) bool{

	if strings.ToLower(b) == "true"{
		return true
	}
	return false

}

// 邮箱格式验证
func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
// 手机格式验证
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
// 用户名格式验证
func CheckUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9_]{4,32}$", strings.TrimSpace(username)); !ok {
		return false
	}
	return true
}
// 自定义模块名称验证
func CheckCustomModuleName(moduleName string) (b bool){
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9_]{1,64}$", strings.TrimSpace(moduleName)); !ok {
		return false
	}
	return true
}

// 任务名称验证
func CheckTaskName(taskName string) (b bool){
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9_]{1,64}$", strings.TrimSpace(taskName)); !ok {
		return false
	}
	return true
}

// API随机字符验证
func CheckRandomStr(r string) (b bool){
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9_]{6,32}$", strings.TrimSpace(r)); !ok {
		return false
	}
	return true
}