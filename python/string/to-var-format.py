import re

def trim_string_space(input_str):
    if not isinstance(input_str, (str,)):
        print("input_str must be a type of string")

    return re.sub(r"[^a-zA-Z\d\_]", " ", input_str).strip().replace(" ", "_")


a = " abc def - 1"
print(trim_string_space(a))


b = trim_string_space(f"    aaaa  bbbb  _  2  ")
print(b)

c = " Abc Def - 2"
print(trim_string_space(c))
