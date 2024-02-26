package perlin_test

import (
	"fmt"
	"math/rand/v2"

	"github.com/FixtureMapper/go-perlin"
)

const (
	alpha        = 2.
	beta         = 2.
	n            = 3
	seed1 uint64 = 100
	seed2 uint64 = 200
)

func ExampleNewPerlinRandSource() {
	p := perlin.NewPerlinRandSource(alpha, beta, n, rand.New(rand.NewPCG(seed1, seed2)))
	for x := 0.; x < 3; x++ {
		fmt.Printf("%0.0f;%0.4f\n", x, p.Noise1D(x/10))
	}
	// Output:
	// 0;0.0000
	// 1;0.1079
	// 2;0.0888
}

func ExamplePerlin_Noise1D() {
	p := perlin.NewPerlin(alpha, beta, n, seed1, seed2)
	for x := 0.; x < 3; x++ {
		fmt.Printf("%0.0f;%0.4f\n", x, p.Noise1D(x/10))
	}
	// Output:
	// 0;0.0000
	// 1;0.1079
	// 2;0.0888
}

func ExamplePerlin_Noise2D() {
	p := perlin.NewPerlin(alpha, beta, n, seed1, seed2)
	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			fmt.Printf("%0.0f;%0.0f;%0.4f\n", x, y, p.Noise2D(x/10, y/10))
		}
	}
	// Output:
	// 0;0;0.0000
	// 0;1;-0.1477
	// 1;0;0.1509
	// 1;1;0.0154
}

func ExamplePerlin_Noise3D() {
	p := perlin.NewPerlin(alpha, beta, n, seed1, seed2)
	for x := 0.; x < 2; x++ {
		for y := 0.; y < 2; y++ {
			for z := 0.; z < 2; z++ {
				fmt.Printf("%0.0f;%0.0f;%0.0f;%0.4f\n", x, y, z, p.Noise3D(x/10, y/10, z/10))
			}
		}
	}
	// Output:
	// 0;0;0;0.0000
	// 0;0;1;0.2477
	// 0;1;0;-0.0149
	// 0;1;1;0.1465
	// 1;0;0;0.0562
	// 1;0;1;0.1828
	// 1;1;0;0.0278
	// 1;1;1;0.0942
}
