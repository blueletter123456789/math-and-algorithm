M = 1000000007
n = int(input())
base = (n*(n+1)//2)%M
total = (base * base)%M
print(total)
