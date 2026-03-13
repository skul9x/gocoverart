# Music Cover Tagger

Ứng dụng desktop để tìm và áp dụng bìa album cho tập hợp nhạc của bạn.

## 🎯 Tính Năng

- **Tìm ảnh bìa cao cấp**: Sử dụng engine Regex mạnh mẽ để tìm ảnh bìa chất lượng cao từ Google
- **Hỗ trợ đa định dạng**: Hỗ trợ các định dạng âm thanh phổ biến (MP3, FLAC, WAV...)
- **Chọn ảnh thủ công**: Giao diện modal để chọn ảnh bìa từ 5 kết quả tìm được
- **Metadata thông minh**: Tự động lấy thông tin từ tên tệp nếu thiếu tag
- **Giao diện hiện đại**: UI Glassmorphism với hiệu ứng Staggered animations

## 🛠️ Công Nghệ

- **Backend**: Go 1.24.0 + Wails v2.11.0
- **Frontend**: React + TypeScript + Vite + Tailwind CSS + Framer Motion
- **Thư viện**: id3v2, tag, flacpicture

## 🚀 Cài Đặt

### Yêu cầu hệ thống:
- Linux, macOS hoặc Windows
- Go 1.24.0 trở lên
- Node.js 20.x trở lên

### Cách cài đặt:
1. Clone repo: `git clone https://github.com/skul9x/gocoverart.git`
2. Chạy: `go mod tidy`
3. Cài đặt frontend: `cd frontend && npm install`
4. Build: `wails build`

## 📦 Build Executable

Để tạo file executable:
```bash
wails build
```

File executable sẽ được tạo ở thư mục `build/bin/`

## 📖 Sử Dụng

1. Mở ứng dụng
2. Chọn thư mục chứa nhạc
3. Chọn ảnh bìa từ kết quả tìm kiếm
4. Áp dụng cho các tệp nhạc

## 📁 Cấu Trúc Dự Án

```
.
├── app.go              # Entry point của ứng dụng
├── backend/            # Logic backend (tìm ảnh, xử lý metadata)
├── frontend/           # Giao diện người dùng React
├── build/              # Thư mục build
├── docs/               # Tài liệu
├── CHANGELOG.md        # Nhật ký thay đổi
└── README.md           # Tài liệu này
```

## 🤝 Đóng Góp

1. Fork repo
2. Tạo branch feature mới
3. Commit thay đổi
4. Push lên remote
5. Tạo Pull Request

## 📄 Giấy Phép

MIT License - Xem file LICENSE để biết thêm chi tiết.
