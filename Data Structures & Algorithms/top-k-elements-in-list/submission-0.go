func topKFrequent(nums []int, k int) []int {
    countNums := make(map[int]int)

    for _, num := range nums {
        countNums[num]++
    }
 
    bucket := make([][]int, len(nums) + 1)

    for key, val := range countNums {
        bucket[val] = append(bucket[val], key)
    }

    result := []int{}

    for i := len(bucket) - 1; i > -1; i-- {
        for _, elm := range bucket[i] {
            result = append(result, elm)
            if len(result) == k {
                return result
            }
        }
    }

    return []int{}
}
