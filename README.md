# CredFinder

Script to look for any possible credential in the file of specific folder

```

> go run .\CredFinder.go -path ./test/ -dictionary ./wordlist.txt

 ____________________________________________________________

   CredFinder - GOlang Credential Finder
 ____________________________________________________________

Start searching possible credential under "./test/"
Dictionary: [password id credential パスワード 認証情報]
Started at: 2025-03-25 11:12:58
=======================Result=========================
Found "password" in: test\pass.txt
Found "パスワード" in: test\pass2.txt
Found "認証情報" in: test\testtest\メモ.txt
Interesting filename: "test\パスワード.txt"
Found "password" in: test\パスワード.txt

=======================Finished=======================
Finished at: 2025-03-25 11:12:58
Execution time: 22.9888ms
```
