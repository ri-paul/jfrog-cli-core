package utils

import (
	"github.com/jfrog/jfrog-cli-core/utils/config"
	"github.com/jfrog/jfrog-client-go/artifactory"
	"github.com/jfrog/jfrog-client-go/utils/io"
)

func CreateUploadServiceManager(serverDetails *config.ServerDetails, threads int, dryRun bool, progressBar io.ProgressMgr) (artifactory.ArtifactoryServicesManager, error) {
	return CreateServiceManagerWithProgressBar(serverDetails, threads, dryRun, progressBar)
}

type UploadConfiguration struct {
	Deb                   string
	Threads               int
	MinChecksumDeploySize int64
	ExplodeArchive        bool
	Retries               int
}
