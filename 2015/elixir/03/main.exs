defmodule Day3 do
  def main() do
    one()
    two()
  end

  def one() do
    File.read!("input.txt")
    |> String.graphemes
    |> Enum.reduce([[0, 0], [[0, 0]]], &travel/2)
    |> Enum.at(1)
    |> Enum.frequencies
    |> Map.keys
    |> Enum.count
    |> IO.inspect(label: "One")
  end

  def travel(action, [[curX, curY], history]) do
    newPos = case action do
      "<" -> [curX-1, curY]
      ">" -> [curX+1, curY]
      "v" -> [curX, curY-1]
      "^" -> [curX, curY+1]
    end
    [newPos, [newPos | history]]
  end

  def two() do
    File.read!("input.txt")
    |> String.graphemes
    |> Enum.chunk_every(2)
    |> Enum.reduce([[], []], fn [a, b], [np, rp] -> [[a | np], [b | rp]] end)
    |> Enum.map(&Enum.reverse/1)
    |> Enum.map(fn x -> Enum.reduce(x, [[0, 0], [[0, 0]]], &travel/2) end)
    |> Enum.map(fn x -> Enum.at(x, 1) end)
    |> List.flatten
    |> Enum.chunk_every(2)
    |> Enum.frequencies
    |> Map.keys
    |> Enum.count
    |> IO.inspect(label: "Two")
  end
end

Day3.main()
