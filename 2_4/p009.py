def main(n, s, l):
    for i in range(2**n):
        t = []
        for j in range(n):
            if (i >> j) & 1:
                t.append(l[j])
        if sum(t) == s:
            print('Yes')
            return
    print('No')

if __name__ == '__main__':
    n, s = map(int, input().split())
    l = list(map(int, input().split()))
    main(n, s, l)
