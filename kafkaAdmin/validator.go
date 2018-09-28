package kafkaadmin

import (
	"fmt"
)

func validateCleanupPolicy(v interface{}, k string) (ws []string, errors []error) {
	validTypes := map[string]struct{}{
		"compact": {},
		"delete":  {},
	}

	value := v.(string)

	if _, ok := validTypes[value]; !ok {
		errors = append(errors, fmt.Errorf(
			"%q must be one of ['compact', 'delete']", k))
	}
	return
}
