class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        minprice = float('inf')
        maxprofit = 0

        for p in prices:
            minprice = min(p,minprice)
            maxprofit = max(maxprofit,p-minprice)
        return maxprofit