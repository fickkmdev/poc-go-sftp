package main

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
)

func main() {
	// ตั้งค่าการเชื่อมต่อ SSH
	sshConfig := &ssh.ClientConfig{
		User: "testuser", // ใส่ชื่อผู้ใช้ SFTP
		Auth: []ssh.AuthMethod{
			ssh.Password("testpass"), // ใช้รหัสผ่าน
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // ไม่ตรวจสอบ HostKey (ไม่ปลอดภัยในโปรดักชัน)
	}

	// เชื่อมต่อไปยังเซิร์ฟเวอร์ SSH
	conn, err := ssh.Dial("tcp", "127.0.0.1:22", sshConfig)
	if err != nil {
		log.Fatalf("Failed to dial SSH: %v", err)
	}
	defer conn.Close()

	// สร้าง Client สำหรับ SFTP
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("Failed to create SFTP client: %v", err)
	}
	defer client.Close()

	// อัปโหลดไฟล์ไปยังเซิร์ฟเวอร์
	srcFile, err := os.Open("test_file.txt") // ไฟล์ที่ต้องการอัปโหลด
	if err != nil {
		log.Fatalf("Failed to open source file: %v", err)
	}
	defer srcFile.Close()

	dstFile, err := client.Create("/upload/test_file.txt") // ไฟล์ปลายทาง
	if err != nil {
		log.Fatalf("Failed to create destination file: %v", err)
	}
	defer dstFile.Close()

	// คัดลอกเนื้อหาไฟล์
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatalf("Failed to copy file: %v", err)
	}

	fmt.Println("File uploaded successfully!")

	// ดาวน์โหลดไฟล์จากเซิร์ฟเวอร์
	remoteFile, err := client.Open("/upload/test_file.txt")
	if err != nil {
		log.Fatalf("Failed to open remote file: %v", err)
	}
	defer remoteFile.Close()

	localFile, err := os.Create("downloaded_file.txt")
	if err != nil {
		log.Fatalf("Failed to create local file: %v", err)
	}
	defer localFile.Close()

	_, err = io.Copy(localFile, remoteFile)
	if err != nil {
		log.Fatalf("Failed to download file: %v", err)
	}

	fmt.Println("File downloaded successfully!")
}