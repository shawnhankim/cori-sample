import caller

# def print_callee():
#     print("This is a calle, ", caller.srcname(), caller.lineno())
    

if __name__ == '__main__':
    print ("hello, this is line number", caller.lineno())
    print 
    print 
    print ("and this is line", caller.srcname(), caller.lineno())    