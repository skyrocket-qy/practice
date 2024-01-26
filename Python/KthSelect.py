def kth(arr, k):
    def quickselect(arr, left, right, k):
        pivot = arr[right]
        i=left
        for j in range(left, right):
            if arr[j]<pivot:
                arr[i], arr[j] = arr[j], arr[i]
                i+=1
        arr[i], arr[right] = arr[right], arr[i]
        if k<i:
            return quickselect(arr, left, i-1, k)
        elif k==i:
            return arr[i]
        else:
            return quickselect(arr, i+1, right, k)
    return quickselect(arr, 0, len(arr)-1, k-1)
