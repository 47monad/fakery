package templater_test

import (
	"fmt"
	"sync"
	"testing"
	"text/template"

	"github.com/47monad/fakery/internal/templater"
)

func TestExecBindsCallerFuncs(t *testing.T) {
	templater.ResetCache()
	tmpl := "{{prefix}} {{firstName}} {{lastName}}"

	got, err := templater.Exec("person", tmpl, template.FuncMap{
		"prefix":    func() string { return "Dr." },
		"firstName": func() string { return "Ada" },
		"lastName":  func() string { return "Lovelace" },
	})
	if err != nil {
		t.Fatalf("Exec returned error: %v", err)
	}
	if want := "Dr. Ada Lovelace"; got != want {
		t.Fatalf("Exec = %q, want %q", got, want)
	}

	// The same cached template must honor a different func binding.
	got, err = templater.Exec("person", tmpl, template.FuncMap{
		"prefix":    func() string { return "Prof." },
		"firstName": func() string { return "Grace" },
		"lastName":  func() string { return "Hopper" },
	})
	if err != nil {
		t.Fatalf("Exec returned error: %v", err)
	}
	if want := "Prof. Grace Hopper"; got != want {
		t.Fatalf("Exec = %q, want %q", got, want)
	}
}

func TestExecParsesOnce(t *testing.T) {
	templater.ResetCache()
	tmpl := "{{firstName}} {{lastName}}"
	for i := range 100 {
		got, err := templater.Exec("person", tmpl, template.FuncMap{
			"firstName": func() string { return fmt.Sprintf("F%d", i) },
			"lastName":  func() string { return "L" },
		})
		if err != nil {
			t.Fatalf("Exec returned error: %v", err)
		}
		if want := fmt.Sprintf("F%d L", i); got != want {
			t.Fatalf("Exec = %q, want %q", got, want)
		}
	}
	if n := templater.ParseCount(); n != 1 {
		t.Fatalf("parsed %d templates, want exactly 1", n)
	}
}

func TestExecNeverPanics(t *testing.T) {
	templater.ResetCache()
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Exec panicked: %v", r)
		}
	}()

	// Malformed template: error, not panic.
	if _, err := templater.Exec("bad", "{{", nil); err == nil {
		t.Fatal("expected an error for a malformed template")
	}
	// A function name outside the cacheable set still works via fallback.
	got, err := templater.Exec("num", "{{num 3}}", template.FuncMap{
		"num": func(n int) int { return n },
	})
	if err != nil {
		t.Fatalf("uncached Exec returned error: %v", err)
	}
	if got != "3" {
		t.Fatalf("Exec = %q, want %q", got, "3")
	}
	// Calling a func that returns an error must surface as an error.
	if _, err := templater.Exec("err", "{{boom}}", template.FuncMap{
		"boom": func() (string, error) { return "", fmt.Errorf("boom") },
	}); err == nil {
		t.Fatal("expected the func error to propagate")
	}
}

func TestExecConcurrent(t *testing.T) {
	templater.ResetCache()
	tmpl := "{{firstName}} {{lastName}}"
	var wg sync.WaitGroup
	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 200 {
				if _, err := templater.Exec("person", tmpl, template.FuncMap{
					"firstName": func() string { return "A" },
					"lastName":  func() string { return "B" },
				}); err != nil {
					t.Errorf("Exec returned error: %v", err)
					return
				}
			}
		}()
	}
	wg.Wait()
}
