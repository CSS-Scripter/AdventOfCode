def one():
  inputLines = open('input.txt', 'r').read().split('\n')
  count = [{"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}]
  for line in inputLines:
    for i in range(0, len(line)):
      count[i][str(line[i])] += 1
  epsilonBin = ''
  gammaBin = ''
  for num in count:
    if num["0"] > num["1"]:
      epsilonBin += '0'
      gammaBin += '1'
    else:
      epsilonBin += '1'
      gammaBin += '0'
  epsilon = int(epsilonBin, 2)
  gamma = int(gammaBin, 2)
  print("Power consumption: " + str(gamma * epsilon))


def two():
  oxygenBins = open('input.txt', 'r').read().split('\n')
  scrubberBins = open('input.txt', 'r').read().split('\n')
  count = {
    "oxygen": [{"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}],
    "scrubber": [{"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}, {"0": 0, "1": 0}]
  }
  oxygen = 0
  scrubber = 0
  for i in range(0, 12):
    for line in oxygenBins:
      count["oxygen"][i][str(line[i])] += 1
    for line in scrubberBins:
      count["scrubber"][i][str(line[i])] += 1
    criteriaOxygen = '0' if count["oxygen"][i]["0"] > count["oxygen"][i]["1"] else '1'
    criteriaScrubber = '0' if count["scrubber"][i]["0"] <= count["scrubber"][i]["1"] else '1'
    oxygenBins = list(filter(lambda bin: bin[i] == criteriaOxygen, oxygenBins))
    scrubberBins = list(filter(lambda bin: bin[i] == criteriaScrubber, scrubberBins))
    if (len(oxygenBins) == 1):
      oxygen = int(oxygenBins[0], 2)
    if (len(scrubberBins) == 1):
      scrubber = int(scrubberBins[0], 2)
  print(oxygen * scrubber)

two()