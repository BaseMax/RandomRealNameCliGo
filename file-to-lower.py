import sys

if len(sys.argv) > 1:
    filepath = sys.argv[1]
else:
    filepath = input("Enter a file name: ")

file = open(filepath, "r")
data = file.read().lower().strip()
file.close()

file = open(filepath, "w")
file.write(data)
file.close()