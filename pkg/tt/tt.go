package tt

import (
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"
)

func AssertTest(t testing.TB, got, want interface{}) {
	t.Helper()
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func AssertTestDeepEqual(t testing.TB, got, want interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func UnmarshalTest(t *testing.T, rr *httptest.ResponseRecorder, dto interface{}) *interface{} {
	if err := json.Unmarshal(rr.Body.Bytes(), &dto); err != nil {
		t.Fatalf("failed to unmarshal bytes: %v", err)
	}
	return &dto
}
