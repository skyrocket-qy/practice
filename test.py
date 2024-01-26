class Solution:
    def numWays(self, steps: int) -> int:

        arrLen = steps // 2 + 1
        print(arrLen)
        dp = [0] * arrLen
        dp[0] = 1

        for step in range(1, steps + 1):
            left = 0
            print(step)
            
            for i in range( min(arrLen, steps - step + 1) ):
                tmp = dp[i]
                dp[i] = left + dp[i] + (dp[i + 1] if i + 1 < arrLen else 0)
                left = tmp
                print(dp)

        return dp[0] % 1000_000_007
    
so = Solution()
so.numWays(4)