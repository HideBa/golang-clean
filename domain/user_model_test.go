package domain

import (
	"reflect"
	"testing"
)

func TestNewUserModel(t *testing.T) {
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name string
		args args
		want *UserModel
	}{
		{name: "It should create new user model with email and username",
			args: args{name: "hoge", email: "hoge@example.com"},
			want: &UserModel{Name: "hoge", Email: "hoge@example.com"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserModel(tt.args.name, tt.args.email); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserModel() = %v, want %v", got, tt.want)
			}
		})
	}
}
