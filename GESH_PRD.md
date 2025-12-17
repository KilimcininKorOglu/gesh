# GESH - Product Requirements Document

> **Gesh** (𒄑) - Sümer dilinde "kalem, yazı aleti" anlamına gelir.
> A minimal TUI text editor written in Go with Bubble Tea.

**Versiyon:** 1.0.0  
**Tarih:** Aralık 2024  
**Durum:** Taslak

---

## 1. Yönetici Özeti

Gesh, Go programlama dili ve Bubble Tea framework'ü kullanılarak geliştirilecek, nano benzeri minimal bir terminal tabanlı metin editörüdür. Hedef, hızlı, hafif ve kullanıcı dostu bir düzenleme deneyimi sunmaktır.

### 1.1 Vizyon

```
┌─ GESH ──────────────────────────────────────────────────────────┐
│  Basitlik + Güç + Hız = Modern Terminal Editörü                 │
└─────────────────────────────────────────────────────────────────┘
```

### 1.2 Hedef Kitle

- Terminal kullanıcıları
- Sistem yöneticileri
- Geliştiriciler (hızlı düzenleme için)
- Nano kullanıcıları (geçiş kolaylığı)

---

## 2. Proje Hedefleri

### 2.1 Birincil Hedefler

| Hedef | Açıklama | Öncelik |
|-------|----------|---------|
| Minimal tasarım | Nano benzeri sade arayüz | P0 |
| Hızlı başlangıç | < 50ms açılış süresi | P0 |
| Düşük bellek | < 10MB RAM kullanımı | P0 |
| Cross-platform | Linux, macOS, Windows | P0 |
| Tek binary | Bağımlılıksız dağıtım | P0 |

### 2.2 İkincil Hedefler

| Hedef | Açıklama | Öncelik |
|-------|----------|---------|
| Syntax highlighting | Temel dil desteği | P1 |
| Tema desteği | Açık/koyu tema | P1 |
| Plugin sistemi | Genişletilebilirlik | P2 |
| LSP entegrasyonu | Kod tamamlama | P2 |

---

## 3. Fonksiyonel Gereksinimler

### 3.1 Temel Düzenleme (MVP)

#### 3.1.1 Dosya İşlemleri

| Özellik | Kısayol | Açıklama |
|---------|---------|----------|
| Yeni dosya | `Ctrl+Alt+N` | Boş buffer oluştur |
| Dosya aç | `Ctrl+O` | Dosya seçici dialog |
| Kaydet | `Ctrl+S` | Mevcut dosyaya kaydet |
| Farklı kaydet | `Ctrl+Shift+S` | Yeni isimle kaydet |
| Çıkış | `Ctrl+X` | Programdan çık (seçim yoksa) |

#### 3.1.2 Metin Düzenleme

| Özellik | Kısayol | Açıklama |
|---------|---------|----------|
| Karakter sil (geri) | `Backspace` | İmleç öncesi sil |
| Karakter sil (ileri) | `Delete` | İmleç sonrası sil |
| Satır sil | `Ctrl+K` | Tüm satırı sil |
| Satır kes | `Ctrl+U` | Satırı kes (clipboard) |
| Yapıştır | `Ctrl+V` | Clipboard'dan yapıştır |
| Geri al | `Ctrl+Z` | Son işlemi geri al |
| Yinele | `Ctrl+Y` | Geri alınanı yinele |

#### 3.1.3 Navigasyon

| Özellik | Kısayol | Açıklama |
|---------|---------|----------|
| İmleç yukarı | `↑` / `Ctrl+P` | Bir satır yukarı |
| İmleç aşağı | `↓` / `Ctrl+N` | Bir satır aşağı |
| İmleç sol | `←` / `Ctrl+B` | Bir karakter sol |
| İmleç sağ | `→` / `Ctrl+F` | Bir karakter sağ |
| Satır başı | `Home` / `Ctrl+A` | Satır başına git (tekrar: tümünü seç) |
| Satır sonu | `End` / `Ctrl+E` | Satır sonuna git |
| Dosya başı | `Ctrl+Home` | Dosya başına git |
| Dosya sonu | `Ctrl+End` | Dosya sonuna git |
| Sayfa yukarı | `PageUp` | Bir sayfa yukarı |
| Sayfa aşağı | `PageDown` | Bir sayfa aşağı |
| Satıra git | `Ctrl+G` | Belirli satıra atla |
| Kelime sol | `Ctrl+←` | Önceki kelimeye |
| Kelime sağ | `Ctrl+→` | Sonraki kelimeye |

#### 3.1.4 Arama & Değiştirme

| Özellik | Kısayol | Açıklama |
|---------|---------|----------|
| Ara | `Ctrl+W` | Metin ara |
| Sonraki | `F3` / `Ctrl+W` (tekrar) | Sonraki eşleşme |
| Önceki | `Shift+F3` | Önceki eşleşme |
| Değiştir | `Ctrl+R` | Bul ve değiştir |
| Tümünü değiştir | `Ctrl+Shift+R` | Tümünü değiştir |

#### 3.1.5 Seçim

| Özellik | Kısayol | Açıklama |
|---------|---------|----------|
| Seçim modu | `Ctrl+Space` | Seçimi başlat/bitir |
| Tümünü seç | `Ctrl+A` (2x) | Tüm metni seç |
| Shift+Ok tuşları | `Shift+↑↓←→` | Seçerek hareket |
| Seçimi kopyala | `Ctrl+C` | Clipboard'a kopyala |
| Seçimi kes | `Ctrl+X` (seçim varken) | Kes |

### 3.2 Gelişmiş Özellikler (v1.1+)

#### 3.2.1 Syntax Highlighting

Desteklenecek diller (öncelik sırasına göre):

| Faz | Diller |
|-----|--------|
| Faz 1 | Go, Python, JavaScript, JSON |
| Faz 2 | Rust, C, C++, Java |
| Faz 3 | HTML, CSS, YAML, TOML, Markdown |
| Faz 4 | PHP, Ruby, Shell/Bash |

#### 3.2.2 Çoklu Buffer

- Tab benzeri buffer yönetimi
- `Ctrl+Tab` ile buffer değiştirme
- Split view (yatay/dikey)

#### 3.2.3 Makrolar

- `Ctrl+M` makro kayıt başlat/durdur
- `Ctrl+Shift+M` makro çalıştır
- Makro kaydetme/yükleme

---

## 4. Teknik Mimari

### 4.1 Sistem Mimarisi

```
┌────────────────────────────────────────────────────────────────────┐
│                         GESH ARCHITECTURE                          │
├────────────────────────────────────────────────────────────────────┤
│                                                                    │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐         │
│  │   Terminal   │───▶│  Bubble Tea  │───▶│    Model     │         │
│  │    Input     │    │   Runtime    │    │    State     │         │
│  └──────────────┘    └──────────────┘    └──────────────┘         │
│                             │                    │                 │
│                             ▼                    ▼                 │
│                      ┌──────────────┐    ┌──────────────┐         │
│                      │    View      │◀───│   Update     │         │
│                      │   Render     │    │   Handler    │         │
│                      └──────────────┘    └──────────────┘         │
│                             │                    │                 │
│                             ▼                    ▼                 │
│  ┌──────────────┐    ┌──────────────┐    ┌──────────────┐         │
│  │   Terminal   │◀───│   Lipgloss   │    │   Buffer     │         │
│  │   Output     │    │   Styling    │    │   Manager    │         │
│  └──────────────┘    └──────────────┘    └──────────────┘         │
│                                                  │                 │
│                                                  ▼                 │
│                                          ┌──────────────┐         │
│                                          │  File I/O    │         │
│                                          │   Manager    │         │
│                                          └──────────────┘         │
│                                                                    │
└────────────────────────────────────────────────────────────────────┘
```

### 4.2 Modül Yapısı

```
gesh/
├── main.go                 # Uygulama giriş noktası
├── go.mod
├── go.sum
├── Makefile
├── README.md
│
├── internal/
│   ├── app/
│   │   ├── app.go          # Ana uygulama yapısı
│   │   ├── model.go        # Bubble Tea model
│   │   ├── update.go       # Update fonksiyonları
│   │   ├── view.go         # View render
│   │   └── commands.go     # Bubble Tea commands
│   │
│   ├── buffer/
│   │   ├── buffer.go       # Buffer veri yapısı
│   │   ├── gap.go          # Gap buffer implementasyonu
│   │   ├── cursor.go       # Cursor yönetimi
│   │   ├── selection.go    # Seçim yönetimi
│   │   └── history.go      # Undo/Redo stack
│   │
│   ├── editor/
│   │   ├── editor.go       # Editor mantığı
│   │   ├── keybindings.go  # Tuş atamaları
│   │   ├── clipboard.go    # Clipboard işlemleri
│   │   └── search.go       # Arama/değiştirme
│   │
│   ├── ui/
│   │   ├── components/
│   │   │   ├── header.go   # Üst bilgi çubuğu
│   │   │   ├── editor.go   # Editor alanı
│   │   │   ├── statusbar.go # Alt durum çubuğu
│   │   │   ├── dialog.go   # Dialog pencereleri
│   │   │   └── prompt.go   # Input prompt
│   │   │
│   │   ├── styles/
│   │   │   ├── theme.go    # Tema tanımları
│   │   │   └── colors.go   # Renk paletleri
│   │   │
│   │   └── layout.go       # Layout yönetimi
│   │
│   ├── file/
│   │   ├── file.go         # Dosya işlemleri
│   │   ├── watcher.go      # Dosya değişiklik takibi
│   │   └── encoding.go     # Encoding desteği
│   │
│   ├── syntax/
│   │   ├── highlighter.go  # Syntax highlighting engine
│   │   ├── lexer.go        # Token lexer
│   │   └── languages/      # Dil tanımları
│   │       ├── go.go
│   │       ├── python.go
│   │       └── ...
│   │
│   └── config/
│       ├── config.go       # Konfigürasyon yapısı
│       ├── keybindings.go  # Tuş konfigürasyonu
│       └── defaults.go     # Varsayılan değerler
│
├── pkg/
│   └── version/
│       └── version.go      # Versiyon bilgisi
│
└── configs/
    ├── gesh.yaml           # Örnek konfigürasyon
    └── themes/
        ├── default.yaml
        ├── dark.yaml
        └── light.yaml
```

### 4.3 Veri Yapıları

#### 4.3.1 Gap Buffer

Gap Buffer, metin editörleri için optimize edilmiş bir veri yapısıdır. İmleç pozisyonunda boşluk (gap) tutarak, yerel düzenlemeleri O(1) karmaşıklığında gerçekleştirir.

```go
// GapBuffer metin düzenleme için optimize edilmiş veri yapısı
type GapBuffer struct {
    data     []rune   // Karakter dizisi
    gapStart int      // Gap başlangıç pozisyonu
    gapEnd   int      // Gap bitiş pozisyonu
    
    // Performans metrikleri
    totalChars int    // Toplam karakter sayısı
}

// Örnek: "Hello World" metni, imleç "Hello" ve "World" arasında
//
// data: ['H','e','l','l','o',' ', _, _, _, _, 'W','o','r','l','d']
//                              ^           ^
//                          gapStart     gapEnd
//
// Kullanıcı 'X' yazarsa:
// data: ['H','e','l','l','o',' ','X', _, _, _, 'W','o','r','l','d']
//                                  ^         ^
//                              gapStart   gapEnd
```

**Gap Buffer Operasyonları:**

| Operasyon | Karmaşıklık | Açıklama |
|-----------|-------------|----------|
| Insert | O(1) amortized | Gap pozisyonuna ekleme |
| Delete | O(1) amortized | Gap genişletme |
| Move cursor | O(n) worst case | Gap'i yeni pozisyona taşı |
| Get char at | O(1) | Pozisyondaki karakteri al |

```go
// GapBuffer interface
type GapBuffer interface {
    // Temel operasyonlar
    Insert(r rune)
    InsertString(s string)
    Delete() rune
    DeleteForward() rune
    
    // Navigasyon
    MoveLeft()
    MoveRight()
    MoveTo(pos int)
    
    // Erişim
    Len() int
    String() string
    RuneAt(pos int) rune
    Slice(start, end int) string
    
    // Satır operasyonları
    LineCount() int
    LineStart(line int) int
    LineEnd(line int) int
    CurrentLine() int
    CurrentColumn() int
}
```

#### 4.3.2 Cursor

```go
// Position editördeki bir pozisyonu temsil eder
type Position struct {
    Line   int  // 0-indexed satır numarası
    Column int  // 0-indexed sütun numarası
    Offset int  // Buffer içindeki absolute offset
}

// Cursor imleç durumunu yönetir
type Cursor struct {
    pos       Position
    
    // Preferred column - dikey hareket için
    // Kısa satırdan uzun satıra geçerken hatırlanır
    preferredColumn int
    
    // Selection
    selecting    bool
    selectionStart Position
}
```

#### 4.3.3 Document

```go
// Document açık bir dosyayı temsil eder
type Document struct {
    // Dosya bilgileri
    path     string
    filename string
    
    // İçerik
    buffer   *GapBuffer
    
    // Durum
    modified bool
    readonly bool
    
    // Encoding
    encoding string  // "utf-8", "latin1", etc.
    lineEnding string // "\n", "\r\n", "\r"
    
    // Metadata
    language string  // Syntax highlighting için
    
    // History
    undoStack *HistoryStack
    redoStack *HistoryStack
}
```

#### 4.3.4 History (Undo/Redo)

```go
// EditOperation bir düzenleme işlemini temsil eder
type EditOperation struct {
    Type      OpType    // Insert, Delete, Replace
    Position  int       // İşlem pozisyonu
    Text      string    // Eklenen/silinen metin
    Timestamp time.Time
}

type OpType int

const (
    OpInsert OpType = iota
    OpDelete
    OpReplace
)

// HistoryStack undo/redo için stack yapısı
type HistoryStack struct {
    operations []EditOperation
    maxSize    int
    
    // Gruplandırma için
    groupID    int
    groupOpen  bool
}

// History manager
type History struct {
    undo *HistoryStack
    redo *HistoryStack
    
    // Benzer operasyonları grupla (hızlı yazım için)
    mergeTimeout time.Duration
}
```

#### 4.3.5 Application Model

```go
// Mode editör modunu belirtir
type Mode int

const (
    ModeNormal Mode = iota  // Normal düzenleme
    ModeSearch              // Arama modu
    ModeReplace             // Değiştirme modu
    ModeGoto                // Satıra git modu
    ModeCommand             // Komut modu
    ModeSaveAs              // Farklı kaydet dialog
    ModeConfirm             // Onay dialog
)

// Model ana Bubble Tea model yapısı
type Model struct {
    // Boyutlar
    width  int
    height int
    
    // Document
    doc *Document
    
    // Viewport
    viewport Viewport
    
    // Cursor
    cursor *Cursor
    
    // Mode
    mode Mode
    
    // UI State
    showLineNumbers bool
    showStatusBar   bool
    
    // Input buffers (search, goto, etc.)
    inputBuffer string
    inputPrompt string
    
    // Search state
    searchQuery   string
    searchMatches []Position
    searchIndex   int
    
    // Messages
    statusMessage string
    messageTime   time.Time
    
    // Config
    config *Config
    
    // Clipboard
    clipboard string
}

// Viewport görünür alanı yönetir
type Viewport struct {
    topLine    int  // Görünür ilk satır
    leftColumn int  // Yatay scroll offset
    
    // Görünür alan boyutları (header/statusbar hariç)
    visibleLines   int
    visibleColumns int
}
```

### 4.4 Bubble Tea Akışı

```go
// Init başlangıç komutu
func (m Model) Init() tea.Cmd {
    return tea.Batch(
        tea.EnterAltScreen,
        tea.EnableMouseCellMotion,
        loadFile(m.doc.path),
    )
}

// Update mesaj işleyici
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        return m.handleKeyPress(msg)
    case tea.MouseMsg:
        return m.handleMouse(msg)
    case tea.WindowSizeMsg:
        return m.handleResize(msg)
    case fileLoadedMsg:
        return m.handleFileLoaded(msg)
    case fileSavedMsg:
        return m.handleFileSaved(msg)
    case errorMsg:
        return m.handleError(msg)
    }
    return m, nil
}

// View render
func (m Model) View() string {
    var b strings.Builder
    
    // Header
    b.WriteString(m.renderHeader())
    b.WriteString("\n")
    
    // Editor area
    b.WriteString(m.renderEditor())
    
    // Status bar
    b.WriteString(m.renderStatusBar())
    
    // Input prompt (if active)
    if m.mode != ModeNormal {
        b.WriteString(m.renderPrompt())
    }
    
    return b.String()
}
```

---

## 5. Kullanıcı Arayüzü

### 5.1 Layout

```
┌─────────────────────────────────────────────────────────────────────┐
│ GESH │ filename.go                              [Modified] UTF-8 LF │ ← Header (1 satır)
├─────────────────────────────────────────────────────────────────────┤
│   1 │ package main                                                  │
│   2 │                                                               │
│   3 │ import (                                                      │
│   4 │     "fmt"                                                     │
│   5 │     "os"                                                      │
│   6 │ )                                                             │
│   7 │                                                               │ ← Editor Area
│   8 │ func main() {                                                 │   (height - 3 satır)
│   9 │     fmt.Println("Hello, World!")█                             │
│  10 │ }                                                             │
│  11 │                                                               │
│  12 │                                                               │
│     │                                                               │
│     │                                                               │
├─────────────────────────────────────────────────────────────────────┤
│ Ln 9, Col 35 │ 12 lines │ Go                         INS           │ ← Status Bar (1 satır)
├─────────────────────────────────────────────────────────────────────┤
│ ^S Save  ^O Open  ^X Exit  ^W Search  ^G Goto  ^K Cut  ^U Paste    │ ← Help Bar (1 satır)
└─────────────────────────────────────────────────────────────────────┘
```

### 5.2 Bileşen Detayları

#### 5.2.1 Header

```
┌─────────────────────────────────────────────────────────────────────┐
│ GESH │ [path/to/]filename.ext              [Modified] [RO] UTF-8 LF │
└─────────────────────────────────────────────────────────────────────┘
  │      │                                    │        │    │     │
  │      │                                    │        │    │     └─ Line ending
  │      │                                    │        │    └─ Encoding
  │      │                                    │        └─ Read-only flag
  │      │                                    └─ Modified flag
  │      └─ Dosya yolu (uzunsa kısaltılır)
  └─ Logo
```

#### 5.2.2 Editor Area

```
┌─────────────────────────────────────────────────────────────────────┐
│   1 │ package main                                                  │
│   2 │                                                               │
│   3 │ import "fmt"                                                  │
│   4 │                                                               │
│   5 │ func main() {                                                 │
│ → 6 │     fmt.Println("Hello")█                                     │ ← Current line marker
│   7 │ }                                                             │
│   ~ │                                                               │ ← Empty line indicator
└─────────────────────────────────────────────────────────────────────┘
  │     │
  │     └─ Metin içeriği (syntax highlighted)
  └─ Satır numaraları (opsiyonel, toggle ile)
```

#### 5.2.3 Status Bar

```
┌─────────────────────────────────────────────────────────────────────┐
│ Ln 9, Col 35 │ 156 lines │ 4.2 KB │ Go                    INS      │
└─────────────────────────────────────────────────────────────────────┘
  │              │            │        │                      │
  │              │            │        │                      └─ Insert/Overwrite mode
  │              │            │        └─ Detected language
  │              │            └─ File size
  │              └─ Total line count
  └─ Current position (line, column)
```

#### 5.2.4 Help Bar

```
┌─────────────────────────────────────────────────────────────────────┐
│ ^S Save  ^O Open  ^X Exit  ^W Search  ^G Goto  ^K Cut  ^U Paste    │
└─────────────────────────────────────────────────────────────────────┘

Modlara göre değişir:
- Normal: Temel kısayollar
- Search: ^W Next  ^Q Cancel  ^R Replace
- Goto:   Enter Confirm  ^Q Cancel
```

### 5.3 Dialog Tasarımları

#### 5.3.1 Search Dialog

```
┌─────────────────────────────────────────────────────────────────────┐
│ Search: pattern█                                     [1/5 matches]  │
└─────────────────────────────────────────────────────────────────────┘
```

#### 5.3.2 Goto Dialog

```
┌─────────────────────────────────────────────────────────────────────┐
│ Go to line: 42█                                       [1-156]       │
└─────────────────────────────────────────────────────────────────────┘
```

#### 5.3.3 Save Confirmation

```
┌────────────────────────────────────────┐
│  Save changes to "filename.go"?        │
│                                        │
│  [Y]es    [N]o    [C]ancel             │
└────────────────────────────────────────┘
```

#### 5.3.4 File Open Dialog

```
┌─ Open File ─────────────────────────────────────────────────────────┐
│ Path: /home/user/projects/█                                         │
├─────────────────────────────────────────────────────────────────────┤
│ > ..                                                                │
│   src/                                                              │
│   main.go                                                           │
│   go.mod                                                            │
│   README.md                                                         │
├─────────────────────────────────────────────────────────────────────┤
│ ↑↓ Navigate  Enter Select  ^Q Cancel  Tab Autocomplete              │
└─────────────────────────────────────────────────────────────────────┘
```

### 5.4 Renk Şeması

#### 5.4.1 Default Dark Theme

```go
var DarkTheme = Theme{
    Name: "dark",
    
    // UI Colors
    Background:     "#1a1a2e",
    Foreground:     "#eaeaea",
    
    // Header
    HeaderBg:       "#16213e",
    HeaderFg:       "#e94560",
    HeaderAccent:   "#0f3460",
    
    // Editor
    EditorBg:       "#1a1a2e",
    EditorFg:       "#eaeaea",
    LineNumberFg:   "#4a4a6a",
    CurrentLineBg:  "#252545",
    SelectionBg:    "#3a3a5a",
    
    // Status Bar
    StatusBg:       "#0f3460",
    StatusFg:       "#eaeaea",
    
    // Help Bar
    HelpBg:         "#16213e",
    HelpKey:        "#e94560",
    HelpText:       "#a0a0c0",
    
    // Syntax (base)
    Keyword:        "#e94560",
    String:         "#7ec8e3",
    Number:         "#f9a825",
    Comment:        "#6a6a8a",
    Function:       "#50fa7b",
    Type:           "#bd93f9",
    Operator:       "#ff79c6",
    
    // Search
    SearchMatch:    "#ffff00",
    SearchCurrent:  "#ff8800",
}
```

#### 5.4.2 Light Theme

```go
var LightTheme = Theme{
    Name: "light",
    
    Background:     "#ffffff",
    Foreground:     "#1a1a1a",
    
    HeaderBg:       "#f0f0f0",
    HeaderFg:       "#d63031",
    
    EditorBg:       "#ffffff",
    EditorFg:       "#1a1a1a",
    LineNumberFg:   "#b0b0b0",
    CurrentLineBg:  "#f5f5f5",
    SelectionBg:    "#b4d7ff",
    
    StatusBg:       "#e0e0e0",
    StatusFg:       "#1a1a1a",
    
    // ... syntax colors
}
```

---

## 6. CLI Arayüzü

### 6.1 Kullanım

```bash
# Temel kullanım
gesh [options] [file]

# Örnekler
gesh                      # Yeni boş dosya
gesh main.go              # Dosya aç
gesh -r config.yaml       # Read-only aç
gesh +42 main.go          # 42. satırda aç
gesh +42:10 main.go       # 42. satır, 10. sütunda aç
```

### 6.2 Seçenekler

| Flag | Uzun | Açıklama |
|------|------|----------|
| `-h` | `--help` | Yardım göster |
| `-v` | `--version` | Versiyon göster |
| `-r` | `--readonly` | Salt okunur aç |
| `-n` | `--norc` | Config dosyasını yüklemeden aç |
| `-t` | `--theme` | Tema seç (dark/light) |
| `+N` | | N. satırda aç |
| `+N:M` | | N. satır, M. sütunda aç |
| | `--no-syntax` | Syntax highlighting kapalı |
| | `--no-line-numbers` | Satır numaraları kapalı |

### 6.3 Exit Codes

| Code | Anlam |
|------|-------|
| 0 | Başarılı |
| 1 | Genel hata |
| 2 | Dosya bulunamadı |
| 3 | İzin hatası |
| 4 | Geçersiz argüman |

---

## 7. Konfigürasyon

### 7.1 Dosya Konumu

```
Linux/macOS: ~/.config/gesh/gesh.yaml
Windows:     %APPDATA%\gesh\gesh.yaml
```

### 7.2 Konfigürasyon Dosyası

```yaml
# gesh.yaml - Gesh Text Editor Configuration

editor:
  # Tab genişliği (spaces)
  tab_width: 4
  
  # Tab yerine space kullan
  expand_tabs: true
  
  # Otomatik girinti
  auto_indent: true
  
  # Satır numaralarını göster
  line_numbers: true
  
  # Satır kaydırma (wrap)
  word_wrap: false
  
  # İmleç stili: "block", "bar", "underline"
  cursor_style: block
  
  # Scroll padding (imleç kenara yaklaşınca kaç satır margin)
  scroll_padding: 5

ui:
  # Tema: "dark", "light", veya custom tema dosyası yolu
  theme: dark
  
  # Status bar göster
  show_status_bar: true
  
  # Help bar göster
  show_help_bar: true
  
  # Animasyonlar (örn. smooth scroll)
  animations: true

file:
  # Varsayılan encoding
  default_encoding: utf-8
  
  # Varsayılan satır sonu: "lf", "crlf", "cr"
  default_line_ending: lf
  
  # Dosya sonunda newline ekle
  final_newline: true
  
  # Trailing whitespace temizle
  trim_trailing_whitespace: true
  
  # Otomatik kaydetme (saniye, 0 = kapalı)
  auto_save: 0
  
  # Backup dosyası oluştur
  create_backup: false

search:
  # Büyük/küçük harf duyarlılığı
  case_sensitive: false
  
  # Regex varsayılan olarak açık
  use_regex: false
  
  # Arama sırasında highlight
  highlight_matches: true

syntax:
  # Syntax highlighting aktif
  enabled: true
  
  # Varsayılan dil (boşsa otomatik algıla)
  default_language: ""

history:
  # Maksimum undo adımı
  max_undo: 1000
  
  # Undo gruplandırma timeout (ms)
  group_timeout: 500

keybindings:
  # Custom key bindings (varsayılanları override eder)
  # Format: action: key
  # Örnek:
  # save: ctrl+s
  # quit: ctrl+q
```

### 7.3 Tema Dosyası

```yaml
# ~/.config/gesh/themes/custom.yaml

name: custom

colors:
  background: "#1e1e2e"
  foreground: "#cdd6f4"
  
  ui:
    header_bg: "#181825"
    header_fg: "#f38ba8"
    status_bg: "#313244"
    status_fg: "#cdd6f4"
    line_number: "#6c7086"
    current_line: "#313244"
    selection: "#45475a"
  
  syntax:
    keyword: "#cba6f7"
    string: "#a6e3a1"
    number: "#fab387"
    comment: "#6c7086"
    function: "#89b4fa"
    type: "#f9e2af"
    operator: "#89dceb"
```

---

## 8. Performans Gereksinimleri

### 8.1 Hedef Metrikler

| Metrik | Hedef | Ölçüm |
|--------|-------|-------|
| Başlangıç süresi | < 50ms | Boş dosya açılışı |
| Bellek (boş) | < 5MB | RSS |
| Bellek (1MB dosya) | < 15MB | RSS |
| Bellek (10MB dosya) | < 50MB | RSS |
| Keystroke latency | < 16ms | Input → Render |
| Scroll FPS | 60fps | Smooth scroll |
| File load (1MB) | < 100ms | Disk → Render |
| File save (1MB) | < 50ms | Buffer → Disk |

### 8.2 Optimizasyon Stratejileri

#### 8.2.1 Lazy Loading

```go
// Büyük dosyalar için lazy loading
type LazyDocument struct {
    file     *os.File
    chunks   map[int]*Chunk  // Yüklenen chunk'lar
    chunkSize int            // Chunk boyutu (örn. 64KB)
}

// Sadece görünür ve yakın chunk'ları bellekte tut
func (d *LazyDocument) EnsureLoaded(startLine, endLine int) {
    // Gerekli chunk'ları yükle
    // Uzak chunk'ları unload et
}
```

#### 8.2.2 Incremental Rendering

```go
// Sadece değişen satırları render et
type RenderCache struct {
    lines      []string  // Render edilmiş satırlar
    dirtyLines []bool    // Değişen satırlar
}

func (c *RenderCache) InvalidateLine(n int) {
    c.dirtyLines[n] = true
}

func (c *RenderCache) Render(doc *Document) string {
    for i, dirty := range c.dirtyLines {
        if dirty {
            c.lines[i] = renderLine(doc, i)
            c.dirtyLines[i] = false
        }
    }
    return strings.Join(c.lines, "\n")
}
```

#### 8.2.3 Syntax Highlighting Cache

```go
// Token cache - sadece değişen satırları yeniden tokenize et
type SyntaxCache struct {
    tokens     [][]Token
    states     []LexerState  // Her satır sonundaki lexer durumu
    dirtyFrom  int           // Bu satırdan itibaren dirty
}
```

---

## 9. Test Stratejisi

### 9.1 Test Kategorileri

#### 9.1.1 Unit Tests

```go
// buffer/gap_test.go
func TestGapBuffer_Insert(t *testing.T) {
    b := NewGapBuffer()
    b.InsertString("Hello")
    assert.Equal(t, "Hello", b.String())
    assert.Equal(t, 5, b.Len())
}

func TestGapBuffer_Delete(t *testing.T) {
    b := NewGapBuffer()
    b.InsertString("Hello")
    b.Delete()
    assert.Equal(t, "Hell", b.String())
}

func TestGapBuffer_MoveCursor(t *testing.T) {
    b := NewGapBuffer()
    b.InsertString("Hello World")
    b.MoveTo(5)
    b.Insert('!')
    assert.Equal(t, "Hello! World", b.String())
}
```

#### 9.1.2 Integration Tests

```go
// integration/editor_test.go
func TestEditor_OpenSaveFile(t *testing.T) {
    // Temp dosya oluştur
    // Editör ile aç
    // Değişiklik yap
    // Kaydet
    // Dosyayı tekrar oku ve doğrula
}

func TestEditor_UndoRedo(t *testing.T) {
    // Birden fazla edit yap
    // Undo ile geri al
    // Redo ile yinele
    // State'i doğrula
}
```

#### 9.1.3 Benchmark Tests

```go
// buffer/benchmark_test.go
func BenchmarkGapBuffer_Insert(b *testing.B) {
    buf := NewGapBuffer()
    for i := 0; i < b.N; i++ {
        buf.Insert('x')
    }
}

func BenchmarkGapBuffer_RandomAccess(b *testing.B) {
    buf := NewGapBuffer()
    buf.InsertString(strings.Repeat("x", 100000))
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        pos := rand.Intn(100000)
        buf.MoveTo(pos)
        buf.Insert('y')
        buf.Delete()
    }
}
```

### 9.2 Test Coverage Hedefi

| Paket | Hedef Coverage |
|-------|----------------|
| buffer | 90%+ |
| editor | 85%+ |
| file | 80%+ |
| ui | 70%+ |
| **Toplam** | **80%+** |

---

## 10. Geliştirme Fazları

### Faz 1: MVP (2-3 hafta)

**Hedef:** Temel metin düzenleme

| Görev | Süre | Öncelik |
|-------|------|---------|
| Proje yapısı ve boilerplate | 1 gün | P0 |
| Gap buffer implementasyonu | 2 gün | P0 |
| Temel Bubble Tea model | 2 gün | P0 |
| Cursor ve navigasyon | 2 gün | P0 |
| Basit UI (header, editor, status) | 2 gün | P0 |
| Dosya açma/kaydetme | 2 gün | P0 |
| Temel keybindings | 1 gün | P0 |
| Undo/Redo | 2 gün | P0 |

**Çıktı:** Çalışan basit metin editörü

### Faz 2: Temel Özellikler (2 hafta)

**Hedef:** Nano eşdeğeri fonksiyonellik

| Görev | Süre | Öncelik |
|-------|------|---------|
| Arama fonksiyonu | 2 gün | P0 |
| Değiştirme fonksiyonu | 1 gün | P0 |
| Seçim (selection) | 2 gün | P0 |
| Clipboard (cut/copy/paste) | 1 gün | P0 |
| Satıra git (goto) | 0.5 gün | P0 |
| Dosya dialogs | 2 gün | P1 |
| Help bar | 0.5 gün | P1 |
| Scroll optimizasyonu | 1 gün | P1 |

**Çıktı:** Nano benzeri tam fonksiyonel editör

### Faz 3: Polish & UX (1-2 hafta)

**Hedef:** Kullanıcı deneyimi iyileştirmeleri

| Görev | Süre | Öncelik |
|-------|------|---------|
| Tema sistemi | 2 gün | P1 |
| Konfigürasyon dosyası | 2 gün | P1 |
| CLI argümanları | 1 gün | P1 |
| Satır numaraları toggle | 0.5 gün | P1 |
| Error handling & messages | 1 gün | P1 |
| Mouse desteği | 2 gün | P2 |
| Word wrap | 1 gün | P2 |

**Çıktı:** Kullanıma hazır v1.0

### Faz 4: Syntax Highlighting (2 hafta)

**Hedef:** Temel syntax highlighting

| Görev | Süre | Öncelik |
|-------|------|---------|
| Highlighter engine | 3 gün | P1 |
| Go syntax | 1 gün | P1 |
| Python syntax | 1 gün | P1 |
| JavaScript/JSON syntax | 1 gün | P1 |
| Diğer diller | 2 gün | P2 |
| Performance optimizasyonu | 2 gün | P1 |

**Çıktı:** Syntax highlighted editör v1.1

### Faz 5: Gelişmiş Özellikler (Opsiyonel)

| Görev | Öncelik |
|-------|---------|
| Çoklu buffer/tab | P2 |
| Split view | P2 |
| Makro sistemi | P2 |
| Plugin API | P3 |
| LSP entegrasyonu | P3 |

---

## 11. Bağımlılıklar

### 11.1 Doğrudan Bağımlılıklar

```go
require (
    github.com/charmbracelet/bubbletea v0.25.0
    github.com/charmbracelet/lipgloss v0.9.1
    github.com/charmbracelet/bubbles v0.17.1
    gopkg.in/yaml.v3 v3.0.1
    github.com/spf13/cobra v1.8.0
    github.com/atotto/clipboard v0.1.4
)
```

### 11.2 Geliştirme Bağımlılıkları

```go
require (
    github.com/stretchr/testify v1.8.4  // Testing
)
```

---

## 12. Dokümantasyon

### 12.1 Kullanıcı Dokümantasyonu

- README.md - Hızlı başlangıç
- INSTALL.md - Kurulum rehberi
- KEYBINDINGS.md - Tüm kısayollar
- CONFIG.md - Konfigürasyon rehberi
- THEMES.md - Tema oluşturma

### 12.2 Geliştirici Dokümantasyonu

- CONTRIBUTING.md - Katkı rehberi
- ARCHITECTURE.md - Mimari dokümantasyonu
- API.md - Internal API dokümantasyonu

---

## 13. Riskler ve Çözümler

| Risk | Olasılık | Etki | Çözüm |
|------|----------|------|-------|
| Performans sorunları büyük dosyalarda | Orta | Yüksek | Lazy loading, chunk-based rendering |
| Cross-platform terminal uyumsuzlukları | Orta | Orta | tcell/bubbletea abstraction, test matrix |
| Kompleks keybinding çakışmaları | Düşük | Orta | Configurable keybindings |
| Encoding sorunları | Orta | Orta | Encoding detection, fallback to UTF-8 |

---

## 14. Başarı Kriterleri

### 14.1 MVP Başarı Kriterleri

- [ ] Dosya açma/kaydetme çalışıyor
- [ ] Temel metin düzenleme (insert, delete)
- [ ] Cursor navigasyonu sorunsuz
- [ ] Undo/Redo çalışıyor
- [ ] 50ms altında başlangıç süresi
- [ ] 10MB RAM altında bellek kullanımı (boş dosya)

### 14.2 v1.0 Başarı Kriterleri

- [ ] Tüm nano temel özelliklerine sahip
- [ ] Arama/değiştirme çalışıyor
- [ ] Konfigürasyon dosyası desteği
- [ ] Dark/Light tema
- [ ] Cross-platform çalışıyor (Linux, macOS, Windows)
- [ ] 80%+ test coverage
- [ ] Kullanıcı dokümantasyonu hazır

---

## 15. Sonraki Adımlar

1. **Proje başlatma**
   - [ ] Repository oluştur
   - [ ] Go module init
   - [ ] Temel dizin yapısı

2. **MVP geliştirme başlangıcı**
   - [ ] Gap buffer implementasyonu
   - [ ] Bubble Tea skeleton

3. **CI/CD kurulumu**
   - [ ] GitHub Actions
   - [ ] Test automation
   - [ ] Release pipeline

---

*Bu PRD, geliştirme sürecinde güncellenecektir.*

**Son güncelleme:** Aralık 2024
