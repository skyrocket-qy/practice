# -*- coding: utf-8 -*-
"""
Created on Fri Nov 19 11:56:21 2021

@author: QingYun
"""

import copy as cp
       
inpu = ['a', 'b', 'c']       
output = ['', 'a','b','c','ab','ac','bc','abc']
final = [[]]

#一次放一個元素進去
for ele1 in inpu:
    tmp = []
    for ele2 in final:
        space = cp.copy(ele2)
        space.append(ele1)
        tmp.append(space)
    final += tmp        
  