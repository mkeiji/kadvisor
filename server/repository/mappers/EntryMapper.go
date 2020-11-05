package mappers

import (
	s "kadvisor/server/repository/structs"
	"kadvisor/server/resources/constants"
	"time"
)

type EntryMapper struct{}

func (e *EntryMapper) MapEntry(entry s.Entry) s.Entry {
	entry = e.mapEntryDate(entry)
	entry = e.mapEntryAmount(entry)
	return entry
}

func (e *EntryMapper) mapEntryDate(
	entry s.Entry,
) s.Entry {
	utc, _ := time.LoadLocation("UTC")
	entry.Date = entry.Date.In(utc)
	return entry
}

func (e *EntryMapper) mapEntryAmount(
	entry s.Entry,
) s.Entry {
	if entry.EntryTypeCodeID == constants.EXPENSE_ENTRY_TYPE {
		if entry.Amount > 0 {
			entry.Amount = entry.Amount * -1
		}
	}

	if entry.EntryTypeCodeID == constants.INCOME_ENTRY_TYPE {
		if entry.Amount < 0 {
			entry.Amount = entry.Amount * -1
		}
	}
	return entry
}
