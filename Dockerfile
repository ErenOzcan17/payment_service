# 1. Go'nun resmi lightweight bir versiyonunu kullan
FROM golang:1.23-alpine

# 2. Çalışma dizinini ayarla
WORKDIR /app

# 3. Bağımlılıkları ve kaynak kodu kopyala
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 4. Uygulamayı derle
RUN go build -o payment_go

# 5. Container çalıştırıldığında başlatılacak komut
CMD ["./payment_go"]
