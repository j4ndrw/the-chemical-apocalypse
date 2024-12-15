package utils

import "log"

func Assert(condition bool, message string) {
	if !condition {
		log.Fatal("Assertion Failed: ", message)
	}
}

func AssertNotNil[T any](ptr *T, message string) {
	if ptr == nil {
		log.Fatal("Assertion Failed - expected non-nil pointer: ", message)
	}
}

func AssertNil[T any](ptr *T, message string) {
	if ptr != nil {
		log.Fatal("Assertion Failed - expected nil pointer: ", message)
	}
}
