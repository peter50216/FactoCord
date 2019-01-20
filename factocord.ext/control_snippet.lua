-- FactoCord.ext modifications start

local function log_message(event, msg)
	print (event.tick .. " [FactoCord] " .. msg)
	-- game.write_file("server.log", msg .. "\n", true)
end

local function log_player_message(event, msg_in)
	local msg = "Player " .. game.players[event.player_index].name .. " " .. msg_in .. "."
	log_message (event, msg)
	-- game.write_file("server.log", msg .. "\n", true)
end

local function log_player_death_message(event, msg_in)
	local cs = event.cause.name
	local msg = "" .. game.players[event.player_index].name .. " has been killed by some " .. cs .. "! ROFL!"
	log_message (event, msg)
	-- game.write_file("server.log", msg .. "\n", true)
end

local function log_research_message(event, msg_in)
	local msg = msg_in .. " \"" .. event.research.name .. "\""
	log_message (event, msg)

	--for _, player in pairs(game.players) do
	--	player.print{event.research.localised_name[1]}
	--end
	-- game.write_file("server.log", msg .. "\n", true)
end

script.on_event(defines.events.on_player_died,			function(event) log_player_death_message(event, "") end)
script.on_event(defines.events.on_research_started,		function(event) log_research_message(event, "Started research of") end)
script.on_event(defines.events.on_research_finished,	function(event) log_research_message(event, "Research finished for") end)
script.on_event(defines.events.on_rocket_launched,		function(event) log_message("A rocket was launched.") end)

-- FactoCord.ext modifications end

