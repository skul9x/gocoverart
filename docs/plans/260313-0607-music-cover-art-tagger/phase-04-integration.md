# Phase 04: Integration & Logic
Status: ⬜ Pending
Dependencies: phase-02-backend.md, phase-03-frontend.md

## Objective
Nối Frontend event `Bấm START` với luồng Go quét file, tải ảnh và cover. Báo cáo % tiến trình real-time qua chuẩn Wails Events.

## Requirements
### Functional
- [ ] Khi Frontend bấm Start, gọi API qua wails: `StartProcessing()`.
- [ ] Backend loop qua danh sách file, đẩy ngược Event `ItemProcessing`, `ItemFinished`, `ItemError` qua JS event bus nội bộ Wails.
- [ ] JS bắt event và update React State list các bài hát để hiện hiệu ứng loading/finished.

### Non-Functional
- [ ] Chống memory rò rỉ trong file loop của Go.
- [ ] Hạn chế tần suất spam Event vào JS thread.

## Implementation Steps
1. [ ] Wrap Wails `context` vào `App` struct để gọi `runtime.EventsEmit`.
2. [ ] Sửa `StartProcessing`: thêm goroutine, có channel hoặc WaitGroup quản lý các file.
3. [ ] Viết hàm JS ListenEvent tại root Component `useEffect` của React.
4. [ ] Mở log console trên cả 2 side.

## Test Criteria
- [ ] Chạy với kho ngẫu nhiên 10 mp3 -> Click start -> Xong tệp nào danh sách trên UI có "tick" tệp đó ngay lập tức (1s/file).

---
Next Phase: phase-05-testing.md
