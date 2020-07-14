package dotenveditor

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var stdin = bufio.NewScanner(os.Stdin)

// Run run dotenveditor.
func Run(in, out string) error {
	if err := validate(in, out); err != nil {
		return err
	}

	inFile, err := os.Open(in)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(out)
	if err != nil {
		return err
	}
	defer outFile.Close()

	readAndInput(inFile, outFile)

	return nil
}

func validate(in, out string) error {
	if _, err := os.Stat(in); err != nil {
		return err
	}
	if _, err := os.Stat(out); err == nil {
		fmt.Printf("%s is exists. Are you sure override it? [y,N] ", out)
		stdin.Scan()
		if in := stdin.Text(); in != "y" {
			return fmt.Errorf("bye")
		}
		os.Remove(out)
	}
	return nil
}

func readAndInput(in, out *os.File) {
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		data := scanner.Bytes()
		if i := bytes.IndexByte(data, '='); i >= 0 {
			fmt.Printf("%s=", data[0:i])
			stdin.Scan()
			var tmp []byte
			tmp = append(tmp, data[0:i+1]...)
			data = append(tmp, stdin.Bytes()...)
		}
		data = append(data, '\n')
		out.Write(data)
	}
}
