package safe

import "log"

// Printf is the global function, will be called if Close method returns an error.
var Printf = log.Printf
