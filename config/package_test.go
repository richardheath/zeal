package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPackageReplaceSettings(t *testing.T) {
	pkg := Package{
		Name:        "test-package",
		Description: "test description",
		Version:     "1.0.0",
		Keep:        "",
		Split:       false,
		Author:      "{$name}",
		Destination: "/temp/{#path}",
		Files: map[string]string{
			"/src/*": "/{#version}/dest",
		},
		Scripts: map[string][]string{
			"echo": []string{"echo /temp/{#path}/{#version}"},
		},
	}

	settings := map[string]string{
		"#path":    "test/path",
		"#version": "1.0.0",
		"$name":    "test-package",
	}

	pkg.ReplaceSettings(settings)

	assert.Equal(t, "test-package", pkg.Author)
	assert.Equal(t, "/temp/test/path", pkg.Destination)
	assert.Equal(t, "echo /temp/test/path/1.0.0", pkg.Scripts["echo"][0])
	assert.Equal(t, "/1.0.0/dest", pkg.Files["/src/*"])
}
