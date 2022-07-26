package thealogrithms

type BinarySearch struct {
	endIndex   int
	startIndex int
	mid        int
	Array      []int
	Target     int
}

func (binary *BinarySearch) calculatMid() int {

	return ((binary.endIndex + binary.startIndex) / 2)
}

func (binary *BinarySearch) fromRight() {

	if binary.Array[binary.mid] > binary.Target {
		binary.endIndex = binary.mid - 1
	} else {
		binary.startIndex = binary.mid + 1
	}
}

func (binary *BinarySearch) fromLeft() {

	if binary.Array[binary.mid] < binary.Target {
		binary.startIndex = binary.mid + 1
	} else {

		binary.endIndex = binary.mid - 1
	}
}

func (binary *BinarySearch) getResult(rtl bool) int {

	if rtl && binary.endIndex != -1 && binary.Array[binary.endIndex] == binary.Target {

		return binary.endIndex
	} else if !rtl && binary.Array[binary.startIndex] == binary.Target {

		return binary.startIndex
	}

	return -1
}

/**
 * Binary Search
 * @param rtl true if search from right
 * @return index of target
 */
func (binary *BinarySearch) Search(rtl bool) int {

	binary.endIndex = len(binary.Array) - 1
	binary.startIndex = 0

	for binary.startIndex <= binary.endIndex {
		binary.mid = binary.calculatMid()

		if binary.Array[binary.mid] == binary.Target {

			return binary.mid
		}

		if rtl {
			binary.fromRight()

		} else {

			binary.fromLeft()
		}

	}

	return binary.getResult(rtl)
}

/**
 * Find Index by Binary Search
 *
 * @param nums []int
 * @param target int
 * @param rtl true if search from right
 * @return index of target
 */
func FindBytBinarySearch(array []int, target int, rtl bool) int {

	result := BinarySearch{Array: array, Target: target}
	return result.Search(rtl)
}
