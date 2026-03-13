# 📖 SPECS: Music Cover Art Auto-Tagger

## 1. Executive Summary
Cover Art Tagger là một ứng dụng Desktop App trên nền tảng Linux, được xây dựng bằng Wails v2 (Go + React TS). Mục đích chính là giúp người dùng chọn một thư mục chứa nhạc, tự động quét siêu dữ liệu (metadata), tìm kiếm ảnh bìa trên Google Images, và nhúng ảnh đó trực tiếp vào các file nhạc (MP3/FLAC).

## 2. User Stories
- **Là một người dùng**, tôi muốn có thể bấm 1 nút để chọn Folder chứa toàn bộ kho nhạc của tôi.
- **Là một người dùng**, tôi muốn nhìn thấy danh sách các bài hát trong Folder đó trên màn hình trước khi thao tác tiếp.
- **Là một người dùng**, tôi muốn bấm nút "START" để app tự động tìm ảnh bìa Google cho từng bài hát một cách tuần tự.
- **Là một người dùng**, tôi muốn ứng dụng tự động thử link thứ 2, thứ 3... nếu link thứ 1 chết, để đảm bảo tỷ lệ thành công cao nhất.
- **Là một người dùng**, tôi muốn giao diện app tối giản, sang trọng (màu tối, hiệu ứng kính mờ - Glassmorphism, chuyển cảnh nhẹ nhàng).

## 3. Database Design / Dữ liệu
Dự án này **không sử dụng Database** truyền thống (SQL/NoSQL) mà thao tác trực tiếp trên bộ nhớ (RAM) và File System.
- **Trạng thái lưu trữ:** State của danh sách nhạc được giữ ở phía Frontend (React State) và Backend (Go Array/Slices) trong suốt vòng đời của 1 phiên làm việc.

```go
type Track struct {
    ID          string `json:"id"`
    FilePath    string `json:"filepath"`
    FileName    string `json:"filename"`
    Title       string `json:"title"`
    Artist      string `json:"artist"`
    Album       string `json:"album"`
    Status      string `json:"status"` // "pending", "processing", "success", "error"
    CoverURL    string `json:"cover_url"` // Phục vụ hiển thị UI tạm sau khi crawl
    ErrorMsg    string `json:"error_msg"`
}
```

## 4. Logic Flowchart (Mermaid)

```mermaid
sequenceDiagram
    participant U as User (UI)
    participant F as Frontend (React)
    participant B as Backend (Go)
    participant G as Google Images
    participant FS as File System

    U->>F: Click "Select Folder"
    F->>B: Call SelectFolder()
    B->>FS: Opendir & Scan .mp3, .flac
    B-->>F: Trả về []Track (Danh sách nhạc)
    F-->>U: Hiển thị Grid danh sách
    U->>F: Click "START"
    F->>B: Call StartProcessing()
    loop Mỗi Track trong Danh sách
        B->>F: Emit Event "TrackProcessing" (ID)
        B->>FS: Đọc Metadata (Title, Artist)
        B->>G: Scrape Search: "Artist Title cover art"
        G-->>B: Trả về Top 5 URLs
        loop Thử từng URL (1-5)
            B->>G: HTTP Get Image
            alt Thành công
                B->>FS: Nhúng Image Bytes vào File (ID3/FLAC tag)
                B->>F: Emit Event "TrackSuccess" (ID, CoverURL)
                break
            else Thất bại
                B->>B: Ghi log, thử URL tiếp theo
            end
        end
        alt Hết 5 URL đều lỗi
            B->>F: Emit Event "TrackError" (ID, Lỗi)
        end
    end
    B-->>F: Emit Event "ProcessComplete"
    F-->>U: Tick xanh, báo hoàn thành 100%.
```

## 5. API Contract (Wails Bindings)
Các hàm Go này sẽ được expose thành JS Promises cho Frontend gọi:
- `SelectMusicFolder() ([]Track, error)`: Mở hộp thoại chọn thư mục và quét sơ bộ.
- `StartProcessing(tracks []Track)`: Kích hoạt tiến trình cào ảnh và ghi file trên Go thread.

**Sự kiện (Events Emit từ Go sang React):**
- `track_processing`: Payload là `track_id` -> JS cập nhật icon spinner.
- `track_success`: Payload là `{track_id, cover_url}` -> JS cập nhật ảnh bìa và icon xanh.
- `track_error`: Payload là `{track_id, error_msg}` -> JS cập nhật icon đỏ sẫm báo lỗi.

## 6. UI Components
- **Core Layout:** Chứa Drag-region cho thanh tiêu đề window native. Frame nền tối kết hợp Glassmorphism.
- **Home/Hero Component:** Nút Select Folder to ở chính giữa với hiệu ứng Glow/Ripple.
- **Track Grid Component:** Scroll ảo (Virtual Scroll) hoặc Grid CSS để render danh sách. Mỗi item là một thẻ chữ nhật có: Khung ảnh bìa (nhỏ), Tên Bài Hát, Trạng thái (Pending / Loading / Đã Xong).
- **Progress Overlay/Header:** Nằm cố định bên dưới hoặc trên cùng, hiện ProgressBar và Text dạng "Processing 12 / 105 files...".

## 7. Third-party Integrations
- Khai thác kết quả của **Google Images**. Cần parser HTML DOM thông minh để lọc đúng tag `<img>` có chứa link thật. Thư viện Go: `PuerkitoBio/goquery`.

## 8. Hidden Requirements & Edge Cases
- **Bảo toàn File:** Ghi đè ID3 tag luôn có rủi ro hỏng file gốc nếu bị crash giữa chừng. Backup file là cần thiết `file.mp3 -> file.mp3.bak` trước khi ghi, ghi xong xóa `.bak`.
- **Giới hạn tỷ lệ (Rate Limit):** Cào Google Image quá nhanh có thể dính ReCAPTCHA hoặc 429 Too Many Requests. Cần thời gian nghỉ (Sleep) khoảng 1-2 giây giữa mỗi bài hát.
- **Link ảo:** Một số kết quả <img> của Google là dạng `data:image/base64,...` hoặc link Tracking. Cần lọc các link bắt đầu bằng `http`/`https` thật sự.
- **File bị khóa:** File có thể đang dc play trên trình nghe nhạc nên cấm ghi. Try/catch cẩn thận.

## 9. Tech Stack
- Frontend: Wails v2, React 18, TypeScript, TailwindCSS, Framer Motion, Lucide React (Icons).
- Backend: Go 1.21+, `wailsapp/wails/v2`, `dhowden/tag`, `goquery`.
- OS Target: Linux (amd64 / arm64).

## 10. Build Checklist
- Thiết lập ảnh Icon Logo cho App (`appicon.png`).
- Biên dịch bằng `wails build -platform linux/amd64`.
- Tạo file chạy độc lập (.bin hoặc AppImage).
