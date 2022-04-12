package destination

import (
	"os"
)

// Destination represents download destination on a drive
type Destination struct {
	Dest string
}

// NewDestination() is a constructor for Destination that also validates destination path
func NewDestination(dest string) (*Destination, error) {
	d := Destination{dest}

	if d.Dest == "" {
		d.Dest = "."
	}

	_, err := os.Stat(d.Dest)

	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(d.Dest, 0777)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return &d, nil
}

func (d *Destination) GetDest() string {
	return d.Dest
}
