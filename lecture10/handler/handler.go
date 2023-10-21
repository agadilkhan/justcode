package handler

import (
	"lecture10/repo"
	"lecture10/repo/cache"
)

type Handler struct {
	Repo repo.Repo
	cache.BookCache
}

func New(repo repo.Repo, bookCache cache.BookCache) Handler {
	return Handler{repo, bookCache}
}
