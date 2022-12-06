// compressaudio project main.go
// compresses Asterisk recording files in /var/spool/asterisk/monitor/
// It converts GSM .wav files to .ogg files
// it can run as crontab job every day to compress yesterday audo files
// Developed by code.sd

package main

import (
	"os"
)

func main() {
	var day string
	// Optional day parameter, if not sent, yesterday date will be used
	if len(os.Args) > 1 {
		day = os.Args[1]
	}
	doCompression(day)
}
