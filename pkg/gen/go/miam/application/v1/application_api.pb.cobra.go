// Code generated by protoc-gen-cobra. DO NOT EDIT.

package applicationv1

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/golang/protobuf/jsonpb"
	proto "github.com/golang/protobuf/proto"
	prettyjson "github.com/hokaccha/go-prettyjson"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"golang.org/x/oauth2"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

// -----------------------------------------------------------------------------
var DefaultClientCommandConfig = newClientCommandConfig()

type ClientCommandConfig struct {
	ServerAddr         string        `envconfig:"SERVER_ADDR" default:"localhost:8080"`
	Timeout            time.Duration `envconfig:"TIMEOUT" default:"10s"`
	TLS                bool          `envconfig:"TLS"`
	ServerName         string        `envconfig:"TLS_SERVER_NAME"`
	InsecureSkipVerify bool          `envconfig:"TLS_INSECURE_SKIP_VERIFY"`
	CACertFile         string        `envconfig:"TLS_CA_CERT_FILE"`
	CertFile           string        `envconfig:"TLS_CERT_FILE"`
	KeyFile            string        `envconfig:"TLS_KEY_FILE"`
	AuthToken          string        `envconfig:"AUTH_TOKEN"`
	AuthTokenType      string        `envconfig:"AUTH_TOKEN_TYPE" default:"Bearer"`
	JWTKey             string        `envconfig:"JWT_KEY"`
	JWTKeyFile         string        `envconfig:"JWT_KEY_FILE"`
}

func newClientCommandConfig() *ClientCommandConfig {
	c := &ClientCommandConfig{}
	envconfig.Process("", c)
	return c
}

func (o *ClientCommandConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&o.ServerAddr, "server-addr", "s", o.ServerAddr, "server address in form of host:port")
	fs.DurationVar(&o.Timeout, "timeout", o.Timeout, "client connection timeout")
	fs.BoolVar(&o.TLS, "tls", o.TLS, "enable tls")
	fs.StringVar(&o.ServerName, "tls-server-name", o.ServerName, "tls server name override")
	fs.BoolVar(&o.InsecureSkipVerify, "tls-insecure-skip-verify", o.InsecureSkipVerify, "INSECURE: skip tls checks")
	fs.StringVar(&o.CACertFile, "tls-ca-cert-file", o.CACertFile, "ca certificate file")
	fs.StringVar(&o.CertFile, "tls-cert-file", o.CertFile, "client certificate file")
	fs.StringVar(&o.KeyFile, "tls-key-file", o.KeyFile, "client key file")
	fs.StringVar(&o.AuthToken, "auth-token", o.AuthToken, "authorization token")
	fs.StringVar(&o.AuthTokenType, "auth-token-type", o.AuthTokenType, "authorization token type")
	fs.StringVar(&o.JWTKey, "jwt-key", o.JWTKey, "jwt key")
	fs.StringVar(&o.JWTKeyFile, "jwt-key-file", o.JWTKeyFile, "jwt key file")
}

// -----------------------------------------------------------------------------

func dial(cfg *ClientCommandConfig) (*grpc.ClientConn, error) {
	// Default client connection options
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTimeout(cfg.Timeout),
	}

	// TLS connection
	if cfg.TLS {
		tlsConfig := &tls.Config{}

		// Validate certificate chain
		if cfg.InsecureSkipVerify {
			tlsConfig.InsecureSkipVerify = true
		}

		// CA given
		if cfg.CACertFile != "" {
			cacert, err := ioutil.ReadFile(cfg.CACertFile)
			if err != nil {
				return nil, fmt.Errorf("ca cert: %v", err)
			}
			certpool := x509.NewCertPool()
			certpool.AppendCertsFromPEM(cacert)
			tlsConfig.RootCAs = certpool
		}

		// Client certificate given
		if cfg.CertFile != "" {
			if cfg.KeyFile == "" {
				return nil, fmt.Errorf("missing key file")
			}
			pair, err := tls.LoadX509KeyPair(cfg.CertFile, cfg.KeyFile)
			if err != nil {
				return nil, fmt.Errorf("cert/key: %v", err)
			}
			tlsConfig.Certificates = []tls.Certificate{pair}
		}

		// Override server name
		if cfg.ServerName != "" {
			tlsConfig.ServerName = cfg.ServerName
		} else {
			addr, _, _ := net.SplitHostPort(cfg.ServerAddr)
			tlsConfig.ServerName = addr
		}

		// tlsConfig.BuildNameToCertificate()
		cred := credentials.NewTLS(tlsConfig)
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		// Fallback to insecure
		opts = append(opts, grpc.WithInsecure())
	}

	// Token given
	if cfg.AuthToken != "" {
		cred := oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: cfg.AuthToken,
			TokenType:   cfg.AuthTokenType,
		})
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}

	// JWT key
	if cfg.JWTKey != "" {
		cred, err := oauth.NewJWTAccessFromKey([]byte(cfg.JWTKey))
		if err != nil {
			return nil, fmt.Errorf("jwt key: %v", err)
		}
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}

	// Load JWT key from file
	if cfg.JWTKeyFile != "" {
		cred, err := oauth.NewJWTAccessFromFile(cfg.JWTKeyFile)
		if err != nil {
			return nil, fmt.Errorf("jwt key file: %v", err)
		}
		opts = append(opts, grpc.WithPerRPCCredentials(cred))
	}

	// Real dial connection
	conn, err := grpc.Dial(cfg.ServerAddr, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func beautify(msg proto.Message) {
	m := &jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	result, err := m.MarshalToString(msg)
	if err != nil {
		log.Fatal("Unable to serialize response")
	}

	out, err := prettyjson.Format([]byte(result))
	if err != nil {
		log.Fatal("Unable to beautify response")
	}

	unicode := []byte{92, 117, 48}
	index := bytes.Index(out, unicode)

	for index != -1 {
		uni := out[index : index+6]
		s, err := strconv.Unquote("\"" + fmt.Sprintf("%s", uni) + "\"")
		if err != nil {
			log.Fatal("Unable to unquote response")
			fmt.Printf("%s", out)
			return
		}

		out = append(append(out[:index], s...), out[index+6:]...)
		index = bytes.Index(out, unicode)
	}
	fmt.Printf("%s", out)
}

// -----------------------------------------------------------------------------

var ApplicationAPIClientCommand = &cobra.Command{
	Use: "applicationAPI",
}

var applicationAPI_CreateClientCommand = &cobra.Command{
	Use:  "create",
	Long: "Create client\n\nYou can use environment variables with the same name of the command flags.\nAll caps and s/-/_, e.g. SERVER_ADDR.",
	Example: `
Save a sample request to a file (or refer to your protobuf descriptor to create one):
	create -p > req.json
Submit request using file:
	create -f req.json
Authenticate using the Authorization header (requires transport security):
	export AUTH_TOKEN=your_access_token
	export SERVER_ADDR=api.example.com:443
	echo '{json}' | create --tls`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var req CreateRequest

		// Get a connection
		conn, err := dial(DefaultClientCommandConfig)
		if err != nil {
			return err
		}
		defer conn.Close()

		// Initialize client wrapper
		grpcClient := NewApplicationAPIClient(conn)

		// Unmarshal request
		if err := jsonpb.Unmarshal(bufio.NewReader(os.Stdin), &req); err != nil {
			return err
		}

		// Prepare context
		ctx := context.Background()

		// Do the call
		res, err := grpcClient.Create(ctx, &req)
		if err != nil {
			return err
		}

		// Beautify result
		beautify(res)

		// no error
		return nil
	},
}

func init() {
	ApplicationAPIClientCommand.AddCommand(applicationAPI_CreateClientCommand)
	DefaultClientCommandConfig.AddFlags(applicationAPI_CreateClientCommand.Flags())
}

var applicationAPI_UpdateClientCommand = &cobra.Command{
	Use:  "update",
	Long: "Update client\n\nYou can use environment variables with the same name of the command flags.\nAll caps and s/-/_, e.g. SERVER_ADDR.",
	Example: `
Save a sample request to a file (or refer to your protobuf descriptor to create one):
	update -p > req.json
Submit request using file:
	update -f req.json
Authenticate using the Authorization header (requires transport security):
	export AUTH_TOKEN=your_access_token
	export SERVER_ADDR=api.example.com:443
	echo '{json}' | update --tls`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var req UpdateRequest

		// Get a connection
		conn, err := dial(DefaultClientCommandConfig)
		if err != nil {
			return err
		}
		defer conn.Close()

		// Initialize client wrapper
		grpcClient := NewApplicationAPIClient(conn)

		// Unmarshal request
		if err := jsonpb.Unmarshal(bufio.NewReader(os.Stdin), &req); err != nil {
			return err
		}

		// Prepare context
		ctx := context.Background()

		// Do the call
		res, err := grpcClient.Update(ctx, &req)
		if err != nil {
			return err
		}

		// Beautify result
		beautify(res)

		// no error
		return nil
	},
}

func init() {
	ApplicationAPIClientCommand.AddCommand(applicationAPI_UpdateClientCommand)
	DefaultClientCommandConfig.AddFlags(applicationAPI_UpdateClientCommand.Flags())
}

var applicationAPI_DeleteClientCommand = &cobra.Command{
	Use:  "delete",
	Long: "Delete client\n\nYou can use environment variables with the same name of the command flags.\nAll caps and s/-/_, e.g. SERVER_ADDR.",
	Example: `
Save a sample request to a file (or refer to your protobuf descriptor to create one):
	delete -p > req.json
Submit request using file:
	delete -f req.json
Authenticate using the Authorization header (requires transport security):
	export AUTH_TOKEN=your_access_token
	export SERVER_ADDR=api.example.com:443
	echo '{json}' | delete --tls`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var req DeleteRequest

		// Get a connection
		conn, err := dial(DefaultClientCommandConfig)
		if err != nil {
			return err
		}
		defer conn.Close()

		// Initialize client wrapper
		grpcClient := NewApplicationAPIClient(conn)

		// Unmarshal request
		if err := jsonpb.Unmarshal(bufio.NewReader(os.Stdin), &req); err != nil {
			return err
		}

		// Prepare context
		ctx := context.Background()

		// Do the call
		res, err := grpcClient.Delete(ctx, &req)
		if err != nil {
			return err
		}

		// Beautify result
		beautify(res)

		// no error
		return nil
	},
}

func init() {
	ApplicationAPIClientCommand.AddCommand(applicationAPI_DeleteClientCommand)
	DefaultClientCommandConfig.AddFlags(applicationAPI_DeleteClientCommand.Flags())
}
