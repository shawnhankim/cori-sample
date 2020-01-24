import tempfile

data = {
    "a": "b",
    "c": "d"
}

with tempfile.NamedTemporaryFile() as fp:
    fp.write(str(data).encode())
    fp.seek(0)
    read_data = fp.read()
    print(f"read_data : {read_data}, fp.name: {fp.name}")

