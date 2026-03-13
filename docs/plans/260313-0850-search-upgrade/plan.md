# Plan: Nâng Cấp Thuật Toán Tìm Ảnh Bìa (Regex Search Mode)
Created: 2026-03-13T08:50:00+07:00
Status: ✅ Complete

## Overview
Dự án Music Cover Tagger sẽ nâng cấp module `scraper.go`. Thay vì dùng `goquery` quét các thẻ `<img>` thô (chất lượng thấp), ta sẽ copy toàn bộ logic từ file `GoogleImageHelper.kt` của dự án Gamedanhvan: giả lập Headers trình duyệt, dùng Regex để tìm các URL ẩn chứa hình ảnh High-Res, decode Unicode và lọc rác (thumbnails).

## Tech Stack
- Backend: `golang` (package: `net/http`, `regexp`, `net/url`, `strings`)
- Regex Engine: Chuỗi 6 bộ lọc regex lấy trực tiếp từ dự án `gamedanhvan`

## Phases

| Phase | Name | Status | Progress |
|-------|------|--------|----------|
| 01 | Setup Architecture & Test Bed | ✅ Complete | 100% |
| 02 | Implement Regex & Helpers | ✅ Complete | 100% |
| 03 | Refactor `SearchGoogleImage` | ✅ Complete | 100% |
| 04 | Integration & Testing | ✅ Complete | 100% |

## Quick Commands
- Start Phase 1: `/code phase-01`
- Check progress: `/next`
