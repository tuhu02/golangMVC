# Golang MVC - Product & Category Management System

Aplikasi web sederhana untuk manajemen produk dan kategori yang dibangun menggunakan arsitektur MVC (Model-View-Controller) dengan Go sebagai backend dan TailwindCSS untuk styling.

## 🚀 Fitur

- **Dashboard** - Halaman utama aplikasi
- **Manajemen Kategori** - CRUD (Create, Read, Update, Delete) kategori produk
- **Manajemen Produk** - CRUD produk dengan relasi ke kategori
- **Responsive Design** - UI yang responsif menggunakan TailwindCSS
- **Database Integration** - Integrasi dengan MySQL database

## 🛠️ Teknologi yang Digunakan

### Backend

- **Go 1.25.1** - Bahasa pemrograman utama
- **MySQL** - Database untuk menyimpan data
- **go-sql-driver/mysql** - Driver MySQL untuk Go
- **html/template** - Template engine bawaan Go

### Frontend

- **TailwindCSS 3.4.10** - Framework CSS untuk styling
- **@tailwindplus/elements** - Komponen UI tambahan
- **Vite 5.0.0** - Build tool untuk frontend assets
- **PostCSS & Autoprefixer** - CSS processing

## 📁 Struktur Project

```
belajar/golangMVC/
├── config/
│   └── database.go          # Konfigurasi koneksi database
├── controllers/
│   ├── categorycontroller/  # Controller untuk kategori
│   ├── homecontroller/      # Controller untuk halaman utama
│   └── productcontroller/   # Controller untuk produk
├── entities/
│   ├── category.go          # Struct Category
│   └── product.go           # Struct Product
├── models/
│   ├── categorymodel/       # Model untuk operasi database kategori
│   └── productmodel/        # Model untuk operasi database produk
├── views/
│   ├── category/            # Template HTML untuk kategori
│   ├── home/                # Template HTML untuk home
│   ├── Product/             # Template HTML untuk produk
│   ├── css/                 # Source CSS files
│   └── dist/                # Compiled CSS & JS files
├── main.go                  # Entry point aplikasi
├── go.mod                   # Go module dependencies
└── package.json             # Node.js dependencies untuk build tools
```

## 🏗️ Arsitektur MVC

### Model

- **entities/** - Definisi struktur data (Product, Category)
- **models/** - Logic untuk operasi database (CRUD operations)

### View

- **views/** - Template HTML untuk tampilan user interface
- Menggunakan Go's `html/template` untuk rendering data

### Controller

- **controllers/** - Logic untuk menangani HTTP requests dan responses
- Menghubungkan Model dan View

## 📋 Prerequisites

Pastikan Anda sudah menginstall:

- **Go 1.25.1 atau lebih baru**
- **MySQL Server**
- **Node.js & npm** (untuk build frontend assets)

## ⚙️ Instalasi & Setup

### 1. Clone Repository

```bash
git clone <repository-url>
cd belajar/golangMVC
```

### 2. Setup Database

Buat database MySQL dan tabel yang diperlukan:

```sql
CREATE DATABASE golang_mvc;
USE golang_mvc;

-- Tabel categories
CREATE TABLE categories (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Tabel products
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INT NOT NULL,
    stock BIGINT NOT NULL DEFAULT 0,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);
```

### 3. Konfigurasi Database

Edit file `config/database.go` sesuai dengan konfigurasi MySQL Anda:

```go
// Sesuaikan dengan konfigurasi database Anda
db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/golang_mvc")
```

### 4. Install Dependencies

**Go Dependencies:**

```bash
go mod tidy
```

**Node.js Dependencies:**

```bash
npm install
```

### 5. Build Frontend Assets

```bash
npm run build
```

Atau untuk development dengan auto-reload:

```bash
npm run dev
```

### 6. Jalankan Aplikasi

```bash
go run main.go
```

Aplikasi akan berjalan di `http://localhost:8000`

## 🌐 Endpoints

### Home

- `GET /` - Halaman dashboard utama

### Categories

- `GET /categories` - Daftar semua kategori
- `GET /categories/add` - Form tambah kategori
- `POST /categories/add` - Proses tambah kategori
- `GET /categories/update?id={id}` - Form edit kategori
- `POST /categories/update` - Proses update kategori
- `GET /categories/delete?id={id}` - Hapus kategori

### Products

- `GET /products` - Daftar semua produk
- `GET /products/detail?id={id}` - Detail produk
- `GET /products/add` - Form tambah produk
- `POST /products/add` - Proses tambah produk
- `GET /products/update?id={id}` - Form edit produk
- `POST /products/update` - Proses update produk
- `GET /products/delete?id={id}` - Hapus produk

## 🎨 Styling

Project ini menggunakan TailwindCSS untuk styling dengan konfigurasi:

- **Responsive design** - Mobile-first approach
- **Dark theme navigation** - Navigation bar dengan tema gelap
- **Form styling** - Form yang konsisten dan user-friendly
- **Table styling** - Tabel data yang rapi dan mudah dibaca

## 📝 Development Scripts

```bash
# Build CSS untuk production
npm run build-css

# Watch CSS changes untuk development
npm run watch

# Build semua assets
npm run build

# Development mode dengan auto-reload
npm run dev
```

## 🔧 Troubleshooting

### Template Parsing Error

Jika mengalami error template parsing, pastikan:

- Template variables berada di konteks yang benar
- Tidak ada karakter khusus yang tidak di-escape
- Struktur HTML valid

### Database Connection Error

Pastikan:

- MySQL server berjalan
- Kredensial database benar
- Database dan tabel sudah dibuat

### CSS Not Loading

Pastikan:

- Jalankan `npm run build` untuk compile CSS
- File CSS ada di `views/dist/main.css`
- Path CSS di template benar

## 🤝 Contributing

1. Fork repository
2. Buat feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push ke branch (`git push origin feature/AmazingFeature`)
5. Buat Pull Request

## 📄 License

Project ini dibuat untuk tujuan pembelajaran dan pengembangan skill dalam Go web development dengan arsitektur MVC.

---

**Happy Coding! 🚀**
