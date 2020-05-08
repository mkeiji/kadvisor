package services

import (
	"kadvisor/server/repository"
	"kadvisor/server/repository/structs"
)

type EntryService struct {
	repository 	repository.EntryRepository
}

func (svc *EntryService) GetManyByUserId(
	userID int, limit int) ([]structs.Entry, error) {
	return svc.repository.FindAllByUserId(userID, limit)
}

func (svc *EntryService) GetManyByClassId(
	classID int, limit int) ([]structs.Entry, error) {
	return svc.repository.FindAllByClassId(classID, limit)
}

func (svc *EntryService) GetOneById(
	id int) (structs.Entry, error) {
	return svc.repository.FindOne(id)
}

func (svc *EntryService) Post(
	entry structs.Entry) (structs.Entry, error) {
	return svc.repository.Create(entry)
}

func (svc *EntryService) Put(
	entry structs.Entry) (structs.Entry, error) {
	return svc.repository.Update(entry)
}

func (svc *EntryService) Delete(
	id int) (int, error) {
	return svc.repository.Delete(id)
}