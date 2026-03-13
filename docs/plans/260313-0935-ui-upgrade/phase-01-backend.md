# Phase 01: Backend - Cover Extraction
Status: ✅ Complete

## Objective
Cập nhật Backend để trích xuất ảnh bìa thực tế từ metadata của file nhạc và gửi về cho Frontend dưới dạng Base64.

## Requirements
### Functional
- [ ] Mở rộng struct `TrackMetadata` để chứa dữ liệu ảnh.
- [ ] Cập nhật hàm `ExtractMetadata` để đọc `Picture()` data.
- [ ] Chuyển đổi dữ liệu ảnh sang định dạng Base64 string thân thiện với HTML/CSS.

### Non-Functional
- [ ] Tối ưu dung lượng: Cân nhắc thu nhỏ ảnh nếu quá lớn (optional).
- [ ] Xử lý lỗi: Không để việc lỗi đọc ảnh làm treo việc lấy metadata.

## Implementation Steps
1. [x] Sửa file `backend/core.go`: Thêm trường `CoverData string` vào `TrackMetadata`.
2. [x] Sửa file `backend/metadata.go`: 
    - [x] Lấy dữ liệu từ `m.Picture()`.
    - [x] Dùng `base64.StdEncoding.EncodeToString` để convert.
    - [x] Gán vào `CoverData`.

## Files to Create/Modify
- `backend/core.go` - Update struct. ✅ Done
- `backend/metadata.go` - Update extraction logic. ✅ Done

---
Next Phase: [Phase 02: Modern List Layout](file:///home/skul9x/Desktop/Test_code/gocoverart-main/docs/plans/260313-0935-ui-upgrade/phase-02-frontend-list.md)
