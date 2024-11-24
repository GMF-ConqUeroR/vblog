package impl_test

import (
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	t.Log(os.Getenv("BLOG_ENV"))
}
