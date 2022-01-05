def dev(n):
    def _dev(n):
        l = []
        r = True
        max_num = int(n**0.5)
        for i in range(2, max_num+1):
            if n % i == 0:
                r = False
                l.append(i)
                l += _dev(int(n/i))
                break
        if r:
            l.append(n)
        return l
    res = _dev(n)
    res.sort()
    print(' '.join([str(i) for i in res]))


if __name__ == '__main__':
    n = int(input())
    dev(n)
