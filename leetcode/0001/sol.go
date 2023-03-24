package sol

func twoSum(nums []int, target int) []int {
	mem := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		cor := target - num

		y, ok := mem[cor]
		if ok {
			return []int{y, i}
		}

		mem[num] = i
	}

	return nil
}
