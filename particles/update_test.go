package particles

import (
	"testing"
	"project-particles/config"
)

func TestAdd_pointeur(t *testing.T) {
	s := NewSystem()
	s.pointeur()
	if len(s.Content) != config.General.InitNumParticles + 1 {
		t.Error()
	}
}

func (systeme *System) pointeur() {
	systeme = Add_pointeur(systeme)
}