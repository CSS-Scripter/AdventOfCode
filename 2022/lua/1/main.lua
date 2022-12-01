local function lines_from(file)
  local lines = {}
  for line in io.lines(file) do
    lines[#lines+1] = line
  end
  return lines
end

local function get_totals_from_lines(lines)
  local totals = {}
  local total = 0
  for _, line in pairs(lines) do
    if line == "" then
      totals[#totals+1] = total
      total = 0
    else
      total = total + tonumber(line)
    end
  end
  totals[#totals+1] = total
  return totals
end

local lines = lines_from("input.txt")
local totals = get_totals_from_lines(lines)
table.sort(totals)
print("One: " .. totals[#totals])
print("Two: " .. (totals[#totals] + totals[#totals-1] + totals[#totals-2]))