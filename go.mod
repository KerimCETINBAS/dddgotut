module github.com/kerimmcetinbas/partizor

go 1.21.1

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.50.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	go.uber.org/dig v1.17.1
	golang.org/x/sys v0.14.0 // indirect
)

replace github.com/kerimcetinbas/partizor.api => ./api
replace github.com/kerimcetinbas/partizor.infrastructure => ./infrastructure
require (
	github.com/gofiber/fiber/v2 v2.51.0
	github.com/kerimcetinbas/partizor.api v1.0.0
    github.com/kerimcetinbas/partizor.infrastructure v1.0.0
)
