package monitor

import (
	"testing"
)

func TestSection(t *testing.T) {
	t.Run("regexp", func(t *testing.T) {
		var req Request
		req.URL = "GET /api/user HTTP/1.0"
		if req.Section() != "/api" {
			t.Errorf(`expected="/api", actual="%s"`, req.Section())
		}
		req.URL = "GET /api HTTP/1.0"
		if req.Section() != "/api" {
			t.Errorf(`expected="/api", actual="%s"`, req.Section())
		}
		req.URL = "GET / HTTP/1.0"
		if req.Section() != "/" {
			t.Errorf(`expected="/", actual="%s"`, req.Section())
		}
		req.URL = "1.0"
		if req.Section() != "" {
			t.Errorf(`expected="", actual="%s"`, req.Section())
		}
	})
}
