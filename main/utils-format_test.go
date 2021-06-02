package main

import (
	"testing"
)

func TestValidateImageName(t *testing.T) {
	testValidateImageName(t, "mariadb", true)
	testValidateImageName(t, "mariadb:10.4", true)
	testValidateImageName(t, "matomo:3.7.0-fpm", true)
	testValidateImageName(t, "replace-php-serialize-safe:master-SNAPSHOT", true)
	testValidateImageName(t, "replace-php-serialize-safe:master_SNAPSHOT", true)
	testValidateImageName(t, "foilen/replace-php-serialize-safe", true)
	testValidateImageName(t, "foilen/replace-php-serialize-safe:1.0.0", true)
	testValidateImageName(t, "mariadb 10.4", false)
	testValidateImageName(t, "mariadb#10.4", false)
	testValidateImageName(t, "mariadb&10.4", false)
	testValidateImageName(t, "mariadb*10.4", false)
	testValidateImageName(t, "foilen;replace-php-serialize-safe:1.0.0", false)
}

func testValidateImageName(t *testing.T, input string, expected bool) {
	if validateImageName(input) != expected {
		t.Errorf("%s should be %t", input, expected)
	}
}
