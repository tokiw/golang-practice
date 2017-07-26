go build clock.go
go build clockwall.go
./clock -port 8001 & ./clockwall localhost:8001 localhost:8001