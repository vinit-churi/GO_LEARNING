package main

import (
	"path/filepath"
	"testing"
)

func TestSaveNote(t *testing.T) {
	path := filepath.Join(t.TempDir(), "note.txt")
	if err := SaveNote(path, "deploy ready"); err != nil {
		t.Fatalf("SaveNote() error = %v", err)
	}
	if got, err := LoadNote(path); err != nil || got != "deploy ready" {
		t.Fatalf("LoadNote() after SaveNote = %q, %v", got, err)
	}
}

func TestLoadNoteTrimsWhitespace(t *testing.T) {
	path := filepath.Join(t.TempDir(), "note.txt")
	if err := SaveNote(path, "  deploy ready\n"); err != nil {
		t.Fatalf("SaveNote() error = %v", err)
	}
	if got, err := LoadNote(path); err != nil || got != "deploy ready" {
		t.Fatalf("LoadNote() = %q, %v", got, err)
	}
}

func TestDuplicateNote(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "src.txt")
	dst := filepath.Join(dir, "dst.txt")
	if err := SaveNote(src, " copied text \n"); err != nil {
		t.Fatalf("SaveNote() error = %v", err)
	}
	if err := DuplicateNote(src, dst); err != nil {
		t.Fatalf("DuplicateNote() error = %v", err)
	}
	if got, err := LoadNote(dst); err != nil || got != "copied text" {
		t.Fatalf("LoadNote(dst) = %q, %v", got, err)
	}
}
