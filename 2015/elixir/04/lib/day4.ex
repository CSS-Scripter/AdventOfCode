defmodule Day4 do
  def main() do
    one()
    two()
  end

  def one() do
    findHash("yzbqklnj", 0, "00000")
    |> IO.inspect(label: "One")
  end

  def two() do
    findHash("yzbqklnj", 0, "000000")
    |> IO.inspect(label: "One")
  end

  def findHash(secret, prepend, prefix) do
    md5key = secret <> to_string(prepend)
    hash = :crypto.hash(:md5, md5key)
    |> Base.encode16
    case String.starts_with?(hash, prefix) do
      true -> prepend
      false -> findHash(secret, prepend+1, prefix)
    end
  end
end
