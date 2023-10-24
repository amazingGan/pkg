package cal

import (
	"sort"

	"github.com/shopspring/decimal"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 泛型支持
func AbsGeneric[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func RoundUpToNearest(value, unit float64) float64 {
	if unit == 0 {
		return value
	}
	dVaule := decimal.NewFromFloat(value)
	dUnit := decimal.NewFromFloat(unit)

	return dVaule.Div(dUnit).Ceil().Mul(dUnit).InexactFloat64()
}

func RoundDownToNearest(value, unit float64) float64 {
	if unit == 0 {
		return value
	}
	dVaule := decimal.NewFromFloat(value)
	dUnit := decimal.NewFromFloat(unit)

	return dVaule.Div(dUnit).Floor().Mul(dUnit).InexactFloat64()
}

type Cup struct {
	ID       uint32 // 主键
	Priority string // 优先级
	MaxVol   int32  // 容量
	Vol      int32  // 已使用量
}

func PourWater(water int32, cups []*Cup) map[uint32]int32 {
	if water == 0 || len(cups) == 0 {
		return nil
	}
	result := make(map[uint32]int32)

	sort.SliceStable(cups, func(i, j int) bool {
		return cups[i].Priority < cups[j].Priority
	})

	for _, cup := range cups {
		if water <= 0 {
			break
		}

		// 计算每个杯子可以装多少水
		vol := cup.MaxVol - cup.Vol
		if vol <= 0 { // 说明当前杯子已经是装满的，跳过这个杯子
			continue
		}
		if vol > water {
			vol = water
		}

		cup.Vol += vol
		water -= vol

		result[cup.ID] = cup.Vol
	}

	return result
}

// 手动加入一个杯子，杯子容量可以是负数
func PourWaterAddCup(cups []*Cup, cup *Cup) map[uint32]int32 {
	// TODO
	return nil
}
