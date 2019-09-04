import os

if os.environ.get('MY_ENV1'):
    my_env = os.environ['MY_ENV1']
    print (my_env)    
else:
    print ("MY_ENV is not defined")
