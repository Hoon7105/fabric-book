# chaincode insall
docker exec cli peer chaincode install -n book -v 1.0 -p github.com/book
# chaincode instatiate
docker exec cli peer chaincode instantiate -n book -v 1.0 -C mychannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")'
sleep 5

echo '-------------------------------------END-------------------------------------'
