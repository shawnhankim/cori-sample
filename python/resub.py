
import re

SAMPLE_FORMAT = '   $abc - $def - [$ghi] jk="$lmn"  '

res = re.sub('$def', '"$jkl"', SAMPLE_FORMAT)
print(res)

res = re.sub('$abc', '', SAMPLE_FORMAT)
print(res)

res = re.sub(r'$abc', r'$jkl', SAMPLE_FORMAT)
print(res)

res = res.strip().replace("$abc", '"$jkl"')
print(res)

res = re.sub('$abc - ($jkl) - $ghi', r'\1', SAMPLE_FORMAT)
print(res)

res = re.sub(r'(?<![a-z])$def(?![a-z])', '$jkl', SAMPLE_FORMAT, flags=re.I)
print(res)

text = 'how are u? umberella u! u. U. U@ U# u '
res = re.sub (' [u|U][s,.,?,!,W,#,@ (^a-zA-Z)]', ' you ', text)
print(res)

res = re.sub(r"[^0-9a-z.\-]$", "", SAMPLE_FORMAT).strip().replace("$def", '"$jkl"')
print(res)

def _get_format(input):
    return (
        re.sub(r"[^0-9a-z.\-]$", "", SAMPLE_FORMAT)
        .strip()
        .replace("$abc", f'"{input}"')
    )
res = _get_format("$jkl")
str = f"{res}##"
print(str)
print(_get_format("$jkl"), "##")

#res = re.sub(SAMPLE_FORMAT).string().replace("$abc", "$ddd")
#print


str = SAMPLE_FORMAT
res1 = str.strip().replace("$abc", "$ooo")
print(res1)

#res1 = re.sub(r"[^0-9a-z.\-]$", "", SAMPLE_FORMAT)
#print(res1)

#res2 = re.sub(r"[^0-9a-z.\-]$", "", SAMPLE_FORMAT).strip()
#print(res2)



