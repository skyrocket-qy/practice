def heapSort(arr):
    def sink(start, end):
        root = start
        while True:
            child = 2*root + 1
            if child > end:
                break
            if child + 1 <= end and arr[child] < arr[child+1]:
                child += 1
            if arr[root] < arr[child]:
                arr[root], arr[child] = arr[child], arr[root]
                root = child
            else:
                break
    for start in range((len(arr)-2)//2, -1, -1):
        sink(start, len(arr)-1)
    for end in range(len(arr)-1, 0, -1):
        arr[0], arr[end] = arr[end], arr[0]
        sink(0, end-1)
    return arr