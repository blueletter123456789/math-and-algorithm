def calc(n):
    res = 1 if n else 0
    while n > 0:
        res *= n % 10
        n //= 10
    return res
# n, b = map(int, input().split())
# ans = 0
# for i in range(n):
#     t = calc(i)
#     if i - t == b:
#         print(i, t, b)

def func(digit, m):
	if digit == 11:
		return {calc(m)}
	
	min_value = m % 10
	ret = set()
	for i in range(min_value, 10):
		r = func(digit + 1, m * 10 + i)
		for j in r:
			ret.add(j)
	return ret

f_list = func(0, 0)
n, b = map(int, input().split())
ans = 0
for fm in f_list:
    m = fm + b
    prd = calc(m)
    if m - prd == b and m <= n:
        ans += 1
print(ans)
