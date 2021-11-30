# Book

## 네트워크 구동
```
.../book/network$ ./generate.sh
.../book/network$ ./network.sh
```

---

## 체인코드 설치 및 배포
```
$ .../book/network$ ./cc.sh
```

---

## 애플리케이션 실행

> 필요한 라이브러리 설치
```
$ .../book/application$ npm install
```

> Identity 파일 생성 (Wallet 구성)
```
$ .../book/application$ node enrollAdmin.js
$ .../book/application$ node registerUser.js
```
> 서버 실행
```
$ .../book/application$ node server.js
```
> 웹브라우저 실행

[로컬 환경에서 실행하는 경우 클릭하세요](localhost:8070)