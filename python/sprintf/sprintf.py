val = 1.1
print('The value of pi is approximately %5.3f.' % val)

str_val = 'The value of pi is approximately %5.3f.' % val
print(str_val)

msg = "Test 1"
str_val = 'The value of pi is approximately %5.3f (%-10s) !!!' % (val, msg)
print(str_val)

msg = "Test 123"
str_val = 'The value of pi is approximately %5.3f (%-10s) !!!' % (val, msg)
print(str_val)
