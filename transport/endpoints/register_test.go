package endpoints

import (
	"context"
	"testing"

	"auth/transport"
)

func TestValidateUser(t *testing.T) {
	ctx := context.Background()
	testCases := []struct {
		desc       string
		u          transport.User
		shouldFail bool
	}{
		{
			desc: "valid user, should pass and return nil err",
			u: transport.User{
				UserName: "sdasfsfsfs",
				FullName: "fsfsdfsdfsfsfs",
				Password: "1fdsfsdfsdfsdfFfd<",
			},
		},
		{
			desc: "invalid user, userName less than 3",
			u: transport.User{
				UserName: "",
				FullName: "fsfsdfsdfsfsfs",
				Password: "fdsfsdfsdfsdfsfd",
			},
			shouldFail: true,
		},
		{
			desc: "invalid user, fullName less than 3",
			u: transport.User{
				UserName: "dfsffdsf",
				FullName: "s",
				Password: "fdsfsdfsdfsdfsfd",
			},
			shouldFail: true,
		},
		{
			desc: "invalid user, invalid password no speacial char and no numeric",
			u: transport.User{
				UserName: "dfsffdsf",
				FullName: "s",
				Password: "bad password",
			},
			shouldFail: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := validateUser(ctx, tC.u)

			if tC.shouldFail {
				if err == nil {
					t.Fail()
					t.Errorf("expected err but got no err")
				}
			} else {
				if err != nil {
					t.Fail()
					t.Errorf("expected no err got %s", err.Error())
				}
			}
		})
	}
}
