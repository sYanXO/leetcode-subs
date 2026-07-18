class Solution:
    def findGCD(self, nums: List[int]) -> int:
        def gcd(a,b):
            while b:
                t = a
                a = b
                b = t % b
            return a
        return gcd(max(nums),min(nums))