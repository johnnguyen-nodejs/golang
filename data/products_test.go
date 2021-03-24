package data
import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "tuan",
		Price: 1,
		SKU: "abc-abc-abc",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}