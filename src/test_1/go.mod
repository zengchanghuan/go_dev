module hello.go

go 1.18

require (
	github.com/myuser/calculator v0.0.0
	rsc.io/quote v1.5.2

)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace github.com/myuser/calculator => ../calculator
