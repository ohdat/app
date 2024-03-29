package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

// DeIn 去掉交集
func DeIn(ids1, ids2 []int) (res1 []int, res2 []int) {
	var ids2Map = make(map[int]int)
	for i := 0; i < len(ids2); i++ {
		ids2Map[ids2[i]] = ids2[i]
	}
	if len(ids2Map) > 0 {
		for i := 0; i < len(ids1); i++ {
			if _, ok := ids2Map[ids1[i]]; ok {
				delete(ids2Map, ids1[i])
				ids1 = append(ids1[:i], ids1[i+1:]...)
				i-- // form the remove item index to start iterate next item
			}
		}
	}
	var ids3 = make([]int, 0)
	if len(ids2Map) > 0 {
		for _, id_ := range ids2Map {
			ids3 = append(ids3, id_)
		}
	}

	return ids1, ids3
}

func StrToInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	id, err := strconv.Atoi(str)
	return id, err
}
func ArrStringToMapInt32(arr []string) map[int32]string {

	var maps = make(map[int32]string)
	for i, v := range arr {
		maps[int32(i)] = v
	}
	return maps
}

func Str2ArrInt(str string) (arr []int, err error) {
	//字符串分割成数组
	arrStr := strings.Split(str, ",")
	for _, v := range arrStr {
		if v == "" {
			continue
		}
		id, errs := StrToInt(v)
		if errs != nil {
			continue
		}
		arr = append(arr, id)
	}
	return
}

// IsDev 是否是开发环境
func IsDev() bool {
	return viper.GetString("environment") == "development"
}

// NotLogin 是否不需要登录
func NotLogin() bool {
	if IsDev() && viper.GetInt("notlogin") == 1 {
		return true
	}
	return false
}

func Struct2struct(in interface{}, out interface{}) {
	s, _ := json.Marshal(in)
	json.Unmarshal(s, out)
}

// StructAssign
// binding type interface 要修改的结构体
// value type interface 有数据的结构体
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			//验证类型
			if tok := bVal.FieldByName(name).Type().AssignableTo(vTypeOfT.Field(i).Type); tok {
				bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
			}
		}
	}
}

// DuplicationArrInt 数组去重
func DuplicationArrInt(arr []int) []int {
	var newArr []int
	for _, v := range arr {
		if !ContainsInt(newArr, v) {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

// ContainsInt 判断数组中是否包含某个值
func ContainsInt(arr []int, target int) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func ArrIntToString(aids []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(aids)), ","), "[]")
}
