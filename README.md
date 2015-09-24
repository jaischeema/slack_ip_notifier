# Slack IP Notifier

Super simple script to call slack webhook with caller's IP address

## Build

`go build`

## Usage

`./slack_ip_notifier --slack_hook_url="http://slack.hook/url"`

This will write the current ip address to `~/.last_ip` and won't make any
redundant calls if the IP hasn't changed last time.

## License

MIT
