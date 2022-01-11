# def main(n, l):
#     cnt = 0
    # for a in range(n):
    #     for b in range(a+1, n):
    #         for c in range(b+1, n):
    #             for d in range(c+1, n):
    #                 for e in range(d+1, n):
    #                     if l[a]+l[b]+l[c]+l[d]+l[e] == 1000:
    #                         cnt += 1

# if __name__ == '__main__':
#     n = int(input())
#     l = list(map(int, input().split()))
#     main(n, l)
N = int(input())
A = list(map(int, input().split()))
 
dp = [[0] * 1001 for _ in range(6)]
dp[0][0] = 1
 
for a in A:
    for k in [4, 3, 2, 1, 0]:
        for v in range(1001 - a):
            dp[k+1][v+a] += dp[k][v]
print(dp[5][1000])