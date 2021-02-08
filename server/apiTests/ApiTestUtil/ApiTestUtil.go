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
