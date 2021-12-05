package console

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

const (
	errorStr      = "ERROR"
	fatalErrorStr = "FATAL ERROR"
	panicStr      = "PANIC"
)

// InitMessage prints the HITO logo in ASCII characters.
func InitMessage() {
	fmt.Print("\n/--------------------\\\n")
	fmt.Println("|~0~0~~000~~000~~000~|")
	fmt.Println("|~000~~~0~~~~0~~~0~0~|")
	fmt.Println("|~0~0~~000~~~0~~~000~|")
	fmt.Print("\\--------------------/\n\n")
}

// ServerInfo prints information about the server.
func ServerInfo(portNumber string) {
	fmt.Printf("\nListening on port http://localhost%s\n\n", portNumber)
}

// ClosingMessage prints closing information.
func ClosingMessage(s string) {
	fmt.Printf("\n%s", formatPrefix(true, "", s))
}

// AssertError asserts an error and prints it.
func AssertError(err error) {
	if err != nil {
		fmt.Println(formatPrefix(false, errorStr, err.Error()))
	}
}

// AssertFatal asserts a fatal error, prints it and exits the program with code 1.
func AssertFatal(err error) {
	if err != nil {
		fmt.Println(formatPrefix(false, fatalErrorStr, err.Error()))
		os.Exit(1)
	}
}

// AssertPanic asserts an error, prints it and then it panics.
func AssertPanic(err error) {
	if err != nil {
		fmt.Println(formatPrefix(false, panicStr, err.Error()))
		panic(err)
	}
}

// Log prints logging information in a custom format.
func Log(s string) {
	fmt.Println(formatPrefix(true, "", s))
}

// Message prints a message.
func Message(s string) {
	fmt.Println(s)
}

func formatPrefix(l bool, t string, s string) string {
	dt := dateTime()
	if l {
		return fmt.Sprintf("[LOG]%s: %s", dt, s)
	}
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		log.Fatalln("It was not possible to retrieve runtime.Caller()")
	}
	fl := fmt.Sprintf("%s::%d:", filepath.Base(file), line)
	return fmt.Sprintf("[%s]%s >>> %s %s", t, dt, fl, s)
}

func dateTime() string {
	dt := time.Now()
	year := dt.Year()
	month := dt.Month()
	day := dt.Day()
	hour := dt.Hour()
	minute := dt.Minute()
	second := dt.Second()
	return fmt.Sprintf("[%d-%d-%dT%d:%d:%d]", year, month, day, hour, minute, second)
}
