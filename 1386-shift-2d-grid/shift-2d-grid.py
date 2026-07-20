class Solution:
    def shiftGrid(self, grid: List[List[int]], k: int) -> List[List[int]]:
        """
        flatten into 1-d array then
        transfrm back into matrix
        """
        x = len(grid)
        y = len(grid[0])

        k = k%(x*y)

        flat = []
        for row in grid:
            flat.extend(row)
        rotated = flat[-k:] + flat[:-k] if k>0 else flat

        res= []
        for i in range(x):
            res.append(rotated[i*y : (i+1)*y])
        return res
