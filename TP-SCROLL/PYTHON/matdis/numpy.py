import numpy as np
# two dimensional arrays
m1 = np.array([[1,5,3],[2,4,2]])
m2 = np.array([[6,7],[2,7],[3,1]])
m3 = np.dot(m1,m2) 
print('Two dimensional array:\n', m3) 

# three dimensional arrays
m1 = ([1, 3, 5],[6 ,7, 2],[6, 23, 2]) 
m2 = ([3, 4, 6],[8, 34, 6],[6, 5, 7]) 
m3 = np.dot(m1,m2) 
print('Three dimensional array:\n', m3)