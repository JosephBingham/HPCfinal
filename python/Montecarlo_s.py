#!/usr/bin/python

import random, sys

val = 0
lim = 2147483647
i = 0
while i < lim:
	y = random.random()
	x = random.random()
	if x**2 + y ** 2 < 1:
		val += 1
	i += 1
print 4.0 * val/lim
