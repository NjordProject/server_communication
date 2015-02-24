import serial
import redis
import sys

r = redis.StrictRedis(host='localhost', port='6379', db=0)
arduino = serial.Serial('/dev/ttyACM1', 9600)

while True:
    msg = arduino.readline()
    msg = [int(m.split(':')[1]) for m in msg[:-2].split(';')]
    r.lpush(sys.argv[1], msg)
