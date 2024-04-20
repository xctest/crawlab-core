package fs

import (
	"github.com/apex/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"path"
)

func init() {
	rootDir, err := homedir.Dir()
	if err != nil {
		log.Warnf("cannot find home directory: %v", err)
		return
	}
	DefaultWorkspacePath = path.Join(rootDir, "crawlab_workspace")

	workspacePath := viper.GetString("workspace")
	if workspacePath == "" {
		viper.Set("workspace", DefaultWorkspacePath)
	}
}

var DefaultWorkspacePath string
