# API v1

| Method | URL            | Info     |
| :----- | -------------- | -------- |
| `GET`  | `/users/login` | 登录跳转 |
| `GET`  | `/users/verify` | 登录认证state并重定向 |
|        |                |          |
|        |                |          |

## `/Users/Verify`

```json
{
    "code": string,
    "state": string,
    "redirectUrl": string
}
```
