package main

import "testing"

func Test_WithNode(t *testing.T) {

	t.Run("given node id greater than zero tsid should contain correct node id", func(t *testing.T) {
		for i := 0; i < 20; i++ {

			nodeBits := NODE_BITS_1024
			shift := RANDOM_BITS - nodeBits
			mask := (1 << nodeBits) - 1

			node := int32(500 & mask)
			tsidFactory, err := TsidFactoryBuilder().
				WithNode(node).
				WithNodeBits(nodeBits).
				NewInstance()
			if err != nil {
				t.Error(err.Error())
				t.FailNow()
			}

			tsid, err := tsidFactory.Generate()
			if err != nil {
				t.Error(err.Error())
				t.FailNow()
			}

			actualNode := int32(uint32(tsid.GetRandom())>>shift) & int32(mask)
			if actualNode != node {
				t.FailNow()
			}
		}
	})

	t.Run("should use default node id as zero", func(t *testing.T) {
		for i := 0; i < 20; i++ {

			nodeBits := NODE_BITS_1024
			shift := RANDOM_BITS - nodeBits
			mask := (1 << nodeBits) - 1

			tsidFactory, err := TsidFactoryBuilder().
				WithNodeBits(nodeBits).
				NewInstance()
			if err != nil {
				t.Error(err.Error())
				t.FailNow()
			}

			tsid, err := tsidFactory.Generate()
			if err != nil {
				t.Error(err.Error())
				t.FailNow()
			}

			actualNode := int32(uint32(tsid.GetRandom())>>shift) & int32(mask)
			if actualNode != 0 {
				t.FailNow()
			}
		}
	})
}

func Test_WithNodeBits(t *testing.T) {

	t.Run("given node bits should use correct node bits in tsid", func(t *testing.T) {
		// possible node bits are from [0, 20]. testing all
		for i := 0; i <= 20; i++ {

			nodeBits := int32(i)
			shift := RANDOM_BITS - nodeBits
			mask := (1 << nodeBits) - 1

			node := int32(500 & mask)
			tsidFactory, err := TsidFactoryBuilder().
				WithNode(node).
				WithNodeBits(nodeBits).
				NewInstance()
			if err != nil {
				t.Error(err.Error())
				t.FailNow()
			}

			tsid, err := tsidFactory.Generate()
			if err != nil {
				t.Error(err.Error())
				t.FailNow()
			}

			actualNode := int32(uint32(tsid.GetRandom())>>shift) & int32(mask)
			if actualNode != node {
				t.FailNow()
			}
		}
	})

	t.Run("should use default node bits in tsid when not provided", func(t *testing.T) {
		defaultNodeBits := 0
		for i := 0; i <= 20; i++ {

			shift := RANDOM_BITS - int32(defaultNodeBits)
			mask := (1 << defaultNodeBits) - 1

			node := int32(500 & mask)
			tsidFactory, err := TsidFactoryBuilder().
				WithNode(node).
				NewInstance()
			if err != nil {
				t.Error(err.Error())
				t.FailNow()
			}

			tsid, err := tsidFactory.Generate()
			if err != nil {
				t.Error(err.Error())
				t.FailNow()
			}

			actualNode := int32(uint32(tsid.GetRandom())>>shift) & int32(mask)
			if actualNode != 0 {
				t.FailNow()
			}
		}
	})
}

func Test_WithRandom(t *testing.T) {

	t.Run("given random should not return error", func(t *testing.T) {

		supplier := NewMathRandomSupplier()
		random := NewIntRandom(supplier)

		tsidFactory, err := TsidFactoryBuilder().
			WithRandom(random).
			NewInstance()
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		_, err = tsidFactory.Generate()
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}
	})

	t.Run("should use default random when not provided", func(t *testing.T) {

		tsidFactory, err := TsidFactoryBuilder().
			NewInstance()
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		_, err = tsidFactory.Generate()
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}
	})
}
