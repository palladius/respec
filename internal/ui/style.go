// Package ui provides consistent colored/emoji terminal output for speck.
package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	successStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).Bold(true)
	errorStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	questionStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("51")).Bold(true)
	infoStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))
	warnStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
)

// Success prints a green, checkmarked status line to stdout.
func Success(format string, a ...any) {
	fmt.Println(successStyle.Render("✅ " + fmt.Sprintf(format, a...)))
}

// Error prints a red, x-marked status line to stderr.
func Error(format string, a ...any) {
	fmt.Fprintln(os.Stderr, errorStyle.Render("❌ "+fmt.Sprintf(format, a...)))
}

// Info prints a dim status line to stdout.
func Info(format string, a ...any) {
	fmt.Println(infoStyle.Render("✨ " + fmt.Sprintf(format, a...)))
}

// Warn prints an amber warning line to stderr.
func Warn(format string, a ...any) {
	fmt.Fprintln(os.Stderr, warnStyle.Render("⚠️  "+fmt.Sprintf(format, a...)))
}

// Question prints a cyan, question-marked prompt to stdout (no trailing newline
// suppressed — callers typically read a line right after).
func Question(text string) {
	fmt.Println(questionStyle.Render("❓ " + text))
}

// Thinking prints a transient "model is working" line to stdout.
func Thinking(format string, a ...any) {
	fmt.Println(infoStyle.Render("🤔 " + fmt.Sprintf(format, a...)))
}
