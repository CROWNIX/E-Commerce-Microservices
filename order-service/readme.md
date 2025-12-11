# 🚀 Go Service Boilerplate

Project ini adalah boilerplate service berbasis **Golang** dengan struktur modular, menggunakan **OpenAPI/Protobuf** sebagai kontrak, **Wire** untuk Dependency Injection, serta **Viper** untuk konfigurasi.  
Struktur ini ditujukan agar codebase lebih terorganisir, scalable, dan maintainable.

---

## 📂 Project Structure

```Bash
├── api/              # Tempat file OpenAPI/Protobuf (contract dengan SDK luar/provider)
├── cmd/              # Entry point Go apps (CLI/API). Menggunakan Viper & Wire (Dependency Injection)
├── internal/         # Detail implementasi aplikasi
│   ├── config/       # Konfigurasi aplikasi
│   ├── dto/          # Data Transfer Objects
│   ├── infra/        # Infrastruktur (client eksternal, dsb.)
│   ├── models/       # Model untuk repositories
│   ├── repositories/ # Layer akses data
│   │   ├── datastore/ # Database/Redis (per table folder: users/, categories/, dll.)
│   │   └── s3/        # Object storage (image, video, dsb.)
│   ├── services/     # Business logic (per domain folder: auth/, order/, menu/, dll.)
│   └── presentations/ # Layer paling luar (API)
│       ├── handler/   # Implementasi interfaces hasil OpenAPI codegen
│       │   └── cms/   # Handler untuk CMS admin panel (create product, notif order, dll.)
│       └── middleware/ # Custom middleware (auth, logging, dll.)
└── scripts/          # Script helper (install, migrations, dll.)
```

---
---

## 📌 Rules & Conventions

### 🗄️ Repositories

- Buat **1 folder per table** di `repositories/datastore`.  
  Contoh:
  - `repositories/datastore/users/`
  - `repositories/datastore/categories/`
- **Wajib** ada file `interfaces.go` di setiap folder table yang berisi interface **Reader** dan **Writer**.
- **Wajib** membuat file `entity.go` yang isinya entity dari table tersebut
- **Naming convention interface**:
  - `<Entity>RepositoryReaderInterfaces`
  - `<Entity>RepositoryWriterInterfaces`

**Contoh `repositories/datastore/users/interfaces.go`:**

```go
package users

type UserRepositoryReaderInterfaces interface {
    // contoh: GetByID(ctx context.Context, id int64) (*User, error)
    // contoh: FindByEmail(ctx context.Context, email string) (*User, error)
}

type UserRepositoryWriterInterfaces interface {
    // contoh: Create(ctx context.Context, u *User) error
    // contoh: Update(ctx context.Context, u *User) error
    // contoh: DeleteByID(ctx context.Context, id int64) error
}
```

### Services

- Buat **1 folder per domain bisnis logic** di `services`.  
Contoh:
  - services/auth/
  - services/order/
  - services/menu/

Wajib ada file interfaces.go berisi interface service untuk domain tersebut.

Naming convention interface:
 • `<Domain>ServiceInterfaces`

untuk parameters di tiap" interface itu harus hanya memiliki 2, params pertama ctx params kedua itu diikut dengan `<NamaMethod>Input`
untuk response itu wajib <= 2, cmn ada `<NamaMethod>Output` dan error
Contoh services/auth/interfaces.go:

```go
package auth

type AuthServiceInterfaces interface {
    // contoh: Login(ctx context.Context, req LoginInput) (resp LoginOutput, err error)
    // contoh: RefreshToken(ctx context.Context, req RefreshTokenInput) (resp RefreshTokenOutput, err error)
}
```

### Presentations

- Semua komunikasi dengan client masuk lewat layer ini.
- **handler/**: implementasi hasil codegen dari OpenAPI.
- **handler/cms/**: khusus untuk CMS Admin Panel (contoh: create product, notif order).
- **middleware/**: custom middleware (contoh: auth, logging).
untuk naming file itu per implementations method, contoh.
`apiV1GetProduct`
`apiV1DeleteProduct`
`apiV1PostOrder`
`apiV1PutOrder`

## 📑 General Conventions

 1. Semua interface disimpan di file interfaces.go pada folder domain/entitas masing-masing.
 2. Suffix Interfaces (jamak) digunakan untuk konsistensi penamaan:
UserRepositoryReaderInterfaces, UserRepositoryWriterInterfaces, AuthServiceInterfaces, dll.
 3. Satu folder = satu domain/table, jangan campur interface lintas domain dalam satu folder.
 4. Implementation dipisah per backend (mis. pg_repository.go, redis_repository.go, s3_repository.go) sesuai lokasi (datastore/, s3/, dll).
 5. Models: untuk request/response atau payload lintas layer (hindari bocor domain).

---

## ⚙️ Scripts

### Install dependency

pastikan kamu mempunyai pip dan minimal go 1.24.0

```bash
./scripts/install_dependency.sh
```

### Create File Migration

```bash
./scripts/create_migration.sh create_user_addresses_table
```

### Apply migration

```bash
./scripts/apply_migration.sh -d 'postgres://user_service_user:user_service_pass@103.31.132.7:1400/user_service_db?sslmode=disable'
```

Ganti -d dengan connection string database milikmu.

## 🛠️ Makefile Commands

- generate_api → Generate API contract dari api/openapi/api.yaml.
- generate_wire → Generate dependency injection (wire).
- clean → Membersihkan file hasil generate/build.
- preview_open_api → Preview OpenAPI contract menggunakan redocly.
- run-api → Menjalankan API.

## 🚀 Quick Start

1. Install dependency:

    ```bash
    ./scripts/install_dependency.sh
    ```

2. Generate API contract:

   ```bash
   make generate_api
   ```

3. Generate DI (wire):

   ```bash
   make generate_wire
   ```

4. jalankan api:

   ```bash
   make run-api
   ```

5. preview open api:

   ```bash
   make preview_open_api
   ```

## 🧰 Tech Stack

- Golang 1.24.0
- Gin
- PostgreSQL
- Redis
- S3 / Object Storage
- OpenAPI / Protobuf (API contract)
- Wire (Dependency Injection)
- Viper (Configuration)
- Redocly (OpenAPI preview)
