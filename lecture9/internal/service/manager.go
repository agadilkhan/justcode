package service

import "lecture9/internal/repository"

type Manager struct {
	Repo repository.Repository
}

func NewManager(repo repository.Repository) *Manager {
	return &Manager{Repo: repo}
}
