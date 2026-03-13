# Phase 01: Backend Metadata Fallback
Status: ⬜ Pending

## Objective
Nâng cấp logic lấy metadata để luôn trả về một tiêu đề có ích (ưu tiên tag, fallback về tên tệp).

## Implementation Steps
1. [ ] Sửa func `ExtractMetadata` trong `backend/metadata.go`.
2. [ ] Nếu `m.Title()` rỗng hoặc lỗi, dùng `filepath.Base(filePath)` và xóa phần mở rộng để làm `Title`.

## Files to Create/Modify
- `backend/metadata.go`
