def moveLastToFirst(num):

    string = str(num)

    mylist = list(string)

    # for position, item in enumerate(mylist):
    #     print "position: %s, item: %s" % (position, item)

    last = mylist[-1]

    length = len(mylist)

    mylist.insert(0,last)

    last = mylist[-1]

    length = len(mylist)

    del mylist[length-1]

    

    mylist = ''.join(mylist)

    num = int(mylist)

    return num



def compare(num):
    double = num * 2
    ltf = moveLastToFirst(num)
   # print "ltf %s" % ltf
    #if num == 2000:
     #    return "found"

    if ltf == double:
         return "found"


i = 1

while True:
    print i
    returnval = compare(i)
    if returnval is "found":
        print "FOUND"
        file = "answer.txt"
        content = open(file, 'r+')
        print "i is %s" % i
        var = str(i)
        content.write(var)
        content.close()
        
        break

    else:
        i = i + 1


# initialize

#num = 10



#double_num = num * 2



#second = moveLastToFirst(num)

# while True: 

#     if 1 < 2:

#         False

#     else:

#         True