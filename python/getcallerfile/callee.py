import caller

# def print_callee():
#     print("This is a calle, ", caller.srcname(), caller.lineno())
    
import inspect
from inspect import getframeinfo, stack

def debuginfo(message, color='green'):
    if color == 'red':
        template = '\033[31m%s\033[0m'
    elif color == 'green':
        template = '\033[32m%s\033[0m'
    elif color == 'yellow':
        template = '\033[33m%s\033[0m'
    elif color == 'magenta':
        template = '\033[35m%s\033[0m'
    elif color == 'cyan':
        template = '\033[36m%s\033[0m'

    caller = getframeinfo(stack()[1][0])
    print (template %("%s:%d - %s" % (caller.filename, caller.lineno, message)))

def grr(arg):
    debuginfo(arg)

if __name__ == '__main__':
    print ("hello, this is line number", caller.lineno())
    print 
    print 
    print ("and this is line", caller.srcname(), caller.lineno())    
    debuginfo("Test")