# n = int(input())
# i = 2
# res = 'Yes'
# while i**2 <= n:
#     if n % i != 0:
#         i += 1
#         continue
#     else:
#         res = 'No'
#         break
# print(res)

def prime(n):
    max_num = int(n**0.5)
    for i in range(2, max_num+1):
        if n % i == 0:
            return False
    return True

if __name__ == '__main__':
    n = int(input())
    if prime(n):
        print('Yes')
    else:
        print('No')
