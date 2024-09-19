# generate tailwind css files
FROM node:22-alpine3.20 AS tailwindbuilder
LABEL stage=tailwindbuilder
WORKDIR /build
COPY package.json package-lock.json tailwind.config.js ./
RUN npm install
COPY controllers/web/components/*.templ ./controllers/web/components/
COPY controllers/web/static/index.css ./controllers/web/static/
RUN npx tailwindcss -i ./controllers/web/static/index.css -o ./static/output.css

# generate _templ.go files from .templ files
FROM golang:alpine AS templbuilder
LABEL stage=templbuilder
WORKDIR /build
RUN go install github.com/a-h/templ/cmd/templ@v0.2.778
COPY controllers/web/components/*.templ ./controllers/web/components/
RUN templ generate

# build standalone golang server binary
FROM golang:alpine AS gobuilder
LABEL stage=gobuilder
ENV CGO_ENABLED 0
WORKDIR /build
RUN apk update --no-cache && apk add --no-cache tzdata
ADD go.mod .
ADD go.sum .
RUN go mod download
COPY cmd ./cmd
COPY controllers ./controllers
COPY --from=templbuilder /build/controllers/web/components/*_templ.go ./controllers/web/components/
RUN go build -ldflags="-s -w" -o /app/unotone ./cmd/server/main.go
RUN chmod +x /app/unotone

# runner image
FROM scratch
WORKDIR /app
ENV UNOTONE_STATIC_DIR /var/www/static
COPY --from=gobuilder /build/controllers/web/static ${UNOTONE_STATIC_DIR}
COPY --from=tailwindbuilder /build/static/output.css ${UNOTONE_STATIC_DIR}/output.css
COPY --from=gobuilder /app/unotone /app/unotone
CMD ["./unotone"]