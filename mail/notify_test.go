package mail

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewNotification_MissingAccount(t *testing.T) {
	_, err := NewNotification("", "secretpassword")
	if err == nil {
		t.Fatal("expected error when account is missing")
	}
}

func TestNewNotification_MissingPassword(t *testing.T) {
	_, err := NewNotification("test@gmail.com", "")
	if err == nil {
		t.Fatal("expected error when password is missing")
	}
}

func TestNewNotification_Success(t *testing.T) {
	m, err := NewNotification("test@gmail.com", "secretpassword")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if m.Sender != "Notification" {
		t.Errorf("expected Sender 'Notification', got %q", m.Sender)
	}
	if m.Authentication == nil {
		t.Error("expected Authentication to be non-nil")
	}
}

func TestValidateAttachments_NonExistentFile(t *testing.T) {
	err := validateAttachments([]string{"/nonexistent/path/to/file.txt"})
	if err == nil {
		t.Fatal("expected error for non-existent attachment")
	}
}

func TestValidateAttachments_ExistingFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(tmpFile, []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}

	err := validateAttachments([]string{tmpFile})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestValidateAttachments_EmptyList(t *testing.T) {
	err := validateAttachments(nil)
	if err != nil {
		t.Fatalf("unexpected error for empty attachments: %s", err)
	}

	err = validateAttachments([]string{})
	if err != nil {
		t.Fatalf("unexpected error for empty attachments: %s", err)
	}
}

func TestValidateAttachments_ExistingAndMissingFiles(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(tmpFile, []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}

	err := validateAttachments([]string{tmpFile, "/nonexistent/file.txt"})
	if err == nil {
		t.Fatal("expected error when one attachment is missing")
	}
}

func TestNewTemplate(t *testing.T) {
	placeholders := map[string]string{"{{name}}": "Alice"}
	tmpl := NewTemplate(placeholders)
	if tmpl.Map["{{name}}"] != "Alice" {
		t.Errorf("expected Map[{{name}}] = Alice, got %q", tmpl.Map["{{name}}"])
	}
}

func TestReplaceContent(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "template.html")
	content := "<html><body>Hello {{name}}</body></html>"
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	tmpl := NewTemplate(map[string]string{"{{name}}": "Alice"})
	result, err := tmpl.ReplaceContent(tmpFile)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	expected := "<html><body>Hello Alice</body></html>"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestReplaceContent_NonExistentFile(t *testing.T) {
	tmpl := NewTemplate(map[string]string{})
	_, err := tmpl.ReplaceContent("/nonexistent/template.html")
	if err == nil {
		t.Fatal("expected error for non-existent template file")
	}
}
