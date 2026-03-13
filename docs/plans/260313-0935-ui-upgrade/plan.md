# Plan: UI Upgrade & Real Cover Display
Created: 2026-03-13 09:35
Status: 🟡 In Progress

## Overview
Nâng cấp giao diện từ Grid View sang List View chuyên nghiệp hơn, đồng thời trích xuất và hiển thị ảnh bìa thực tế từ metadata của file nhạc (.mp3, .flac). Hỗ trợ cập nhật ảnh bìa trên UI ngay lập tức sau khi "Apply" thành công.

## Tech Stack
- **Frontend**: React (Vite), TailwindCSS, Framer Motion, Lucide React.
- **Backend**: Go (Wails v2), `github.com/dhowden/tag`, `github.com/go-flac/go-flac`.

## Phases

| Phase | Name | Status | Progress |
|-------|------|--------|----------|
| 01 | Backend: Cover Extraction | ✅ Complete | 100% |
| 02 | Frontend: Modern List Layout | ✅ Complete | 100% |
| 03 | Real-time Image Update Logic | ✅ Complete | 100% |
| 04 | Performance & Polish | ✅ Complete | 100% |

## Quick Commands
- Start Phase 1: `/code phase-01`
- Check progress: `/next`
- Save context: `/save_brain`
