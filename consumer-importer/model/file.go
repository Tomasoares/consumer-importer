package model

import (
	"strconv"
	"time"
)

//File struct that contains data related to imported files
type File struct {
	ID         *int64
	Name       string
	ImportDate time.Time
	Successful bool
}

func (f *File) String() string {
	return "[id=" + strconv.FormatInt(*f.ID, 2) +
		", name=" + f.Name +
		", ImportDate=" + f.ImportDate.String() +
		", Successful=" + strconv.FormatBool(f.Successful) + "]"
}
