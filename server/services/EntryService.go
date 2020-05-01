package services

import (
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
)

type EntryService struct {
	repository 	repository.EntryRepository
}

func (svc *EntryService) GetManyByUserId(
	userID int) ([]structs.Entry, error) {
	return svc.repository.FindAllByUserId(userID)
}

func (svc *EntryService) GetManyByClassId(
	classID int) ([]structs.Entry, error) {
	return svc.repository.FindAllByClassId(classID)
}

func (svc *EntryService) GetManyBySubClassId(
	subclassID int) ([]structs.Entry, error) {
	return svc.repository.FindAllBySubClassId(subclassID)
}

func (svc *EntryService) GetManyByClassAndSubClassId(
	classID int, subclassID int) ([]structs.Entry, error) {
	return svc.repository.FindAllByClassIdAndSubClassId(classID, subclassID)
}

func (svc *EntryService) GetOneById(
	id int) (structs.Entry, error) {
	return svc.repository.FindOne(id)
}

func (svc *EntryService) Post(
	entry structs.Entry) (structs.Entry, error) {
	return svc.repository.Create(entry)
}

func (svc *EntryService) Delete(
	id int) (int, error) {
	return svc.repository.Delete(id)
}