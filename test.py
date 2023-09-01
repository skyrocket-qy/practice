import time

n = 100000000
old = time.time()

for _ in range(n):
    _ = 99999999 * 0.01

print(time.time()-old)

old = time.time()

for _ in range(n):
    _ = 99999999 / 100

print(time.time()-old)