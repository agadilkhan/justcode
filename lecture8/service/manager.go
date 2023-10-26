package service

import "lecture8/repository"

type Manager struct {
	Repo repository.Repository
}

func New(repo repository.Repository) *Manager {
	return &Manager{repo}
}
