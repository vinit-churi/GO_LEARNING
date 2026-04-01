package inputoutput

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// captureStdout captures output written to stdout
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// TestBasicPrint tests the BasicPrint function
func TestBasicPrint(t *testing.T) {
	output := captureStdout(func() {
		BasicPrint()
	})

	want := "Hello World!\n"

	if output != want {
		t.Errorf("BasicPrint() output = %q, want %q", output, want)
	}
}

// TestFormattedOutput tests the FormattedOutput function with various inputs
func TestFormattedOutput(t *testing.T) {
	tests := []struct {
		name      string
		inputName string
		age       int
		height    float64
		want      string
	}{
		{
			name:      "Alice example",
			inputName: "Alice",
			age:       30,
			height:    1.65,
			want:      "Name: Alice, Age: 30, Height: 1.65m\n",
		},
		{
			name:      "Bob example",
			inputName: "Bob",
			age:       25,
			height:    1.80,
			want:      "Name: Bob, Age: 25, Height: 1.80m\n",
		},
		{
			name:      "Charlie example",
			inputName: "Charlie",
			age:       35,
			height:    1.75,
			want:      "Name: Charlie, Age: 35, Height: 1.75m\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureStdout(func() {
				FormattedOutput(tt.inputName, tt.age, tt.height)
			})

			if output != tt.want {
				t.Errorf("FormattedOutput(%q, %d, %.2f) output = %q, want %q",
					tt.inputName, tt.age, tt.height, output, tt.want)
			}
		})
	}
}

// TestFormatVerbsPractice tests the FormatVerbsPractice function
func TestFormatVerbsPractice(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{
			name:  "string value",
			value: "hello",
			want:  "Value: hello\nType: string\nString: \"hello\"\n",
		},
		{
			name:  "integer value",
			value: 42,
			want:  "Value: 42\nType: int\n",
		},
		{
			name:  "float value",
			value: 3.14,
			want:  "Value: 3.14\nType: float64\n",
		},
		{
			name:  "boolean value",
			value: true,
			want:  "Value: true\nType: bool\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureStdout(func() {
				FormatVerbsPractice(tt.value)
			})

			if output != tt.want {
				t.Errorf("FormatVerbsPractice(%v) output = %q, want %q", tt.value, output, tt.want)
			}
		})
	}
}

// TestReadUserInput tests the ReadUserInput function
func TestReadUserInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantName string
		wantAge  int
		wantErr  bool
	}{
		{
			name:     "valid input",
			input:    "Alice 30",
			wantName: "Alice",
			wantAge:  30,
			wantErr:  false,
		},
		{
			name:     "another valid input",
			input:    "Bob 25",
			wantName: "Bob",
			wantAge:  25,
			wantErr:  false,
		},
		{
			name:     "valid input with newline",
			input:    "Charlie 35\n",
			wantName: "Charlie",
			wantAge:  35,
			wantErr:  false,
		},
		{
			name:     "invalid input - missing age",
			input:    "Dave",
			wantName: "",
			wantAge:  0,
			wantErr:  true,
		},
		{
			name:     "invalid input - wrong type for age",
			input:    "Eve notanumber",
			wantName: "",
			wantAge:  0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			gotName, gotAge, err := ReadUserInput(reader)

			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUserInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if gotName != tt.wantName {
					t.Errorf("ReadUserInput() name = %q, want %q", gotName, tt.wantName)
				}
				if gotAge != tt.wantAge {
					t.Errorf("ReadUserInput() age = %d, want %d", gotAge, tt.wantAge)
				}
			}
		})
	}
}

// TestReadUserInputWithDifferentReaders tests ReadUserInput with various io.Reader types
func TestReadUserInputWithDifferentReaders(t *testing.T) {
	t.Run("bytes.Buffer reader", func(t *testing.T) {
		var buf bytes.Buffer
		buf.WriteString("TestUser 42")

		name, age, err := ReadUserInput(&buf)
		if err != nil {
			t.Fatalf("ReadUserInput() unexpected error: %v", err)
		}

		if name != "TestUser" {
			t.Errorf("name = %q, want %q", name, "TestUser")
		}
		if age != 42 {
			t.Errorf("age = %d, want %d", age, 42)
		}
	})

	t.Run("strings.Reader", func(t *testing.T) {
		reader := strings.NewReader("AnotherUser 99")

		name, age, err := ReadUserInput(reader)
		if err != nil {
			t.Fatalf("ReadUserInput() unexpected error: %v", err)
		}

		if name != "AnotherUser" {
			t.Errorf("name = %q, want %q", name, "AnotherUser")
		}
		if age != 99 {
			t.Errorf("age = %d, want %d", age, 99)
		}
	})
}

// BenchmarkBasicPrint benchmarks the BasicPrint function
func BenchmarkBasicPrint(b *testing.B) {
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = old }()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BasicPrint()
	}
}

// BenchmarkFormattedOutput benchmarks the FormattedOutput function
func BenchmarkFormattedOutput(b *testing.B) {
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = old }()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FormattedOutput("Alice", 30, 1.65)
	}
}

// ExampleBasicPrint demonstrates the BasicPrint function
func ExampleBasicPrint() {
	BasicPrint()
	// Output: Hello World!
}

// ExampleFormattedOutput demonstrates the FormattedOutput function
func ExampleFormattedOutput() {
	FormattedOutput("Alice", 30, 1.65)
	// Output: Name: Alice, Age: 30, Height: 1.65m
}

// ExampleFormatVerbsPractice_string demonstrates FormatVerbsPractice with a string
func ExampleFormatVerbsPractice_string() {
	FormatVerbsPractice("hello")
	// Output:
	// Value: hello
	// Type: string
	// String: "hello"
}

// ExampleFormatVerbsPractice_integer demonstrates FormatVerbsPractice with an integer
func ExampleFormatVerbsPractice_integer() {
	FormatVerbsPractice(42)
	// Output:
	// Value: 42
	// Type: int
}

// ExampleReadUserInput demonstrates the ReadUserInput function
func ExampleReadUserInput() {
	reader := strings.NewReader("Alice 30")
	name, age, err := ReadUserInput(reader)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	// Output: Name: Alice, Age: 30
}
