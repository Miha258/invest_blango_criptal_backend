package service

import (
	"errors"
	"invest_blango_criptal_backend/models"
	"invest_blango_criptal_backend/repository"
)



type AcountService struct {
	repo repository.Account
}


func NewAcountService(repo repository.Account) *AcountService {
	return &AcountService{repo: repo}
}


func (s *AcountService) ChangePassword(userId int, newPassword string) error {
	hashedPassword, err := hashPassword(newPassword)

	if err != nil {
		return errors.New(err.Error())
	}

	return s.repo.ChangePassword(userId, hashedPassword)
}


func (s *AcountService) EditUserData(userId int, newDocs models.UserDocs) error {
	return s.repo.EditUserData(userId, newDocs)
}


func (s *AcountService) UpdateBalance(userId int, amount int64) error {
	return s.repo.UpdateBalance(userId, amount)
}


func (s *AcountService) CreatePromocode(userId int, promocodeName string) error {
	return s.repo.CreatePromocode(userId, promocodeName)
}


func (s * AcountService) AcceptPromocodeUsage(userId int) error {
	return s.repo.AcceptPromocodeUsage(userId)
}