local tester = require("tester")

local reports = {}
for line in io.lines("input") do
	--print(line)
	local levels = {}
	for number in string.gmatch(line, "%d+") do
		table.insert(levels, tonumber(number))
	end
	table.insert(reports, levels)
end

local safe = 0
for _, report in ipairs(reports) do
	--print("---------------------------")
	if tester(report) then
		safe = safe + 1
	end
end

print(safe)