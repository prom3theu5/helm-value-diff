package diff

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Run(basePath, changedPath string) error {
	base, err := loadYaml(basePath)
	if err != nil {
		return fmt.Errorf("failed to load base file: %w", err)
	}
	changed, err := loadYaml(changedPath)
	if err != nil {
		return fmt.Errorf("failed to load changed file: %w", err)
	}
	diff := extractDiff(base, changed)
	return writeYaml(diff)
}

func loadYaml(path string) (map[string]any, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result map[string]any
	err = yaml.Unmarshal(data, &result)
	return result, err
}

func extractDiff(base, changed map[string]any) map[string]any {
	diff := map[string]any{}
	for key, changedVal := range changed {
		baseVal, exists := base[key]
		switch changedValTyped := changedVal.(type) {
		case map[string]any:
			if baseMap, ok := baseVal.(map[string]any); ok && exists {
				subDiff := extractDiff(baseMap, changedValTyped)
				if len(subDiff) > 0 {
					diff[key] = subDiff
				}
			} else {
				diff[key] = changedVal
			}
		default:
			if !exists || fmt.Sprintf("%v", baseVal) != fmt.Sprintf("%v", changedVal) {
				diff[key] = changedVal
			}
		}
	}
	return diff
}

func writeYaml(data any) error {
	encoder := yaml.NewEncoder(os.Stdout)
	defer func() {
		if err := encoder.Close(); err != nil {
			log.Fatalf("failed to close YAML encoder: %v", err)
		}
	}()
	encoder.SetIndent(2)
	return encoder.Encode(data)
}
