import queue

r, c = map(int, input().split())
s = list(map(int, input().split()))
g = list(map(int, input().split()))
c_map = [list(input()) for _ in range(r)]
dis = [[-1]*(c) for _ in range(r)]

dis[s[0]-1][s[1]-1] = 0
que = queue.Queue()
que.put((s[0]-1, s[1]-1))
while not que.empty():
    y, x = que.get()
    if y == g[0]-1 and x == g[1]-1:
        break
    if c_map[y][x-1] == '.' and dis[y][x-1] == -1:
        dis[y][x-1] = dis[y][x] + 1
        que.put((y, x-1))
    if c_map[y-1][x] == '.' and dis[y-1][x] == -1:
        dis[y-1][x] = dis[y][x] + 1
        que.put((y-1, x))
    if c_map[y][x+1] == '.' and dis[y][x+1] == -1:
        dis[y][x+1] = dis[y][x] + 1
        que.put((y, x+1))
    if c_map[y+1][x] == '.' and dis[y+1][x] == -1:
        dis[y+1][x] = dis[y][x] + 1
        que.put((y+1, x))

print(dis[g[0]-1][g[1]-1])


print('##############################################')

r, c = map(int, input().split())
s = list(map(int, input().split()))
g = list(map(int, input().split()))
c_map = [list(input()) for _ in range(r)]
dis = [[3000]*(c+1) for _ in range(r+1)]

dis[s[0]-1][s[1]-1] = 0
que = queue.Queue()
que.put((s[0]-1, s[1]-1))
while not que.empty():
    y, x = que.get()
    if y == g[0]-1 and x == g[1]-1:
        break
    if c_map[y][x-1] == '.' and dis[y][x-1] > dis[y][x]:
        dis[y][x-1] = dis[y][x] + 1
        que.put((y, x-1))
    if c_map[y-1][x] == '.' and dis[y-1][x] > dis[y][x]:
        dis[y-1][x] = dis[y][x] + 1
        que.put((y-1, x))
    if c_map[y][x+1] == '.' and dis[y][x+1] > dis[y][x]:
        dis[y][x+1] = dis[y][x] + 1
        que.put((y, x+1))
    if c_map[y+1][x] == '.' and dis[y+1][x] > dis[y][x]:
        dis[y+1][x] = dis[y][x] + 1
        que.put((y+1, x))

print(dis[g[0]-1][g[1]-1])
