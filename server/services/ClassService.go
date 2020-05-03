package services

import (
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
)

type ClassService struct {
	repository 	repository.ClassRepository
}

func (svc *ClassService) GetManyByUserId(
	userID int, preloaded bool) ([]structs.Class, error) {
	return svc.repository.FindAllByUserId(userID, preloaded)
}

func (svc *ClassService) GetOneById(
	id int, preloaded bool) (structs.Class, error) {
	return svc.repository.FindOne(id, preloaded)
}

func (svc *ClassService) Post(
	class structs.Class) (structs.Class, error) {
	return svc.repository.Create(class)
}

func (svc *ClassService) Put(
	class structs.Class) (structs.Class, error) {
	return svc.repository.Update(class)
}

func (svc *ClassService) Delete(
	id int) (int, error) {
	return svc.repository.Delete(id)
}