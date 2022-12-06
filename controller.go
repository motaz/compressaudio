package main

import (
	"bytes"
	"errors"
	"io/ioutil"

	"os/exec"
	"time"

	"github.com/motaz/codeutils"
)

func listFiles(dir string) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		println("Error: " + err.Error())
	} else {
		for _, file := range files {
			if file.Size() > 10000 {
				afile := dir + file.Name()
				println(afile, file.Size())
				_, err := execShell("sox " + afile + " /tmp/" + file.Name() + ".ogg")
				if err == nil {
					_, err = execShell("mv /tmp/" + file.Name() + ".ogg" + " " + afile)
				}
				//os.Remove("/tmp/" + file.Name() + ".ogg")
			}
		}
	}

}

func doCompression(aday string) {
	var day time.Time
	if aday == "" {
		day = time.Now().AddDate(0, 0, -1)
	} else {
		day, _ = time.Parse("2006/01/02", aday)
	}
	dir := "/var/spool/asterisk/monitor/" + day.Format("2006/01/02/")
	codeutils.WriteToLog("Compress: "+dir, "compress")

	listFiles(dir)
	codeutils.WriteToLog("finished", "compress")
}

func execShell(command string) (result string, err error) {

	var out bytes.Buffer
	var errBuf bytes.Buffer

	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = &out
	cmd.Stderr = &errBuf
	cmd.Run()

	if errBuf.String() != "" {
		err = errors.New(errBuf.String())
		println("Error:  ", err.Error())

	}
	result = out.String()
	println("Result: ", out.String())

	return
}
