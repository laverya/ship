package fs

import (
	"log"
	"os"

	"github.com/spf13/afero"
)

// NewBaseFilesystem creates a new Afero OS filesystem
func NewBaseFilesystem() afero.Afero {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return afero.Afero{Fs: afero.NewBasePathFs(afero.NewOsFs(), workingDir)}
}
