# -*- coding: utf-8 -*-

import json

input_file = open ('swagger.json')
json_body  = json.load(input_file)

for item in json_body['paths']:
    print("path: %s" % item)
    for method in json_body['paths'][item]:
        print("  - method: %s" % method)
        if   method == "post" or method == "put":
            print("    req: %s" % json_body['paths'][item][method]['parameters'])
        print("    res: %s"     % json_body['paths'][item][method]['responses'])

with open('swagger_out1.json', 'w') as outfile:
    json.dump(json_body, outfile)
    
with open('swagger_out2.json', 'w', encoding='utf-8') as f:
    json.dump(json_body, f, ensure_ascii=False, indent=4)    