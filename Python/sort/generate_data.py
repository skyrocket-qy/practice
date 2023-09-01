import random

def gen_increse(n):
    return [i for i in range(n)]

def gen_decrease(n):
    return [i for i in range(n-1, -1, -1)]

def gen_normal(n):
    re = []
    if n % 2 == 0:
        for i in range(0, n, 2):
            re.append(i)
        for i in range(n-1, 0, -2):
            re.append(i)
    else:
        for i in range(0, n, 2):
            re.append(i)
        for i in range(n-2, 0, -2):
            re.append(i)
    return re

def gen_random(n):
    re = [i for i in range(n)]
    random.shuffle(re)
    return re

def gen_rep_normal(n):
    re = []
    m = n // 2
    s = n // 6
    for i in range(n):
        re.append(int(random.gauss(m, s)))
    return re
        