def main(n, l):
    d = {k: 0 for k in range(1, 4)}
    cnt = 0
    for i in l:
        d[i] += 1
    for i in range(1, 4):
        s = d[i]
        cnt += (s*(s-1))//2
    print(cnt)

if __name__ == '__main__':
    n = int(input())
    l = list(map(int, input().split()))
    main(n, l)
