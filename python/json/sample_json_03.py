# -*- coding: utf-8 -*-

import json

input_file = open ('example_03.json')
json_array = json.load(input_file)
store_list = []

for item in json_array:
    store_details = {"name":None, "city":None, "location":None}
    store_details['name'    ] = item['name'    ]
    store_details['city'    ] = item['city'    ]
    store_details['location'] = item['location']['lat']
    store_list.append(store_details)

print(store_list)