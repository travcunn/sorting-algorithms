def sort(list):
    for i in range(1, len(list)):
        x = list[i]
        j = i -1
        while j >= 0 and list[j] > x:
            list[j+1] = list[j]
            j = j - 1
        list[j+1] = x

    return list
