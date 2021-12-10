defmodule D1 do
  def one() do
    getInput("input.txt")
    |> Enum.map(&calcFuel/1)
    |> Enum.reduce(0, fn x, acc -> x + acc end)
  end

  def two() do
    getInput("input.txt")
    |> Enum.map(fn x -> 
      x
      |> calcFuel
      |> calcFuelOfFuel(0) 
      end)
    |> Enum.reduce(0, fn x, acc -> x + acc end)
  end

  def getInput(filename) do
    {:ok, input} = File.read(filename)
    input
    |> String.split("\n")
    |> Enum.map(&String.to_integer/1)
  end

  def calcFuel(mass) do
    mass
    |> Math.divide(3)
    |> Float.floor
    |> Math.substract(2)
    |> trunc
  end

  def calcFuelOfFuel(mass, acc) when mass > 0 do
    acc = mass
    |> calcFuel
    |> calcFuelOfFuel(acc)
    |> Math.addition(mass)
    acc
  end

  def calcFuelOfFuel(_, acc) do
    acc
  end
end

defmodule Math do
  def divide(a, b) do
    a / b
  end

  def substract(a, b) do
    a - b
  end

  def addition(a, b) do
    a + b
  end
end

IO.puts("One: #{D1.one()}")
# D1.two()
IO.puts("Two: #{D1.two()}")