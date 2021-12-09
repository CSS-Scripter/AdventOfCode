f = open('./input.txt', 'r')
input = f.read()

values = list(map(int, input.split('\n')))

amountOfIncrements = 0
previousValue = values[0]
for val in values[1:]:
  if (val > previousValue):
    amountOfIncrements += 1
  previousValue = val
print('Part 1: ' + str(amountOfIncrements))

amountOfIncrements = 0
previousValue = sum(values[0:3])
for i in range(1, len(values) - 2):
  val = sum(values[i:i+3])
  if (val > previousValue):
    amountOfIncrements += 1
  previousValue = val

print('Part 2: ' + str(amountOfIncrements))