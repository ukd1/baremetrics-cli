# baremetrics-cli

A super simple, pretty terrible client for [Baremetrics](https://baremetrics.io/?utm_source=rhs&utm_medium=github&utm_campaign=baremetrics-cli). This is un-official.

I use it with [uebersicht](https://github.com/felixhageloh/uebersicht) and a custom widget I should release. It just dumps JSON, so you can use it for whatever.

## Building

```bash
go build baremetrics-cli.go
mv baremetrics-cli /usr/local/bin/
```

## Usage

You'll need to fetch your cookie from Baremetrics using your browser:

1. Open [the dashboard](https://dashboard.baremetrics.io/?utm_source=rhs&utm_medium=github&utm_campaign=baremetrics-cli) in Chrome
2. Open the developer tools (cmd + opt + i)
3. Reload
4. Select the first entry
5. Copy the cookie out of the request headers

```bash
$ baremetrics-cli -cookie '<the cookie text you copied from step 5>'
```

Enjoy fresh JSON from baremetrics - currently only supports 30 day MRR.

## Contributing

1. Fork
2. Pull request
3. Teach me how to test things plz

kthxbai
