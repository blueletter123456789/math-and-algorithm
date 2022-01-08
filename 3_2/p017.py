n = int(input())
l = list(map(int, input().split()))

a = x = l[0]
b = c = 0
for i in range(1, n):
    b = l[i]
    while a and b:
        if a >= b:
            a %= b
        else:
            b %= a
    if a < b:
        a = b
    x = (l[i]*x)//a
print(x)

# # 最大公約数を返す関数
# def GCD(A, B):
# 	while A >= 1 and B >= 1:
# 		if A < B:
# 			B = B % A  # A < B の場合、大きい方 B を書き換える
# 		else:
# 			A = A % B  # A >= B の場合、大きい方 A を書き換える
# 	if A >= 1:
# 		return A
# 	return B

# # 最小公倍数を返す関数
# def LCM(A, B):
# 	return int(A / GCD(A, B)) * B

# # 入力
# N = int(input())
# A = list(map(int, input().split()))

# # 答えを求める
# R = LCM(A[0], A[1])
# for i in range(2, N):
# 	R = LCM(R, A[i])

# # 出力
# print(R)
