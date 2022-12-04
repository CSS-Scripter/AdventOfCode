defmodule Day2 do
  def main() do
    one()
    two()
  end

  def one() do
    scoremap = %{
      "A X" => 4,
      "A Y" => 8,
      "A Z" => 3,
      "B X" => 1,
      "B Y" => 5,
      "B Z" => 9,
      "C X" => 7,
      "C Y" => 2,
      "C Z" => 6
    }
    getInput("input.txt")
    |> Enum.map(fn x -> scoremap[x] end)
    |> Enum.sum
    |> IO.inspect(label: "One")
  end

  def two() do
    scoremap = %{
      "A X" => 3,
      "A Y" => 4,
      "A Z" => 8,
      "B X" => 1,
      "B Y" => 5,
      "B Z" => 9,
      "C X" => 2,
      "C Y" => 6,
      "C Z" => 7
    }
    getInput("input.txt")
    |> Enum.map(fn x -> scoremap[x] end)
    |> Enum.sum
    |> IO.inspect(label: "Two")
  end

  def getInput(file) do
    File.read!(file)
    |> String.split("\n")
  end
end

Day2.main()
