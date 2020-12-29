package password

import (
	"testing"

	"github.com/S0obi/k8s-secret-auditor/pkg/config"
)

func TestIsCompliant(t *testing.T) {
	conf := config.NewConfig("../../config.yaml")

	passNotCompliant := NewPassword("password", "P4ssw0rd")
	passCompliant := NewPassword("password", "4M2&+(fW@3<K!3V!")

	shouldBeNotCompliant := passNotCompliant.IsCompliant(conf)
	shouldBeCompliant := passCompliant.IsCompliant(conf)

	if shouldBeNotCompliant {
		t.Errorf("passNotCompliant : result was incorrect, got: %t, want: %t.", shouldBeNotCompliant, false)
	}

	if !shouldBeCompliant {
		t.Errorf("passCompliant : result was incorrect, got: %t, want: %t.", shouldBeCompliant, true)
	}
}

func TestIsPassword(t *testing.T) {
	conf := config.NewConfig("../../config.yaml")

	thisIsAPassword := IsPassword("thisIsAPassword", conf)
	notAPassword := IsPassword("configValue", conf)

	if notAPassword {
		t.Errorf("notAPassword : result was incorrect, got: %t, want: %t.", notAPassword, false)
	}

	if !thisIsAPassword {
		t.Errorf("thisIsAPassword : result was incorrect, got: %t, want: %t.", thisIsAPassword, true)
	}
}

func TestComputeEntropy(t *testing.T) {
	a := computeEntropy("password")
	b := computeEntropy("p4ssw0rd")
	c := computeEntropy("P4ssw0rd")
	d := computeEntropy("P@ssw0rd")

	if a == 37.603518 {
		t.Errorf("password : result was incorrect, got: %f, want: %f.", a, 37.603518)
	}

	if b == 41.359400 {
		t.Errorf("p4ssw0rd : result was incorrect, got: %f, want: %f.", b, 41.359400)
	}

	if c == 47.633570 {
		t.Errorf("P4ssw0rd : result was incorrect, got: %f, want: %f.", c, 47.633570)
	}

	if d == 52.917679 {
		t.Errorf("P@ssw0rd : result was incorrect, got: %f, want: %f.", d, 40.0)
	}
}
