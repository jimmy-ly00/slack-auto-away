# Slack Auto Away 
Set Slack to away using legacy API tokens with no need to install apps or use webhooks. Once compiled to a binary, you can set it up to be compatible with Windows, Linux and macOS.

## Instructions 
1. Log in to your Slack team (SLACK_URL=https://[TEAMNAME].slack.com/)
2. Get API_TOKEN by following the following steps:
    1. Go to https://[TEAMNAME].slack.com/customize/emoji
    2. Press F12 to get Developer Tools
    3. Go to JavScript "Console" tab and run to get API TOKEN window.prompt("your api token is: ", TS.boot_data.api_token)
    References: (Method 4) https://github.com/erroneousboat/slack-term/wiki#running-slack-term-without-legacy-tokens and https://github.com/yuya373/emacs-slack#how-to-get-token
3. git clone https://github.com/jimmy-ly00/slack-auto-away
4. Modify SLACK_URL and API_TOKEN variables
5. go get -u -v github.com/akamensky/argparse
6. go build main.go
7. Set up Task Scheduler for Windows or CronJob for Linux and macOS

## Windows

## Linux and macOS
