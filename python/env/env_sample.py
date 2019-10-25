import os

a = os.environ.get("CTL_SAMPLE_ENV1")
print("1. CTL_SAMPLE_ENV1 : %r" % a)

b = os.environ.get("CTL_SAMPLE_ENV2", "false").lower()
print("2. CTL_SAMPLE_ENV2 : %s" % b)

c = True if os.environ.get("CTL_SAMPLE_ENV3", "false").lower() == "true" else False
if c:
    print("3. CTL_SAMPLE_ENV3 : True")
else:
    print("3. CTL_SAMPLE_ENV3 : False")

d = lambda x: True if x == True else False
if d(True):
    print("4. x == True : True")
else:
    print("4. x != True : False")
