n, k = map(int, input().split())
l = [abs(int(i)) for i in input().split()]
sum_num = sum(l)
if sum_num <= k and sum_num % 2 == k % 2:
    print('Yes')
else:
    print('No')
