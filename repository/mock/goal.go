package mock

import (
	"errors"

	"github.com/ehardi19/bitment-backend/domain"
)

type GoalMockRepository struct {
	db map[int64]domain.Goal
}

func NewGoalMockRepository() *GoalMockRepository {
	return &GoalMockRepository{
		db: map[int64]domain.Goal{},
	}
}

func (g *GoalMockRepository) GetAllGoalByUser(userID int64) []domain.Goal {
	var goals []domain.Goal

	for _, goal := range g.db {
		if goal.UserID == userID {
			goals = append(goals, goal)
		}
	}

	return goals
}

func (g *GoalMockRepository) CreateGoal(goal domain.Goal) (domain.Goal, error) {
	_, exist := g.db[goal.ID]
	if exist {
		return domain.Goal{}, errors.New("goal with that id already existed")
	}

	g.db[goal.ID] = goal

	return goal, nil
}

func (g *GoalMockRepository) GetGoalByID(id int64) (domain.Goal, error) {
	goal, ok := g.db[id]
	if ok {
		return goal, nil
	}

	return domain.Goal{}, errors.New("not found")
}
