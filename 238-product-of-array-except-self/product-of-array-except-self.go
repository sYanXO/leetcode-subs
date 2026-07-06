func productExceptSelf(nums []int) []int {
    n := len(nums)
    ans := make([]int,n)

    ans[n-1] = 1 // test right pass first
    for i:= n-2; i>=0 ;i--{
        ans[i] = ans[i+1]*nums[i+1]
    }

    leftprod := 1
    for i:=0; i< n;i++{
        ans[i] = leftprod*ans[i]
        leftprod *= nums[i]
    }
    return ans
    
}