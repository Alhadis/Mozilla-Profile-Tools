package cityhash

import (
	"testing"
)

func TestCityHash64WithHelloWorld(t *testing.T) {
	// input for CityHash64()
	var (
		input     []byte = []byte("Hello world")
		inputSize uint32 = uint32(len(input))
	)

	// expected output hash value
	const (
		expectedOutput uint64 = 13182503335733000132
	)

	// output from CityHash64(input, inputSize)
	var actualOutput uint64 = CityHash64(input, inputSize)

	if actualOutput != expectedOutput {
		t.Errorf("hash does not match expected output")
	}
}

func TestCityHash64WithSeedHelloWorld(t *testing.T) {
	// input for CityHash64WithSeed()
	var (
		input     []byte = []byte("Hello world")
		inputSize uint32 = uint32(len(input))
		seed      uint64 = 8 //sizeof(uint64)
	)

	// expected output hash value
	const (
		expectedOutput uint64 = 4936126472499377365
	)

	// output from CityHash64WithSeed(input, inputSize, seed)
	var actualOutput uint64 = CityHash64WithSeed(input, inputSize, seed)

	if actualOutput != expectedOutput {
		t.Errorf("hash does not match expected output")
	}
}

func TestCityHash64WithSeedsHelloWorld(t *testing.T) {
	// input for CityHash64WithSeeds()
	var (
		input     []byte = []byte("Hello world")
		inputSize uint32 = uint32(len(input))
		seed_1    uint64 = 1 //sizeof(uint8)
		seed_2    uint64 = 8 //sizeof(uint64)
	)

	// expected output hash value
	const (
		expectedOutput uint64 = 10995333839364115670
	)

	// output from CityHash64WithSeeds(input, inputSize, seed_1, seed_2)
	var actualOutput uint64 = CityHash64WithSeeds(input, inputSize, seed_1, seed_2)

	if actualOutput != expectedOutput {
		t.Errorf("hash does not match expected output")
	}
}

func TestCityHash128WithHelloWorld(t *testing.T) {
	// input for CityHash128()
	var (
		input     []byte = []byte("Hello world")
		inputSize uint32 = uint32(len(input))
	)

	// expected output hash values for uint128 first/lower and second/higher
	const (
		expectedOutputFirst  uint64 = 14676461730635999210
		expectedOutputSecond uint64 = 18211519499011055500
	)

	// output from CityHash128(input, inputSize)
	var (
		actualOutput              = CityHash128(input, inputSize)
		actualOutputFirst  uint64 = actualOutput.Lower64()
		actualOutputSecond uint64 = actualOutput.Higher64()
	)

	if actualOutputFirst != expectedOutputFirst || actualOutputSecond != expectedOutputSecond {
		t.Errorf("hash does not match expected output")
	}
}

func TestCityHash128WithSeedWithHelloWorld(t *testing.T) {
	// input for CityHash128WithSeed()
	var (
		input     []byte = []byte("Hello world")
		inputSize uint32 = uint32(len(input))
		seed             = Uint128{1, 8} // sizeof(uint8) & sizeof (uint64)
	)

	// expected output hash values for uint128 first/lower and second/higher
	const (
		expectedOutputFirst  uint64 = 15656991209745481333
		expectedOutputSecond uint64 = 7719688212359498183
	)

	// output from CityHash128WithSeed(input, inputSize, seed)
	var (
		actualOutput              = CityHash128WithSeed(input, inputSize, seed)
		actualOutputFirst  uint64 = actualOutput.Lower64()
		actualOutputSecond uint64 = actualOutput.Higher64()
	)

	if actualOutputFirst != expectedOutputFirst || actualOutputSecond != expectedOutputSecond {
		t.Errorf("hash does not match expected output")
	}
}
