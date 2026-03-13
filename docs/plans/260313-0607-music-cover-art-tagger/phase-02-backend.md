# Phase 02: Backend Metadata & Scraper
Status: ✅ Complete
Dependencies: phase-01-setup.md

## Objective
Viết logic Go để đọc ID3 tags từ file âm thanh, cào ảnh từ Google Images và nhúng (embed) ảnh vào ngược file âm thanh.

## Requirements
### Functional
- [x] `ScanFolder`: Hàm quét thư mục và tìm ra toàn bộ tập tin `.mp3` và `.flac`.
- [x] `ExtractMetadata`: Hàm nhận path file, trả về `Artist`, `Title`, `Album`.
- [x] `SearchGoogleImage`: Hàm nhận text (VD: "Artist Title cover art"), lấy Top 5 URLs ảnh từ Google Images.
- [x] `EmbedCoverArt`: Tải ảnh đó về ram -> Dùng thư viện `id3v2/tag` để nhúng dữ liệu bit của ảnh vào original file. Backup file lỗi trước khi ghi nếu có rủi ro hỏng tệp.

### Non-Functional
- [x] Quét nhanh, không block UI thread.
- [x] Handle panic ở Go nếu ảnh corrupt / scrape bị HTTP 429 chặn.

## Implementation Steps
1. [x] Tạo struct `TrackInfo` chứa path, id, metadata, state.
2. [x] Viết bộ quét file trong thư mục đệ quy.
3. [x] Viết Web Scraper giả lập User-Agent, bóc tách `img` tag từ DOM kết quả tìm kiếm Google.
4. [x] Lặp theo logic ưu tiên (Top 1 lỗi HTTP -> Thử tiếp Top 2).
5. [x] Đóng file mp3 -> Nhúng Blob (JPG) làm `AttachedPicture` Frame cho mp3.

## Test Criteria
- [ ] Đưa vào 1 file nhạc `.mp3` fake, Go code cào được URL và nhúng được ảnh. Mở app VLC thấy có ảnh cover.

---
Next Phase: phase-03-frontend.md
