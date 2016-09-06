./arch2vizceral -arch edge | python -m json.tool  > edge_vizceral.json
./arch2vizceral -arch test | python -m json.tool  > test_vizceral.json
cp *.json ../../vizceral-example/src
