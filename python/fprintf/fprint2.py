
def main():

    f = open("abc.txt", "a+")

    for i in range(3):
        f.write("abcd: %d\r\n" % (i + 1))

    f.close()

 
    f = open("def.txt", "w")
    for i in range(3):
        f.write("abcd: %d\r\n" % (i + 2))

    f.close()


if __name__=="__main__":
    main()
