import jinja2

YML_TEMPLATE = """
version: '2.5'
services:
  real: {{is_test}}
  image: {{image_name}}
{% if is_test != True %}
etc.: real
{% endif %}
"""

def generate_yml(file_name, yml_data):
  f = open(file_name, "w")
  f.write(yml_data)
  f.close()

test_yml = jinja2.Template(YML_TEMPLATE).render(image_name="test image", is_test=True)
print(test_yml)
#generate_yml(file_name="test.yml", yml_data=test_yml)

real_yml = jinja2.Template(YML_TEMPLATE).render(image_name="real image", is_test=False)
print("\n%s" % real_yml)
#generate_yml(file_name="real.yml", yml_data=real_yml)

