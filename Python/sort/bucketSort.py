def bucketSort(arr):
    length = len(arr)
    if length < 2: return arr
    if length < 10: return bubbleSort(arr)
    bucket = int(math.sqrt(length))
    max_ = max(arr)
    lst = []
    for i in range(bucket+1):
        lst.append([])
    for ele in arr:
        lst[(ele*bucket)//max_].append(ele)
    re = []
    for i in range(bucket+1):
        re += bucketSort(lst[i])
    return re