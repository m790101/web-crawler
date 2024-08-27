package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	cases := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:     "absolute URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="https://blog.boot.dev">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev"},
		},
		{
			name:     "relative URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no href",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a>
			<span>Boot.dev></span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "bad HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html body>
	<a href="path/one">
		<span>Boot.dev></span>
	</a>
</html body>
`,
			expected: []string{"https://blog.boot.dev/path/one"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "handle invalid base URL",
			inputURL: `:\\invalidBaseURL`,
			inputBody: `
<html>
	<body>
		<a href="/path">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected:      nil,
			errorContains: "couldn't parse base URL",
		},
	}
	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URLs %v, got URLs %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}

// // TestHelloName calls greetings.Hello with a name, checking
// // for a valid return value.
// func TestGetUrl(t *testing.T) {

// 	data := `
// 	<html>
// 		<body>
// 			<a href="/path/one">
// 				<span>Boot.dev</span>
// 			</a>
// 			<a href="https://other.com/path/one">
// 				<span>Boot.dev</span>
// 			</a>
// 		</body>
// 	</html>
// 	`

// 	urls, err := getURLsFromHTML(data, "https://blog.boot.dev")
// 	if err != nil {
// 		t.Fatalf("Expected no error, but got %s", err)
// 	}
// 	expectedURLs := []string{"www.google.com", "https://other.com/path/one"}
// 	expectedURLs2 := []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"}
// 	isEqual := reflect.DeepEqual(urls, expectedURLs)
// 	isEqual2 := reflect.DeepEqual(urls, expectedURLs2)

// 	if isEqual {
// 		t.Fatalf("%s not equal to  %s", expectedURLs, urls)
// 	} else {
// 		t.Logf("Expected %s, got %s", expectedURLs, urls)
// 		t.Logf("Expected %s, got %s", expectedURLs, urls)
// 		t.Logf("Expected %s, got %s", expectedURLs, urls)
// 	}

// 	if !isEqual2 {
// 		t.Fatalf("%s not equal to  %s", expectedURLs2, urls)
// 	} else {
// 		t.Logf("Expected %s, got %s", expectedURLs2, urls)
// 		t.Logf("Expected %s, got %s", expectedURLs2, urls)
// 		t.Logf("Expected %s, got %s", expectedURLs2, urls)
// 	}

// }
