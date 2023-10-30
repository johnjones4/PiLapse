# PiLapse

A simple and easy timelapse tool for the Raspberry Pi.

## Installation

From a freshly installed copy of Raspberry Pi OS running on a Pi or Pi Zero:

1. Download the latest `pilapse.tar.gz` release from https://github.com/johnjones4/pilapse/releases: `wget https://github.com/johnjones4/pilapse/releases/download/0.0.1-alpha-9/pilapse.tar.gz`
2. Extract the archive: `tar zxvf pilapse.tar.gz`
3. Enter the distribution directory: `cd pilapse`
4. Install: `sudo ./install.sh`

## Configuration

There is a configuration file where you can specify credentials for FTP and Telegram integration. That is located at `/etc/pilapse/config.json`. Any time you edit that file, restart Pilapse by running `sudo systemctl restart pilapse.service`.

```json
{
	"ftp": {
		"host": "",
		"username": "",
		"password": "",
		"path": ""
	},
	"telegram": {
		"botId": "",
		"chatId": 0,
		"sendInterval": ""
	}
}
```

### FTP

Every time there is a new image captured from the timelapse, PiLapse will push the image to the FTP location specified here.

### Telegram

You can also configure PiLapse to regularly send captured frames over Telegram. To set this up, create a new bot and specify the bot's token in the `botId` field. (To discover a Chat ID, message the bot from your own Telegram account and then call the Telegram API's `getUpdates` function. In the reponse there will be a chat object with an ID field.) Finally, you can specify how often you get Telegram messages instead of every time there's a new frame. To do that, specify a duration in the `sendInterval` field using Go duration formatting. (ie `10s`, `1m30s`, etc).

## Use

In your browser, go to `http://<Raspberry Pi IP or Hostname>/`. From there, you can configure a capture job, and then once it's running, see stats and the latest capture from the job.
