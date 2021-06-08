package model_test

import (
	"testing"

	"github.com/ShunyaNagashige/ca-tech-dojo-golang/model"
)

func TestGetAll(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name     string
		expected []model.User
		err      error
	}{
		{
			name: "ok",
			expected: []model.User{
				{User_id: 1, User_name: "test", Token: "5555"},
			},
			err: nil,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual, err := model.GetAllUser()

			if err != tt.err {
				t.Errorf("want %#v, got %#v", tt.err, err)
			}

			if len(actual) != len(tt.expected) {
				t.Errorf("len(actual)=%d is different from len(tt.expected)=%d", len(actual), len(tt.expected))
			}

			for i := 0; i < len(tt.expected); i++ {
				if actual[i] != tt.expected[i] {
					t.Errorf("want %#v, got %#v", tt.expected[i], actual[i])
				}
			}
		})
	}
}
