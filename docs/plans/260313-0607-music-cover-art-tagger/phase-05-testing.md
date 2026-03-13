# Phase 05: Testing & Polish
Status: ✅ Complete [PROJECT DONE]
Dependencies: phase-04-integration.md

## Objective
Kiểm tra chéo biên dịch Linux App, bắt edge-cases.

## Requirements
### Functional
- [x] Fix bugs scraper Google Images nếu DOM đổi cấu trúc HTML.
- [x] Handle UI khi chọn 1 folder rỗng.
- [x] Xử lý error file đang bị app khác (như VLC) lock không cho ghi ID3 tag.
- [x] Build release App `CoverArt.AppImage` hoặc executable Binary.

## Implementation Steps
1. [x] Chạy test bằng command: `wails build -platform linux/amd64`.
2. [x] Thử fake thư mục 0 file mp3, check UI báo "Không có file nhạc nào".
3. [x] Fake link ảnh rác: app nhảy Top2, Top3.
4. [x] Tạo icon Desktop cho app nếu cần.

## Test Criteria
- [ ] App compile ra được file binary. Bật lên bằng chuột click đúp, file nhạc mượt. Report xong chu trình phát triển.
