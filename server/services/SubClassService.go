package services

import (
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
)

type SubClassService struct {
	repository 	repository.SubClassRepository
}

func (svc *SubClassService) GetManyByUserId(
	userID int) ([]structs.SubClass, error) {
	return svc.repository.FindAllByUserId(userID)
}

func (svc *SubClassService) GetManyByClassId(
	classID int) ([]structs.SubClass, error) {
	return svc.repository.FindAllByClassId(classID)
}

func (svc *SubClassService) GetOneById(
	subclassID int) (structs.SubClass, error) {
	return svc.repository.FindOne(subclassID)
}

func (svc *SubClassService) Post(
	subclass structs.SubClass) (structs.SubClass, error) {
	return svc.repository.Create(subclass)
}

func (svc *SubClassService) Delete(
	id int) (int, error) {
	return svc.repository.Delete(id)
}
