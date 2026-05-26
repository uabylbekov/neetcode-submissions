func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)

    for i, num := range nums {
        complement := target - num
        if _, ok := seen[complement]; ok {
            return []int{seen[complement], i}
        }
        seen[num] = i
    }

    return []int{}
}
