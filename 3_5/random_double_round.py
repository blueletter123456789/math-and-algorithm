import random

avg_cnt = 0
for _ in range(5):
    cnt = 0
    for _ in range(10**6):
        x = random.random()*6
        y = random.random()*9
        dist1 = ((x-3)**2 + (y-3)**2)**0.5
        dist2 = ((x-3)**2 + (y-7)**2)**0.5
        if dist1 <= 3 or dist2 <= 2:
            cnt += 1
    avg_cnt += cnt
avg_cnt //= 5
print(f"{avg_cnt=}")

area = (6*9)*(avg_cnt/10**6)
print(f"{area=}")
