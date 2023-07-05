# 腾讯云 “E证通” 中间件
当前版本：`v1.0.1` 作者：Prk

腾讯云官网给的代码示例中，仅有 Go 和 Java 两种语言的。对于 Node.js 或 PHP 等语言来说，想要使用相同的方法较为复杂。

比如使用 PHP，就要引入比如 GMSSL 和 PHPKMS 的类库（因为它们支持 SM2、SM4 两种国密算法）。这显然是一件麻烦的事情，我还是打算使用 Go 语言做一个中间件。

如果你也觉得中间件方便那就一起用吧！


## 如何使用

其实很简单，只需要使用 `GET` 方法请求传 Query 参数即可！

```url
/?key=xxx&des_key=xxx&user_info=xxx
```

其中：`key` 为十六进制的私钥证书。`des_key` 和 `user_info` 为接口传回内容（详情请参考“[腾讯云官方文档](https://cloud.tencent.com/document/api/1007/54090)”）。


## 值得注意的是

如果你是 PHP，然后你的私钥文件 `CAkey.pem` 是下面的格式：

```pem
-----BEGIN EC PARAMETERS-----
xxx
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
xxxxxxx
-----END EC PRIVATE KEY-----
```

那么很正常。你可以使用下面的函式来获取十六进制做 Key 传。

```php
function getCAKeyPEMHex(string $pem): string {
    return bin2hex(
        openssl_pkey_get_details(
            openssl_pkey_get_private(
                $pem
            )
        )['ec']['d']
    );
}
```

在你请求本优雅的中间件的时候，别忘记使用 `urlencode()` 函式，避免腾讯传回的字符串的 `+` 被识别成空格 ` ` 等问题的产生。
