import json

data = json.load(
    open("data.json", "r")
)

max_length = 0
for person in data:
    length = len(data[person])
    if length > max_length:
        max_length = length
    
people = list(data.keys())

s = ""
for person in people:
    s += person + ","
s += "\n"

for i in range(max_length):
    for person in people:
        if i < len(data[person]):
            s += str(data[person][i]) + ","
        else:
            s += str(data[person][-1]) + ","
    s += "\n"

with open("data", "w") as f:
    f.write(s)