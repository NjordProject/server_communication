#!usr/bin/python3.4
#-*-coding:utf-8-*

import matplotlib.pyplot as plt
import numpy as np
from matplotlib import cm

data = list()

#Just, in order to test vizualisation. Waiting for Redis
for i in range(0, 255):
	row = list()
	for j in range(0, 255):
		row.append(np.random.randint(256))
	data.append(row)

fig, ax = plt.subplots()
###Different colormaps:
#cm.spectral
#cm.rainbow
#cm.autumn
#cm.cool
cax = ax.imshow(data, interpolation="nearest", cmap=cm.rainbow) #See matplotlib documentation for others interpolation methods
ax.set_title("Classroom's topography")
cbar = fig.colorbar(cax, ticks=[0, 127, 255])
cbar.ax.set_yticklabels(["< 0", "127", " > 255"])

plt.show()
