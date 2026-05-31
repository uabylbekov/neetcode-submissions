import "slices"

func threeSum(nums []int) [][]int {
    triplets := make([][]int, 0)
    slices.Sort(nums)

    for i := 0; i < len(nums)-2; i++ {
        if i > 0 &&  nums[i - 1] == nums[i] {
            continue
        } 

        l := i + 1
        r := len(nums) - 1

        for l < r {
            if nums[i] + nums[l] + nums[r] == 0 {
                triplets = append(triplets, []int{nums[i], nums[l], nums[r]})
                l++
                r--
                for l < r && nums[l - 1] == nums[l] {
                    l++
                }
                for l < r && nums[r + 1] == nums[r] {
                    r--
                }
            } else if nums[i] + nums[l] + nums[r] < 0 {
                l++
            } else {
                r--
            }
        }
    }

    return triplets
}