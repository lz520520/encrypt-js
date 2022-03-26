一个前端加密相关的处理小工具，只是为了省点事

提供了Makefile,自行编译，不需要第三方依赖
```
make win
make lin
make mac
```

使用步骤
1. 明文密码写入in_plain_pwd.txt
2. 执行encrypt-js.exe js, 输出out_js.txt
3. out_js.txt里的加密算法部分自行替换，然后在F12 console里执行，执行前最好清除历史
4. 执行后全选复制console里的结果到in_console_encrypt.txt
5. 执行 encrypt-js.exe enc, 输出out_encrypt_pwd.txt
6. out_encrypt_pwd.txt里拿去爆破，结果可以在in_console_encrypt.txt索引相应明文