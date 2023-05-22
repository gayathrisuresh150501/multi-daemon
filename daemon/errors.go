package daemon

import "errors"

var ErrInvalidNodeCount = errors.New("node count less than or equal to zero")
var ErrInvalidDatatype = errors.New("node count not an integer type")
var ErrDirNotFound = errors.New("directory not found at the specified path")
var ErrCouldNotCreateRootInst = errors.New("could not create root instance")
var ErrCouldNotCreateInst = errors.New("could not create instance")
var ErrUnableToSetPath = errors.New("could not set path for the new instances")
var ErrUnableToRunDaemon = errors.New("could not run ipfs daemon")
var ErrCouldNotFindOpenPort= errors.New("could not find open port")