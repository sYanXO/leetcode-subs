func longestConsecutive(nums []int) int {
    // idea - The key: check num-1 first
    //If num-1 is in the set, you'll find this sequence later.
    //Only iterate chains starting from numbers with no predecessor


    if len(nums) == 0{
        return 0
    }
    
    numbers := make(map[int]bool)

    for _,num := range nums{
        numbers[num] = true;
    }

    ans :=0
    for num:= range numbers{
        if !numbers[num-1]{
            curr := num
            length := 1 
        
        // extend
        for
        {
            if !numbers[curr+1]{
                break
            }
            curr++
            length++
        }

            if length > ans{
                ans = length
            }
        }
    }
    return ans
}