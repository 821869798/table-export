


local TablePostProcess = {}

---@class global_postprocess
--全局配置后处理,在所有配置预处理之后
local function global_postprocess(database)
	--处理校验或者其他逻辑
	
	
	--也可以删除包内不要的表，database.xx = nil即可（相当于仅做预处理）
	
end


return global_postprocess