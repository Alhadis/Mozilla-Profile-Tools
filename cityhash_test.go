package cityhash

import (
	"reflect"
	"testing"
)

func TestCityHash64(t *testing.T) {
	type args struct {
		s      []byte
		length uint32
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "Hello world",
			args: args{
				s:      []byte("Hello world"),
				length: uint32(len([]byte("Hello world"))),
			},
			want: 13182503335733000132,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CityHash64(tt.args.s, tt.args.length); got != tt.want {
				t.Errorf("CityHash64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCityHash64WithSeed(t *testing.T) {
	type args struct {
		s      []byte
		length uint32
		seed   uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "Hello world",
			args: args{
				s:      []byte("Hello world"),
				length: uint32(len([]byte("Hello world"))),
				seed:   8, //sizeof(uint64)
			},
			want: 4936126472499377365,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CityHash64WithSeed(tt.args.s, tt.args.length, tt.args.seed); got != tt.want {
				t.Errorf("CityHash64WithSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCityHash64WithSeeds(t *testing.T) {
	type args struct {
		s      []byte
		length uint32
		seed0  uint64
		seed1  uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "Hello world",
			args: args{
				s:      []byte("Hello world"),
				length: uint32(len([]byte("Hello world"))),
				seed0:  1, //sizeof(uint8)
				seed1:  8, //sizeof(uint64)
			},
			want: 10995333839364115670,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CityHash64WithSeeds(tt.args.s, tt.args.length, tt.args.seed0, tt.args.seed1); got != tt.want {
				t.Errorf("CityHash64WithSeeds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCityHash128(t *testing.T) {
	type args struct {
		s      []byte
		length uint32
	}
	tests := []struct {
		name string
		args args
		want Uint128
	}{
		{
			name: "Hello world",
			args: args{
				s:      []byte("Hello world"),
				length: uint32(len([]byte("Hello world"))),
			},
			want: Uint128{14676461730635999210, 18211519499011055500},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CityHash128(tt.args.s, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CityHash128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCityHash128WithSeed(t *testing.T) {
	type args struct {
		s      []byte
		length uint32
		seed   Uint128
	}
	tests := []struct {
		name string
		args args
		want Uint128
	}{
		{
			name: "Hello world",
			args: args{
				s:      []byte("Hello world"),
				length: uint32(len([]byte("Hello world"))),
				seed:   Uint128{1, 8}, // sizeof(uint8) & sizeof (uint64)
			},
			want: Uint128{15656991209745481333, 7719688212359498183},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CityHash128WithSeed(tt.args.s, tt.args.length, tt.args.seed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CityHash128WithSeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
