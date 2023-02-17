from time import sleep
from threading import Timer

def sleepSort(arr):
    re = []
    def add1(x):
        re.append(x)
    mx = arr[0]
    for v in arr:
        if mx < v: mx = v
        Timer(v*.02, add1, [v]).start()
    sleep((mx+1)*.02)
    return re  