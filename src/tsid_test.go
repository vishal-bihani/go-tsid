package main

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	LOOP_MAX = 1000
)

func Test_GetUnixMillis(t *testing.T) {

	t.Run("should return correct time", func(t *testing.T) {
		start := time.Now().UnixMilli()

		tsidFactory, _ := TsidFactoryBuilder().
			NewInstance()
		assert.NotNil(t, tsidFactory)

		tsid, _ := tsidFactory.Generate()
		assert.NotNil(t, tsid)

		middle := tsid.GetUnixMillis()
		end := time.Now().UnixMilli()

		if middle < start || (middle > end) {
			t.Fail()
		}
	})

	t.Run("given custom time should return correct time", func(t *testing.T) {
		bound := math.Pow(2, 42)

		for i := 0; i < LOOP_MAX; i++ {

			// generate random value
			random := rand.New(rand.NewSource(time.Now().UnixNano())).
				Int63n(int64(bound))

			// ensuring date is generated after TSID_EPOCH
			millis := random + TSID_EPOCH
			time := time.UnixMilli(millis)

			// int random supplier func
			intRandomSupplierFunc := func() (int32, error) {
				return 0, nil
			}

			intRandom := NewIntRandomWithSupplierFunc(intRandomSupplierFunc)

			tsidFactory, _ := TsidFactoryBuilder().
				WithClock(time).
				WithRandom(intRandom).
				NewInstance()
			assert.NotNil(t, tsidFactory)

			tsid, _ := tsidFactory.Generate()
			assert.NotNil(t, tsid)

			result := tsid.GetUnixMillis()
			assert.Equal(t, millis, result)
		}
	})

	t.Run("given custom epoch should return correct time", func(t *testing.T) {

		epoch := time.Date(1984, time.January, 1, 0, 0, 0, 0, time.UTC).
			UnixMilli()

		start := time.Now().UnixMilli()

		tsidFactory, _ := TsidFactoryBuilder().
			WithCustomEpoch(epoch).
			NewInstance()
		assert.NotNil(t, tsidFactory)

		tsid, _ := tsidFactory.Generate()
		assert.NotNil(t, tsid)

		middle := tsid.GetUnixMillisWithCustomEpoch(epoch)
		end := time.Now().UnixMilli()

		if middle < start || (middle > end) {
			t.Fail()
		}
	})
}
