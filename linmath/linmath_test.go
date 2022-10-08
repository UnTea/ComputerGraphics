package linmath

import (
	"testing"
)

func TestLength(t *testing.T) {
	tests := []struct {
		inputVector    *vector3
		expectedLength float64
	}{
		{NewVector3(0, 0, 0), NewVector3(0, 0, 0).Length()},
		{NewVector3(1, 0, 0), NewVector3(1, 0, 0).Length()},
		{NewVector3(1, 1, 0), NewVector3(1, 1, 0).Length()},
		{NewVector3(1, -1, -1), NewVector3(1, -1, -1).Length()},
		{NewVector3(1, -1, 2), NewVector3(1, -1, 2).Length()},
		{NewVector3(1, -2, 2), NewVector3(1, -2, 2).Length()},
		{NewVector3(-2, -2, 2), NewVector3(-2, -2, 2).Length()},
	}

	for _, ts := range tests {
		vector := ts.inputVector
		length := vector.Length()

		if length != ts.expectedLength {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedLength, length)
		}
	}
}

func TestNormal(t *testing.T) {
	tests := []struct {
		inputVector    *vector3
		expectedNormal *vector3
	}{
		{NewVector3(1, 0, -1), NewVector3(1, 0, -1).Normal()},
		{NewVector3(1, -1, 0), NewVector3(1, -1, 0).Normal()},
		{NewVector3(1, 1, -1), NewVector3(1, 1, -1).Normal()},
		{NewVector3(-1, 1, -2), NewVector3(-1, 1, -2).Normal()},
		{NewVector3(1, -2, 2), NewVector3(1, -2, 2).Normal()},
		{NewVector3(-2, 2, 2), NewVector3(-2, 2, 2).Normal()},
	}

	for _, ts := range tests {
		vector := ts.inputVector
		normal := vector.Normal()

		if normal.x != ts.expectedNormal.x || normal.y != ts.expectedNormal.y || normal.z != ts.expectedNormal.z {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedNormal, normal)
		}
	}
}

func TestNegative(t *testing.T) {
	tests := []struct {
		inputVector      *vector3
		expectedNegative *vector3
	}{
		{NewVector3(1, 0, 0), NewVector3(1, 0, 0).Negative()},
		{NewVector3(1, -1, 0), NewVector3(1, -1, 0).Negative()},
		{NewVector3(1, 1, -1), NewVector3(1, 1, -1).Negative()},
		{NewVector3(1, 1, 2), NewVector3(1, 1, 2).Negative()},
		{NewVector3(1, -2, 2), NewVector3(1, -2, 2).Negative()},
		{NewVector3(-2, 2, 2), NewVector3(-2, 2, 2).Negative()},
	}

	for _, ts := range tests {
		vector := ts.inputVector
		negative := vector.Negative()

		if negative.x != ts.expectedNegative.x || negative.y != ts.expectedNegative.y || negative.z != ts.expectedNegative.z {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedNegative, negative)
		}
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		inputVector1     *vector3
		inputVector2     *vector3
		expectedMultiply *vector3
	}{
		{NewVector3(1, 1, 1), NewVector3(1, 1, 1), NewVector3(1, 1, 1).Multiply(NewVector3(1, 1, 1))},
		{NewVector3(1, 1, 1), NewVector3(2, 1, 2), NewVector3(1, 1, 1).Multiply(NewVector3(2, 1, 2))},
		{NewVector3(1, 1, 1), NewVector3(1, 2, 1), NewVector3(1, 1, 1).Multiply(NewVector3(1, 2, 1))},
		{NewVector3(1, 1, 1), NewVector3(3, 3, 3), NewVector3(1, 1, 1).Multiply(NewVector3(3, 3, 3))},
		{NewVector3(1, 1, 1), NewVector3(-1, -2, -3), NewVector3(1, 1, 1).Multiply(NewVector3(-1, -2, -3))},
		{NewVector3(1, 1, 1), NewVector3(-2, -1, 0), NewVector3(1, 1, 1).Multiply(NewVector3(-2, -1, 0))},
		{NewVector3(1, 1, 1), NewVector3(-3, -3, -3), NewVector3(1, 1, 1).Multiply(NewVector3(-3, -3, -3))},
		{NewVector3(1, 1, 1), NewVector3(0, 0, 0), NewVector3(1, 1, 1).Multiply(NewVector3(0, 0, 0))},
	}

	for _, ts := range tests {
		vector1 := ts.inputVector1
		vector2 := ts.inputVector2
		multiply := vector1.Multiply(vector2)

		if multiply.x != ts.expectedMultiply.x || multiply.y != ts.expectedMultiply.y || multiply.z != ts.expectedMultiply.z {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedMultiply, multiply)
		}
	}
}

func TestMultiplyOnScalar(t *testing.T) {
	tests := []struct {
		inputVector              *vector3
		inputScalar              float64
		expectedMultiplyOnScalar *vector3
	}{
		{NewVector3(0, 0, 0), 5.1, NewVector3(0, 0, 0).MultiplyOnScalar(5.1)},
		{NewVector3(0, 0, 0), -5.4, NewVector3(0, 0, 0).MultiplyOnScalar(-5.4)},
		{NewVector3(1, 1, 1), 11.3, NewVector3(1, 1, 1).MultiplyOnScalar(11.3)},
		{NewVector3(1, 1, 1), 0.8, NewVector3(1, 1, 1).MultiplyOnScalar(0.8)},
		{NewVector3(1, 1, 1), 3.6, NewVector3(1, 1, 1).MultiplyOnScalar(3.6)},
		{NewVector3(1, 1, 1), -4.2, NewVector3(1, 1, 1).MultiplyOnScalar(-4.2)},
		{NewVector3(1, 1, 1), 7.7, NewVector3(1, 1, 1).MultiplyOnScalar(7.7)},
		{NewVector3(1, 1, 1), 8.5, NewVector3(1, 1, 1).MultiplyOnScalar(8.5)},
		{NewVector3(1, 1, 1), -0.1, NewVector3(1, 1, 1).MultiplyOnScalar(-0.1)},
		{NewVector3(1, 1, 1), 0.01, NewVector3(1, 1, 1).MultiplyOnScalar(0.01)},
		{NewVector3(1, 1, 1), -0.001, NewVector3(1, 1, 1).MultiplyOnScalar(-0.001)},
		{NewVector3(1, 1, 1), 14.2, NewVector3(1, 1, 1).MultiplyOnScalar(14.2)},
		{NewVector3(1, 1, 1), -0.2, NewVector3(1, 1, 1).MultiplyOnScalar(-0.2)},
	}

	for _, ts := range tests {
		vector := ts.inputVector
		multiplyOnScalar := vector.MultiplyOnScalar(ts.inputScalar)

		if multiplyOnScalar.x != ts.expectedMultiplyOnScalar.x || multiplyOnScalar.y != ts.expectedMultiplyOnScalar.y || multiplyOnScalar.z != ts.expectedMultiplyOnScalar.z {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedMultiplyOnScalar, multiplyOnScalar)
		}
	}
}

func TestDot(t *testing.T) {
	tests := []struct {
		inputVector1 *vector3
		inputVector2 *vector3
		expectedDot  float64
	}{
		{NewVector3(0, 0, 1.2), NewVector3(-1, -1, -1), NewVector3(0, 0, 1.2).Dot(NewVector3(-1, -1, -1))},
		{NewVector3(0, -1.7, 0), NewVector3(-1, -1, 1), NewVector3(0, -1.7, 0).Dot(NewVector3(-1, -1, 1))},
		{NewVector3(1.1, 2.2, -3.3), NewVector3(1, -1, -1), NewVector3(1.1, 2.2, -3.3).Dot(NewVector3(1, -1, -1))},
		{NewVector3(3, 5, 8), NewVector3(-1, 1, -1), NewVector3(3, 5, 8).Dot(NewVector3(-1, 1, -1))},
		{NewVector3(-4.5, 1, 9), NewVector3(-1, 1, 1), NewVector3(-4.5, 1, 9).Dot(NewVector3(-1, 1, 1))},
		{NewVector3(3.7, 2, 8), NewVector3(1, -1, -1), NewVector3(3.7, 2, 8).Dot(NewVector3(1, -1, -1))},
		{NewVector3(7.5, 3, 7), NewVector3(1, -1, 1), NewVector3(7.5, 3, 7).Dot(NewVector3(1, -1, 1))},
		{NewVector3(8.4, 4, 6), NewVector3(-1, 1, 1), NewVector3(8.4, 4, 6).Dot(NewVector3(-1, 1, 1))},
		{NewVector3(9.3, 5, 5), NewVector3(1, -1, 1), NewVector3(9.3, 5, 5).Dot(NewVector3(1, -1, 1))},
		{NewVector3(-6.4, 6, 4), NewVector3(1, 1, -1), NewVector3(-6.4, 6, 4).Dot(NewVector3(1, 1, -1))},
		{NewVector3(2.3, 7, 3), NewVector3(-1, 1, 1), NewVector3(2.3, 7, 3).Dot(NewVector3(-1, 1, 1))},
		{NewVector3(1.6, 8, 2), NewVector3(1, -1, 1), NewVector3(1.6, 8, 2).Dot(NewVector3(1, -1, 1))},
		{NewVector3(8.7, 9, 1), NewVector3(-1, 1, -1), NewVector3(8.7, 9, 1).Dot(NewVector3(-1, 1, -1))},
	}

	for _, ts := range tests {
		vector1 := ts.inputVector1
		vector2 := ts.inputVector2
		dot := vector1.Dot(vector2)

		if dot != ts.expectedDot {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedDot, dot)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		inputVector1 *vector3
		inputVector2 *vector3
		expectedAdd  *vector3
	}{
		{NewVector3(0, 0, 1.2), NewVector3(-1, -1, -1), NewVector3(0, 0, 1.2).Add(NewVector3(-1, -1, -1))},
		{NewVector3(0, -1.7, 0), NewVector3(-1, -1, 1), NewVector3(0, -1.7, 0).Add(NewVector3(-1, -1, 1))},
		{NewVector3(1.1, 2.2, -3.3), NewVector3(1, -1, -1), NewVector3(1.1, 2.2, -3.3).Add(NewVector3(1, -1, -1))},
		{NewVector3(3, 5, 8), NewVector3(-1, 1, -1), NewVector3(3, 5, 8).Add(NewVector3(-1, 1, -1))},
		{NewVector3(-4.5, 1, 9), NewVector3(-1, 1, 1), NewVector3(-4.5, 1, 9).Add(NewVector3(-1, 1, 1))},
		{NewVector3(3.7, 2, 8), NewVector3(1, -1, -1), NewVector3(3.7, 2, 8).Add(NewVector3(1, -1, -1))},
		{NewVector3(7.5, 3, 7), NewVector3(1, -1, 1), NewVector3(7.5, 3, 7).Add(NewVector3(1, -1, 1))},
		{NewVector3(8.4, 4, 6), NewVector3(-1, 1, 1), NewVector3(8.4, 4, 6).Add(NewVector3(-1, 1, 1))},
		{NewVector3(9.3, 5, 5), NewVector3(1, -1, 1), NewVector3(9.3, 5, 5).Add(NewVector3(1, -1, 1))},
		{NewVector3(-6.4, 6, 4), NewVector3(1, 1, -1), NewVector3(-6.4, 6, 4).Add(NewVector3(1, 1, -1))},
		{NewVector3(2.3, 7, 3), NewVector3(-1, 1, 1), NewVector3(2.3, 7, 3).Add(NewVector3(-1, 1, 1))},
		{NewVector3(1.6, 8, 2), NewVector3(1, -1, 1), NewVector3(1.6, 8, 2).Add(NewVector3(1, -1, 1))},
		{NewVector3(8.7, 9, 1), NewVector3(-1, 1, -1), NewVector3(8.7, 9, 1).Add(NewVector3(-1, 1, -1))},
	}

	for _, ts := range tests {
		vector1 := ts.inputVector1
		vector2 := ts.inputVector2
		add := vector1.Add(vector2)

		if add.x != ts.expectedAdd.x || add.y != ts.expectedAdd.y || add.z != ts.expectedAdd.z {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedAdd, add)
		}
	}
}

func TestSubtraction(t *testing.T) {
	tests := []struct {
		inputVector1        *vector3
		inputVector2        *vector3
		expectedSubtraction *vector3
	}{
		{NewVector3(0, 0, 1.2), NewVector3(-1, -1, -1), NewVector3(0, 0, 1.2).Subtraction(NewVector3(-1, -1, -1))},
		{NewVector3(0, -1.7, 0), NewVector3(-1, -1, 1), NewVector3(0, -1.7, 0).Subtraction(NewVector3(-1, -1, 1))},
		{NewVector3(1.1, 2.2, -3.3), NewVector3(1, -1, -1), NewVector3(1.1, 2.2, -3.3).Subtraction(NewVector3(1, -1, -1))},
		{NewVector3(3, 5, 8), NewVector3(-1, 1, -1), NewVector3(3, 5, 8).Subtraction(NewVector3(-1, 1, -1))},
		{NewVector3(-4.5, 1, 9), NewVector3(-1, 1, 1), NewVector3(-4.5, 1, 9).Subtraction(NewVector3(-1, 1, 1))},
		{NewVector3(3.7, 2, 8), NewVector3(1, -1, -1), NewVector3(3.7, 2, 8).Subtraction(NewVector3(1, -1, -1))},
		{NewVector3(7.5, 3, 7), NewVector3(1, -1, 1), NewVector3(7.5, 3, 7).Subtraction(NewVector3(1, -1, 1))},
		{NewVector3(8.4, 4, 6), NewVector3(-1, 1, 1), NewVector3(8.4, 4, 6).Subtraction(NewVector3(-1, 1, 1))},
		{NewVector3(9.3, 5, 5), NewVector3(1, -1, 1), NewVector3(9.3, 5, 5).Subtraction(NewVector3(1, -1, 1))},
		{NewVector3(-6.4, 6, 4), NewVector3(1, 1, -1), NewVector3(-6.4, 6, 4).Subtraction(NewVector3(1, 1, -1))},
		{NewVector3(2.3, 7, 3), NewVector3(-1, 1, 1), NewVector3(2.3, 7, 3).Subtraction(NewVector3(-1, 1, 1))},
		{NewVector3(1.6, 8, 2), NewVector3(1, -1, 1), NewVector3(1.6, 8, 2).Subtraction(NewVector3(1, -1, 1))},
		{NewVector3(8.7, 9, 1), NewVector3(-1, 1, -1), NewVector3(8.7, 9, 1).Subtraction(NewVector3(-1, 1, -1))},
	}

	for _, ts := range tests {
		vector1 := ts.inputVector1
		vector2 := ts.inputVector2
		subtraction := vector1.Subtraction(vector2)

		if subtraction.x != ts.expectedSubtraction.x || subtraction.y != ts.expectedSubtraction.y || subtraction.z != ts.expectedSubtraction.z {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedSubtraction, subtraction)
		}
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		inputVector1   *vector3
		inputVector2   *vector3
		expectedDivine *vector3
	}{
		{NewVector3(1, 1, 1), NewVector3(1, 1, 1), NewVector3(1, 1, 1).Divide(NewVector3(1, 1, 1))},
		{NewVector3(1, 1, 1), NewVector3(2, 1, 2), NewVector3(1, 1, 1).Divide(NewVector3(2, 1, 2))},
		{NewVector3(1, 1, 1), NewVector3(1, 2, 1), NewVector3(1, 1, 1).Divide(NewVector3(1, 2, 1))},
		{NewVector3(1, 1, 1), NewVector3(3, 3, 3), NewVector3(1, 1, 1).Divide(NewVector3(3, 3, 3))},
		{NewVector3(1, 1, 1), NewVector3(-1, -2, -3), NewVector3(1, 1, 1).Divide(NewVector3(-1, -2, -3))},
		{NewVector3(1, 1, 1), NewVector3(-2, -1, 0), NewVector3(1, 1, 1).Divide(NewVector3(-2, -1, 0))},
		{NewVector3(1, 1, 1), NewVector3(-3, -3, -3), NewVector3(1, 1, 1).Divide(NewVector3(-3, -3, -3))},
		{NewVector3(1, 1, 1), NewVector3(0, 0, 0), NewVector3(1, 1, 1).Divide(NewVector3(0, 0, 0))},
	}

	for _, ts := range tests {
		vector1 := ts.inputVector1
		vector2 := ts.inputVector2
		divine := vector1.Divide(vector2)

		if divine.x != ts.expectedDivine.x || divine.y != ts.expectedDivine.y || divine.z != ts.expectedDivine.z {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedDivine, divine)
		}
	}
}

func TestDivideOnScalar(t *testing.T) {
	tests := []struct {
		inputVector            *vector3
		inputScalar            float64
		expectedDivineOnScalar *vector3
	}{
		{NewVector3(0, 0, 0), 5.1, NewVector3(0, 0, 0).DivideOnScalar(5.1)},
		{NewVector3(0, 0, 0), -5.4, NewVector3(0, 0, 0).DivideOnScalar(-5.4)},
		{NewVector3(1, 1, 1), 11.3, NewVector3(1, 1, 1).DivideOnScalar(11.3)},
		{NewVector3(1, 1, 1), 0.8, NewVector3(1, 1, 1).DivideOnScalar(0.8)},
		{NewVector3(1, 1, 1), 3.6, NewVector3(1, 1, 1).DivideOnScalar(3.6)},
		{NewVector3(1, 1, 1), -4.2, NewVector3(1, 1, 1).DivideOnScalar(-4.2)},
		{NewVector3(1, 1, 1), 7.7, NewVector3(1, 1, 1).DivideOnScalar(7.7)},
		{NewVector3(1, 1, 1), 8.5, NewVector3(1, 1, 1).DivideOnScalar(8.5)},
		{NewVector3(1, 1, 1), -0.1, NewVector3(1, 1, 1).DivideOnScalar(-0.1)},
		{NewVector3(1, 1, 1), 0.01, NewVector3(1, 1, 1).DivideOnScalar(0.01)},
		{NewVector3(1, 1, 1), -0.001, NewVector3(1, 1, 1).DivideOnScalar(-0.001)},
		{NewVector3(1, 1, 1), 14.2, NewVector3(1, 1, 1).DivideOnScalar(14.2)},
		{NewVector3(1, 1, 1), -0.2, NewVector3(1, 1, 1).DivideOnScalar(-0.2)},
	}

	for _, ts := range tests {
		vector := ts.inputVector
		divineOnScalar := vector.DivideOnScalar(ts.inputScalar)

		if divineOnScalar.x != ts.expectedDivineOnScalar.x || divineOnScalar.y != ts.expectedDivineOnScalar.y || divineOnScalar.z != ts.expectedDivineOnScalar.z {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedDivineOnScalar, divineOnScalar)
		}
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		inputVector1  *vector3
		inputMin      float64
		inputMax      float64
		expectedClamp *vector3
	}{
		{NewVector3(0, 0, 1.2), -1, 1, NewVector3(0, 0, 1.2).Clamp(-1, 1)},
		{NewVector3(0, -1.7, 0), -2, 2, NewVector3(0, -1.7, 0).Clamp(-2, 2)},
		{NewVector3(1.1, 2.2, -3.3), -3, 3, NewVector3(1.1, 2.2, -3.3).Clamp(-3, 3)},
		{NewVector3(3, 5, 8), -4, 4, NewVector3(3, 5, 8).Clamp(-4, 4)},
		{NewVector3(-4.5, 1, 9), -2, 6, NewVector3(-4.5, 1, 9).Clamp(-2, 6)},
		{NewVector3(3.7, 2, 8), -2, 3, NewVector3(3.7, 2, 8).Clamp(-2, 3)},
		{NewVector3(7.5, 3, 7), -5, 1, NewVector3(7.5, 3, 7).Clamp(-5, 1)},
		{NewVector3(8.4, 4, 6), -7, 8, NewVector3(8.4, 4, 6).Clamp(-7, 8)},
		{NewVector3(9.3, 5, 5), -9, 9, NewVector3(9.3, 5, 5).Clamp(-9, 9)},
		{NewVector3(-6.4, 6, 4), -2, 2, NewVector3(-6.4, 6, 4).Clamp(-2, 2)},
		{NewVector3(2.3, 7, 3), -4, 4, NewVector3(2.3, 7, 3).Clamp(-4, 4)},
		{NewVector3(1.6, 8, 2), -1, 3, NewVector3(1.6, 8, 2).Clamp(-1, 3)},
		{NewVector3(8.7, 9, 1), -7, 5, NewVector3(8.7, 9, 1).Clamp(-7, 5)},
	}

	for _, ts := range tests {
		vector := ts.inputVector1
		clamp := vector.Clamp(ts.inputMin, ts.inputMax)

		if clamp.x != ts.expectedClamp.x || clamp.y != ts.expectedClamp.y || clamp.z != ts.expectedClamp.z {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedClamp, clamp)
		}
	}
}

func TestPower(t *testing.T) {
	tests := []struct {
		inputVector   *vector3
		inputScalar   float64
		expectedPower *vector3
	}{
		{NewVector3(0.4, 0.5, 1.2), 5.1, NewVector3(0.4, 0.5, 1.2).Power(5.1)},
		{NewVector3(0.1, 1.7, 0.9), 5.4, NewVector3(0.1, 1.7, 0.9).Power(5.4)},
		{NewVector3(1.1, 2.2, 3.3), 1.3, NewVector3(1.1, 2.2, 3.3).Power(1.3)},
		{NewVector3(3, 5, 8), 0.8, NewVector3(3, 5, 8).Power(0.8)},
		{NewVector3(4.5, 1, 9), 3.6, NewVector3(4.5, 1, 9).Power(3.6)},
		{NewVector3(3.7, 2, 8), 4.2, NewVector3(3.7, 2, 8).Power(4.2)},
		{NewVector3(7.5, 3, 7), 7.7, NewVector3(7.5, 3, 7).Power(7.7)},
		{NewVector3(8.4, 4, 6), 3.5, NewVector3(8.4, 4, 6).Power(3.5)},
		{NewVector3(9.3, 5, 5), 0.1, NewVector3(9.3, 5, 5).Power(0.1)},
		{NewVector3(6.4, 6, 4), 0.01, NewVector3(6.4, 6, 4).Power(0.01)},
		{NewVector3(2.3, 7, 3), 0.001, NewVector3(2.3, 7, 3).Power(0.001)},
		{NewVector3(1.6, 8, 2), 1.2, NewVector3(1.6, 8, 2).Power(1.2)},
		{NewVector3(8.7, 9, 1), 0.2, NewVector3(8.7, 9, 1).Power(0.2)},
	}

	for _, ts := range tests {
		vector := ts.inputVector
		scalar := vector.Power(ts.inputScalar)

		if scalar.x != ts.expectedPower.x || scalar.y != ts.expectedPower.y || scalar.z != ts.expectedPower.z {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedPower, scalar)
		}
	}
}

func TestRadians(t *testing.T) {
	tests := []struct {
		inputDegrees    float64
		expectedRadians float64
	}{
		{180, Radians(180)},
		{165, Radians(165)},
		{150, Radians(150)},
		{135, Radians(135)},
		{120, Radians(120)},
		{105, Radians(105)},
		{90, Radians(90)},
		{75, Radians(75)},
		{60, Radians(60)},
		{45, Radians(45)},
		{30, Radians(30)},
		{15, Radians(15)},
	}

	for _, ts := range tests {
		degree := Radians(ts.inputDegrees)

		if degree != ts.expectedRadians {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedRadians, degree)
		}
	}
}

func TestSplat(t *testing.T) {
	tests := []struct {
		inputDegrees  float64
		expectedSplat *vector3
	}{
		{1.5, Splat(1.5)},
		{-1.5, Splat(-1.5)},
	}

	for _, ts := range tests {
		splat := Splat(ts.inputDegrees)

		if splat.x != ts.expectedSplat.x || splat.y != ts.expectedSplat.y || splat.z != ts.expectedSplat.z {
			t.Fatalf("expected [%v] but have [%v]", ts.expectedSplat, splat)
		}
	}
}
