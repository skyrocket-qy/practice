def insertionSort(arr):
    for i in range(1, len(arr)): #
        val = arr[i]
        for j in range(i-1, -1, -1):
            if val >= arr[j]:
                arr[j] = val
                break
            elif j == 0:
                arr[j+1], arr[j] = arr[j], val
            else:
                arr[j+1] = arr[j]
    return arr
    
    
