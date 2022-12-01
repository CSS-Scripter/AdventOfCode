defmodule Day1 do
  def main() do
    sortedInput = getRawInput("input.txt")
    |> String.split("\n\n")
    |> Enum.map(fn x ->
      x
      |> String.split("\n")
      |> Enum.map(fn y ->
        y
        |> Integer.parse(10)
        |> Tuple.to_list
        |> Enum.at(0)
      end)
      |> sumList
    end)
    |> Enum.sort(:desc)

    one = sortedInput
    |> Enum.at(0)
    |> Integer.to_string(10)

    two = sortedInput
    |> Enum.take(3)
    |> Enum.sum
    |> Integer.to_string(10)

    IO.puts "One: " <> one
    IO.puts "Two: " <> two
  end

  def getRawInput(filename) do
    File.read!(filename)
  end

  def sumList([]) do
    0
  end

  def sumList([h|t]) do
    h + sumList(t)
  end
end

Day1.main()
