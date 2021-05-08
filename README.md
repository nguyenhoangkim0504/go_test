# App test

# URL test
GET: http://{host}/positions?room={room}
POST:
```
curl --location --request POST 'http://{host}/positions?room=1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "seats": [
        {
            "row": 0,
            "col": 14
        }
    ]
}'
```

# Test result
- Hiển thị các ghế còn trống khi cần đặt n(seat) theo group
- Hiển thị các ghế còn trống khi cần đặt 1 seat
- Cho phép đặt 1 seat
- Cho phép đặt nhiều seat 1 lần
- Cách test:
    - Dùng API get để lấy thông tin suitable empty seats
    - Dùng API post để đăng ký seat (theo các thông số row, col được lấy từ API trên)

# Note
- DB seats: quản lý thông tin của 1 seat
- DB room: chứa info của 1 room (room number, rows, cols)
- Không có API update
