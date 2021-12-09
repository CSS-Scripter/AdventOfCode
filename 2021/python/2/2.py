def one():
  horizontal = 0
  depth = 0

  inputLines = open('input.txt', 'r').read().split('\n')
  for line in inputLines:
    [motion, amount] = line.split(' ')
    amount = int(amount)
    if motion == 'forward':
      horizontal += amount
      continue
    if motion == 'down':
      depth += amount
      continue
    if motion == 'up':
      depth -= amount
      continue
    print('Illegal statement reached')
    print(line)

  print('Final depth: ' + str(depth))
  print('Final horizontal: ' + str(horizontal))
  print('Final answer: ' + str(depth * horizontal))

def two():
  horizontal = 0
  depth = 0
  aim = 0

  inputLines = open('input.txt', 'r').read().split('\n')
  for line in inputLines:
    [motion, amount] = line.split(' ')
    amount = int(amount)
    if motion == 'forward':
      horizontal += amount
      depth += (amount * aim)
      continue
    if motion == 'down':
      aim += amount
      continue
    if motion == 'up':
      aim -= amount
      continue
    print('Illegal statement reached')
    print(line)

  print('Final depth: ' + str(depth))
  print('Final horizontal: ' + str(horizontal))
  print('Final answer: ' + str(depth * horizontal))

two()