package components

import "testing"

func TestAttributes_Encode(t *testing.T) {
	classAttr := NewAttributes().
		Set("class", "btn").
		Add("class", "btn-primary").
		Set("disabled", "")

	tests := []struct {
		name string
		a    Attributes
		want string
	}{
		{
			name: "Add class names",
			a:    classAttr,
			want: `class="btn btn-primary" disabled`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Encode()

			if got != tt.want {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}

			t.Logf("Got: %s", got)
		})
	}
}
