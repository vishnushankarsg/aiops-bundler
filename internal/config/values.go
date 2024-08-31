package config

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/AO-Metaplayer/aiops-bundler/pkg/modules/entities"
	"github.com/AO-Metaplayer/aiops-bundler/pkg/signer"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Values struct {
	// Documented variables.
	PrivateKey                   string
	EthClientUrl                 string
	Port                         int
	DataDirectory                string
	SupportedAiMiddlewares       []common.Address
	MaxVerificationGas           *big.Int
	MaxBatchGasLimit             *big.Int
	MaxOpTTL                     time.Duration
	OpLookupLimit                uint64
	Beneficiary                  string
	NativeBundlerCollectorTracer string
	NativeBundlerExecutorTracer  string
	ReputationConstants          *entities.ReputationConstants

	// Searcher mode variables.
	EthBuilderUrls    []string
	BlocksInTheFuture int

	// Observability variables.
	OTELServiceName      string
	OTELCollectorHeaders map[string]string
	OTELCollectorUrl     string
	OTELInsecureMode     bool

	// Alternative mempool variables.
	AltMempoolIPFSGateway string
	AltMempoolIds         []string

	// Rollup related variables.
	IsOpStackNetwork   bool
	IsRIP7212Supported bool
	IsArbStackNetwork  bool

	// Undocumented variables.
	DebugMode bool
	GinMode   string
}

func envKeyValStringToMap(s string) map[string]string {
	out := map[string]string{}
	for _, pair := range strings.Split(s, "&") {
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			break
		}
		out[kv[0]] = kv[1]
	}
	return out
}

func envArrayToAddressSlice(s string) []common.Address {
	env := strings.Split(s, ",")
	slc := []common.Address{}
	for _, ep := range env {
		slc = append(slc, common.HexToAddress(strings.TrimSpace(ep)))
	}

	return slc
}

func envArrayToStringSlice(s string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, ",")
}

func variableNotSetOrIsNil(env string) bool {
	return !viper.IsSet(env) || viper.GetString(env) == ""
}

// GetValues returns config for the bundler that has been read in from env vars. See
func GetValues() *Values {
	// Default variables
	viper.SetDefault("aiops_bundler_port", 4337)
	viper.SetDefault("aiops_bundler_data_directory", "/tmp/aiops_bundler")
	viper.SetDefault("aiops_bundler_supported_ai_middleware", "0x") //TODO: add aiMiddleware Contract address
	viper.SetDefault("aiops_bundler_max_verification_gas", 6000000)
	viper.SetDefault("aiops_bundler_max_batch_gas_limit", 18000000)
	viper.SetDefault("aiops_bundler_max_op_ttl_seconds", 180)
	viper.SetDefault("aiops_bundler_op_lookup_limit", 2000)
	viper.SetDefault("aiops_bundler_blocks_in_the_future", 6)
	viper.SetDefault("aiops_bundler_otel_insecure_mode", false)
	viper.SetDefault("aiops_bundler_is_op_stack_network", false)
	viper.SetDefault("aiops_bundler_is_arb_stack_network", false)
	viper.SetDefault("aiops_bundler_is_rip7212_supported", false)
	viper.SetDefault("aiops_bundler_debug_mode", false)
	viper.SetDefault("aiops_bundler_gin_mode", gin.ReleaseMode)

	// Read in from .env file if available
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			// Can ignore
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	// Read in from environment variables
	_ = viper.BindEnv("aiops_bundler_eth_client_url")
	_ = viper.BindEnv("aiops_bundler_private_key")
	_ = viper.BindEnv("aiops_bundler_port")
	_ = viper.BindEnv("aiops_bundler_data_directory")
	_ = viper.BindEnv("aiops_bundler_supported_ai_middleware")
	_ = viper.BindEnv("aiops_bundler_beneficiary")
	_ = viper.BindEnv("aiops_bundler_native_bundler_collector_tracer")
	_ = viper.BindEnv("aiops_bundler_native_bundler_executor_tracer")
	_ = viper.BindEnv("aiops_bundler_max_verification_gas")
	_ = viper.BindEnv("aiops_bundler_max_batch_gas_limit")
	_ = viper.BindEnv("aiops_bundler_max_op_ttl_seconds")
	_ = viper.BindEnv("aiops_bundler_op_lookup_limit")
	_ = viper.BindEnv("aiops_bundler_eth_builder_urls")
	_ = viper.BindEnv("aiops_bundler_blocks_in_the_future")
	_ = viper.BindEnv("aiops_bundler_otel_service_name")
	_ = viper.BindEnv("aiops_bundler_otel_collector_headers")
	_ = viper.BindEnv("aiops_bundler_otel_collector_url")
	_ = viper.BindEnv("aiops_bundler_otel_insecure_mode")
	_ = viper.BindEnv("aiops_bundler_alt_mempool_ipfs_gateway")
	_ = viper.BindEnv("aiops_bundler_alt_mempool_ids")
	_ = viper.BindEnv("aiops_bundler_is_op_stack_network")
	_ = viper.BindEnv("aiops_bundler_is_arb_stack_network")
	_ = viper.BindEnv("aiops_bundler_is_rip7212_supported")
	_ = viper.BindEnv("aiops_bundler_debug_mode")
	_ = viper.BindEnv("aiops_bundler_gin_mode")

	// Validate required variables
	if variableNotSetOrIsNil("aiops_bundler_eth_client_url") {
		panic("Fatal config error: aiops_bundler_eth_client_url not set")
	}

	if variableNotSetOrIsNil("aiops_bundler_private_key") {
		panic("Fatal config error: aiops_bundler_private_key not set")
	}

	if !viper.IsSet("aiops_bundler_beneficiary") {
		s, err := signer.New(viper.GetString("aiops_bundler_private_key"))
		if err != nil {
			panic(err)
		}
		viper.SetDefault("aiops_bundler_beneficiary", s.Address.String())
	}

	switch viper.GetString("mode") {
	case "searcher":
		if variableNotSetOrIsNil("aiops_bundler_eth_builder_urls") {
			panic("Fatal config error: aiops_bundler_eth_builder_urls not set")
		}
	}

	// Validate O11Y variables
	if viper.IsSet("aiops_bundler_otel_service_name") &&
		variableNotSetOrIsNil("aiops_bundler_otel_collector_url") {
		panic("Fatal config error: aiops_bundler_otel_service_name is set without a collector URL")
	}

	// Validate Alternative mempool variables
	if viper.IsSet("aiops_bundler_alt_mempool_ids") &&
		variableNotSetOrIsNil("aiops_bundler_alt_mempool_ipfs_gateway") {
		panic("Fatal config error: aiops_bundler_alt_mempool_ids is set without specifying an IPFS gateway")
	}

	// Return Values
	privateKey := viper.GetString("aiops_bundler_private_key")
	ethClientUrl := viper.GetString("aiops_bundler_eth_client_url")
	port := viper.GetInt("aiops_bundler_port")
	dataDirectory := viper.GetString("aiops_bundler_data_directory")
	supportedAiMiddlewares := envArrayToAddressSlice(viper.GetString("aiops_bundler_supported_ai_middleware"))
	beneficiary := viper.GetString("aiops_bundler_beneficiary")
	nativeBundlerCollectorTracer := viper.GetString("aiops_bundler_native_bundler_collector_tracer")
	nativeBundlerExecutorTracer := viper.GetString("aiops_bundler_native_bundler_executor_tracer")
	maxVerificationGas := big.NewInt(int64(viper.GetInt("aiops_bundler_max_verification_gas")))
	maxBatchGasLimit := big.NewInt(int64(viper.GetInt("aiops_bundler_max_batch_gas_limit")))
	maxOpTTL := time.Second * viper.GetDuration("aiops_bundler_max_op_ttl_seconds")
	opLookupLimit := viper.GetUint64("aiops_bundler_op_lookup_limit")
	ethBuilderUrls := envArrayToStringSlice(viper.GetString("aiops_bundler_eth_builder_urls"))
	blocksInTheFuture := viper.GetInt("aiops_bundler_blocks_in_the_future")
	otelServiceName := viper.GetString("aiops_bundler_otel_service_name")
	otelCollectorHeader := envKeyValStringToMap(viper.GetString("aiops_bundler_otel_collector_headers"))
	otelCollectorUrl := viper.GetString("aiops_bundler_otel_collector_url")
	otelInsecureMode := viper.GetBool("aiops_bundler_otel_insecure_mode")
	altMempoolIPFSGateway := viper.GetString("aiops_bundler_alt_mempool_ipfs_gateway")
	altMempoolIds := envArrayToStringSlice(viper.GetString("aiops_bundler_alt_mempool_ids"))
	isOpStackNetwork := viper.GetBool("aiops_bundler_is_op_stack_network")
	isArbStackNetwork := viper.GetBool("aiops_bundler_is_arb_stack_network")
	isRIP7212Supported := viper.GetBool("aiops_bundler_is_rip7212_supported")
	debugMode := viper.GetBool("aiops_bundler_debug_mode")
	ginMode := viper.GetString("aiops_bundler_gin_mode")
	return &Values{
		PrivateKey:                   privateKey,
		EthClientUrl:                 ethClientUrl,
		Port:                         port,
		DataDirectory:                dataDirectory,
		SupportedAiMiddlewares:       supportedAiMiddlewares,
		Beneficiary:                  beneficiary,
		NativeBundlerCollectorTracer: nativeBundlerCollectorTracer,
		NativeBundlerExecutorTracer:  nativeBundlerExecutorTracer,
		MaxVerificationGas:           maxVerificationGas,
		MaxBatchGasLimit:             maxBatchGasLimit,
		MaxOpTTL:                     maxOpTTL,
		OpLookupLimit:                opLookupLimit,
		ReputationConstants:          NewReputationConstantsFromEnv(),
		EthBuilderUrls:               ethBuilderUrls,
		BlocksInTheFuture:            blocksInTheFuture,
		OTELServiceName:              otelServiceName,
		OTELCollectorHeaders:         otelCollectorHeader,
		OTELCollectorUrl:             otelCollectorUrl,
		OTELInsecureMode:             otelInsecureMode,
		AltMempoolIPFSGateway:        altMempoolIPFSGateway,
		AltMempoolIds:                altMempoolIds,
		IsOpStackNetwork:             isOpStackNetwork,
		IsArbStackNetwork:            isArbStackNetwork,
		IsRIP7212Supported:           isRIP7212Supported,
		DebugMode:                    debugMode,
		GinMode:                      ginMode,
	}
}
