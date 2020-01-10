import calendar
import time

ltime = calendar.timegm(time.gmtime())
print(ltime)

epoch_time = int(time.time())
print(epoch_time)

after_1yr = 60*60*24*365*10
expiry_time = epoch_time + after_1yr
print(expiry_time)

expiry_datetime = time.gmtime(expiry_time)
print(expiry_datetime)

print(time.gmtime(1577836800))
