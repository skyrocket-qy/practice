def radixSort(arr):
    length = len(arr)
    max_ = max(arr)
    maxDigit = 0
    while max_ > 0:
        maxDigit += 1
        max_ //= 10
    bucket = [[] for i in range(10)]
    digit = 1
    while maxDigit > 0:
        for i in range(length):
            val = arr.pop(0)
            bucket[val // digit % 10].append(val)
        for i in bucket:
            for j in range(len(i)):
                val = i.pop(0)
                arr.append(val)
        maxDigit -= 1
        digit *= 10
    return arr