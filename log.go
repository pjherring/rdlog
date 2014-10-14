package rdlog

import (
	"log"
	"os"
)

var (
	TRACE   *log.Logger
	INFO    *log.Logger
	WARNING *log.Logger
	ERROR   *log.Logger
)

func init() {

	TRACE = log.New(os.Stdout,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Llongfile)

	INFO = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Llongfile)

	WARNING = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Llongfile)

	ERROR = log.New(os.Stderr,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Llongfile)
}
