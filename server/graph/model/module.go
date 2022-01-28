package model

type Module struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Synopsis string `json:"synopsis"`
}

type LearningObjective struct {
	ID       string `json:"id"`
	ModuleID string `json:"module"`
	Name     string `json:"name"`
}
