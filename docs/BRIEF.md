# 💡 BRIEF: UI Upgrade & Real Cover Display

**Ngày tạo:** 2026-03-13
**Brainstorm cùng:** skul9x

---

## 1. VẤN ĐỀ CẦN GIẢI QUYẾT
- UI hiện tại dùng dạng Grid, khó quản lý danh sách nhạc dài và không hiển thị được đầy đủ tên bài hát.
- Người dùng không nhìn thấy được ảnh bìa hiện tại của file nhạc (chỉ thấy icon check/music).
- Không có phản hồi thị giác ngay lập tức khi thay đổi ảnh bìa thành công (phải dựa vào icon check).

## 2. GIẢI PHÁP ĐỀ XUẤT
- Chuyển đổi giao diện chính sang dạng **List View** (tương tự Spotify/Apple Music list).
- Trích xuất và hiển thị dữ liệu ảnh bìa thực tế từ metadata của file nhạc.
- Cơ chế cập nhật UI theo thời gian thực (Real-time update) ngay khi file được tag ảnh thành công.

## 3. ĐỐI TƯỢNG SỬ DỤNG
- Người dùng yêu thích nghe nhạc offline, muốn quản lý thư viện nhạc cá nhân đẹp mắt và đầy đủ thông tin.

## 4. TÍNH NĂNG

### 🚀 MVP (Bắt buộc có):
- [ ] **Real Cover Display**: Hiển thị ảnh bìa thực tế từ file .mp3 / .flac lên list.
- [ ] **Modern List View**: Thay thế Grid bằng danh sách hàng ngang, ưu tiên hiển thị text đầy đủ.
- [ ] **Instant UI Update**: Cập nhật ảnh bìa trên UI ngay sau khi "Apply" thành công mà không cần refresh.
- [ ] **Smart Fallback**: Hiển thị ảnh mặc định chuyên nghiệp khi file chưa có cover.

### 🎁 Phase 2 (Làm sau):
- [ ] **Lazy Loading Images**: Chỉ load ảnh khi scroll đến để tối ưu hiệu năng cho folder lớn.
- [ ] **Sorting & Filtering**: Sắp xếp theo tên, artist hoặc trạng thái (đã có/chưa có cover).
- [ ] **Drag & Drop**: Kéo thả ảnh trực tiếp vào bài hát để apply.

## 5. ƯỚC TÍNH SƠ BỘ
- **Độ phức tạp:** Trung bình (Cần sửa cả Backend để lấy Data và Frontend để hiển thị).
- **Rủi ro:** Folder chứa hàng nghìn bài hát có thể làm UI lag nếu gửi quá nhiều dữ liệu Base64 cùng lúc.

## 6. BƯỚC TIẾP THEO
→ Chạy `/plan` để lên thiết kế chi tiết các thay đổi trong `backend/metadata.go`, `app.go` và `App.tsx`.
