package cal

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundUpToNearest(t *testing.T) {
	type args struct {
		value float64
		unit  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{
				value: 1.8,
				unit:  0.5,
			},
			want: 2.0,
		},
		{
			name: "",
			args: args{
				value: 1.7,
				unit:  0.5,
			},
			want: 2.0,
		},
		{
			name: "",
			args: args{
				value: 0.6,
				unit:  0.5,
			},
			want: 1.0,
		},
		{
			name: "",
			args: args{
				value: 0.4,
				unit:  0.5,
			},
			want: 0.5,
		},
		{
			name: "",
			args: args{
				value: 0.5,
				unit:  0.5,
			},
			want: 0.5,
		},
		{
			name: "",
			args: args{
				value: 1.0,
				unit:  0.5,
			},
			want: 1.0,
		},
		{
			name: "",
			args: args{
				value: 3.1,
				unit:  0.5,
			},
			want: 3.5,
		},
		{
			name: "",
			args: args{
				value: 4.4,
				unit:  1.0,
			},
			want: 5.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundUpToNearest(tt.args.value, tt.args.unit); got != tt.want {
				t.Errorf("RoundUpToNearest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoundDownToNearest(t *testing.T) {
	type args struct {
		value float64
		unit  float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{
				value: 1.3,
				unit:  0.5,
			},
			want: 1.0,
		},
		{
			name: "",
			args: args{
				value: 1.51695123216021,
				unit:  0.5,
			},
			want: 1.5,
		},
		{
			name: "",
			args: args{
				value: 1.9,
				unit:  0.5,
			},
			want: 1.5,
		},
		{
			name: "",
			args: args{
				value: 3.13901101,
				unit:  0.5,
			},
			want: 3.0,
		},
		{
			name: "",
			args: args{
				value: 0.663,
				unit:  0.5,
			},
			want: 0.5,
		},
		{
			name: "",
			args: args{
				value: 0.1003,
				unit:  0.5,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoundDownToNearest(tt.args.value, tt.args.unit); got != tt.want {
				t.Errorf("RoundDownToNearest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PourWater(t *testing.T) {
	var cups = []*Cup{
		{
			ID:       5,
			Priority: "5",
			MaxVol:   5,
			Vol:      0,
		},
		{
			ID:       3,
			Priority: "3",
			MaxVol:   3,
			Vol:      0,
		},
		{
			ID:       2,
			Priority: "2",
			MaxVol:   2,
			Vol:      0,
		},
	}

	type args struct {
		req int32
		cup *Cup
	}
	tests := []struct {
		name string
		args args
		want map[uint32]int32
	}{
		{
			name: "",
			args: args{
				req: 6,
				cup: nil,
			},
			want: map[uint32]int32{
				2: 2,
				3: 3,
				5: 1,
			},
		},
		{
			name: "",
			args: args{
				req: 0,
				cup: &Cup{
					ID:       6,
					Priority: "6",
					MaxVol:   6,
					Vol:      0,
				},
			},
			want: nil,
		},
		{
			name: "",
			args: args{
				req: 5,
				cup: nil,
			},
			want: map[uint32]int32{
				5: 5,
				6: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.cup != nil {
				cups = append(cups, tt.args.cup)
			}

			if got := PourWater(tt.args.req, cups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PourWater() = %v, want %v", got, tt.want)
			}
			bts, _ := json.MarshalIndent(cups, "", "\t")
			t.Logf("cups=%s", string(bts))
		})
	}
}

func Test_AddCup_PourWater(t *testing.T) {
	var cups = []*Cup{
		{
			ID:       1,
			Priority: "1",
			MaxVol:   5,
			Vol:      5,
		},
		{
			ID:       3,
			Priority: "3",
			MaxVol:   3,
			Vol:      1,
		},
		{
			ID:       5,
			Priority: "5",
			MaxVol:   2,
			Vol:      0,
		},
	}
	type args struct {
		req int32
		cup *Cup
	}
	tests := []struct {
		name string
		args args
		want map[uint32]int32
	}{
		{
			name: "",
			args: args{
				req: 0,
				cup: &Cup{
					ID:       2,
					Priority: "2",
					MaxVol:   -2,
					Vol:      0,
				},
			},
			want: map[uint32]int32{
				3: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cups = append(cups, tt.args.cup)
			if got := PourWater(tt.args.req, cups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PourWater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	assert.Equal(t, int32(1), AbsGeneric[int32](int32(-1)))
	assert.Equal(t, int8(1), AbsGeneric[int8](int8(-1)))
	assert.Equal(t, int64(1), AbsGeneric[int64](int64(-1)))
	assert.Equal(t, float32(1), AbsGeneric[float32](float32(-1)))
	assert.Equal(t, float64(1), AbsGeneric[float64](float64(-1)))
}
