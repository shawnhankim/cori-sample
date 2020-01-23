from urllib.parse import urlsplit, urlunsplit

split_url = urlsplit('http://127.0.0.1/asdf/login.php?q=abc#stackoverflow')

# You now have:
# split_url.scheme   "http"
# split_url.netloc   "127.0.0.1" 
# split_url.path     "/asdf/login.php"
# split_url.query    "q=abc"
# split_url.fragment "stackoverflow"

# Use all the path except everything after the last '/' 
clean_path = "".join(split_url.path.rpartition("/")[:-1])
print(f"clean_path : {clean_path}")

# "/asdf/"

# urlunsplit joins a urlsplit tuple
clean_url = urlunsplit(split_url)
print(f"clean_url : {clean_url}")

# "http://127.0.0.1/asdf/login.php?q=abc#stackoverflow"


# A more advanced example 
advanced_split_url = urlsplit('http://foo:bar@127.0.0.1:5000/asdf/login.php?q=abc#stackoverflow')
print(f"advanced_split_url : {advanced_split_url}")


abc = urlsplit('http://test.com/api/v1')
print(f"abc : {abc}")
print(f"base : {abc.path}")


abc = urlsplit('http://abc.com').path
if abc:
    print(abc)
else:
    print(False)

abc = urlsplit('http://abc.com/').path
print(abc)

abc = urlsplit('http://abc.com//').path
print(abc)


# You now have *in addition* to the above:
# advanced_split_url.username   "foo"
# advanced_split_url.password   "bar"
# advanced_split_url.hostname   "127.0.0.1"
# advanced_split_url.port       "5000"

