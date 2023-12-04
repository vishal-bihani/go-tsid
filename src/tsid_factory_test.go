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

// func Test(t *testing.T) {

// 	randomSupplier := NewMathRandomSupplier()
// 	random := NewIntRandom(randomSupplier)

// 	tsidFactory, _ := TsidFactoryBuilder().
// 		WithNode(1).
// 		WithRandom(random).
// 		WithTime(time.Now()).
// 		WithNodeBits(10).
// 		WithCustomEpoch(TSID_EPOCH).
// 		Build()

// 	value, _ := tsidFactory.Generate()

// 	t.Log(value.ToLong())
// 	t.Log(value.ToString())
// 	t.Log(value.ToBytes())

// 	// 121706220661772722 -> 03C3356RR04DJ -> [1 176 99 41 177 128 17 178]
// 	tsid := FromString("03C3356RR04DJ")
// 	t.Log(tsid.ToLong())

// 	bytes := []byte{1, 176, 99, 41, 177, 128, 17, 178}
// 	tsid = FromBytes(bytes)
// 	t.Log(tsid.ToLong())

// }
