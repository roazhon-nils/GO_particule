package particles

import (
	"testing"
	"project-particles/config"
)

func TestCreation(t *testing.T) {
	s := NewSystem()
	if len(s.Content) != config.General.InitNumParticles {
		t.Error()
	}
}

func TestAdd(t *testing.T) {
	s := NewSystem()
	s = Add(s)
	if len(s.Content) != config.General.InitNumParticles + 1 {
		t.Error()
	}
}