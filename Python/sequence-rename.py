# -*- coding: utf-8 -*-
import os


def Rename_sort(path, index = 1, sort = "random"):
    if sort == 'name':
        lis = os.listdir(path)
        lis.sort(key= lambda x: int(x[: -4]))
        index = 1
        for img_name in lis:
            os.rename(os.path.join(path, img_name), os.path.join(path, str(index) + ".png"))
            index += 1
        return None
    os.chdir(path)
    index_copy = index 
    for i in os.listdir(path):
        new_name = "fpzee" + str(index_copy) + ".png"
        os.rename(i, new_name)
        index_copy += 1
    index_copy = index
    for i in os.listdir(path):
        new_name = str(index_copy) + ".png"
        os.rename(i, new_name)
        index_copy += 1   
    
if __name__ == '__main__':        
    path = input("input the address:")
    Rename_sort(path)
    

'''
To rename the file in the folder, start from index. 
If sort == 'name', the function will sort file with the previous name if previous name is int.
'''    
    
    