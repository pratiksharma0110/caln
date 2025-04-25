# caln 🗓️ - Nepali Calendar CLI

`caln` is a simple CLI tool written in Go that displays the Nepali calendar based on the current system date. It aims to provide a lightweight and offline-friendly way to access Bikram Sambat (BS) calendar data directly from the terminal.

---

## 📦 Features

- Shows the **current month** in the Nepali calendar
- Parses JSON data files for date mapping
- Modular structure for easy maintenance

---

## 🚀 Getting Started

### Build

```bash
go build -o caln
```

### Run

```bash
./caln
```

This will display the current Nepali calendar month based on today's date.

---

## 📁 Project Structure

```
.
├── calendar/
│   ├── calendar.go     # Core calendar logic
│   └── loader.go       # Loads data from JSON
├── data/
│   └── 2082.json       # Nepali calendar data
├── main.go             # Entry point
├── go.mod              # Go module file
└── README.md
```

---

## 🌟 Future Updates

- [ ] Query a **specific month and year** (e.g. `caln 2082 5`)
- [ ] Highlight current date in the month
---

## 📚 Data Source

Calendar data is parsed from static JSON files stored in the `data/` directory. Each file corresponds to a specific year in the Bikram Sambat calendar.

---

## 🛠️ Contribution

Pull requests are welcome! If you have an idea, improvement, or fix, feel free to open an issue or PR.

---



## 📦 Installation (Optional)

If you want to run `caln` from anywhere in the terminal like a regular command, you can move the compiled binary to `/usr/bin`.

### Step 1: Build the binary

```bash
go build -o caln
```

### Step 2: Move it to `/usr/bin` (requires sudo)

```bash
sudo mv caln /usr/bin/
```

Now you can just run:

```bash
caln
```

from any directory 

---

## 📁 Keeping the Project in `~/Desktop/proj`

Right now, you can keep the project in `~/Desktop/proj` and just recompile it anytime you make changes.

```bash
cd ~/Desktop/proj/caln
go build -o caln
./caln
```




## 🙏 Credits

Special thanks to [sajanm/nepali-lunar-calendar-events](https://github.com/sajanm/nepali-lunar-calendar-events) for the base calendar data used in the `data/` directory.

