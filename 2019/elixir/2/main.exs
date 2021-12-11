defmodule D2 do
  def one() do
    intCode = getInput("input_1.txt")
    zero = loopOne(intCode, 0)
    IO.inspect("One: #{zero}")
  end

  def two() do
    intCode = getInput("input.txt")
    arr = 0..length(intCode)
    try do 
      Enum.each(arr, fn noun -> 
        Enum.each(arr, fn verb -> 
          changedIntCode = List.replace_at(intCode, 1, noun)
          changedIntCode = List.replace_at(changedIntCode, 2, verb)
          zeroAddress = loopTwo(changedIntCode, 0)
          if (zeroAddress == 19690720) do
            IO.inspect("Two: #{100 * noun + verb}")
            throw(:break)
          end
        end)
      end)
    catch
      :break -> :broken
    end
  end

  def getInput(filename) do
    {:ok, input} = File.read(filename)
    input
    |> String.split(",")
    |> Enum.map(&String.to_integer/1)
  end

  def loopOne(intCode, i) do
    case Enum.at(intCode, i) do
      1 -> 
        i1 = Enum.at(intCode, Enum.at(intCode, i+1))
        i2 = Enum.at(intCode, Enum.at(intCode, i+2))
        intCode = List.replace_at(intCode, Enum.at(intCode, i+3), i1 + i2)
        loopOne(intCode, i+4)
      2 -> 
        i1 = Enum.at(intCode, Enum.at(intCode, i+1))
        i2 = Enum.at(intCode, Enum.at(intCode, i+2))
        intCode = List.replace_at(intCode, Enum.at(intCode, i+3), i1 * i2)
        loopOne(intCode, i+4)
      99 -> 
        Enum.at(intCode, 0)
      _ -> 
        IO.inspect("ERROR #{i}")
    end
  end

  def loopTwo(intCode, i) do
    try do
      case Enum.at(intCode, i) do
        1 -> 
          i1 = Enum.at(intCode, Enum.at(intCode, i+1))
          i2 = Enum.at(intCode, Enum.at(intCode, i+2))
          intCode = List.replace_at(intCode, Enum.at(intCode, i+3), i1 + i2)
          loopTwo(intCode, i+4)
        2 -> 
          i1 = Enum.at(intCode, Enum.at(intCode, i+1))
          i2 = Enum.at(intCode, Enum.at(intCode, i+2))
          intCode = List.replace_at(intCode, Enum.at(intCode, i+3), i1 * i2)
          loopTwo(intCode, i+4)
        99 -> 
          Enum.at(intCode, 0)
        _ -> 
          IO.inspect("ERROR #{i}")
      end
    rescue
      ArithmeticError -> 0
    end
  end
end

D2.one()
D2.two()