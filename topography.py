#!usr/bin/python3.4
#-*-coding:utf-8-*

import matplotlib.pyplot as plt
import numpy as np
from matplotlib import cm
import redis
import sys
from numpy.random import randn

name_area = sys.argv[1]
nb_drone = sys.argv[2]
r = redis.StrictRedis(host = "localhost", port = 6379, db = 0)

list_drone = [[False, [0, 0, 0]] for _ in range(int(nb_drone))]

def get_drone_free():
	drone_free = []
	for i, elt in enumerate(list_drone):
		if elt[0] == False:
			drone_free.append(i)
	return drone_free

def send_message_drone():
	drone_free = get_drone_free()
	if len(drone_free) > 0:
		for i in range(0, len(data)):
			for j in range(0, len(data[i])):
				print("i = " + str(i) + " j = " + str(j))
				print data[i][j]
				if data[i][j] == 0:
					num_drone = drone_free.pop()
					list_drone[num_drone][0] = True
					list_drone[num_drone][1][0] = i
					list_drone[num_drone][1][1] = j
					list_drone[num_drone][1][2] = 150
					print("Drone n: " + str(num_drone) + " go to work")
					r.lpush(name_area + "_target", str(num_drone) + ";" + str(i) + ";" + str(j) + ";150;0")
					if len(drone_free) == 0:
						return

data = np.zeros([1, 1])
fig, ax = plt.subplots()
cax = ax.imshow(data, interpolation = "nearest" , cmap = cm.hot_r)
ax.set_title(name_area + "'s topography")
cbar = fig.colorbar(cax, ax=ax)
#cbar.ax.set_yticklabels(['< 0', '150', '> 300'])
#cbar = fig.colorbar(cax)
#cbar.ax.get_yaxis().set_ticks([])
#cbar.ax.text(1.2, 0, "> 0", ha="left", va="center")
#cbar.ax.text(1.2, 0.5, "150", ha="left", va="center")
#cbar.ax.text(1.2, 1, "< 300", ha="left", va="center")
#ax.patch.set_facecolor('white')
plt.show(block = False)

#Read from redis list and add to plot each new value available
while True:
	e = r.blpop(name_area)[1][1:-1].replace(' ', '').split(',')
	e = [int(i) for i in e]
	drone = list_drone[e[0]-1]
	if(drone[1][0] == e[1]) and (drone[1][1] == e[2]) and (drone[1][2] == e[3]):
		list_drone[e[0]-1][0] = False
	if(e[1] > data.shape[0]): #We need to resize the matrix
		tmp = np.zeros([e[1], data.shape[1]])
		tmp[0:data.shape[0], 0:data.shape[1]] = data
		data = tmp
	if(e[2] > data.shape[1]):
		tmp = np.zeros([data.shape[0], e[2]])
		tmp[0:data.shape[0], 0:data.shape[1]] = data
		data = tmp
	if e[3] - e[9] < 0: #If distance > Max distance
		data[e[1] - 1, e[2] - 1] = 0
	else:
		data[e[1] - 1, e[2] - 1] = e[3] - e[9]
	cax.set_data(data)
	cax.autoscale()
	plt.draw()
	send_message_drone()
