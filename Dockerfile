FROM golang:1.22.5-alpine

WORKDIR /app

# Menyalin go.mod dan go.sum untuk dependency
COPY go.mod go.sum ./

# Mengunduh dependensi Go
RUN go mod download
RUN go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go mod tidy 

# Menetapkan environment variables
ENV APP_ENV="dev"
ENV PORT="6543"
ENV DB_HOST="aws-0-ap-southeast-1.pooler.supabase.com"
ENV DB_USER="postgres.oczurxscwnpcgcyzgdtt"
ENV PASS="Wanjaywkwkwk12!"
ENV DB_NAME="postgres"

# Menyalin seluruh kode sumber ke dalam container
COPY . .

# Mengkompilasi kode Go
RUN go clean
RUN go build -o main ./cmd

# Mengekspos port yang digunakan oleh aplikasi
EXPOSE 6543

# Menjalankan aplikasi
CMD ["./main"]
