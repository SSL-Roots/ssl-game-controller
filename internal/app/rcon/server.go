package rcon

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/RoboCup-SSL/ssl-game-controller/pkg/refproto"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

type Server struct {
	Clients           map[string]*Client
	TrustedKeys       map[string]*rsa.PublicKey
	ConnectionHandler func(net.Conn)
}

type Client struct {
	Id                 string
	Conn               net.Conn
	Token              string
	PubKey             *rsa.PublicKey
	VerifiedConnection bool
}

func NewServer() (s *Server) {
	s = new(Server)
	s.Clients = map[string]*Client{}
	s.TrustedKeys = map[string]*rsa.PublicKey{}
	return
}

func (s *Server) Listen(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Print("Failed to listen on ", address)
		return
	}
	log.Print("Listening on ", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("Could not accept connection: ", err)
		} else {
			go s.ConnectionHandler(conn)
		}
	}
}

func (s *Server) ListenTls(address string) {

	if _, err := os.Stat("server.crt"); os.IsNotExist(err) {
		log.Println("Missing certificate for TLS endpoint. Put a server.crt in the working dir.")
		return
	}
	if _, err := os.Stat("server.key"); os.IsNotExist(err) {
		log.Println("Missing certificate key for TLS endpoint. Put a server.key in the working dir.")
		return
	}

	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Printf("Could not load X509 key pair: %v", err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	listener, err := tls.Listen("tcp", address, config)
	if err != nil {
		log.Printf("Failed to listen on %v: %v", address, err)
		return
	}
	log.Print("Listening on ", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print("Could not accept connection: ", err)
		} else {
			go s.ConnectionHandler(conn)
		}
	}
}

func (s *Server) CloseConnection(conn net.Conn, id string) {
	delete(s.Clients, id)
	log.Printf("Connection to %v closed", id)
}

func (c *Client) Ok() (reply refproto.ControllerReply) {
	reply.StatusCode = new(refproto.ControllerReply_StatusCode)
	*reply.StatusCode = refproto.ControllerReply_OK
	c.addVerification(&reply)
	return
}

func (c *Client) addVerification(reply *refproto.ControllerReply) {
	reply.Verification = new(refproto.ControllerReply_Verification)
	if c.Token != "" {
		reply.NextToken = new(string)
		*reply.NextToken = c.Token
		*reply.Verification = refproto.ControllerReply_VERIFIED
		c.VerifiedConnection = true
	} else {
		*reply.Verification = refproto.ControllerReply_UNVERIFIED
		c.VerifiedConnection = false
	}
}

func (c *Client) Reject(reason string) (reply refproto.ControllerReply) {
	log.Print("Reject connection: " + reason)
	reply.StatusCode = new(refproto.ControllerReply_StatusCode)
	*reply.StatusCode = refproto.ControllerReply_REJECTED
	reply.Reason = new(string)
	*reply.Reason = reason
	return
}

func (s *Server) LoadTrustedKeys(trustedKeysDir string) {
	files, err := ioutil.ReadDir(trustedKeysDir)
	if err != nil {
		log.Print("Could not read trusted keys: ", err)
		return
	}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".pub.pem") {
			pubKey := readPublicKey(trustedKeysDir + "/" + file.Name())
			if pubKey != nil {
				name := strings.Replace(file.Name(), ".pub.pem", "", 1)
				s.TrustedKeys[name] = pubKey
				log.Printf("Loaded public key for %v", name)
			}
		}
	}
	log.Printf("Loaded %v public keys", len(s.TrustedKeys))
}

func readPublicKey(filename string) *rsa.PublicKey {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print("Could not find private key at ", filename)
		return nil
	}
	p, rest := pem.Decode(b)
	if p == nil {
		log.Print("Could not decode public key. Remaining data: ", string(rest))
		return nil
	}
	if p.Type != "PUBLIC KEY" {
		log.Print("Public key type is wrong: ", p.Type)
		return nil
	}
	publicKey, err := x509.ParsePKIXPublicKey(p.Bytes)
	if err != nil {
		log.Print("Failed to parse public key: ", err)
		return nil
	}
	switch pub := publicKey.(type) {
	case *rsa.PublicKey:
		return pub
	default:
		log.Print("Unsupported public key: ", pub)
		return nil
	}
}

func VerifySignature(key *rsa.PublicKey, message proto.Message, signature []byte) error {
	messageBytes, err := proto.Marshal(message)
	if err != nil {
		log.Fatal("Failed to marshal message: ", err)
	}
	hash := sha256.New()
	hash.Write(messageBytes)
	d := hash.Sum(nil)
	return rsa.VerifyPKCS1v15(key, crypto.SHA256, d, signature)
}
