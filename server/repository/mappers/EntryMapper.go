package mappers

import (
	"kadvisor/server/libs/KeiGenUtil"
	s "kadvisor/server/repository/structs"
	"kadvisor/server/resources/constants"
)

type EntryMapper struct{}

func (this *EntryMapper) MapEntry(entry s.Entry) s.Entry {
	entry = this.MapEntryDate(entry)
	entry = this.mapEntryAmount(entry)
	return entry
}

func (this *EntryMapper) MapEntryDate(
	entry s.Entry,
) s.Entry {
	return this.MapEntriesDates([]s.Entry{entry})[0]
}

func (this *EntryMapper) MapEntriesDates(
	entries []s.Entry,
) []s.Entry {
	if len(entries) == 0 {
		return entries
	}

	var updatedEntries []s.Entry
	for _, entry := range entries {
		entry.Date = KeiGenUtil.DateToUTCISO8601(entry.Date)
		updatedEntries = append(updatedEntries, entry)
	}

	return updatedEntries
}

func (this *EntryMapper) mapEntryAmount(
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
