package commits

import "commits-manager-service/internal/storage/db"
import "commits-manager-service/internal/constants/models"

type CommitsManagerService struct {
	CommitsPersistence db.CommitRepository
}

func NewCommitsManagerService(commitsPersistence db.CommitRepository) CommitsManagerService {
	return CommitsManagerService{CommitsPersistence: commitsPersistence}
}

func (rc CommitsManagerService) GetCommitsByRepositoryName(repoName string) ([]*models.Commit, error) {
	return rc.CommitsPersistence.GetCommitsByRepoName(repoName)
}

func (rc CommitsManagerService) GetTopCommitAuthors(limit int) ([]*models.CommitAuthor, error) {
	return rc.CommitsPersistence.GetTopCommitAuthors(limit)
}
func (rc CommitsManagerService) GetTopCommitAuthorsByRepoName(repoName string, limit int) ([]*models.CommitAuthor, error) {
	return rc.CommitsPersistence.GetTopCommitAuthorsByRepo(repoName, limit)
}
