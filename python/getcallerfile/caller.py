"""This provides a lineno() function to make it easy to grab the line
number that we're on.

Danny Yoo (dyoo@hkn.eecs.berkeley.edu)
"""

import inspect
from inspect import currentframe, getframeinfo

def lineno():
    """Returns the current line number in our program."""
    return inspect.currentframe().f_back.f_lineno

def srcname():
    """Returns the current source code name in our program."""
    cf = inspect.currentframe().f_back
    filename = getframeinfo(cf).filename
    return filename

# if __name__ == '__main__':
#     print ("hello, this is line number", lineno())
#     print 
#     print 
#     print ("and this is line", srcname(), lineno())