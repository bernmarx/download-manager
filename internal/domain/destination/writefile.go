package destination

import (
	"os"

	"github.com/bernmarx/download-manager/internal/infrastructure/logger"
)

func (d *Destination) WriteFile(filename string, data []byte) error {
	flags := os.O_CREATE | os.O_WRONLY
	filepath := d.Dest + "/" + filename

	logger.Logger().Info("filepath: " + filepath)

	fileInfo, err := os.Stat(filepath)

	if err == nil {
		if fileInfo.Size() > 0 {
			flags = os.O_APPEND | os.O_WRONLY
		} else {
			flags = os.O_WRONLY
		}
	}

	fileWriter, err := os.OpenFile(filepath, flags, 0666)

	if err != nil {
		return err
	}

	_, err = fileWriter.Write(data)

	return err
}
