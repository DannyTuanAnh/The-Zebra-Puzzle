# The Zebra Puzzle (Einstein's Riddle)

Đây là một chương trình giải quyết **The Zebra Puzzle** - bài toán logic nổi tiếng được gọi là **Einstein's Riddle**.

## Mô tả bài toán

Bài toán này liên quan đến việc xác định ai sở hữu con ngựa vằn và ai uống nước, dựa trên một tập hợp các gợi ý logic.

**5 ngôi nhà** với các đặc điểm:

- **Màu sắc**: Vàng, Xanh lam, Đỏ, Xanh lá, Ngà
- **Quốc tịch**: Anh, Tây Ban Nha, Ukraina, Na Uy, Nhật Bản
- **Thú cưng**: Ngựa, Chó, Ốc sên, Cáo, Ngựa vằn
- **Đồ uống**: Cà phê, Trà, Sữa, Nước cam, Nước
- **Thuốc lá**: Old Gold, Kools, Chesterfields, Parliaments, Lucky Strike

## Hướng dẫn chạy chương trình

### Yêu cầu

- Go 1.16 hoặc cao hơn nếu chạy trực tiếp từ file Go

### Cách 1: Chạy file thực thi

#### Trên Windows

1. Mở Command Prompt hoặc PowerShell
2. Di chuyển đến thư mục chứa file thực thi
3. Chạy lệnh:

```bash
.\the-zebra-puzzle.exe
```

Hoặc đơn giản: Double-click vào file `the-zebra-puzzle.exe`

#### Trên Mac

1. Mở Terminal
2. Di chuyển đến thư mục chứa file thực thi (thay thế bằng đường dẫn thực tế):

```bash
cd ~/path/to/The\ Zebra\ Puzzle
```

Ví dụ: Nếu file nằm trong thư mục Downloads:

```bash
cd ~/Downloads/The\ Zebra\ Puzzle
```

3. Cấp quyền chạy file (chỉ cần làm một lần):

```bash
chmod +x the-zebra-puzzle
```

4. Chạy chương trình:

```bash
./the-zebra-puzzle
```

### Cách 2: Chạy trực tiếp từ Go

```bash
go run .
```

hoặc

```bash
go run main.go
```

## Kết quả đầu ra

Chương trình sẽ:

1. **Áp dụng AC3** - Giảm miền giá trị của các biến bằng cách loại bỏ các giá trị không tương thích
2. **Sử dụng Backtracking Search với MAC + MRV + Degree + LCV** - Tìm giải pháp thỏa mãn tất cả các ràng buộc
3. **In kết quả** - Hiển thị các biến được sắp xếp theo thứ tự vị trí (1-5)

### Ví dụ kết quả:

```
Domain of each variable after using AC3:
[In ra miền giá trị của mỗi biến sau khi được cắt giảm]

Using MAC in backtrack
Result:
Kools:1 Yellow:1 Norwegian:1 Water:1 Fox:1
Tea:2 Ukrainian:2 Horse:2 Chesterfields:2 Blue:2
Britt:3 Red:3 Snails:3 OldGold:3 Milk:3
Spaniard:4 OrangeJuice:4 LuckyStrike:4 Dog:4 Ivory:4
Japanese:5 Coffee:5 Zebra:5 Parliaments:5 Green:5

The owner of the zebra lives in house 5.
The owner who drinks water lives in house 1.
```
