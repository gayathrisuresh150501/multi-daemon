module rundaemons

go 1.20

replace daemon => ./daemon

require daemon v0.0.0-00010101000000-000000000000

require github.com/stretchr/testify v1.8.4 // indirect
