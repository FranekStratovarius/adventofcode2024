local function check_direction(report, level)
	if report.direction == 0 then
		if report.last_level < level then
			report.direction = 1
		elseif report.last_level > level then
			report.direction = -1
		end
	else
		if report.last_level < level and report.direction ~= 1 then
			--print("direction > flagged", report.last_level..">"..level)
			report.unsafe = true
		elseif report.last_level > level and report.direction ~= -1 then
			--print("direction < flagged", report.last_level.."<"..level)
			report.unsafe = true
		end
	end
	--print("direction", report.direction, report.last_level, level)
end

local function check_difference(report, level)
	local difference = math.abs(report.last_level - level)
	-- level too low
	if difference < 1 then
		--print("difference < flagged", difference, report.last_level.." "..level)
		report.unsafe = true
	-- level too high
	elseif difference > 3 then
		--print("difference > flagged", difference, report.last_level.." "..level)
		report.unsafe = true
	end
end

local function tester(report)
    --print("-", table.concat(report, ", "))
    report.direction = 0
    report.unsafe = false

    for i, level in ipairs(report) do
        --print(">", level, "|", report.last_level)
        if report.last_level then
            check_direction(report, level)
            check_difference(report, level)
        end
        report.last_level = level
        if report.unsafe then
            --print("unsafe")
            return false
        end
    end

    --print("safe")
    return true
end

return function (report)
    --print("#", table.concat(report, ", "))

    -- if safe, just return
    local safe = tester(report)
    if safe then
        --print("#", "safe")
        return true
    end

    -- retry by removing one number at a time
    --print("#", "dampener used")
    for i = 1, #report do
        local next_report = {}
        for j, level in ipairs(report) do
            if j ~= i then
                table.insert(next_report, level)
            end
        end

        -- test the modified report
        if tester(next_report) then
            --print("#", "almost unsafe but safe with removal")
            return true
        end
    end

    --print("#", "unsafe")
    return false
end