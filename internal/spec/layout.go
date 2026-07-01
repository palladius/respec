package spec

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var nonSlugChars = regexp.MustCompile(`[^a-z0-9]+`)

// Sanitize turns arbitrary model output into a lowercase, hyphenated,
// filesystem-safe path segment. Empty or entirely-invalid input falls back
// to "misc" so callers never end up with an empty path segment.
func Sanitize(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = nonSlugChars.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	if s == "" {
		return "misc"
	}
	return s
}

// ResolvePath turns a model-suggested category/slug into a concrete
// destination directory and SPEC.md path under baseDir.
func ResolvePath(baseDir, category, slug string) (dir, specPath string) {
	dir = filepath.Join(baseDir, Sanitize(category), Sanitize(slug))
	specPath = filepath.Join(dir, "SPEC.md")
	return dir, specPath
}

// ErrExists is returned by CheckOverwrite when a SPEC.md already exists at
// the destination and force was not requested.
var ErrExists = errors.New("SPEC.md already exists at destination (use --force to overwrite)")

// CheckOverwrite refuses to proceed if specPath already exists, unless force is set.
func CheckOverwrite(specPath string, force bool) error {
	if force {
		return nil
	}
	if _, err := os.Stat(specPath); err == nil {
		return ErrExists
	}
	return nil
}
