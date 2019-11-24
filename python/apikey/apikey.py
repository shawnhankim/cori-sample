import secrets

api_key_1 = secrets.token_urlsafe(16)
print(api_key_1)

api_key_2 = secrets.token_hex(16)
print(api_key_2)

