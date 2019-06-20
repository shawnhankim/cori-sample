# #!/usr/bin/env python3
# # -*- coding: utf-8 -*-
# """
# Created on Thu Jun 20 10:09:55 2019

# @author: hankim
# """

# package main

# import "C"

# //export add
# func add(left, right int) int {
# 	return left + right
# }

# func main() {
# }

"""Sample code to call Golang from Python

"""
from ctypes import *

# Load library
lib = cdll.LoadLibrary("./sample_go_lib.so")

# Call Add() Golang function from Python
lib.Add.argtypes = [c_longlong, c_longlong]
print ("%-25s = %d" % ("sample_go_lib.Add(5, 3)", lib.Add(5, 3)))

# Call Cosine() Golang function from Python
lib.Cosine.argtypes = [c_double]
lib.Cosine.restype  =  c_double 
cos = lib.Cosine(90)
print ("%-25s = %f" % ("sample_go_lib.Cosine(90)", cos))

# Call Sort() Golang function from Python
class GoSlice(Structure):
    _fields_ = [("data", POINTER(c_void_p)), 
                ("len", c_longlong), ("cap", c_longlong)]
nums = GoSlice((c_void_p * 5)(74, 4, 122, 9, 12), 5, 5)
lib.Sort.argtypes = [GoSlice]
lib.Sort.restype = None
lib.Sort(nums)

# Call Log() Golang function from Python
class GoString(Structure):
    _fields_ = [("p", c_char_p), ("n", c_longlong)]
lib.Log.argtypes = [GoString]
msg = GoString(b"I am trying to call Golang function from Python!", 48)
lib.Log(msg)
