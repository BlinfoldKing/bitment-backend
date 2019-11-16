package services

import (
	"time"

	"github.com/ehardi19/bitment-backend/domain"
)

func (s *Services) GetAllGoalByUser(userID int64) ([]domain.Goal, error) {
	return s.GoalRepo.GetAllGoalByUser(userID), nil
}

func (s *Services) CreateGoal(goal domain.Goal) (domain.Goal, int, error) {
	goal.ID = time.Now().UnixNano()

	goal, err := s.GoalRepo.CreateGoal(goal)
	if err != nil {
		return domain.Goal{}, 400, err
	}

	return goal, 201, nil
}

func (s *Services) GetGoalByID(id int64) (domain.Goal, error) {
	return s.GoalRepo.GetGoalByID(id)
}
