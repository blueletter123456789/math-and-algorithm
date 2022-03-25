M = 998244353
def calc_base(n):
    return (n*(n+1)//2)%M

a, b, c = map(int, input().split())
total = (calc_base(a)*calc_base(b)*calc_base(c))%M
print(total)
