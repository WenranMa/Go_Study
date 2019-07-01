

web或移动端 上传视频 上传服务器 存储 下行服务器 下行视频 web或移动端

前后端分离
最著名的是MVC 


## 接口设计：
### vlog 播放接口
    接口描述：获取一个vlog视频
    请求方法：Get
    请求地址：http://domain/video/{vidoe.mp4}
    状态码：200

### vlog 上传接口
    接口描述：上传一个vlog视频
    请求方法：Post multipart/form-data
    请求地址：http://domain/api/upload
    请求参数：参数名uploadFile 类型file
    返回状态：200成功 500失败

### mine 列表接口
    接口描述：查看vlog列表
    请求方法：Get
    请求地址：http://domain/api/list
    返回状态：200成功 500失败
    返回参数：
        {
            "http://domain/static/video1.mp4",
            "http://domain/static/video2.mp4",
            ...
        }

## 代码
    net/http包

    客户端：
    Client，Response

    服务端：
    ResponseWriter
    Handler/HandlerFunc
    Server
    ServerMux

    公用：
    Header
    Request
    Cookie


