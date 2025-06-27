package utils

import (
	"log/slog"
	"os/exec"
	"strings"
)

func Execute(command []string) ([]byte, error) {
	/*
		Example:

		command := []string{
		    "/<path>/yourscript.sh",
		    "arg1=val1",
		    "arg2=val2",
		}

		Execute(command)
	*/
	script := command[0]

	// cmd := &exec.Cmd{
	//	Path:   script,
	//	Args:   command,
	//	Stdout: os.Stdout,
	//	Stderr: os.Stderr,
	//}

	cmd := exec.Command(script, command[1:]...)

	/*
		log.WithFields(log.Fields{
		    "animal": "walrus",
		    "number": 1,
		    "size":   10,
		  }).Info("A walrus appears")
		Output:
		time="2015-09-07T08:48:33Z" level=info msg="A walrus appears" animal=walrus number=1 size=10
	*/

	logger := slog.New(AbleHandler)
	logger.With(slog.String("cmd", strings.Join(command, ", ")))
	logger.Debug("Executing command ", "cmd", command)

	//output, err := cmd.CombinedOutput()
	output, err := cmd.Output()
	if err != nil {
		return output, err
	}
	//
	//err = cmd.Wait()
	//if err != nil {
	//	return output, err
	//}

	return output, nil
}
