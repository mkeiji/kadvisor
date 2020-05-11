package mappers

import "kadvisor/server/repository/structs"

type ForecastEntryMapper struct {}

func (f *ForecastEntryMapper) MapForecastEntry(
	entry structs.ForecastEntry) structs.ForecastEntry {
	entry = f.mapIncome(entry)
	entry = f.mapExpense(entry)
	return entry
}

func (f *ForecastEntryMapper) mapIncome(
	entry structs.ForecastEntry) structs.ForecastEntry {
	if entry.Income < 0 {
		entry.Income = entry.Income * -1
	}
	return entry
}

func (f *ForecastEntryMapper) mapExpense(
	entry structs.ForecastEntry) structs.ForecastEntry {
	if entry.Expense > 0 {
		entry.Expense = entry.Expense * -1
	}
	return entry
}