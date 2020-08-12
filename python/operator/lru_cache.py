from functools import lru_cache

@lru_cache(maxsize=None)
def fib(n):
    if n < 2: return n
    return fib(n-1) + fib(n-2)

a = [fib(n) for n in range(16)]
b = fib.cache_info()

print(a)
print(b)


