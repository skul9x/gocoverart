# Phase 03: Frontend UI & Bridge
Status: ✅ Complete
Dependencies: phase-01-setup.md

## Objective
Thiết kế giao- [x] Dựng khung layout Glassmorphism (Sidebar + Main Content).
- [x] Component danh sách nhạc với Grid & Animations (Framer Motion).
- [x] Modal chọn ảnh bìa (Cover Picker).
- [x] Kết nối bridge JS với các hàm Go (`GetTracks`, `ApplyCover`).
- [x] Khi chọn xong, hiện thành Grid List các danh sách file chờ xử lý.
- [x] Mỗi mục grid song hiện `[Title] - [Artist]` và một trạng thái [Trắng].
- [x] Khu vực nút bấm siêu bự "START" style premium để kích hoạt.

### Non-Functional
- [x] Background blur nhẹ nhõm, chuyển cảnh page mượt mà bằng `framer-motion`.
- [x] UI không lag khi render list danh sách lên đến nghìn bài hát (dùng lazy render nếu cần).

## Implementation Steps
1. [x] Xây dựng Tailwind Utility cho Glassmorphism.
2. [x] Tạo `App.tsx` quản lý state (tracks, directory, selections).
3. [x] Tích hợp `lucide-react` icons.
4. [x] Viết logic gọi hàm bridge thông qua alias `wailsjs/go/main/App`.
5. [x] Phối hợp animation entry/exit cho các card nhạc.iển thị X/Y files processed).
5. [x] Define TypeScript interface kết nối với Go Structs sinh từ wails.

## Test Criteria
- [x] Giao diện click nút fake ra danh sách mượt, hover đẹp, scroll êm.

---
Next Phase: phase-04-integration.md
