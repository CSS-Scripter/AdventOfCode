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
    |> Enum.filter(fn c -> pocket2 =~ to_string([c]) end)
    |> List.first
    |> getScoreForCharacter
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
    |> Enum.map(fn [sack1, sack2, sack3] ->
      sack1
      |> String.to_charlist
      |> Enum.filter(fn c -> sack2 =~ to_string([c]) end)
      |> Enum.filter(fn c -> sack3 =~ to_string([c]) end)
      |> List.first
      |> getScoreForCharacter
    end)
    |> Enum.sum
    |> prefixOutput("Two: ")
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
