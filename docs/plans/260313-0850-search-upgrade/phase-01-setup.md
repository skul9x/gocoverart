# Phase 01: Setup Architecture & Test Bed
Status: ✅ Complete

## Objective
Tạo sẵn môi trường và file testing để có thể test độc lập thuật toán lấy ảnh bìa từ Google (trước khi tích hợp vào backend chính).

## Requirements
### Functional
- [ ] Chạy được 1 script go đơn giản lấy HTML từ Google.

## Implementation Steps
1. [ ] Tạo file `backend/scraper_test.go` hoặc một cmd nhỏ để test func `SearchGoogleImage`.
2. [ ] Fake bộ Headers giống hệt `GoogleImageHelper.kt` lưu vào một biến hằng (constant map).
3. [ ] Viết hàm gọi HTTP GET với Headers đó để in ra độ dài nội dung HTML trả về (đảm bảo không bị Google block).

## Files to Create/Modify
- `backend/scraper.go` - Cập nhật headers chuẩn.
- `backend/scraper_test.go` - Script cho việc test offline.

---
Next Phase: `phase-02-implement.md`
