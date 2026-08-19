class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        last_seen = {}
        l = 0
        ans = 0

        for r,c in enumerate(s):
            if c in last_seen:
                l = max(l,last_seen[c]+1)
            last_seen[c] = r
            ans = max(ans,r-l+1)
        return ans