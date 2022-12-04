defmodule Day1 do
  def main() do
    one()
    two()
  end

  def one() do
    input = File.read!("input.txt")
    downs = input
    |> String.graphemes
    |> Enum.count(& &1 == ")")

    floor = String.length(input) - (downs*2)
    IO.inspect(floor, label: "One")
  end

  def two() do
    File.read!("input.txt")
    |> String.graphemes
    |> Enum.reduce_while([0, 0], fn x, acc ->
      [index, floor] = acc
      case {x, floor} do
        {_, -1} -> {:halt, index}
        {"(", _} -> {:cont, [index+1, floor+1]}
        {")", _} -> {:cont, [index+1, floor-1]}
      end
    end)
    |> IO.inspect(label: "Two")
  end
end

Day1.main()
