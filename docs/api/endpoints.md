# API Documentation - Music Cover Tagger

**Base**: Wails Internal Bridge (Go to JS)

---

## 🎵 App Struct (Go)

Các hàm export từ `app.go` có thể gọi từ Frontend qua `window.go.main.App`.

### `SelectDirectory()`
Mở hộp thoại Linux Native để chọn thư mục.
- **Return**: `string` (đường dẫn thư mục)

### `GetTracks(dirPath string)`
Quét thư mục và lấy Metadata.
- **Param**: `dirPath` - Đường dẫn thư mục.
- **Return**: `[]TrackMetadata`

### `SearchCovers(query string)`
Tìm kiếm ảnh bìa trên Google Images.
- **Param**: `query` - Từ khóa (vd: "Artist Title").
- **Return**: `[]CoverResult` (Top 5 kết quả).

### `ApplyCover(filePath string, imageURL string)`
Tải và nhúng ảnh vào file.
- **Param**: 
  - `filePath`: Đường dẫn file nhạc.
  - `imageURL`: Link ảnh từ Google.
- **Return**: `error`

### `StartAutoTagging(tracks []TrackMetadata)`
Kích hoạt chế độ xử lý hàng loạt thông qua Goroutine.
- **Param**: Danh sách tracks cần xử lý.

---

## 📡 Events (Wails Event Bus)

Frontend lắng nghe các sự kiện này để cập nhật UI real-time.

| Event | Data | Mô tả |
|-------|------|-------|
| `track_processing` | `filePath` | Bắt đầu xử lý file này. |
| `track_updated` | `filePath` | Đã nhúng ảnh thành công. |
| `track_error` | `filePath` | Gặp lỗi khi xử lý (mạng/lock file). |
| `batch_finished` | `true` | Đã xử lý xong toàn bộ danh sách. |
