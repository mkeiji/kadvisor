package KeiPassUtil_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"kadvisor/server/libs/KeiPassUtil"
	s "kadvisor/server/repository/structs"
)

var _ = Describe("KeiPassUtil", func() {
	const password = "testPwd"
	var user s.User

	BeforeEach(func() {
		user = s.User{
			Login: s.Login{
				Password: password,
			},
		}
	})

	Describe("HashAndSalt", func() {
		It("should set hashed password", func() {
			KeiPassUtil.HashAndSalt(&user)
			Expect(password).ToNot(Equal(user.Login.Password))
		})
	})

	Describe("IsValidPassword", func() {
		It("should return true if password match", func() {
			hashedPwd := "$2a$04$2XR0oPLu9ezz9YB3Pcztgu4asG53C3ywLQicQTcrMcS1FLO5R/vLG"
			Expect(KeiPassUtil.IsValidPassword(
				hashedPwd,
				user.Login.Password,
			)).To(BeTrue())
		})

		It("should return false if password does not match", func() {
			hashedPwd := "wrong pwd"
			Expect(KeiPassUtil.IsValidPassword(
				hashedPwd,
				user.Login.Password,
			)).To(BeFalse())
		})
	})
})
