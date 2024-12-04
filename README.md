# poc-go-sftp

* create sftp server with docker
docker run -d -p 22:22 -e SFTP_USERS="`<<user>>`:`<<pass>>`:::`<<directory>>`" atmoz/sftp

* test to connect to sftp server
sftp -P 22 `<<user>>`@localhost