n = int(input())
score1 = [0]*(n+1)
score2 = [0]*(n+1)
for i in range(1, n+1):
    c, p = map(int, input().split())
    if c == 1:
        score1[i] = p
    else:
        score2[i] = p
    score1[i] += score1[i-1]
    score2[i] += score2[i-1]
 
q = int(input())
for _ in range(q):
    l, r = map(int, input().split())
    print(score1[r]-score1[l-1], score2[r]-score2[l-1])
