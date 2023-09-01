def countingSort(arr):
    dct = {}
    re = []
    for ele in arr:
        if dct.get(ele)==None:
            dct[ele]=1
        else:
            dct[ele] += 1
    for key in sorted(dct):
        val = dct[key]
        while val > 0:
            re.append(key)
            val -= 1
    return re