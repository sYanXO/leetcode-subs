class Solution:
    def gcdValues(self, nums: List[int], queries: List[int]) -> List[int]:
        m = max(nums)
        cnt = Counter(nums)
        exact = [0]*(m+1)
        for i in range(m,0,-1):
            values = 0
            for multiple in range(i,m+1,i):
                values += cnt[multiple]
                exact[i] = exact[i] - exact[multiple]
            exact[i] += values*(values-1)//2
        prefix = list(accumulate(exact)) 
        return [bisect_right(prefix,q) for q in queries]