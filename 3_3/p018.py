def main(n, l):
    d = {k*100: 0 for k in range(1, 5)}
    for i in l:
        d[i] += 1
    cnt = 0
    cnt += d[100] * d[400]
    cnt += d[200] * d[300]
    print(cnt)

if __name__ == '__main__':
    n = int(input())
    l = list(map(int, input().split()))
    main(n, l)
