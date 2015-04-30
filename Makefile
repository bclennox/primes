primes: primes.go
	go build -o build/primes primes.go

clean:
	rm -f build/primes
