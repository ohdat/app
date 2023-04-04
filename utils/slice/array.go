package utils

// Duplication  数组去重
func Duplication[T int | int32 | int64 | string](arr []T) []T {
	var newArr []T
	for _, v := range arr {
		if !Contains(newArr, v) {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

// Contains  判断数组中是否包含某个值
func Contains[T int | int32 | int64 | string](arr []T, target T) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
