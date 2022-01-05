def main(n):
    max_num = int(n**0.5)
    for i in range(1, max_num+1):
        if n % i != 0:
            continue
        print(i)
        if n/i != i:
            print(int(n/i))

if __name__ == '__main__':
    n = int(input())
    main(n)
