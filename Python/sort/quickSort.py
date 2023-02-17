def quickSort_O(arr):
    if len(arr)==0: return []
    if len(arr)==1: return arr
    pivotList = []
    left = []
    right =[]
    pivot = arr[0]
    for i in range(1, len(arr)):
        if arr[i] < pivot:
            left.append(arr[i])
        else:
            right.append(arr[i])
    pivotList.append(pivot)
    left = quickSort_O(left)
    left.extend(pivotList)
    right = quickSort_O(right)
    left.extend(right)
    return left
    
def quickSort_I(arr, left, right):
    if left >= right: return arr
    pivot=arr[left]
    i=left+1
    j=right
    while i < j:
        while arr[i] < pivot and i < right:
            i += 1
        while arr[j] > pivot and j > left:
            j -= 1
        if i < j:
            arr[i], arr[j] = arr[j], arr[i]
        i+=1
        j-=1
    
    arr[left], arr[j] = arr[j], arr[left]
    quickSort_I(arr, left, j-1)
    quickSort_I(arr, j+1, right)
    return arr

def quickSort_IP(arr, left, right):
    if left >= right: return arr

    def partition(arr, left, right):
        pivot = arr[right]
        i=left-1
        for j in range(left, right):
            if arr[j] < pivot:
                i+=1
                arr[i], arr[j] = arr[j], arr[i]
        i+=1
        arr[i], arr[right] = arr[right], arr[i]
        return i
    pi = partition(arr, left, right)
    quickSort_IP(arr, left, pi-1)
    quickSort_IP(arr, pi+1, right)
    return arr

def quickSort_IFMB(arr, left, right):
    if left >= right: return arr
    if (arr[right] - arr[left])*(arr[right] - arr[(left+right)//2]) < 0:
        arr[right], arr[left] = arr[left], arr[right]
    elif (arr[(left+right)//2] - arr[left])*(arr[(left+right)//2] - arr[right]) < 0:
        arr[(left+right)//2], arr[left] = arr[left], arr[(left+right)//2] 
    quickSort_I(arr, left, right)
    return arr