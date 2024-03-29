package arrays

const (
	LOW  = 0
	HIGH = 1
)

// BinarySearch 二分查找
func BinarySearch(arr []uint32, target uint32) bool {
	if arr == nil {
		return false
	}

	low := 0
	high := len(arr) - 1
	if arr[low] == target {
		return true
	}

	if arr[high] == target {
		return true
	}

	for low < high {
		mid := (low + high) >> 1
		if arr[mid] == target {
			return true
		}

		if arr[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return arr != nil && arr[low] == target
}

func ArrayUint32Exists(arr []uint32, target uint32) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}

func ArrayStringExists(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

// MergeArrayUint32 合并两个数组
func MergeArrayUint32(target []uint32, source []uint32) []uint32 {

	for _, val := range source {
		if !BinarySearch(target, val) {
			target = append(target, val)
		}
	}
	return target
}

func Find(arr []uint32, target uint32) int {
	for index, v := range arr {
		if v == target {
			return index
		}
	}
	return -1
}
