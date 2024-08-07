# Introduction to Bundlers
Bundlers power ERC-4337.

## Suggest Edits
In ERC-4337, a Bundler is the core infrastructure component that allows account abstraction to work on any EVM network. On the highest level, its purpose is to work with a mempool of Ai Operations to get the transaction to be included on-chain.


# Modes

An instance of the bundler supports different modes out of the box. Although all modes produce the same outcome, they vary in three areas:

    1. Mempool support: A private mempool is only visible to an individual bundler whereas a P2P mempool has AiOperations that have been propagated to all bundlers in the network.

    2. Networks: Some modes rely on specific features (e.g. mev-boost) that makes it only available to a certain set of EVM networks.

# Installation

A quick guide to self-hosting an ERC-4337 bundler for handling AiOperations.

## Install

### Docker

    ```
    Shell

    docker build -t aiops-bundler:latest

    docker run -p 4337:4337 aiops-bundler:latest
    ```

### Build from source

If you have Go setup in your environment, you can clone and build the bundler from source.

```
Shell

git clone git@gitlab.com:quantum-warriors/aiops-bundler.git
cd aiops-bundler

This will build a binary from source and output it to the ./tmp directory.
You can move the binary to a different directory in your PATH (e.g. /usr/local/bin).

go build -o ./tmp/aiops-bundler main.go
```

### Running the bundler

Make sure to have your environment variables configured before running the bundler. See the configuration section for details.

Run an instance in private mode:

Binary
```
aiops-bundler start --mode private
```

Docker
```
docker run -d --name bundler -p 4337:4337 \
  -e AIOPS_BUNDLER_ETH_CLIENT_URL=... \
  aiops-bundler:latest \
  /app/aiops-bundler start --mode private
```

For a description on the CLI commands and other supported modes:

Binary
```
aiops-bundler start --help
```

Docker
```
docker run aiops-bundler:latest /app/aiops-bundler --help
```


# Environment Variables

## Base variables

These variables are generally relevant to all bundler modes.

### Required

| Environment Variable	        | Description                       |
| :---------------------------- | :-------------------------------- | 
| AIOPS_BUNDLER_ETH_CLIENT_URL  | RPC url to the execution client.  |
| AIOPS_BUNDLER_PRIVATE_KEY	    | The private key for the EOA used to relay Ai Operation bundles to the AiMiddleware. |

### Optional

| Environment Variable          | Description	            | Default Value |
| :---------------------------- | :-----------------------: | :------------ |
| AIOPS_BUNDLER_PORT	        | Port to run the HTTP server on. | 4337 |
| AIOPS_BUNDLER_DATA_DIRECTORY	| Directory to store the embedded database.	| /tmp/aiopps_bundler|
| AIOPS_BUNDLER_SUPPORTED_AI_MIDDLEWARE | Comma separated AiMiddleware addresses to support. The first address is the preferred AiMiddleware. | Depends on the major version. See 🗺️ Entity Addresses |
| AIOPS_BUNDLER_BENEFICIARY | Address to send gas cost refunds for relaying Ai Operation bundles.	| Defaults to the public address of required private key. |
| AIOPS_BUNDLER_
NATIVE_BUNDLER_COLLECTOR_TRACER	| The name of the native tracer to use during validation. |	 Defaults to nil and will fallback to using the reference JS tracer. |
| AIOPS_BUNDLER_MAX_VERIFICATION_GAS |	The maximum verificationGasLimit on a received AiOperation. |	6,000,000 gas |
| AIOPS_BUNDLER_MAX_BATCH_GAS_LIMIT	| The maximum gas limit that can be submitted per AiOperation batch. |	18,000,000 gas |
| AIOPS_BUNDLER_MAX_OP_TTL_SECONDS	| The maximum duration that a AiOp can stay in the mempool before getting dropped. |	180 seconds |
| AIOPS_BUNDLER_OP_LOOKUP_LIMIT |	The maximum block range when looking up a Ai Operation with eth_getAiOperationReceipt or eth_getAiOperationByHash. Higher limits allow for fetching older Ai Operations but will result in higher request latency due to additional compute on the underlying node. | 2,000 blocks |
| AIOPS_BUNDLER_IS_OP_STACK_NETWORK |	A boolean value for bundlers on an OP stack network to properly account for the L1 callData cost. |	false |
| AIOPS_BUNDLER_IS_ARB_STACK_NETWORK |	A boolean value for bundlers on an Arbitrum stack network to properly account for the L1 callData cost. |	false |
| AIOPS_BUNDLER_IS_RIP7212_SUPPORTED |	A boolean value for bundlers on a network that supports RIP-7212 precompile for secp256r1 signature verification. |	false |


### Observability variables

Aiops Bundler supports tracers and metrics via OpenTelemetry.

| Environment Variable                  |	Description                           |
| :-------------------                  |    :-----------------------------------: |
| AIOPS_BUNDLER_OTEL_SERVICE_NAME	    | The name for this service (e.g. aiops-bundler). |
| AIOPS_BUNDLER_OTEL_COLLECTOR_URL      | The URL to forward OpenTelemetry signals to.    |
| AIOPS_BUNDLER_OTEL_COLLECTOR_HEADERS  | Optional collector request headers. This must be in the form of key1=value1&key2=value2. |
| AIOPS_BUNDLER_OTEL_INSECURE_MODE	    | Optional flag to disable transport security for the exporter's gRPC connection. Defaults to false. |
