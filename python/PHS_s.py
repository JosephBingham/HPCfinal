#!/usr/bin/python
import numpy as np

def PHS_sort(arr):
	if len(arr) < 2:
		return arr
	max = arr[0]
	min = arr[0]
	for i in arr:
		if i < min:
			min = i
		if i > max:
			max = i
	temp = [0] * (max - min + 1)
	for i in arr:
		temp[i-min] = i
	j = 0
	for i in temp:
		if i != 0:
			arr[j] = i
			j += 1
	return arr



data = np.arange(1,214748364)

np.random.shuffle(data)

result = PHS_sort(data)

for i in result:
	print i
