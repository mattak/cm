module github.com/mattak/cm

go 1.14

require (
	cloud.google.com/go/firestore v1.5.0 // indirect
	github.com/mattak/cm/cmd v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.1.3
)

replace github.com/mattak/cm/cmd => ./cmd
