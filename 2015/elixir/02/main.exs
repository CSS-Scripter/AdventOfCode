defmodule Day2 do
  def main() do
    one()
    two()
  end

  def one() do
    File.read!("input.txt")
    |> String.split("\n")
    |> Enum.map(&parseDimensions/1)
    |> Enum.map(&calculateRequiredWrappingPaper/1)
    |> Enum.sum
    |> IO.inspect(label: "One")
  end

  def parseDimensions(line) do
    line
    |> String.split("x")
    |> Enum.map(&String.to_integer/1)
  end

  def calculateRequiredWrappingPaper([x, y, z]) do
    sides = [x*y, x*z, y*z]
    smallestSide = sides
    |> Enum.sort(:asc)
    |> Enum.at(0)

    (Enum.sum(sides) * 2) + smallestSide
  end

  def two() do
    File.read!("input.txt")
    |> String.split("\n")
    |> Enum.map(&parseDimensions/1)
    |> Enum.map(&calculateRibbon/1)
    |> Enum.sum
    |> IO.inspect(label: "Two")
  end

  def calculateRibbon(lengths) do
    ribbonWrap = lengths
    |> Enum.sort(:asc)
    |> Enum.take(2)
    |> Enum.sum

    ribbon = calculateBow(lengths)
    ribbonWrap * 2 + ribbon
  end

  def calculateBow([x, y, z]) do
    x * y * z
  end
end

Day2.main()
