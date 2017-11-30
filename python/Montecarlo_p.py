#!/usr/bin/python
import thread
import random
import Queue
import threading

def run(end, q):
	i = 0
	qlock.acquire()
	val = q.get()
	qlock.release()
	while i < self.getName():
		i += 1
		x = random.random()
		y = random.random()
		if x**2 + y**2 <= 1:
			val += 1
	qlock.acquire()
	if not q.empty():
		val += q.get()
		q.enqueue(val)
	else:
		q.enqueue(val)
	qlock.release()
	return

if __name__ == '__main__':
	q = Queue.Queue(0)
	qlock = threading.Lock()
	lim = 2147483647
	threadCount = 4
	end = 0
	for x in range(threadCount):
		if x != 3:
			end = lim/(threadCount - 1)
		else:
			end = lim % threadCount
		t = thread.start_new_thread(run, (end,q))
	val = q.get()
	print 4.0 * val/lim
