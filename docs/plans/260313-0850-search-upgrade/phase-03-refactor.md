# Phase 03: Refactor SearchGoogleImage
Status: ✅ Complete
Dependencies: phase-02-implement.md

## Objective
Sửa đổi hàm `SearchGoogleImage` cũ, thay thế phân tích `goquery` bằng Regexp engine vừa tạo ở Phase 02.

## Requirements
### Functional
- [ ] Gửi request HTTP GET với query được encode (vd: url.QueryEscape).
- [ ] Chạy mảng 6 Regex trên body HTML trả về.
- [ ] Lọc ra tối đa 5 kết quả tốt nhất.

## Implementation Steps
1. [ ] Cập nhật `SearchGoogleImage(query string) ([]CoverResult, error)`.
2. [ ] Map unique urls có được thành `[]CoverResult`. Rank theo chuẩn thứ tự regex (tìm được trước -> rank cao hơn).

## Files to Create/Modify
- `backend/scraper.go` - Clean up goquery code, return array of `CoverResult`.

---
Next Phase: `phase-04-integration.md`
