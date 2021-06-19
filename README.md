# gotification
Central notification server written in Go.  
Idea is to have one central server setup with all the config and secrets and setup with the third party libraries,
then just run `curl localhost:8080/discord/channel` to send a message.  

Useful for long commands you don't know when will finish: `$ run-something && curl -d '{"content":"finished running command"}' localhost:8080/discord/channel`
or for cron jobs, for example checking if a website is up/down and pinging you the results, or scraping the weather report.

Currently supports:
- [x] Discord

# Running the service
`go run main.go` will look for `config.toml` in the current directory.
You can also pass in another config file with `-f file.toml` or with standard in like `cat file.toml | go run main.go -f -`
  
See [example.toml](example.toml) for the configuration needed to run with discord.  
To run without a service like Discord, either:
- ignore the config block from your TOML file
- set `disable = true` in the TOML file

Override the settings with environment variables like `DISCORD_TOKEN=<your token> DISCORD_GUILD=<your guild> PORT=5000 go run main.go`.