defmodule Day1 do
  def main() do
    sortedInput = getRawInput("input.txt")
    |> parseInput()
    |> Enum.sort(:desc)

    one(sortedInput)
    two(sortedInput)
  end

  def one(sortedInput) do
    one = sortedInput
    |> Enum.at(0)
    |> Integer.to_string(10)

    IO.puts "One: " <> one
  end

  def two(sortedInput) do
    two = sortedInput
    |> Enum.take(3)
    |> Enum.sum
    |> Integer.to_string(10)

    IO.puts "Two: " <> two
  end

  def parseInput(rawInput) do
    rawInput
    |> String.split("\n\n")
    |> Enum.map(&getTotalCalories/1)
  end

  def getTotalCalories(sublist) do
    sublist
    |> String.split("\n")
    |> Enum.map(&String.to_integer/1)
    |> Enum.sum
  end

  def getRawInput(filename) do
    File.read!(filename)
  end
end

Day1.main()
