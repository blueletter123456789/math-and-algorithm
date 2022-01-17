from heapq import merge
import random

# source1
n = 10
def calc(n):
    if n < 2:
        return int(n)
    return int(n) * calc(int(n-1))
print(calc(n))

# source2
def func_u(a, b):
    if b == 0:
        return a
    return func_u(b, a%b)
print(func_u(12, 24))

# source3
l = [random.randint(1, 10) for _ in range(9)]
def func_s(left, right):
    if right - left < 2:
        return l[left]
    middle = (left + right) // 2
    sum_left = func_s(left, middle)
    sum_right = func_s(middle, right)
    return sum_left + sum_right
print(func_s(0, len(l)))

# source4
def func_merge(l1, l2):
    l3 = [0 for _ in range(len(l1) + len(l2))]
    i = j = k = 0
    while len(l1) > i and len(l2) > j:
        if l1[i] > l2[j]:
            l3[k] = l2[j]
            j += 1
        else:
            l3[k] = l1[i]
            i += 1
        k += 1

    while len(l1) > i:
        l3[k] = l1[i]
        i += 1
        k += 1
    while len(l2) > j:
        l3[k] = l2[j]
        j += 1
        k += 1
    return l3

l1 = [13,34,50,75]
l2 = [11,20,28,62]
print(func_merge(l1, l2))

def merge_sort(l):
    if len(l) <= 1:
        return l
    middle = len(l)//2
    left = l[middle:]
    right = l[:middle]
    merge_sort(left)
    merge_sort(right)
    i = j = k = 0
    while len(left) > i and len(right) > j:
        if left[i] > right[j]:
            l[k] = right[j]
            j += 1
        else:
            l[k] = left[i]
            i += 1
        k += 1

    while len(left) > i:
        l[k] = left[i]
        i += 1
        k += 1
    while len(right) > j:
        l[k] = right[j]
        j += 1
        k += 1
    return l

n = int(input())
l = list(map(int, input().split()))
ans = merge_sort(l)
print(*ans)


"""
回答
"""
def MergeSort(A):
	# 長さが 1 の場合、すでにソートされているので何もしない
	if len(A) == 1:
		return A
	
	# 2 つに分割した後、小さい配列をソート
	m = len(A) // 2
	A_Dash = MergeSort(A[0:m])
	B_Dash = MergeSort(A[m:len(A)])
	
	# この時点で以下の 2 つの配列がソートされている：
	# 列 A' に相当するもの [A_Dash[0], A_Dash[1], ..., A_Dash[m-1]]
	# 列 B' に相当するもの [B_Dash[0], B_Dash[1], ..., B_Dash[len(A)-m-1]]
	# 以下が Merge 操作となります。
	c1 = 0
	c2 = 0
	C = []
	while (c1 < len(A_Dash) or c2 < len(B_Dash)):
		if c1 == len(A_Dash):
			# 列 A' が空の場合
			C.append(B_Dash[c2])
			c2 += 1
		elif c2 == len(B_Dash):
			# 列 B' が空の場合
			C.append(A_Dash[c1])
			c1 += 1
		else:
			# そのいずれでもない場合
			if A_Dash[c1] <= B_Dash[c2]:
				C.append(A_Dash[c1])
				c1 += 1
			else:
				C.append(B_Dash[c2])
				c2 += 1
	
	# 列 A', 列 B' を合併した配列 C を返す
	return C

# 以下、メイン部分
N = int(input())
A = list(map(int, input().split()))

# マージソート → 答えの出力
Answer = MergeSort(A)
print(*Answer)
