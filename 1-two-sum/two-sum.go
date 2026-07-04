func twoSum(nums []int, target int) []int {
    seen := make(map[int]int) // val-> index

    for i,n := range nums{
        complement := target - nums[i]
        if idx,ok := seen[complement]; ok{ //map[key] returns 2 values
            return []int{idx,i}
        }
        seen[n] = i 
    }
    return []int{}
}