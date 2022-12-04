defmodule Day4 do
  def main() do
    one()
    two()
  end

  def one() do
    getInput("input.txt")
    |> parseRangePairs
    |> Enum.map(&doesRangePairContain/1)
    |> Enum.sum
    |> IO.inspect(label: "One")
  end

  def getInput(file) do
    File.read!(file)
    |> String.split("\n")
  end

  def parseRangePairs(lines) do
    lines
    |> Enum.map(&parseRangePair/1)
  end

  def parseRangePair(line) do
    line
    |> String.split(",")
    |> Enum.map(&parseRange/1)
  end

  def parseRange(rangeString) do
    rangeString
    |> String.split("-")
    |> Enum.map(&String.to_integer/1)
  end

  def doesRangePairContain(rangePair) do
    [[min1, max1], [min2, max2]] = rangePair
    case (min1 <= min2 && max1 >= max2) || (min2 <= min1 && max2 >= max1) do
      true -> 1
      false -> 0
    end
  end

  def two() do
    getInput("input.txt")
    |> parseRangePairs
    |> Enum.map(&doesRangePairOverlap/1)
    |> Enum.sum
    |> IO.inspect(label: "Two")
  end

  def doesRangePairOverlap(rangePair) do
    [[min1, max1], [min2, max2]] = rangePair
    case (max1 >= min2 && min1 <= max2) || (max2 >= min1 && min2 <= max1) do
      true -> 1
      false -> 0
    end
  end
end

Day4.main()
