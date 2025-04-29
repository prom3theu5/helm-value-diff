package diff

import (
	"reflect"
	"testing"
)

func TestExtractDiff(t *testing.T) {
	tests := []struct {
		name     string
		base     map[string]any
		changed  map[string]any
		expected map[string]any
	}{
		{
			name: "simple value change",
			base: map[string]any{
				"deployment": map[string]any{
					"kind": "Deployment",
				},
			},
			changed: map[string]any{
				"deployment": map[string]any{
					"kind": "DaemonSet",
				},
			},
			expected: map[string]any{
				"deployment": map[string]any{
					"kind": "DaemonSet",
				},
			},
		},
		{
			name: "nested partial changes",
			base: map[string]any{
				"hub": map[string]any{
					"providers": map[string]any{
						"enabled": false,
					},
				},
			},
			changed: map[string]any{
				"hub": map[string]any{
					"providers": map[string]any{
						"enabled": true,
						"cache":   false,
					},
				},
			},
			expected: map[string]any{
				"hub": map[string]any{
					"providers": map[string]any{
						"enabled": true,
						"cache":   false,
					},
				},
			},
		},
		{
			name: "no changes",
			base: map[string]any{
				"deployment": map[string]any{
					"kind": "Deployment",
				},
			},
			changed: map[string]any{
				"deployment": map[string]any{
					"kind": "Deployment",
				},
			},
			expected: map[string]any{},
		},
		{
			name: "new key added",
			base: map[string]any{
				"deployment": map[string]any{},
			},
			changed: map[string]any{
				"deployment": map[string]any{
					"newField": "newValue",
				},
			},
			expected: map[string]any{
				"deployment": map[string]any{
					"newField": "newValue",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractDiff(tt.base, tt.changed)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("extractDiff() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func BenchmarkExtractDiff(b *testing.B) {
	base := map[string]any{
		"deployment": map[string]any{
			"kind":     "Deployment",
			"replicas": 1,
		},
		"service": map[string]any{
			"type": "ClusterIP",
		},
	}
	changed := map[string]any{
		"deployment": map[string]any{
			"kind":     "DaemonSet",
			"replicas": 2,
		},
		"service": map[string]any{
			"type": "ClusterIP",
		},
		"newField": map[string]any{
			"value": "new",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = extractDiff(base, changed)
	}
}
