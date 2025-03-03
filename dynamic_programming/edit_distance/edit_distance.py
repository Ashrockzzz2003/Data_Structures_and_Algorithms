def editDistance(i, j, s1, s2, dp):
    if i < 0:
        return j + 1
    if j < 0:
        return i + 1

    if dp[i][j] != -1:
        return dp[i][j]

    if s1[i] == s2[j]:
        dp[i][j] = editDistance(i - 1, j - 1, s1, s2, dp)
        return dp[i][j]

    insert = 1 + editDistance(i, j - 1, s1, s2, dp)
    delete = 1 + editDistance(i - 1, j, s1, s2, dp)
    replace = 1 + editDistance(i - 1, j - 1, s1, s2, dp)

    dp[i][j] = min(insert, delete, replace)
    return dp[i][j]


def minDistance(s1, s2) -> int:
    dp = [[-1 for _ in range(len(s2))] for _ in range(len(s1))]
    return editDistance(len(s1) - 1, len(s2) - 1, s1, s2, dp)

s1 = "intention"
s2 = "execution"
print(minDistance(s1, s2))

s1 = "horse"
s2 = "ros"
print(minDistance(s1, s2))


