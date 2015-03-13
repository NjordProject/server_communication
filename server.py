import serial
import redis
import sys

r = redis.StrictRedis(host='localhost', port='6379', db=0)
arduino = serial.Serial('/dev/ttyACM0', 9600)

while True:
	msg = arduino.readline()
	msg = [int(m.split(':')[1]) for m in msg[:-2].split(';')]
	print msg
	r.lpush(sys.argv[1], msg)
	msg = r.lpop(sys.argv[1] + "_target")
	if msg is not None:
		print "Writes order on serial"
		#arduino.write(msg)
