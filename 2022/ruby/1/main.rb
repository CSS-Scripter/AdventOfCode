elfs = File.read("input.txt").split("\n\n").map { |e| e.split("\n").map(&:to_i).sum() }.sort_by{ |e| -e }
puts("One: " + elfs[0].to_s + "\nTwo: " + (elfs[0] + elfs[1] + elfs[2]).to_s)