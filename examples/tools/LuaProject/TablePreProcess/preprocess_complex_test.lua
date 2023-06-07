
local function preprocess_complex_test(data_table)
	for k1,v1 in pairs(data_table) do
		for k2,v2 in pairs(v1) do
			v2.number = v2.number + 1
		end
	end
	
	--双key结构开启这个才能支持默认值metatable
	data_table.__isdoublekeydata = true
	
	return data_table
end

return preprocess_complex_test