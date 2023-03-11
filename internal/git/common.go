package git

import (
	"os"
	"path/filepath"

	"github.com/FantomeBeignet/ugiti/internal/utils"
	"github.com/go-git/go-git/v5"
)

// Open a git repository from a path. Thow an error if the path is not a directory.
func GetRepository(path string) (*git.Repository, error) {
	logger := utils.GetLogger()
	cleanPath := filepath.Clean(path)
	stat, err := os.Stat(cleanPath)	
	if err != nil {
		logger.Fatal(err)
	}
	if !stat.IsDir() {
		logger.Fatal("Path is not a directory")
	}
	repo, err := git.PlainOpen(cleanPath)
	return repo, err
}