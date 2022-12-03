defmodule Day3 do
  def main() do
    file = "input.txt"
    one(file)
    two(file)
  end

  def one(file) do
    file
    |> readFile
    |> Enum.map(&getScoreForSack/1)
    |> Enum.sum
    |> prefixOutput("One: ")
  end

  def getScoreForSack(sack) do
    {pocket1, pocket2} = sack
    |> String.split_at(trunc(String.length(sack) / 2))

    pocket1
    |> String.to_charlist
    |> List.first
    |> getScoreForCharacter
  end

  def getDuplicates(l, s) do
    Enum.filter(l, fn c -> s =~ to_string([c]) end)
  end

  def getScoreForCharacter(character) do
    String.to_charlist("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    |> Enum.find_index(fn s -> s == character end)
    |> Math.add(1)
  end

  def two(file) do
    file
    |> readFile
    |> Enum.chunk_every(3)
    |> Enum.map(&getScoreForThreeSacks/1)
    |> Enum.sum
    |> prefixOutput("Two: ")
  end

  def getScoreForThreeSacks([sack1, sack2, sack3]) do
    sack1
    |> String.to_charlist
    |> getDuplicates(sack2)
    |> getDuplicates(sack3)
    |> List.first
    |> getScoreForCharacter
  end

  def readFile(file) do
    File.read!(file)
    |> String.split("\n")
  end

  def prefixOutput(out, prefix \\ "") do
    IO.puts(prefix <> to_string(out))
  end
end

defmodule Math do
  def add(a, b) do
    a + b
  end
end

Day3.main()
