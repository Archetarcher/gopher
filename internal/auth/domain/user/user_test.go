package user

import (
	"errors"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		login     string
		password  string
		firstname string
		lastname  string
	}
	tests := []struct {
		name        string
		args        args
		want        User
		expectedErr error
	}{
		{
			name: "#1 [positive test]: with all fields",
			args: args{
				login:     "test",
				password:  "test",
				firstname: "test",
				lastname:  "test",
			},
			expectedErr: nil,
		},
		{
			name: "#2 [negative test]: without all fields",
			args: args{
				login:     "login",
				password:  "",
				firstname: "",
				lastname:  "",
			},
			expectedErr: ErrInvalidPerson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewUser(tt.args.login, tt.args.password, tt.args.firstname, tt.args.lastname)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.expectedErr)
				return
			}

		})
	}
}
