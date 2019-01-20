<!--p align="center"><img src="http://i.imgur.com/QrhAeBe.png"></p>
<p align="center">FactoCord a Factorio to Discord bridge bot for Linux</p>
<p align="center">
<a href="https://goreportcard.com/report/github.com/FactoKit/FactoCord"><img src="https://goreportcard.com/badge/github.com/FactoKit/FactoCord" alt="Go Report Card"></a>
<a href="https://travis-ci.org/FactoKit/FactoCord"><img src="https://travis-ci.org/FactoKit/FactoCord.svg?branch=master"</img></a>
</p-->

# WORK IN PROGRESS

*Run at your own risk!*

*This readme should be updated later :-P*

# About

This is a modified version of [FactoCord](https://github.com/FactoKit/FactoCord).

A bot that acts as a gateway between Factorio server and Discord. Additionally - with some basic management possibilities.

Some additions to "vanilla FactoCord":

- send Lua commands to the server
- pass Factorio server output to Discord
- pass &lg;server&gt; chat messages to Discord
- give the Factorio server time to save the game before quit
- option to read death / upgrade / rocket launched messages

To get the death and research notifications to Discord, you will have to modify control.lua of the save file you would like to use. It should keep the "vanilla" achievements in a working condition (verified on ONE map, so use with caution). It is slightly modified version of [mlogger 0.2](http://maruohon.kapsi.fi/factorio/mods/mlogger/) and probably can be adapted to be a mod.

# Before you install

Please, keep in mind...

## Security

You are connecting your game to an external service. FactoCord relies on the information from Discord (p.e. admin user IDs) an does not have any additional authentication mechanisms. If your server on any admin account gets compromised, you may have trouble.

Every person that is granted access to the Console channel can read every message from server - possibly leaking technical information that should not be public.

If you run Factorio on a production server and you have the knowledge / resources, keep the FactoCord and Factorio isolated (docker at least).

## Privacy

Every person that is granted access to your Factorio channel can read every message you write in-game. Keep that in mind. You should let all players know, that the game is sending their chat to a Discord server.

## Work in progress

Yeah, this... It should not but it may (maybe some programming mistake or anything) eat all your data.

# Original readme follows...

## Running
*Make sure you have your .env file in the same directory as the executable/binary, you can use .envexample as a template*

There are two ways of starting FactoCord

1. Using the start.sh bash script (bash start.sh or ./start.sh) (make sure you chmod +x the script first)
2. Manually running the binary (./FactoCord)

## Installing as a service

To install FactoCord as a service so that it can run on startup, you can use the provided service.sh

*Note you must run service.sh as root/sudo to install it as a service*

Example of running service.sh:
`./service.sh factorio /home/facotrio/factocord/`


## Compiling

`Requires go 1.8 or above`

FactoCord uses the following packages:

- [DiscordGo](https://github.com/bwmarrin/discordgo)
- [godotenv](https://github.com/joho/godotenv/)
- [tails](https://github.com/hpcloud/tail)

You will need to add these lib as go get:

- `go get github.com/bwmarrin/discordgo`
- `go get github.com/joho/godotenv`
- `go get github.com/hpcloud/tail/...`

To compile just do `go build`


## Error reporting

When FactoCord encounters an error will log to error.log within the same directory as itself.

If you are having an issue make sure to check the error.log to see what the problem is.

If you are unable to solve the issue yourself, please post an issue containing the error.log and I will review and attempt to solve what the problem is.


## Windows Support?

Currently I haven't had any luck getting FactoCord to run correctly on Windows, [see this](https://github.com/FactoKit/FactoCord/issues/3) for information

If a way is found to fix this problem, then Windows support will be added.


## Screenshots

<p><img src="http://i.imgur.com/JsLOVst.png" alt="restart command"></p>
<p><img src="http://i.imgur.com/1cxq54P.png" alt="mod list command"></p>
<p><img src="http://i.imgur.com/qN3NsO6.png" alt="stop command"></p>
<p><img src="http://i.imgur.com/cxjvFG8.png" alt="save command"></p>
<p><img src="http://i.imgur.com/dztOTrk.png" alt="in-game chat being sent to discord, notice how you can mention discord members"></p>
<p><img src="http://i.imgur.com/Npl0vBb.png" alt="discord chat being sent to in-game"></p>


## Special Thanks

  - Brett and MajesticFudgie for making the logo!
  - [UclCommander](https://github.com/UclCommander) for finding me the tails library which made this a lot easier to build.
