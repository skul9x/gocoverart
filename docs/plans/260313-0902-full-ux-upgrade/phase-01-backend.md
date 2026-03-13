# Phase 01: Backend Metadata & Logic
Status: ⬜ Pending

## Objective
Nâng cấp logic lấy thông tin bài hát và đảm bảo backend cung cấp đủ dữ liệu cho tính năng chọn ảnh thủ công.

## Implementation Steps
1. [ ] Cập nhật `backend/metadata.go`: Lấy tên tệp (bỏ extension) nếu `Title` trống.
2. [ ] Đảm bảo `backend/scraper.go` luôn trả về 5 kết quả (đã làm ở bước trước, chỉ cần verify).

## Files to Create/Modify
- `backend/metadata.go`
