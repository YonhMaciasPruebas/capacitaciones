package test

import (
	"regexp"
	"testing"

	"github.com/YonhMaciasPruebas/capacitaciones/Aprendiendo-Golang/ejercicio_modulos"
)

func TestHelloName(t *testing.T) {
	name := "Yonh"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := ejercicio_modulos.Hola("Yonh")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Yonh") = %q,%v,quiere concidencia para %#q,nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := ejercicio_modulos.Hola("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q,%v,quiere "",error`, msg, err)
	}
}
