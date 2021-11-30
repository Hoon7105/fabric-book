// ExpressJS Setup
const express = require('express');
const app = express();

// Hyperledger Bridge
const { FileSystemWallet, Gateway } = require('fabric-network');
const fs = require('fs');
const path = require('path');
const ccpPath = path.resolve(__dirname, '..', 'network' ,'connection.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

// Constants
const PORT = 8070;
const HOST = '0.0.0.0';

// use static file
app.use(express.static(path.join(__dirname, 'views')));
app.use(express.static(path.join(__dirname, 'public')));


// configure app to use body-parser
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

// main page routing
app.get('/', (req, res)=>{
    res.sendFile(__dirname + '/views/index.html');
})
app.get('/curator', (req, res)=>{

//    const bookList = require("./BookRecList");
    const bookList = require(process.cwd() + "/data/BookRecList");
    

    // json생성 query all books -> random or 고르든 지지고 볶아(아마 유저의 책 선정 말씀하신 것 같아요)
//변수로 책 제목들 지정 후 추천 가능한지 확인
var parseData = JSON.parse(JSON.stringify(bookList));
  const index = Math.floor(Math.random() * parseData.length);
  // parseData[index].name
  const myobj = {result: "success", name: parseData[index].name}
  res.status(200).json(myobj) 
})
 async function cc_call(fn_name, args){
     
     const walletPath = path.join(process.cwd(), 'wallet');
     const wallet = new FileSystemWallet(walletPath);
 
     const userExists = await wallet.exists('user1');
     if (!userExists) {
         console.log('An identity for the user "user1" does not exist in the wallet');
         console.log('Run the registerUser.js application before retrying');
         return;
     }
     const gateway = new Gateway();
     await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });
     const network = await gateway.getNetwork('mychannel');
     const contract = network.getContract('book');
 
     var result;
     
     if(fn_name == 'setUser')
     {
 	user=args;
         result = await contract.submitTransaction('setUser', user);
     }
     else if( fn_name == 'readbook')
     {
         user=args[0];
         isbn=args[1];
         name=args[2];
         cat=args[3];
         result = await contract.submitTransaction('readbook', user, isbn, name, cat);
     }
 //    else if(fn_name == 'readRating')
 //        result = await contract.evaluateTransaction('readRating', args);
     else
         result = 'not supported function'
 
     return result;
 }
 
 //create user
app.post('/user', async(req, res)=>{
  const user = req.body.user;
  console.log("add user name: " + user);
  result = cc_call('setUser', user)
  const myobj = {result: "success"}
  res.status(200).json(myobj) 
})

// add score
app.post('/readbook', async(req, res)=>{
    const user = req.body.user;
    const isbn = req.body.isbn;
    const name = req.body.name;
    const cat = req.body.cat;

    console.log("add read book isbn: " + isbn);
    //console.log("add read book title: " + name); //이거 name 중복에 title 주제 = cat 카테고리 랑 같아서 바꿔야할거같습니다 
    console.log("add read book name: " + name);
    console.log("add read book cat: " + cat);

    var args=[user, isbn, name, cat];
    result = cc_call('readbook', args)

    console.log("Transaction has been submitted.");

    const myobj = {result: "success"}
    res.status(200).json(myobj) 
})

// find user
app.post('/readbook/:user', async (req,res)=>{
    const user = req.body.user;
    console.log("user: " + user);
    const walletPath = path.join(process.cwd(), 'wallet');
    const wallet = new FileSystemWallet(walletPath);
    console.log(`Wallet path: ${walletPath}`);

    // Check to see if we've already enrolled the user.
    const userExists = await wallet.exists('user1');
    if (!userExists) {
        console.log('An identity for the user "user1" does not exist in the wallet');
        console.log('Run the registerUser.js application before retrying');
        return;
    }
    const gateway = new Gateway();
    await gateway.connect(ccp, { wallet, identity: 'user1', discovery: { enabled: false } });
    const network = await gateway.getNetwork('mychannel');
    const contract = network.getContract('book');
    const result = await contract.evaluateTransaction('getuserBookinfo', user);
    const myobj = JSON.parse(result)
    res.status(200).json(myobj)
    console.log(myobj);
    // res.status(200).json(result)

});

// server start
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);