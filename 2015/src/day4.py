# Python 3 code to demonstrate the
# working of MD5 (byte - byte)

import hashlib


def partOne(key):
    print(hashlib.md5(bytes(key)))


if __name__ == "__main__":
    with open("./input/day4.txt", "w") as f:
        key = f.readline()
        print(partOne(key))
