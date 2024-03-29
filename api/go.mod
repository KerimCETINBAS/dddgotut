module github.com/kerimcetinbas/partizor.api

go 1.21.1

require github.com/gofiber/fiber/v2 v2.51.0

replace github.com/kerimcetinbas/partizor.contracts => ../contracts
replace github.com/kerimcetinbas/partizor.infrastructure => ../infrastructure
replace github.com/kerimcetinbas/partizor.application => ../application


require github.com/kerimcetinbas/partizor.contracts v1.0.0
require github.com/kerimcetinbas/partizor.infrastructure v1.0.0
require github.com/kerimcetinbas/partizor.application v1.0.0

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.16.0 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/klauspost/compress v1.16.7 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.50.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.8.0 // indirect
)
