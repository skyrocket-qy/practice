def mergeSort(arr):
    def merge(x, y):
        if x == []: return y
        if y == []: return x
        return [x[0]] + merge(x[1: ], y) if x[0]<y[0] else [y[0]] + merge(x, y[1: ])
    def sort(arr):
        n = len(arr)
        m = n//2
        return arr if n<=1 else merge(mergeSort(arr[:m]), mergeSort(arr[m: ]))
    return sort(arr)