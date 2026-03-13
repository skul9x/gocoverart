# Music Cover Tagger

Ứng dụng desktop hiện đại để quản lý ảnh bìa và metadata cho bộ sưu tập nhạc của bạn. Được xây dựng với Go (Wails) và React.

## 🎯 Tính Năng Chính

- **Tìm ảnh bìa cao cấp**: Sử dụng engine thông minh để tìm ảnh bìa chất lượng cao từ internet.
- **Chỉnh sửa Metadata thủ công (MỚI)**: 
  - Giao diện chỉnh sửa trực quan ngay trong ứng dụng.
  - Hỗ trợ đầy đủ các trường: **Title, Artist, Album, Year, Genre, Comment**.
  - Validation trường Title (bắt buộc) để đảm bảo tính toàn vẹn dữ liệu.
  - Hộp thoại xác nhận an toàn trước khi ghi đè dữ liệu vào tệp.
- **Hỗ trợ đa định dạng**: Làm việc tốt với các tệp **MP3** và **FLAC**.
- **Chọn ảnh thủ công**: Giao diện modal chuyên nghiệp để chọn ảnh bìa ưng ý nhất.
- **Giao diện hiện đại**: UI Glassmorphism (hiệu ứng kính mờ) tối màu cực đẹp với các hiệu ứng chuyển động mượt mà từ Framer Motion.
- **Metadata thông minh**: Tự động gợi ý thông tin từ tên tệp nếu tag bị thiếu.

## 🛠️ Công Nghệ

- **Backend**: Go 1.24.0 + [Wails v2](https://wails.io/)
- **Frontend**: React + TypeScript + Vite + Tailwind CSS + Framer Motion
- **Thư viện âm thanh**: 
  - `github.com/bogem/id3v2/v2` (MP3)
  - `github.com/go-flac/go-flac` & `flacvorbis` (FLAC)
  - `github.com/dhowden/tag` (Metadata extraction)

## 🚀 Cài Đặt

### Yêu cầu hệ thống:
- Go 1.24.0+
- Node.js 20.x+
- Wails CLI (`go install github.com/wailsapp/wails/v2/cmd/wails@latest`)

### Cách chạy:
1. Clone repo: `git clone https://github.com/skul9x/gocoverart.git`
2. Cài đặt dependency: `go mod tidy`
3. Chạy ở chế độ phát triển:
   ```bash
   wails dev
   ```

## 📦 Build Executable

Để tạo file ứng dụng hoàn chỉnh cho hệ điều hành của bạn:
```bash
wails build
```
File kết quả sẽ nằm trong thư mục `build/bin/`.

## 📖 Hướng Dẫn Sử Dụng

1. **Chọn thư mục**: Nhấn "Open Library" để tải danh sách nhạc.
2. **Tìm ảnh bìa**: Nhấp vào một bài nhạc để tìm ảnh bìa, sau đó chọn ảnh bạn thích để lưu.
3. **Sửa thông tin**: Rê chuột vào bài nhạc, nhấn icon **Cây bút (Edit)** ở cuối dòng để sửa Title, Artist...
4. **Auto Tag**: Sử dụng tính năng "Auto Tag Missing" để tự động tìm ảnh bìa cho các bài còn thiếu.

## 🤝 Đóng Góp

Mọi đóng góp nhằm cải thiện ứng dụng đều được hoan nghênh. Hãy Fork repo và tạo Pull Request!

## 📄 Giấy Phép

MIT License - skun9x
