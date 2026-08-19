class Solution:
    def longestNiceSubarray(self, nums: List[int]) -> int:
        l = 0
        used = 0
        ans = 0

        for r in range(len(nums)):
            while used & nums[r]:
                used = used ^ nums[l]
                l += 1
            used = used | nums[r]
            ans = max(ans,r-l+1)
        return ans
