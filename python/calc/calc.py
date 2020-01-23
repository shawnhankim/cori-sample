
for i in range(101):
    print(f"2**{i} = {2**i}")

def get_memory_str(server_cnt):
    extra  = 1 if server_cnt % 25 > 0 else 0
    memory = server_cnt // 25 + extra
    return f"{64 * memory}K"

servers = [1, 25, 26, 49, 50, 51, 75, 76, 100, 101, 125, 126, 150]

for server_cnt in servers:
    print(f"servers: {server_cnt}, memory: {get_memory_str(server_cnt)}")
