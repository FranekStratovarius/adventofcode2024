local left = {}
local right = {}
for line in io.lines("input") do
	local i, j = string.find(line, "   ")
	table.insert(left, tonumber(string.sub(line, 0, i-1)))
	table.insert(right, tonumber(string.sub(line, j+1, string.len(line))))
end

table.sort(left)
table.sort(right)

local distances = {}

for k, v in ipairs(left) do
	table.insert(distances, math.abs(v - right[k]))
end

local sum = 0
for _, v in ipairs(distances) do
	sum = sum + v
end

print(sum)