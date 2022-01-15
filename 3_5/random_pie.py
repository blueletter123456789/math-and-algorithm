import random

loop_num = random.randint(1, 10**9)
print(loop_num)
cnt = 0
for _ in range(loop_num):
    x = random.random()
    y = random.random()
    if (x**2 + y**2) <= 1:
        cnt += 1
print((4*cnt)/loop_num)
