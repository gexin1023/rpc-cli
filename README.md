json-rpc cli
------

```$xslt
./rpc-cli post --url http://127.0.0.1:6791 --method eth_getBlock --param 17

./rpc-cli post --url http://127.0.0.1:6791 --method eth_getTransaction --param "0x1234...222"

./rpc-cli post --url http://127.0.0.1:6791 --method eth_getTransactionReceipt --param "0x1234...222"

./rpc-cli post --url http://127.0.0.1:6791 --method personal_unlockAccount --param '["0x1234...222"]'
``` 