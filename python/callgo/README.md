# How to call Golang function from Python
- Implement the function to expose (i.e. sample_go_lib.go)
- Compile using the -buildmode=c-shared build flag to create the shared object binary:
```bash
   $ go build -o sample_go_lib.so -buildmode=c-shared sample_go_lib.go
```
- Upon completion, the compiler outputs two files:sample_go_lib.h, a C header file and sample_go_lib.so, the shared object file, shown below:
```bash
   -rw-r--r--  1 shawnkim   Users     1708 Jun 20 13:12 sample_go_lib.h
   -rw-r--r--  1 shawnkim   Users  2552808 Jun 20 13:12 sample_go_lib.so   
```
- Implement the Python code to call Golang functions (i.e. sample_call_go.py)
- Run the Python code:
```bash
   $ python sample_call_go.py
   sample_go_lib.Add(5, 3)   = 8
   sample_go_lib.Cosine(90)  = -0.448074
   Golang logging message : I am trying to call Golang function from Python!
```