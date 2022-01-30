import math

def calc(x):
    ans = 0
    max_x = int(math.pow(x, 0.5))
    for i in range(1, max_x+1):
        if x % i == 0:
            ans += 1
    ans *= 2
    if max_x**2 == x:
        ans -= 1
    return ans

def main(n):
    res = 0
    for i in range(1, n+1):
        res += i * calc(i)
    print(res)

if __name__ == '__main__':
    n = int(input())
    main(n)


########################################
n = int(input())
l = [0]*(n+1)
for i in range(1, n+1):
    for j in range(i, n+1, i):
        l[j] += 1
ans = 0
for i in range(1, n+1):
    ans += i*l[i]
print(ans)

########################################
n=int(input())
ans=0
for i in range(1,n+1):
    d=n//i
    ans+=i*(d*(d+1)//2)
print(ans)
