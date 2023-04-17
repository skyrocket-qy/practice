class Solution(object):
    def wiggleSort(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        
        n=len(nums)
        #med=median(nums)
        import statistics
        med=statistics.median(nums)
        me=n//2
        i=0
        j=n-1
        if n%2==0:
            while i<me and j>me:
                while i<me and nums[i]<=med:
                    i+=1
                while j>me and nums[j]>med:
                    j-=1
                if i<j:
                    nums[i], nums[j]=nums[j], nums[i]
        else:
            while i<me and j>me:
                while i<me and nums[i]<med:
                    i+=1
                while j>me and nums[j]>=med:
                    j-=1
                if i<j:
                    nums[i], nums[j]=nums[j], nums[i]
        i=0
        j=n-1
        k=0
        if n%2==0:
            n+=1
        def nxt(a):
            return (1+2*a)%n
        while k<=j:
            if nums[nxt(k)]>med:
                nums[nxt(k)], nums[nxt(i)]=nums[nxt(i)], nums[nxt(k)]
                i+=1
                k+=1
            elif nums[nxt(k)]<med:
                nums[nxt(j)], nums[nxt(k)]=nums[nxt(k)], nums[nxt(j)]
                j-=1
            else:
                k+=1
            
        
    
        
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
    
def median(arr):
    n=len(arr)
    if n%2==1:
        return quickselect(arr, 0, n-1, n//2)
    else:
        return 0.5*(quickselect(arr, 0, n-1, n//2)+quickselect(arr, 0, n-1, n//2-1))



