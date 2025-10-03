# Simple Kafka Producer & Consumer go

## Studi Kasus

Terdapat 2 service dalam project ini

1. [`AuthService`](/auth-service) =>
   - Implementasi kafka producer
   - Menggunakan Gin sebagai rest api sederhana
   - 2 end point : forgot-password dan send-otp
   - setiap end point nantinya mengirim message ke masing masing topics
2. [`EmailService`](/email-service) =>
   - Implementasi kafka consumer
   - Hanya menerima message dari setiap topics dan memprint hasil data yang dikirim dengan format sesuai topics

Topics yang digunakan

```go
topics := []string{"forgot-password", "otp"}
```

## Running the Project Locally

Buat file `.env` dengan menyalin file `sampleenv`disetiap service, isi variable yang tersedia sesuai dengan data yang diperlukan

1. [`Env AuthService`](/auth-service/sampleenv)
2. [`Env EmailService`](/email-service/sampleenv)

### 1. Optional run kafka development server (docker)

Skip step ini jika sudah memiliki kafka server yang sudah berjalan

Sesuaikan user dan password di [`kafka_server_jaas.conf`](/messaging-server/kafka_server_jaas.conf)<br>
User dan password ini nantinya digunakan di variable env

```bash
cd messaging-server
docker compose up -d
# pastikan semua container sudah berjalan
```

- kafka akan berjalan di 127.0.0.1:9093
- Buat topic `forgot-password` dan `otp` di [`http://localhost:9000/topic/create`](http://localhost:9000/topic/create)

### 2. Run service

```bash
# install dependencies
go mod tidy

# optional (arahkan ke mingw64)
set PATH=C:\msys64\mingw64\bin;%PATH%

# run AuthService
cd auth-service
go run .

# run EmailService
cd email-service
go run .

```

### 3. Test AuthService / Producer

Akses [`http://localhost:8080/swagger/index.html`](http://localhost:8080/swagger/index.html) . Testing AuthService / producer bisa dilakukan di halaman tersebut

### 3. Test EmailService / Consumer

Service akan memprint setiap message yang diterima sesuai dengan topic nya
