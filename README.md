# dhsp-dkhp

### Cùng vượt cạn rời ĐHSP nào

Vì thời gian mở đăng ký học phần của trường ĐHSP TP.HCM quá tuỳ hứng, lúc được lúc không khiến cần có một giải pháp cấp thiết để giúp chúng ta tránh việc đăng ký trễ, lỡ môn và ra trường chậm hơn bạn bè. Liên tục refresh trang đkhp thì cũng là một giải pháp đó nhưng nó quá mất thời gian và nghiệp dư, từ đó đã tạo tiền đề ra đời cho app này, được tạo nên để giải cứu những sinh viên chuyên trễ đăng ký môn như chính bản thân tôi.

## Quick start

```sh
# Pull the source code *for unix env only*
git clone https://github.com/phuwn/dhsp-dkhp.git
```

Create a bot on telegram app. Check out [here](https://core.telegram.org/bots#6-botfather) for more information.

Create a `credential.json` file including:

- mssv(string): student id
- password(string): account password
- bot_id(string): telegram bot_id
- chat_id(int): telegram chat_id of you and your bot

```sh
# Build executable file
make

# Run the app
bin/dkhp
```

And then we wait. When the dkhp session is ready, your bot will send you a message. Good luck!!!
