package admin

import (
	"log"
	"path/filepath"

	"github.com/qor/assetfs"
)

func adminAssetFS() (assetfs.Interface, error) {
	abs, err := filepath.Abs("./templates/admin")
	if err != nil {
		return nil, err
	}
	log.Printf("setting assetfs templates path to %s", abs)
	afs := &assetfs.AssetFileSystem{}
	afs.RegisterPath(abs)
	return afs, nil
}
