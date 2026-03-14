go install github.com/arch-go/arch-go@latest
arch-go --color no
go test -v ./test/architecture/... -run TestArchitecture -count=1
