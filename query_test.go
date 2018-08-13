package request

import (
	"net/http"
	"testing"
)

func TestQueryInt64Param(t *testing.T) {
	r, _ := http.NewRequest("GET", "/my/path?my_param=1", nil)

	if res := QueryInt64Param(r, "my_param", 999); res != 1 {
		t.Fatalf(`Unexpected result, got: %q`, res)
	}

	if res := QueryInt64Param(r, "my_param_missing", 999); res != 999 {
		t.Fatalf(`Unexpected result, got: %q`, res)
	}

	r, _ = http.NewRequest("GET", "/my/path?my_param=a", nil)
	if res := QueryInt64Param(r, "my_param", 999); res != 999 {
		t.Fatalf(`Unexpected result, got: %q`, res)
	}
}

func TestQueryIntParam(t *testing.T) {
	r, _ := http.NewRequest("GET", "/my/path?my_param=1", nil)
	if res := QueryIntParam(r, "my_param", 999); res != 1 {
		t.Fatalf(`Unexpected result, got: %q`, res)
	}
}
