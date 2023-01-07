package particles

import (
	"container/list"
	"project-particles/config"
	"testing"
)

func TestMain(m *testing.M) {
	config.Get("../config.json")
	m.Run()
}
func Test1(t *testing.T) {
	l := list.New()
	addParticle(l, true, "")
	if l.Len() == 0 {
		t.Error("la fonction n'ajoute pas de particule")
	}
}
func Test2(t *testing.T) {
	p := (&Particle{
		PositionX: config.General.MargeMaxX + 1.0,
		PositionY: config.General.MargeMaxY + 1.0,
	})
	if !isDead(p) {
		t.Error("la fonction isDead return true alors que la particule est en dehors des marge")
	}
}
func Test3(t *testing.T){
	
}
