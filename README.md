## summaly-go
Misskey（とその派生）で使えるGoで書かれた軽量なsummaly-proxy  
中身は[別ライブラリ（summergo）として](https://github.com/nexryai/summergo)公開しています


### run with docker
```
services:
  summaly:
    image: docker.io/nexryai/summaly-go:latest
    restart: always
    ports:
      - 127.0.0.1:3000:3000
```