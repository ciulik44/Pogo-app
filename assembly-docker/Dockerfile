FROM golang:1.16-alpine

WORKDIR /app

# Kopiowanie plików go.mod i go.sum z katalogu nadrzędnego
COPY ../go.mod ../go.sum ./
RUN go mod download

# Kopiowanie reszty plików z katalogu nadrzędnego
COPY ../ ./

# Kompilacja aplikacji
RUN go build -o main .

# Ustawienie komendy startowej
CMD ["./main"]
