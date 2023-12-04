package main

import "testing"

func Test_NewIntRandomWithSupplierFunc(t *testing.T) {

	t.Run("given supplier func nextInt should use supplier func to generate random values", func(t *testing.T) {
		randomValue := 5
		supplierFunc := func() (int32, error) {
			return 5, nil
		}

		intRandom := NewIntRandomWithSupplierFunc(supplierFunc)

		for i := 0; i < 20; i++ {
			value, err := intRandom.NextInt()
			if err != nil || randomValue != int(value) {
				t.FailNow()
			}
		}
	})

	t.Run("given supplier func nextBytes should use supplier func to generate random values", func(t *testing.T) {
		randomValue := 10
		supplierFunc := func() (int32, error) {
			return 10, nil
		}

		intRandom := NewIntRandomWithSupplierFunc(supplierFunc)

		for i := 0; i < 20; i++ {
			var number int32 = 0

			// generate random bytes
			bytes, err := intRandom.NextBytes(INTEGER_BYTES_32)

			// converting bytes to number
			for j := 0; j < INTEGER_BYTES_32; j++ {
				number = int32(byte(number<<BYTE_SIZE) | (bytes[j] & 0xff))
			}

			if err != nil || randomValue != int(number) {
				t.FailNow()
			}
		}
	})
}
