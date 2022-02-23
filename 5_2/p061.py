n = int(input())
flg = True
for i in range(65):
    if n == (1 << i) - 1:
        flg = False
if flg:
    print('First')
else:
    print('Second')
