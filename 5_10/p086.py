n = int(input())
s = input()
cnt = 0
for i in s:
    if i == '(':
        cnt += 1
    elif i == ')':
        cnt -= 1
    if cnt < 0:
        break
if cnt == 0:
    print('Yes')
else:
    print('No')
