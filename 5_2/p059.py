n = int(input())
mod_n = n % 4
if mod_n == 1:
    print(2)
elif mod_n == 2:
    print(4)
elif mod_n == 3:
    print(8)
else:
    print(6)
