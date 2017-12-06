def f(x):
	sum = 0
	n = int(x)
	for i in xrange(n/10):
		sum += i
	return sum * 1.0

def trap(a, b, h):
	return (f(a) + f(b))*h

sum = 0
n = 200
h = 100000000./n
mesh = [i*h for i in xrange(n+1)]
for i in xrange(n):
	sum += trap(mesh[i], mesh[i+1], h)
print sum

