t = int(input())
n = int(input())
diff = [0]*(t+1)
cs = [0]*(t+1)
for _ in range(n):
    l, r = map(int, input().split())
    diff[l] += 1
    diff[r] -= 1
cs[0] = diff[0]
print(cs[0])
for i in range(1, t):
    cs[i] = cs[i-1] + diff[i]
    print(cs[i])
