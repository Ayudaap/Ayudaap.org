package organizaciones

import (
	"fmt"
	"testing"
)

func TestGetAllOrganizaciones(t *testing.T) {
	want := 5
	got := GetAllOrganizaciones()
	if len(got) <= 0 {
		t.Errorf("Obtune = %v | Queria = %v", got, want)
	}
	fmt.Printf("%#v", got[0:1])
}
