import sys

def fprintf(stream, fmt, *args):
    stream.write(fmt % args)

fprintf(sys.stdout, "test %d\n", 1)
fprintf(sys.stdout, "test %d\n", 2)
