import queue

r, c = map(int, input().split())
s = list(map(int, input().split()))
g = list(map(int, input().split()))
c_map = [list(input()) for _ in range(r)]

dis = [[0]*(c+1) for _ in range(r+1)]
dis[s[0]][s[1]] = 0
que = queue.Queue()
que.put([s[0], s[1]])
while not que.empty():
    cur = que.get()
    print(cur)
