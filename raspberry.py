import urllib2
import time
import json
import RPi.GPIO as GPIO
GPIO.setmode(GPIO.BCM)
GPIO.setup(12, GPIO.OUT)

p = GPIO.PWM(12, 50)
p.start(0)

try:
    while 1:
         headers = { 'Authorization':'<Token>' }
         req = urllib2.Request('http://<ServerIP>/lights/get/<LightID>', None, headers)
         response = urllib2.urlopen(req)
         html = response.read()
         result = json.loads(html)
         current = result.get('result').get('currentvalue')
         maxvalue = result.get('result').get('maxvalue')                          
         p.ChangeDutyCycle(current/maxvalue)
         time.sleep(1)         
except KeyboardInterrupt:
   pass
p.stop()
GPIO.cleanup()

