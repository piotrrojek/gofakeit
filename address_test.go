package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleAddress() {
	Seed(11)
	fmt.Println(Address())
	// Output: 872 East Rapids borough, Andrestad, New Jersey 74853-6757
}

func BenchmarkAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Address()
	}
}

func ExampleStreet() {
	Seed(11)
	fmt.Println(Street())
	// Output: 872 East Rapids borough
}

func BenchmarkStreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Street()
	}
}

func ExampleStreetNumber() {
	Seed(11)
	fmt.Println(StreetNumber())
	// Output: 28727
}

func BenchmarkStreetNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetNumber()
	}
}

func ExampleStreetPrefix() {
	Seed(11)
	fmt.Println(StreetPrefix())
	// Output: Lake
}

func BenchmarkStreetPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetPrefix()
	}
}

func ExampleStreetName() {
	Seed(11)
	fmt.Println(StreetName())
	// Output: View
}

func BenchmarkStreetName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetName()
	}
}

func ExampleStreetSuffix() {
	Seed(11)
	fmt.Println(StreetSuffix())
	// Output: land
}

func BenchmarkStreetSuffix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StreetSuffix()
	}
}

func ExampleCity() {
	Seed(11)
	fmt.Println(City())
	// Output: Marcelside
}

func BenchmarkCity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		City()
	}
}

func ExampleState() {
	Seed(11)
	fmt.Println(State())
	// Output: Hawaii
}

func BenchmarkState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		State()
	}
}

func ExampleStateAbr() {
	Seed(11)
	fmt.Println(StateAbr())
	// Output: OR
}

func BenchmarkStateAbr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StateAbr()
	}
}

func ExampleZip() {
	Seed(11)
	fmt.Println(Zip())
	// Output: 28727
}

func BenchmarkZip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Zip()
	}
}

func ExampleCountry() {
	Seed(11)
	fmt.Println(Country())
	// Output: Tajikistan
}

func BenchmarkCountry(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Country()
	}
}

func ExampleLatitude() {
	Seed(11)
	fmt.Println(Latitude())
	// Output: -73.53405629980608
}

func BenchmarkLatitude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Latitude()
	}
}

func ExampleLongitude() {
	Seed(11)
	fmt.Println(Longitude())
	// Output: -147.06811259961216
}

func BenchmarkLongitude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Longitude()
	}
}
