#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import json

# read file
with open('./example_02.json', 'r') as myfile:
    data = myfile.read()

# parse file
input_file = open ('example_02.json')
json_obj = json.load(input_file)

# show values
print("id : %d" % json_obj['id'])
print("location[lat] : %d" % json_obj['location']['lat'])
print("location[lon] : %d" % json_obj['location']['lon'])

