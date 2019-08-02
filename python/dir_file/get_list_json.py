# -*- coding: utf-8 -*-

import os

# Get list in the directory
arr = os.listdir()
print(arr)

# Get JSON file list in the directory
for file in os.listdir("./"):
    if file.endswith(".json"):
        print(file)
        #print(os.path.join("./", file))