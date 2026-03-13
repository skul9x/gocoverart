# Phase 03: Real-time Image Update Logic
Status: ✅ Complete
Dependencies: Phase 02

## Objective
Đảm bảo khi người dùng nhấn "Apply" một ảnh bìa mới, UI sẽ cập nhật ảnh đó ngay lập tức mà không cần load lại toàn bộ thư mục.

## Requirements
### Functional
- [x] Backend: Sau khi ghi ảnh vào file thành công, trích xuất lại ảnh đó và gửi kèm trong Event.
- [x] Frontend: Nhận Event `track_updated` chứa cả đường dẫn file và dữ liệu ảnh mới.
- [x] Frontend: Cập nhật state bài hát cụ thể trong danh sách.

## Implementation Steps
1. [x] Sửa `app.go`: Hàm `ApplyCover` khi thành công sẽ gọi hàm trích xuất metadata để lấy Base64 ảnh mới.
2. [x] Emit Event `track_updated` kèm Payload `{ filePath, coverData }`.
3. [x] Sửa `App.tsx`: Cập nhật `EventsOn("track_updated", ...)` để update state `tracks`.

## Files to Create/Modify
- `app.go` - Logic gửi Event. ✅ Done
- `frontend/src/App.tsx` - Logic nhận Event. ✅ Done

---
Next Phase: [Phase 04: Performance & Polish](file:///home/skul9x/Desktop/Test_code/gocoverart-main/docs/plans/260313-0935-ui-upgrade/phase-04-polish.md)
