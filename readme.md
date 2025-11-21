# E-Commerce Microservices

Repository ini merupakan proyek latihan untuk membangun arsitektur **E-Commerce berbasis microservices** menggunakan berbagai komponen modern seperti Go, Redis, MySQL, Kafka, Kong API Gateway, dan gRPC.

Tujuan utama proyek ini adalah mempelajari bagaimana membangun sistem terdistribusi yang scalable, loosely-coupled, dan mudah dikembangkan.

---

## 🎯 Tujuan Proyek

Proyek ini dibuat sebagai bahan pembelajaran untuk memahami konsep dan implementasi:

* Microservices Architecture
* Inter-service communication (gRPC, Kafka)
* API Gateway (Kong)
* Distributed system patterns
* Authentication & Authorization
* Scalability & Maintainability

---

## 🏗️ Arsitektur Singkat

Saat ini, proyek menggunakan pendekatan **shared database**, tetapi ke depannya setiap service akan memiliki **database masing-masing (1 service 1 database)** untuk mengikuti prinsip microservice yang ideal.

Komponen utama dalam arsitektur:

* **Auth Service** – menangani login, token, dan authorization.
* **Product Service** – mengelola produk.
* **User Service** – mengelola data user.
* **Kafka** – digunakan sebagai message broker untuk event-driven flow.
* **Kong API Gateway** – sebagai pintu masuk ke seluruh service.
* **Redis** – sebagai cache dan penyimpanan token.
* **gRPC** – komunikasi antar service.

---

## 🧰 Tech Stack

Berikut adalah teknologi yang digunakan sejauh ini:

### **Backend:**

* **Golang** (Go) sebagai bahasa utama seluruh service
* **gRPC** untuk komunikasi antar microservices
* **Kafka** sebagai message broker untuk pub/sub, event sourcing, dan async processing

### **Database & Storage:**

* **MySQL** (Saat ini masih shared DB)
* **Redis** untuk caching dan session/token store
* **(Rencana)** Elasticsearch untuk teks pencarian cepat
* **(Rencana)** MongoDB untuk data yang membutuhkan schema fleksibel

### **API Gateway:**

* **Kong Gateway** untuk routing, rate limit, auth, dan traffic control

### **Infrastructure:**

* Docker & Docker Compose
* Makefile automation
* OpenAPI Spec
* Wire Dependency Injection

---

## 🚧 Status Proyek

Saat ini proyek masih dalam tahap awal:

* [x] Setup auth service
* [x] Kafka skeleton
* [x] API Gateway (sebagian)
* [ ] Product service development
* [ ] User service development
* [ ] Migrasi ke "one service one database"
* [ ] Implementasi Elasticsearch
* [ ] Implementasi MongoDB

---

## 📦 Struktur Repository (Ringkas)

```
E-Commerce-Microservices/
  ├── auth-service/
  ├── product-service/
  ├── user-service/
  ├── api-gateway/
  ├── docker-compose.yml
  ├── Makefile
  └── README.md
```

---

## 🚀 Cara Menjalankan (akan ditambahkan)

Instruksi lengkap untuk menjalankan seluruh service akan ditambahkan setelah setiap service stabil.

---

## 📚 Catatan Pengembangan

Proyek ini akan terus dikembangkan seiring pembelajaran dan eksperimen terhadap berbagai pola microservices.

Jika Anda ingin ikut berkontribusi atau memberikan masukan, sangat dipersilakan.

---

## 📄 Lisensi

Proyek ini dibuat untuk pembelajaran dan eksperimen. Tidak ada lisensi formal saat ini.
