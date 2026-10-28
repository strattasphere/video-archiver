air: # default - runs with `just`
  air -c air.toml

run:
  cd cmd && go run main.go

build:
  @go test -v ./... 

test:
  @go test -v ./... 

install-tailwind:
  @curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
  @chmod 744 tailwindcss-linux-x64
  @mv tailwindcss-linux-x64 bin/

install-daisy-css:
  @curl -sL -o internal/dist/assets/css/daisyui.css https://cdn.jsdelivr.net/npm/daisyui@5/daisyui.css
  @curl -sL -o internal/dist/assets/css/daisy-themes.css https://cdn.jsdelivr.net/npm/daisyui@5/themes.css

install-htmx:
  @curl

tw:
  @./bin/tailwindcss-linux-x64
