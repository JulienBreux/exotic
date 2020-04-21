package path_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/julienbreux/exotic/pkg/path"
)

func TestPath(t *testing.T) {
	var p path.Path

	expected := "/home/jbx/.kube/config"

	assert.NoError(t,
		p.Decode("//home/julien/../jbx/.kube//config"),
	)

	assert.Equal(
		t,
		expected,
		string(p),
	)

	assert.Equal(
		t,
		expected,
		p.String(),
	)
}
