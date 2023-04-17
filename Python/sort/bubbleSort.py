def bubbleSort(arr):
    length = len(arr)
    for i in range(length-1):
        flag = 0
        for j in range(length-1-i):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
                flag = 1
        if not flag:
            break
    return arr
