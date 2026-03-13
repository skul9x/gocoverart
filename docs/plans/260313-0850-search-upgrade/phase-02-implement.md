# Phase 02: Implement Regex & Helpers
Status: ✅ Complete
Dependencies: phase-01-setup.md

## Objective
Chuyển đổi hoàn toàn 6 biểu thức chính quy (Regex) logic tìm kiếm và giải mã Unicode từ `GoogleImageHelper.kt` sang ngôn ngữ `Go`.

## Requirements
### Functional
- [ ] Parse được các URL ẩn trong thuộc tính `ou`, `data-src`, JavaScript array `AF_initDataCallback`.
- [ ] Loại bỏ khoảng trắng và thực hiện hàm unescape Unicode/HTML entities (như `\u003d` -> `=`, `&amp;` -> `&`).
- [ ] Lọc bỏ các URL rác như `encrypted-tbn0` hoặc `gstatic.com`.

## Implementation Steps
1. [ ] Cập nhật thư viện `regexp` thay vì dùng `goquery` cho Google.
2. [ ] Biên dịch sẵn (Compile) 6 Regex patterns dùng chung.
3. [ ] Viết hàm `decodeUnicode(input string) string`.
4. [ ] Viết hàm duyệt Regex lấy tất cả URLs, đẩy vào `map[string]bool` để lấy Unique URLs.

## Files to Create/Modify
- `backend/scraper.go` - Regex logic & helper functions.

---
Next Phase: `phase-03-refactor.md`
