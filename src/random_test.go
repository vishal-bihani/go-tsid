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

func Test_NewByteRandomWithSupplierFunc(t *testing.T) {

	t.Run("given supplier func nextInt should use supplier func to generate random values", func(t *testing.T) {
		randomBytes := []byte{0, 0, 0, 15}
		supplierFunc := func(length int32) ([]byte, error) {
			return randomBytes, nil
		}

		byteRandom := NewByteRandomWithSupplierFunc(supplierFunc)

		for i := 0; i < 20; i++ {
			var number int32 = 0

			// generating random bytes
			bytes, err := byteRandom.NextBytes(INTEGER_BYTES_32)

			// converting bytes to number
			for j := 0; j < INTEGER_BYTES_32; j++ {
				number = int32(byte(number<<BYTE_SIZE) | (bytes[j] & 0xff))
			}

			if err != nil || number != 15 {
				t.FailNow()
			}
		}
	})

	t.Run("given supplier func nextBytes should use supplier func to generate random values", func(t *testing.T) {
		randomBytes := []byte{0, 0, 0, 25}

		var randomValue int32 = 0
		// converting returned bytes to number
		for i := 0; i < INTEGER_BYTES_32; i++ {
			randomValue = int32(byte(randomValue<<BYTE_SIZE) | (randomBytes[i] & 0xff))
		}

		supplierFunc := func(length int32) ([]byte, error) {
			return randomBytes, nil
		}

		byteRandom := NewByteRandomWithSupplierFunc(supplierFunc)

		for i := 0; i < 20; i++ {
			var actualNumber int32 = 0

			// generate random bytes
			bytes, err := byteRandom.NextBytes(INTEGER_BYTES_32)

			// converting returned bytes to number
			for j := 0; j < INTEGER_BYTES_32; j++ {
				actualNumber = int32(byte(actualNumber<<BYTE_SIZE) | (bytes[j] & 0xff))
			}

			if err != nil || randomValue != actualNumber {
				t.FailNow()
			}
		}
	})
}
