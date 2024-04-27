# go-zero-tts
基于微软edge大声朗读接口开发的语音合成服务, 后端 go-zero, 前端 vuetify

## 启动
### 前端
```shell
cd web
npm i
npm run dev
```

### 后端
```shell
cd server
go mod tidy
go run main.go
```

## 访问
访问 http://localhost:3000/

## 界面
![image](https://github.com/wujunwei928/go-zero-tts/assets/3396697/aa42ae4b-1ea5-417a-83df-cd73ef6f65c6)


## 使用说明
* 支持切换地区, 选择对应地区的语音
* 支持设置音量, 音高, 语速
* 支持在线试听, 支持下载音频(播放器上点右键)
