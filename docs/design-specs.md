# Design Specifications: Cover Art Tagger

## 🎨 Color Palette (Dark Glassmorphism)
| Name | Hex | Usage |
|------|-----|-------|
| Primary | `#3b82f6` | Nút bấm chính, Progress Bar, Active states |
| Primary Glow | `rgba(59, 130, 246, 0.5)` | Box-shadow cho nút bấm |
| Secondary | `#10b981` | Trạng thái Success (Tìm thấy ảnh) |
| Danger | `#ef4444` | Trạng thái Error (Không tìm thấy ảnh sau 5 lần) |
| Background | `#0f172a` | Màu nền gốc của App (Slate 900) |
| Glass Surface | `rgba(30, 41, 59, 0.7)` | Nền của các Card, List Item (cần kết hợp `backdrop-blur`) |
| Glass Border | `rgba(255, 255, 255, 0.1)` | Viền mỏng cho các khối Glass |
| Text Main | `#f8fafc` | Chữ chính (Tên bài hát) |
| Text Muted | `#94a3b8` | Chữ phụ (Tên ca sĩ, Trạng thái) |

## 📝 Typography
| Element | Font | Size | Weight |
|---------|------|------|--------|
| App Title | Inter | 24px | 700 (Bold) |
| Track Title | Inter | 16px | 600 (Semibold)|
| Track Artist| Inter | 14px | 400 (Regular)|
| Button Text | Inter | 16px | 600 (Semibold)|
| Status Text | Inter | 12px | 500 (Medium) |

*Lưu ý: Font Inter được load trực tiếp qua Google Fonts hoặc nhúng local.*

## 📐 Layout & Spacing
- **Window Size Layout:** Desktop cố định khoảng `900px x 650px` (Không cho resize quá nhỏ để tránh vỡ grid).
- **Padding Container:** `24px` (p-6)
- **Gap giữa các Track:** `12px` (gap-3)

## 🔲 Border Radius & Shapes
| Element | Value (Tailwind) |
|---------|------------------|
| Track Item | `rounded-xl` (12px) |
| Cover Image | `rounded-md` (6px) |
| Main Button | `rounded-full` (9999px) |

## ✨ Animations & Effects
- **Glass Effect:** Thêm CSS `backdrop-filter: blur(16px); background: rgba(30, 41, 59, 0.7);`
- **Glow Effect (Nút Start):** `box-shadow: 0 0 20px rgba(59, 130, 246, 0.5);`
- **Fade In (Framer motion):** Khi list bài hát xuất hiện, list item trồi lên nhẹ nhàng (stagger children).

## 🚀 Màn Hình UI Component Breakdown

### 1. Màn hình rỗng (Empty State)
- Chính giữa màn hình là Icon thư mục (Folder) thật to, phát sáng nhẹ.
- Text: "Kéo thả hoặc nhấn để chọn thư mục nhạc".
- Nút "Select Folder": màu Primary, bo tròn viên thuốc (pill).

### 2. Màn hình danh sách (Track List)
- **Top Bar (Drag region):** Nơi để kéo thả cửa sổ App. Ghi chữ "Cover Art Tagger".
- **Main Area:** Cuộn dọc.
  - Mỗi Track là 1 thẻ nằm ngang (Row). 
  - Bên trái là ô vuông hiển thị Cover (ban đầu là logo đĩa than xám).
  - Giữa là `[Tên bài hát]` (to) \n `[Tên ca sĩ]` (nhỏ hơn).
  - Bên phải là Icon trạng thái (⏳ Đang đợi).
- **Bottom Bar (Fixed):** 
  - Nút **"START TAGGING"** cực bự, màu xanh dương với hiệu ứng lan tỏa (Pulse/Glow).
  - Chữ nhỏ bên dưới báo số lượng (vd: "0/150 files").

### 3. Trạng thái Đang Chạy (Processing State)
- Các Item trong danh sách đang xử lý sẽ có Spinner xoay.
- Nút "START" biến thành nút "STOP" (Màu Đỏ) hoặc mờ đi không cho bấm.
- Bottom Bar hiện Progress Bar chạy từ 0 đến 100%. Mượt mà.
- Khi một file xong, ô vuông bên trái lật hiệu ứng ra Cover thật vừa bóc từ Google. Giữ Glassmục đẹp.
