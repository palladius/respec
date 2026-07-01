// Package spec renders SPEC.md documents: YAML frontmatter for provenance,
// a default section structure for the body, and output path resolution.
package spec

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// MaxInlineIdeaChars is the longest idea text embedded directly in
// frontmatter; longer ideas are written to a sidecar file instead.
const MaxInlineIdeaChars = 2000

// TokenUsage records token accounting for a generation, for provenance.
type TokenUsage struct {
	Prompt int32 `yaml:"prompt"`
	Output int32 `yaml:"output"`
	Total  int32 `yaml:"total"`
}

// Frontmatter is the YAML metadata block at the top of every generated SPEC.md.
type Frontmatter struct {
	SpeckVersion     string     `yaml:"speck_version"`
	Mode             string     `yaml:"mode"`
	Idea             string     `yaml:"idea,omitempty"`
	IdeaFile         string     `yaml:"idea_file,omitempty"`
	InspirationDir   string     `yaml:"inspiration_dir,omitempty"`
	InspirationFiles []string   `yaml:"inspiration_files,omitempty"`
	TranscriptFile   string     `yaml:"transcript_file,omitempty"`
	CreatedAt        string     `yaml:"created_at"`
	Model            string     `yaml:"model"`
	Tokens           TokenUsage `yaml:"tokens"`
}

// SetIdea stores idea on the frontmatter, inline if short enough. If the
// idea is too long to embed inline, it sets IdeaFile instead and returns the
// idea text for the caller to write to that sidecar file.
func (f *Frontmatter) SetIdea(idea, ideaFileName string) (sidecarContent string, needsSidecar bool) {
	if len(idea) <= MaxInlineIdeaChars {
		f.Idea = idea
		return "", false
	}
	f.IdeaFile = ideaFileName
	return idea, true
}

// Render renders the full SPEC.md document: YAML frontmatter, a title, and the body.
func Render(fm Frontmatter, title, body string) (string, error) {
	yamlBytes, err := yaml.Marshal(fm)
	if err != nil {
		return "", fmt.Errorf("marshal frontmatter: %w", err)
	}
	var b strings.Builder
	b.WriteString("---\n")
	b.Write(yamlBytes)
	b.WriteString("---\n\n")
	if title != "" {
		b.WriteString("# " + title + "\n\n")
	}
	b.WriteString(strings.TrimRight(body, "\n"))
	b.WriteString("\n")
	return b.String(), nil
}

// ParseFrontmatter splits a rendered SPEC.md document back into its
// frontmatter and body. Used by tests to round-trip Render.
func ParseFrontmatter(doc string) (Frontmatter, string, error) {
	var fm Frontmatter
	if !strings.HasPrefix(doc, "---\n") {
		return fm, "", fmt.Errorf("document has no frontmatter delimiter")
	}
	rest := doc[len("---\n"):]
	end := strings.Index(rest, "\n---\n")
	if end == -1 {
		return fm, "", fmt.Errorf("document has no closing frontmatter delimiter")
	}
	yamlPart := rest[:end]
	body := strings.TrimPrefix(rest[end+len("\n---\n"):], "\n")
	if err := yaml.Unmarshal([]byte(yamlPart), &fm); err != nil {
		return fm, "", fmt.Errorf("unmarshal frontmatter: %w", err)
	}
	return fm, body, nil
}
