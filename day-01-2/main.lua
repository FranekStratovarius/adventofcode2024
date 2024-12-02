local left = {}
local right = {}
for line in io.lines("input") do
	--print(line)
	local i, j = string.find(line, "   ")
	table.insert(left, tonumber(string.sub(line, 0, i-1)))
	table.insert(right, tonumber(string.sub(line, j+1, string.len(line))))
end

table.sort(left)
table.sort(right)

--print(table.concat(left, ", "))
--print(table.concat(right, ", "))

local distances = {}

local second_iterator = 1
for k, v in ipairs(left) do
	local similarity = 0
	--print("["..v.."]")
	while right[second_iterator] and right[second_iterator] < v do
		--print("< "..right[second_iterator])
		second_iterator = second_iterator + 1
	end

	while right[second_iterator] == v do
		--print("> "..right[second_iterator])
		second_iterator = second_iterator + 1
		similarity = similarity + 1
	end

	if k > 1 and left[k] == left[k-1] then
		table.insert(distances, distances[k-1])
	end
	table.insert(distances, similarity * v)
end
--print(table.concat(distances, ", "))

local sum = 0
for _, v in ipairs(distances) do
	sum = sum + v
end

print(sum)