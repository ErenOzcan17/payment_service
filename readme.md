## komutları çalıştır:
    docker build -t payment_go .
    docker run --name payment_go -p 50051:50051 payment_go

## container oluşturduysan:
    docker run payment_go