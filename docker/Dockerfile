FROM golang:1.22.5-alpine AS base

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY config/ config/
COPY internal/ internal/
COPY pkg/ pkg/
COPY migrations/ migrations/


FROM base as api
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api/*.go

FROM base as migrate
RUN CGO_ENABLED=0 GOOS=linux go build -o migrate ./cmd/migrate/*.go


FROM gcr.io/distroless/static-debian12:latest AS final

ENV TZ=Asia/Jakarta

COPY --from=api /build/api /api
COPY --from=migrate /build/migrate /migrate
COPY --from=base /build/migrations/*.sql /data/migrations/