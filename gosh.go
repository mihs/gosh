package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usage())
		os.Exit(2)
	}
	execute(computeRunPath(os.Args[1]))
}

func computeRunPath(scriptPath string) string {
	f, err := os.Open(scriptPath)
	if err != nil {
		fatal(2, "Error reading file: %v\n", err)
	}
	defer f.Close()

	shebang := make([]byte, 2)
	_, err = f.Read(shebang)
	if err != nil && err != io.EOF {
		fatal(2, "Error reading file: %v\n", err)
	}
	if !bytes.Equal([]byte("#!"), shebang) {
		return scriptPath
	}

	fileCopy, err := ioutil.TempFile("", "gosh*.go")
	if err != nil {
		fatal(3, "Error creating temporary file: %v\n", err)
	}
	_, err = fileCopy.WriteString("//#!")
	if err != nil {
		fatal(3, "Error writing to temporary file: %v\n", err)
	}
	tee := io.TeeReader(f, fileCopy)
	_, err = ioutil.ReadAll(tee)
	if err != nil {
		fatal(3, "Error writing to temporary file: %v\n", err)
	}
	err = fileCopy.Close()
	if err != nil {
		fatal(3, "Error writing temporary file: %v\n", err)
	}
	return fileCopy.Name()
}

func execute(filePath string) {
	goExec, err := exec.LookPath("go")
	if err != nil {
		fatal(4, "Error locating the go executable: %v\n", err)
	}
	if err := syscall.Exec(goExec, []string{"go", "run", filePath}, os.Environ()); err != nil {
		fatal(5, "Error executing: %v\n", err)
	}
}

func fatal(code int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "gosh: "+format, args...)
	os.Exit(code)
}

func usage() string {
	desc := `Enables direct execution of Golang source code at <path>.
This program is useful in supporting shebang directives inside Go source files.
	`
	return "usage: " + os.Args[0] + " <path>\n\n" + desc
}
