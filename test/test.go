package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func Getkeys() {
	//得到私钥
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	x509_Privatekey := x509.MarshalPKCS1PrivateKey(privateKey)
	//创建一个用来保存私钥的以.pem结尾的文件
	fp, _ := os.Create("private.pem")
	defer fp.Close()
	//将私钥字符串设置到pem格式块中
	pem_block := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509_Privatekey,
	}
	//转码为pem并输出到文件中
	pem.Encode(fp, &pem_block)

	//处理公钥,公钥包含在私钥中
	publickKey := privateKey.PublicKey
	//接下来的处理方法同私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	x509_PublicKey, _ := x509.MarshalPKIXPublicKey(&publickKey)
	pem_PublickKey := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509_PublicKey,
	}
	file, _ := os.Create("PublicKey.pem")
	defer file.Close()
	//转码为pem并输出到文件中
	pem.Encode(file, &pem_PublickKey)

}

//使用公钥进行加密
func RSA_encrypter(path string, msg []byte) []byte {
	//首先从文件中提取公钥
	fp, _ := os.Open(path)
	defer fp.Close()
	//测量文件长度以便于保存
	fileinfo, _ := fp.Stat()
	buf := make([]byte, fileinfo.Size())
	fp.Read(buf)
	//下面的操作是与创建秘钥保存时相反的
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码,得到一个interface类型的pub
	pub, _ := x509.ParsePKIXPublicKey(block.Bytes)
	//加密操作,需要将接口类型的pub进行类型断言得到公钥类型
	cipherText, _ := rsa.EncryptPKCS1v15(rand.Reader, pub.(*rsa.PublicKey), msg)
	return cipherText
}

//使用私钥进行解密
func RSA_decrypter(path string, cipherText []byte) []byte {
	//同加密时，先将私钥从文件中取出，进行二次解码
	fp, _ := os.Open(path)
	defer fp.Close()
	fileinfo, _ := fp.Stat()
	buf := make([]byte, fileinfo.Size())
	fp.Read(buf)
	block, _ := pem.Decode(buf)
	PrivateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	//二次解码完毕，调用解密函数
	afterDecrypter, _ := rsa.DecryptPKCS1v15(rand.Reader, PrivateKey, cipherText)
	return afterDecrypter
}

//使用私钥签名，path是私钥路径，msg是要签名的信息
func Signname(path string, msg []byte) []byte {
	//签名函数中需要的数据散列值
	//首先从文件中提取公钥
	fp, _ := os.Open(path)
	defer fp.Close()
	fileinfo, _ := fp.Stat()
	buf := make([]byte, fileinfo.Size())
	fp.Read(buf)
	block, _ := pem.Decode(buf)
	PrivateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	//加密操作,需要将接口类型的pub进行类型断言得到公钥类型

	hash := sha256.Sum256(msg)
	//调用签名函数,填入所需四个参数，得到签名
	sign, _ := rsa.SignPKCS1v15(rand.Reader, PrivateKey, crypto.SHA256, hash[:])
	fmt.Printf("sign:%x\n", sign)
	return sign

}

func VerifySign(path string, signText []byte, plainText []byte) string {
	//首先从文件中提取公钥
	fp, _ := os.Open(path)
	defer fp.Close()
	//测量文件长度以便于保存
	fileinfo, _ := fp.Stat()
	buf := make([]byte, fileinfo.Size())
	fp.Read(buf)
	//下面的操作是与创建秘钥保存时相反的
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码,得到一个interface类型的pub
	pub, _ := x509.ParsePKIXPublicKey(block.Bytes)
	//签名函数中需要的数据散列值
	hash := sha256.Sum256(plainText)
	//验证签名
	err := rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA256, hash[:], signText)
	if err != nil {
		return "认证失败"
	} else {
		return "认证成功"
	}

}

func main() {
	// Getkeys()
	// //尝试调用
	// msg := []byte("RSA非对称加密很棒")
	// ciphertext := RSA_encrypter("csdn_PublicKey.pem", msg)
	// //转化为十六进制方便查看结果
	// fmt.Println(hex.EncodeToString(ciphertext))
	// result := RSA_decrypter("csdn_private.pem", ciphertext)
	// fmt.Println(string(result))

	signstr := Signname("private.pem", []byte("I love you"))
	check := VerifySign("PublicKey.pem", signstr, []byte("I love you"))
	fmt.Print(check)
}
