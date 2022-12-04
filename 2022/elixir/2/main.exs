defmodule Day2 do
  def main() do
    one()
    two()
  end

  def one() do
    scoremap = %{
      "A X" => 4, #(1 + 3),
      "A Y" => 8, #(2 + 6),
      "A Z" => 3, #(3 + 0),
      "B X" => 1, #(1 + 0),
      "B Y" => 5, #(2 + 3),
      "B Z" => 9, #(3 + 6),
      "C X" => 7, #(1 + 6),
      "C Y" => 2, #(2 + 0),
      "C Z" => 6  #(3 + 3)
    }
    getInput("input.txt")
    |> Enum.map(fn x -> scoremap[x] end)
    |> Enum.sum
    |> IO.inspect(label: "One")
  end

  def two() do
    scoremap = %{
      "A X" => 3, #(3 + 0),
      "A Y" => 4, #(1 + 3),
      "A Z" => 8, #(2 + 6),
      "B X" => 1, #(1 + 0),
      "B Y" => 5, #(2 + 3),
      "B Z" => 9, #(3 + 6),
      "C X" => 2, #(2 + 0),
      "C Y" => 6, #(3 + 3),
      "C Z" => 7  #(1 + 6)
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
