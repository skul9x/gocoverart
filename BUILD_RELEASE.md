# 📦 Hướng Dẫn Build & Release

## Build Files

Để build file `.deb` và `.exe`:

```bash
./build.sh
```

Kết quả trong `build/bin/`:
- `music-cover-tagger_1.0.0_amd64.deb` - Package cho Debian/Ubuntu
- `music-cover-tagger.exe` - File cho Windows
- `music-cover-tagger` - Binary cho Linux

## Upload lên GitHub Releases

### Cách 1: Dùng script tự động (khuyến nghị)

```bash
./release.sh v1.0.0 "Release notes của bạn"
```

### Cách 2: Thủ công

1. Vào GitHub repo: https://github.com/skul9x/gocoverart
2. Click **Releases** → **Draft a new release**
3. Tạo tag mới (vd: `v1.0.0`)
4. Upload 2 files:
   - `music-cover-tagger_1.0.0_amd64.deb`
   - `music-cover-tagger.exe`
5. Viết release notes
6. Click **Publish release**

## Cài Đặt

### Linux (Debian/Ubuntu)
```bash
sudo dpkg -i music-cover-tagger_1.0.0_amd64.deb
```

### Windows
Double-click file `music-cover-tagger.exe`

## Yêu Cầu

- **Build**: Wails CLI, Go 1.24+, Node.js 20+
- **Release**: GitHub CLI (`gh`) hoặc upload thủ công
