# Phase 01: Setup Environment
Status: ✅ Complete
Dependencies: N/A

## Objective
Khởi tạo cấu trúc dự án Wails v2 với Go làm backend và React + TS làm frontend, cài đặt các package UI cần thiết.

## Requirements
### Functional
- [x] Khởi tạo Wails v2 project (`wails init -n cover -t react-ts`).
- [x] Cài đặt TailwindCSS và cấu hình tại frontend.
- [x] Cài đặt Framer Motion cho giao diện frontend.
- [x] Cài đặt các thư viện Go (Backend): `dhowden/tag` và `PuerkitoBio/goquery`.

### Non-Functional
- [x] Cấu trúc folder chuẩn bị tách bạch logic.
- [x] Project chạy mượt mà ở chế độ dev (`wails dev`).

## Implementation Steps
1. [x] Chạy lệnh wails tạo base app `React-TS`.
2. [x] `cd frontend` và setup TailwindCSS (Install `tailwindcss`, `postcss`, `autoprefixer` -> Gen config -> Update `index.css`).
3. [x] `cd frontend` và install `framer-motion`, `lucide-react` (icons).
4. [x] `cd ..` (về thư mục root Go) -> Chạy `go get github.com/dhowden/tag` và `go get github.com/PuerkitoBio/goquery`.
5. [x] Tạo file mô phỏng tính năng Backend rỗng (`backend/core.go`).

## Files to Create/Modify
- `wails.json` - config project.
- `frontend/tailwind.config.js` - cấu hình style.
- `frontend/src/index.css` - chứa style base.
- `go.mod` & `go.sum` - danh sách dependencies của Backend.

## Test Criteria
- [ ] `wails dev` build thành công, cửa sổ React trắng xuất hiện có class TailwindCSS.

---
Next Phase: phase-02-backend.md
