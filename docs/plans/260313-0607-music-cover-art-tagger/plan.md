# Plan: Music Cover Art Auto-Tagger
Created: 2026-03-13T06:07:32
Status: 🟡 In Progress

## Overview
Desktop App tự động quét thư mục nhạc, lấy thông tin bài hát (metadata) và sử dụng Google Image để tìm và gán Cover Art (Top 1 đến 5 fallback).

## Tech Stack
- **Framework:** Wails v2
- **Frontend:** React + TypeScript + TailwindCSS + Framer Motion
- **Backend:** Go (Golang)
- **Core Libs:** `dhowden/tag` (đọc ID3 tags), `goquery` (scrape Google Images).

## Phases

| Phase | Name | Status | Progress |
|-------|------|--------|----------|
| 01 | Setup Environment | ✅ Complete | 100% |
| 02 | Backend Metadata & Scraper | ✅ Complete | 100% |
| 03 | Frontend UI & Bridge | ✅ Complete | 100% |
| 04 | Integration & Logic | ✅ Complete | 100% |
| 05 | Testing & Polish | ✅ Complete | 100% |

## Quick Commands
- Start Phase 1: `/code phase-01`
- Check progress: `/next`
- Save context: `/save-brain`
