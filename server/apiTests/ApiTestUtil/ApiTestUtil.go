package ApiTestUtil

import (
	s "kadvisor/server/repository/structs"
)

func CreateTestUsers() []s.User {
	return []s.User{
		{
			FirstName: "testAdmin",
			LastName:  "tester",
			IsPremium: true,
			Phone:     "111-111-1111",
			Address:   "test_address",
			Login: s.Login{
				RoleID:   1,
				Email:    "admin@test.com",
				UserName: "testAdmin",
				Password: "admin",
			},
		},
		{
			FirstName: "testUser",
			LastName:  "user",
			IsPremium: true,
			Phone:     "222-222-2222",
			Address:   "test_address",
			Login: s.Login{
				RoleID:   2,
				Email:    "user@test.com",
				UserName: "testUser",
				Password: "password",
			},
		},
	}
}

func CreateTestForecast(
	userID int,
	year int,
	defaultIncome float64,
	defaultExpense float64,
) s.Forecast {
	var entries []s.ForecastEntry
	forecast := s.Forecast{
		UserID: userID,
		Year:   year,
	}

	for i := 1; i < 13; i++ {
		newEntry := s.ForecastEntry{
			Month:   i,
			Income:  defaultIncome,
			Expense: defaultExpense,
		}
		entries = append(entries, newEntry)
	}

	forecast.Entries = entries
	return forecast
}
