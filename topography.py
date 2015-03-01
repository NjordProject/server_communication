#!usr/bin/python3.4
#-*-coding:utf-8-*

import matplotlib.pyplot as plt
import numpy as np
from matplotlib import cm
import redis
import sys

name_area = sys.argv[1]
r = redis.StrictRedis(host = "localhost", port = 6379, db = 0)

data = np.zeros([0, 0])

fig, ax = plt.subplots()
#cax = ax.imshow(data, interpolation = "nearest", cmap = cm.hot_r, origin = "lower")
cax = ax.imshow(data, cmap = cm.hot_r, origin = "lower")
ax.set_title(name_area + "'s topography")
cbar = fig.colorbar(cax, ticks = [0, 100, 200])
cbar.ax.set_yticklabels(["< 0", "100", "> 200"])
ax.patch.set_facecolor('white')
plt.show(block = False)

#Read from redis list and add to plot each new value available
while True:
	e = r.blpop(name_area)[1][1:-1].replace(' ', '').split(',')
	e = [int(i) for i in e]
	if(e[1] > data.shape[0]): #We need to resize the matrix
		tmp = np.zeros([e[1], data.shape[1]])
		tmp[0:data.shape[0], 0:data.shape[1]] = data
		data = tmp
	if(e[2] > data.shape[1]):
		tmp = np.zeros([data.shape[0], e[2]])
		tmp[0:data.shape[0], 0:data.shape[1]] = data
		data = tmp
	if e[3] - e[9] < 0: #If distance > 5m
		data[e[1] - 1, e[2] - 1] = 0
	else:
		data[e[1] - 1, e[2] - 1] = e[3] - e[9]
	cax = ax.imshow(data, cmap = cm.hot_r, origin = "lower", interpolation = "nearest")
	plt.draw()
