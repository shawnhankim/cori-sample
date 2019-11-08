import json

init_string = {"a": {"aa": 1}, "b": [{"bb": 2}], "c": {"cc":3}}

json_str = json.dumps(init_string)
print("Init string: ", json_str)
print(" - type    : ", type(json_str))

json_dic = json.loads(json_str)
print("Json Dictionary : ", str(json_dic))
print(" - type         : ", type(json_dic))
print(" - json[a]      : ", json_dic['a'])


json_data = '{"a": "1", "b": "2", "c": "3"}'

parsed_json = json.loads(json_data)
json_dump   = json.dumps(parsed_json)

print(json.dumps(parsed_json))
print(json_data[0:5])

