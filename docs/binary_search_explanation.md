# Giải Thích Chi Tiết Hàm Binary Search trong Go

## Tổng Quan

Hàm [`binarySearch`](main.go:5) trong file [`main.go`](main.go:1) là một triển khai kinh điển của thuật toán tìm kiếm nhị phân (binary search) với độ phức tạp O(log n).

## Code Hàm

```go
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2 // 0 + (9-0)/2 = 4

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
```

## Phân Tích Chi Tiết Từng Dòng

### Dòng 5: Khai báo hàm
```go
func binarySearch(nums []int, target int) int {
```
- `func binarySearch`: Khai báo hàm tên là `binarySearch`
- `nums []int`: Tham số đầu vào là một slice các số nguyên đã được sắp xếp
- `target int`: Số nguyên cần tìm
- `int`: Kiểu trả về là vị trí (index) của `target` trong `nums`

### Dòng 6: Khởi tạo biến
```go
left, right := 0, len(nums)-1
```
- `left := 0`: Biên trái của không gian tìm kiếm, bắt đầu từ index 0
- `right := len(nums)-1`: Biên phải của không gian tìm kiếm, kết thúc tại index cuối cùng
- Với ví dụ trong `main()`: `left = 0`, `right = 9` (vì có 10 phần tử từ index 0-9)

### Dòng 8: Bắt đầu vòng lặp
```go
for left <= right {
```
- Điều kiện lặp: tiếp tục khi `left` vẫn nhỏ hơn hoặc bằng `right`
- Khi `left > right` nghĩa là đã tìm hết không gian mà không tìm thấy `target`

### Dòng 9: Tính vị trí giữa
```go
mid := left + (right-left)/2 // 0 + (9-0)/2 = 4
```
- `mid`: Vị trí giữa của không gian tìm kiếm hiện tại
- Công thức `left + (right-left)/2` tránh trường hợp overflow khi `left` và `right` quá lớn
- Với ví dụ ban đầu: `mid = 0 + (9-0)/2 = 4`

### Dòng 11-12: Trường hợp tìm thấy
```go
if nums[mid] == target {
	return mid
```
- Nếu phần tử tại vị trí `mid` bằng `target`, trả về vị trí `mid`
- Đây là trường hợp thành công, kết thúc hàm sớm

### Dòng 13-14: Target ở nửa phải
```go
else if nums[mid] < target {
	left = mid + 1
```
- Nếu `nums[mid] < target`, nghĩa là `target` nằm ở nửa phải của `mid`
- Cập nhật `left = mid + 1` để loại bỏ nửa trái (bao gồm cả `mid`)
- Không gian tìm kiếm giảm đi một nửa

### Dòng 15-17: Target ở nửa trái
```go
else {
	right = mid - 1
```
- Nếu `nums[mid] > target`, nghĩa là `target` nằm ở nửa trái của `mid`
- Cập nhật `right = mid - 1` để loại bỏ nửa phải (bao gồm cả `mid`)
- Không gian tìm kiếm giảm đi một nửa

### Dòng 20: Không tìm thấy
```go
return -1
```
- Nếu vòng lặp kết thúc mà không tìm thấy `target`, trả về -1
- -1 là giá trị quy ước cho "không tìm thấy"

## Ví Dụ Minh Họa

Với input từ hàm `main()`:
- `nums = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]`
- `target = 7`

### Lần lặp 1:
- `left = 0`, `right = 9`
- `mid = 0 + (9-0)/2 = 4`
- `nums[4] = 5`
- `5 < 7` → `left = 4 + 1 = 5`

### Lần lặp 2:
- `left = 5`, `right = 9`
- `mid = 5 + (9-5)/2 = 7`
- `nums[7] = 8`
- `8 > 7` → `right = 7 - 1 = 6`

### Lần lặp 3:
- `left = 5`, `right = 6`
- `mid = 5 + (6-5)/2 = 5`
- `nums[5] = 6`
- `6 < 7` → `left = 5 + 1 = 6`

### Lần lặp 4:
- `left = 6`, `right = 6`
- `mid = 6 + (6-6)/2 = 6`
- `nums[6] = 7`
- `7 == 7` → `return 6`

Kết quả: Hàm trả về 6, đúng là vị trí của số 7 trong mảng.

## Độ Phức Tạp Thời Gian

- **Best Case**: O(1) - khi `target` nằm ngay ở vị trí `mid` đầu tiên
- **Worst Case**: O(log n) - khi phải chia không gian tìm kiếm đến khi còn 1 phần tử
- **Average Case**: O(log n)

## Độ Phức Tạp Không Gian

- **O(1)** - chỉ sử dụng một số biến cố định (`left`, `right`, `mid`)

## Điều Kiện Tiên Quyết

Thuật toán này chỉ hoạt động đúng khi:
1. Mảng `nums` đã được sắp xếp (tăng dần trong trường hợp này)
2. Các phần tử trong mảng là duy nhất (nếu có trùng lặp, sẽ trả về một trong các vị trí)

## Lưu Ý Quan Trọng

1. **Tránh overflow**: Công thức `left + (right-left)/2` an toàn hơn `(left+right)/2`
2. **Index trong Go**: Go sử dụng index bắt đầu từ 0, khác với một số ngôn ngữ khác
3. **Edge cases**: Hàm xử lý đúng các trường hợp:
   - Mảng rỗng: `len(nums) = 0` → `right = -1` → vòng lặp không chạy → trả về -1
   - Target nhỏ hơn phần tử nhỏ nhất
   - Target lớn hơn phần tử lớn nhất
   - Target không tồn tại trong mảng

## Biến Thể

Có thể sửa đổi hàm để:
- Tìm vị trí đầu tiên/cuối cùng của phần tử trùng lặp
- Làm việc với các kiểu dữ liệu khác (string, float, struct)
- Sắp xếp giảm dần
