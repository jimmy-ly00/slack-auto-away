# Slack Auto Away 
Set Slack to away using legacy API tokens without the need to install apps or use webhooks. Once compiled to a binary, you can set up schedules to be away and auto (active) at specific times. Uses go which is cross-compatible with with Windows, Linux and macOS.

## Instructions
1. Log in to your Slack team (https://[TEAMNAME].slack.com/) and get your SLACK_API token by following the following steps:
    1. Go to https://[TEAMNAME].slack.com/customize/emoji
    2. Press F12 to get Developer Tools
    3. Go to JavaScript "Console" tab and run `window.prompt("Your SLACK_API token is: ", TS.boot_data.api_token)` [[1]](https://github.com/erroneousboat/slack-term/wiki#running-slack-term-without-legacy-tokens)[[2]](https://github.com/yuya373/emacs-slack#how-to-get-token)
    
2. Download [releases](https://github.com/jimmy-ly00/slack-auto-away/releases)

3. Set SLACK_API environment variables then restart the Terminals

    * Windows

      ```sh
      setx SLACK_API "xoxs-....."
      ```

    * Linux 

      ```sh
      echo "export SLACK_API="xoxs-....."" >> ~/.bashrc
      ```

    * macOS

      ```sh
      echo "export SLACK_API="xoxs-....."" >> ~/.zshrc
      ```
    
4. Test if it's working e.g. `./slackaway -s away`

5. Set up tasks on Task Scheduler for Windows or cron jobs for Linux and macOS

## Build (optional)
```sh
git clone https://github.com/jimmy-ly00/slack-auto-away 
go get -u -v github.com/akamensky/argparse
go build -o slackaway main.go
```

## Testing

```sh
jly@JIMMY-LAPTOP:/mnt/c/Users/Jimmy/GitHub/slack-auto-away$ ./slackaway -h
usage: slack-auto-away [-h|--help] -s|--status (active|away)

                       Set a slack status

Arguments:

  -h  --help    Print help information
  -s  --status  Set a slack status
jly@JIMMY-LAPTOP:/mnt/c/Users/Jimmy/GitHub/slack-auto-away$ ./slackaway -s away
jly@JIMMY-LAPTOP:/mnt/c/Users/Jimmy/GitHub/slack-auto-away$ ./slackaway -s auto
```

## Scheduling

### Windows

Add two separate scheduled tasks. Weekdays sets auto (active) at 9:00am and sets away at 18:00.

```sh
C:\Users\Jimmy\>schtasks /create /sc weekly /D mon,tue,wed,thu,fri /tn "SlackAutoAway\auto" /tr "'c:\Google Drive\Misc\Scripts\slackaway-windows-amd64.exe' -s auto" /st 09:00

C:\Users\Jimmy\>schtasks /create /sc weekly /D mon,tue,wed,thu,fri /tn "SlackAutoAway\away" /tr "'c:\Google Drive\Misc\Scripts\slackaway-windows-amd64.exe' -s away" /st 18:00
```

### Linux and macOS

Edit crontab to add two jobs. Weekdays sets auto (active) at 9:00am and sets away at 18:00.

```sh
jimmy@Jimmys-Mac ~/> crontab -e
0 9 * * 1-5 /bin/slackaway -s auto
0 18 * * 1-5 /bin/slackaway -s away
jimmy@Jimmys-Mac ~/> crontab -l
0 9 * * 1-5 /bin/slackaway -s auto
0 18 * * 1-5 /bin/slackaway -s away
```

