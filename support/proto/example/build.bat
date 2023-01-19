set example_server=..\..\example_server\internal\messages\
set example_clientgo=..\..\example_clientgo\internal\messages\

rm -r .\messages\
protoc --go_out=. msgid.proto
protoc --go_out=. msgping.proto

rm -r %example_server%
mkdir %example_server%
copy .\messages\*.* %example_server%
copy .\messagesjson\*.* %example_server%

rm -r %example_clientgo%
mkdir %example_clientgo%
copy .\messages\*.* %example_clientgo%
copy .\messagesjson\*.* %example_clientgo%