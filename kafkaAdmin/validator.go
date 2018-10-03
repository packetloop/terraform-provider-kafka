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

func validateGreaterThanZero(v interface{}, k string) (ws []string, errors []error) {
	value := v.(int)
	if value < 1 {
		errors = append(errors, fmt.Errorf(
			"%q must be greater than 1 \"%d\"",
			k, value))
	}
	return
}

func validateSegmentBytes(v interface{}, k string) (ws []string, errors []error) {
	value := v.(int)
	if value < 13 {
		errors = append(errors, fmt.Errorf(
			"%q must be greater than 13 \"%d\"",
			k, value))
	}
	return
}
