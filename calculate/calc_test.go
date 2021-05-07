package calculate

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"small positive integers", args{2, 3}, 5},
		{"large positive integers", args{895436, 654}, 895436 + 654},
		{"small negative integers", args{-2, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

type calcCase struct {
	Name     string
	A, B     int
	Expected int
}

func createMulTestCase(t *testing.T, c *calcCase) {
	t.Helper()
	t.Run(c.Name, func(t *testing.T) {
		if ans := Mul(c.A, c.B); ans != c.Expected {
			t.Fatalf("%d * %d expected %d, but %d got",
				c.A, c.B, c.Expected, ans)
		}
	})
}

func TestMul(t *testing.T) {
	createMulTestCase(t, &calcCase{"small positive", 2, 3, 6})
	createMulTestCase(t, &calcCase{"small negative", 2, -3, -6})
	createMulTestCase(t, &calcCase{"zero case", 2, 0, 1}) // wrong case
}

func TestMain(m *testing.M) {

}
