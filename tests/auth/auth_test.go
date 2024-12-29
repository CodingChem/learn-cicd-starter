package auth_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey(t *testing.T) {
	type testCase struct {
		name  string
		input http.Header
		want  string
	}
	testCases := []testCase{
		{"simple", http.Header{"Authorization": {"ApiKey somekey"}}, "somekey"},
	}
	for _, tc := range testCases {
		result, err := auth.GetAPIKey(tc.input)
		if err != nil {
			t.Fatalf("%s case: GetAPIKey returned error: %s", tc.name, err.Error())
		}
		if !reflect.DeepEqual(result, tc.want) {
			t.Fatalf("expected: %v, got: %v", result, tc.want)
		}
	}
}
