# Documents of tinyurl

## 产品功能
1. 用户输入网址，返回对应的短网址。本质是返回一个简短、不重复的字符串
2. 用户输入短网址，跳转到对应的真实网址

## 前端
TODO

## 后端

### API
1. 生成短网址 POST /v1/url/{origin_url}/shorten
payload:
```json
{
    "origin_url": "http://tinyurl.tk/abc/def/?name=test&age=12"
}
```
2. 解析短网址 GET /n/{short_url}
response:
3. 获取短网址对应域名 GET /v1/url/{short_url}/parse
response:
```json
{
    "origin_url": "http://tinyurl.tk/abc/def/?name=test&age=12"
}
```
origin_url为空，前端不做操作，不为空，前端做跳转，后端不再做跳转
4. 删除短网址 DELETE /v1/url/{short_url}/delete