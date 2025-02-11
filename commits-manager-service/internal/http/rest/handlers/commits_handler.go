package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"commits-manager-service/internal/module/commits"

	"github.com/go-chi/chi/v5"
)

type CommitsHandler struct {
	CommitsManagerService commits.CommitsManagerService
}

func NewCommitsHandler(commitPersistence commits.CommitsManagerService) *CommitsHandler {
	return &CommitsHandler{
		CommitsManagerService: commitPersistence,
	}
}

func (h *CommitsHandler) GetAllCommits(w http.ResponseWriter, r *http.Request) {
	repoName := chi.URLParam(r, "repositoryName")

	commits, err := h.CommitsManagerService.GetCommitsByRepositoryName(repoName)
	if err != nil {
		errorJSON(w, errors.New("failed to fetch commits"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "commits",
		Data:    commits,
	}

	writeJSON(w, http.StatusOK, payload)
}

func (h *CommitsHandler) GetTopCommitAuthors(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errorJSON(w, errors.New("invalid limit parameter"), http.StatusBadRequest)
		return
	}

	authors, err := h.CommitsManagerService.GetTopCommitAuthors(limit)
	if err != nil {
		errorJSON(w, errors.New("failed to fetch top commit authors"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "top commit authors",
		Data:    authors,
	}

	writeJSON(w, http.StatusOK, payload)
}

func (h *CommitsHandler) GetTopCommitAuthorsByRepo(w http.ResponseWriter, r *http.Request) {
	repoName := chi.URLParam(r, "repositoryName")
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errorJSON(w, errors.New("invalid limit parameter"), http.StatusBadRequest)
		return
	}

	authors, err := h.CommitsManagerService.GetTopCommitAuthorsByRepoName(repoName, limit)
	if err != nil {
		errorJSON(w, errors.New("failed to fetch top commit authors for repository"), http.StatusBadRequest)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "top commit authors for repository",
		Data:    authors,
	}

	writeJSON(w, http.StatusOK, payload)
}
