"""
remove front
remove back
remove second - influences direction

dampened = true

9 7 8 4 2 3 1 
      ^ ^
7 8 4 2 3 1 

"""

x = [1, 2, 3, 4, 5]
y = [5, 4, 3, 2, 1]

# 1 vs 2, 2 vs 3

for i, j in zip(x[:-1], x[1:]):
    print(i, j)

print()

for i, j in zip(x, y):
    print(i, j)
