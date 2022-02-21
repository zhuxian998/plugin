// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.9.1
// source: config.proto

package types

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SyncTxConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain33Host         string `protobuf:"bytes,1,opt,name=chain33host,proto3" json:"chain33host,omitempty"`
	PushHost            string `protobuf:"bytes,2,opt,name=pushHost,proto3" json:"pushHost,omitempty"`
	PushName            string `protobuf:"bytes,3,opt,name=pushName,proto3" json:"pushName,omitempty"`
	PushBind            string `protobuf:"bytes,4,opt,name=pushBind,proto3" json:"pushBind,omitempty"`
	MaturityDegree      int32  `protobuf:"varint,5,opt,name=maturityDegree,proto3" json:"maturityDegree,omitempty"`
	FetchHeightPeriodMs int64  `protobuf:"varint,9,opt,name=fetchHeightPeriodMs,proto3" json:"fetchHeightPeriodMs,omitempty"`
	StartSyncHeight     int64  `protobuf:"varint,10,opt,name=startSyncHeight,proto3" json:"startSyncHeight,omitempty"`
	StartSyncSequence   int64  `protobuf:"varint,11,opt,name=startSyncSequence,proto3" json:"startSyncSequence,omitempty"`
	StartSyncHash       string `protobuf:"bytes,12,opt,name=startSyncHash,proto3" json:"startSyncHash,omitempty"`
}

func (x *SyncTxConfig) Reset() {
	*x = SyncTxConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncTxConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncTxConfig) ProtoMessage() {}

func (x *SyncTxConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncTxConfig.ProtoReflect.Descriptor instead.
func (*SyncTxConfig) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *SyncTxConfig) GetChain33Host() string {
	if x != nil {
		return x.Chain33Host
	}
	return ""
}

func (x *SyncTxConfig) GetPushHost() string {
	if x != nil {
		return x.PushHost
	}
	return ""
}

func (x *SyncTxConfig) GetPushName() string {
	if x != nil {
		return x.PushName
	}
	return ""
}

func (x *SyncTxConfig) GetPushBind() string {
	if x != nil {
		return x.PushBind
	}
	return ""
}

func (x *SyncTxConfig) GetMaturityDegree() int32 {
	if x != nil {
		return x.MaturityDegree
	}
	return 0
}

func (x *SyncTxConfig) GetFetchHeightPeriodMs() int64 {
	if x != nil {
		return x.FetchHeightPeriodMs
	}
	return 0
}

func (x *SyncTxConfig) GetStartSyncHeight() int64 {
	if x != nil {
		return x.StartSyncHeight
	}
	return 0
}

func (x *SyncTxConfig) GetStartSyncSequence() int64 {
	if x != nil {
		return x.StartSyncSequence
	}
	return 0
}

func (x *SyncTxConfig) GetStartSyncHash() string {
	if x != nil {
		return x.StartSyncHash
	}
	return ""
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Loglevel        string `protobuf:"bytes,1,opt,name=loglevel,proto3" json:"loglevel,omitempty"`
	LogConsoleLevel string `protobuf:"bytes,2,opt,name=logConsoleLevel,proto3" json:"logConsoleLevel,omitempty"`
	LogFile         string `protobuf:"bytes,3,opt,name=logFile,proto3" json:"logFile,omitempty"`
	MaxFileSize     uint32 `protobuf:"varint,4,opt,name=maxFileSize,proto3" json:"maxFileSize,omitempty"`
	MaxBackups      uint32 `protobuf:"varint,5,opt,name=maxBackups,proto3" json:"maxBackups,omitempty"`
	MaxAge          uint32 `protobuf:"varint,6,opt,name=maxAge,proto3" json:"maxAge,omitempty"`
	LocalTime       bool   `protobuf:"varint,7,opt,name=localTime,proto3" json:"localTime,omitempty"`
	Compress        bool   `protobuf:"varint,8,opt,name=compress,proto3" json:"compress,omitempty"`
	CallerFile      bool   `protobuf:"varint,9,opt,name=callerFile,proto3" json:"callerFile,omitempty"`
	CallerFunction  bool   `protobuf:"varint,10,opt,name=callerFunction,proto3" json:"callerFunction,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *Log) GetLoglevel() string {
	if x != nil {
		return x.Loglevel
	}
	return ""
}

func (x *Log) GetLogConsoleLevel() string {
	if x != nil {
		return x.LogConsoleLevel
	}
	return ""
}

func (x *Log) GetLogFile() string {
	if x != nil {
		return x.LogFile
	}
	return ""
}

func (x *Log) GetMaxFileSize() uint32 {
	if x != nil {
		return x.MaxFileSize
	}
	return 0
}

func (x *Log) GetMaxBackups() uint32 {
	if x != nil {
		return x.MaxBackups
	}
	return 0
}

func (x *Log) GetMaxAge() uint32 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

func (x *Log) GetLocalTime() bool {
	if x != nil {
		return x.LocalTime
	}
	return false
}

func (x *Log) GetCompress() bool {
	if x != nil {
		return x.Compress
	}
	return false
}

func (x *Log) GetCallerFile() bool {
	if x != nil {
		return x.CallerFile
	}
	return false
}

func (x *Log) GetCallerFunction() bool {
	if x != nil {
		return x.CallerFunction
	}
	return false
}

type EthRelayerCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthChainName        string   `protobuf:"bytes,1,opt,name=ethChainName,proto3" json:"ethChainName,omitempty"`
	EthProvider         []string `protobuf:"bytes,2,rep,name=ethProvider,proto3" json:"ethProvider,omitempty"`
	EthProviderCli      []string `protobuf:"bytes,3,rep,name=ethProviderCli,proto3" json:"ethProviderCli,omitempty"`
	BridgeRegistry      string   `protobuf:"bytes,4,opt,name=bridgeRegistry,proto3" json:"bridgeRegistry,omitempty"`
	EthMaturityDegree   int32    `protobuf:"varint,5,opt,name=ethMaturityDegree,proto3" json:"ethMaturityDegree,omitempty"`
	EthBlockFetchPeriod int32    `protobuf:"varint,6,opt,name=ethBlockFetchPeriod,proto3" json:"ethBlockFetchPeriod,omitempty"`
}

func (x *EthRelayerCfg) Reset() {
	*x = EthRelayerCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EthRelayerCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EthRelayerCfg) ProtoMessage() {}

func (x *EthRelayerCfg) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EthRelayerCfg.ProtoReflect.Descriptor instead.
func (*EthRelayerCfg) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *EthRelayerCfg) GetEthChainName() string {
	if x != nil {
		return x.EthChainName
	}
	return ""
}

func (x *EthRelayerCfg) GetEthProvider() []string {
	if x != nil {
		return x.EthProvider
	}
	return nil
}

func (x *EthRelayerCfg) GetEthProviderCli() []string {
	if x != nil {
		return x.EthProviderCli
	}
	return nil
}

func (x *EthRelayerCfg) GetBridgeRegistry() string {
	if x != nil {
		return x.BridgeRegistry
	}
	return ""
}

func (x *EthRelayerCfg) GetEthMaturityDegree() int32 {
	if x != nil {
		return x.EthMaturityDegree
	}
	return 0
}

func (x *EthRelayerCfg) GetEthBlockFetchPeriod() int32 {
	if x != nil {
		return x.EthBlockFetchPeriod
	}
	return 0
}

type Chain33RelayerCfg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SyncTxConfig            *SyncTxConfig `protobuf:"bytes,1,opt,name=syncTxConfig,proto3" json:"syncTxConfig,omitempty"`
	BridgeRegistryOnChain33 string        `protobuf:"bytes,2,opt,name=bridgeRegistryOnChain33,proto3" json:"bridgeRegistryOnChain33,omitempty"`
	ChainName               string        `protobuf:"bytes,3,opt,name=chainName,proto3" json:"chainName,omitempty"`
	ChainID4Chain33         int32         `protobuf:"varint,4,opt,name=chainID4Chain33,proto3" json:"chainID4Chain33,omitempty"`
}

func (x *Chain33RelayerCfg) Reset() {
	*x = Chain33RelayerCfg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chain33RelayerCfg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chain33RelayerCfg) ProtoMessage() {}

func (x *Chain33RelayerCfg) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chain33RelayerCfg.ProtoReflect.Descriptor instead.
func (*Chain33RelayerCfg) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

func (x *Chain33RelayerCfg) GetSyncTxConfig() *SyncTxConfig {
	if x != nil {
		return x.SyncTxConfig
	}
	return nil
}

func (x *Chain33RelayerCfg) GetBridgeRegistryOnChain33() string {
	if x != nil {
		return x.BridgeRegistryOnChain33
	}
	return ""
}

func (x *Chain33RelayerCfg) GetChainName() string {
	if x != nil {
		return x.ChainName
	}
	return ""
}

func (x *Chain33RelayerCfg) GetChainID4Chain33() int32 {
	if x != nil {
		return x.ChainID4Chain33
	}
	return 0
}

type RelayerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title             string             `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	JrpcBindAddr      string             `protobuf:"bytes,2,opt,name=jrpcBindAddr,proto3" json:"jrpcBindAddr,omitempty"` // Jrpc服务地址
	EthRelayerCfg     []*EthRelayerCfg   `protobuf:"bytes,3,rep,name=ethRelayerCfg,proto3" json:"ethRelayerCfg,omitempty"`
	Chain33RelayerCfg *Chain33RelayerCfg `protobuf:"bytes,4,opt,name=chain33RelayerCfg,proto3" json:"chain33RelayerCfg,omitempty"`
	Log               *Log               `protobuf:"bytes,5,opt,name=log,proto3" json:"log,omitempty"`
	Dbdriver          string             `protobuf:"bytes,6,opt,name=dbdriver,proto3" json:"dbdriver,omitempty"` //数据库类型
	DbPath            string             `protobuf:"bytes,7,opt,name=dbPath,proto3" json:"dbPath,omitempty"`     //数据库存储目录
	DbCache           int32              `protobuf:"varint,8,opt,name=dbCache,proto3" json:"dbCache,omitempty"`  //数据库缓存大小
	ProcessWithDraw   bool               `protobuf:"varint,9,opt,name=processWithDraw,proto3" json:"processWithDraw,omitempty"`
	RemindUrl         string             `protobuf:"bytes,10,opt,name=remindUrl,proto3" json:"remindUrl,omitempty"` // 代理打币地址金额不够时发生提醒短信的 url
}

func (x *RelayerConfig) Reset() {
	*x = RelayerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelayerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelayerConfig) ProtoMessage() {}

func (x *RelayerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelayerConfig.ProtoReflect.Descriptor instead.
func (*RelayerConfig) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

func (x *RelayerConfig) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *RelayerConfig) GetJrpcBindAddr() string {
	if x != nil {
		return x.JrpcBindAddr
	}
	return ""
}

func (x *RelayerConfig) GetEthRelayerCfg() []*EthRelayerCfg {
	if x != nil {
		return x.EthRelayerCfg
	}
	return nil
}

func (x *RelayerConfig) GetChain33RelayerCfg() *Chain33RelayerCfg {
	if x != nil {
		return x.Chain33RelayerCfg
	}
	return nil
}

func (x *RelayerConfig) GetLog() *Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *RelayerConfig) GetDbdriver() string {
	if x != nil {
		return x.Dbdriver
	}
	return ""
}

func (x *RelayerConfig) GetDbPath() string {
	if x != nil {
		return x.DbPath
	}
	return ""
}

func (x *RelayerConfig) GetDbCache() int32 {
	if x != nil {
		return x.DbCache
	}
	return 0
}

func (x *RelayerConfig) GetProcessWithDraw() bool {
	if x != nil {
		return x.ProcessWithDraw
	}
	return false
}

func (x *RelayerConfig) GetRemindUrl() string {
	if x != nil {
		return x.RemindUrl
	}
	return ""
}

type SyncTxReceiptConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chain33Host       string   `protobuf:"bytes,1,opt,name=chain33host,proto3" json:"chain33host,omitempty"`
	PushHost          string   `protobuf:"bytes,2,opt,name=pushHost,proto3" json:"pushHost,omitempty"`
	PushName          string   `protobuf:"bytes,3,opt,name=pushName,proto3" json:"pushName,omitempty"`
	PushBind          string   `protobuf:"bytes,4,opt,name=pushBind,proto3" json:"pushBind,omitempty"`
	StartSyncHeight   int64    `protobuf:"varint,5,opt,name=startSyncHeight,proto3" json:"startSyncHeight,omitempty"`
	StartSyncSequence int64    `protobuf:"varint,6,opt,name=startSyncSequence,proto3" json:"startSyncSequence,omitempty"`
	StartSyncHash     string   `protobuf:"bytes,7,opt,name=startSyncHash,proto3" json:"startSyncHash,omitempty"`
	Contracts         []string `protobuf:"bytes,8,rep,name=contracts,proto3" json:"contracts,omitempty"`
}

func (x *SyncTxReceiptConfig) Reset() {
	*x = SyncTxReceiptConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncTxReceiptConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncTxReceiptConfig) ProtoMessage() {}

func (x *SyncTxReceiptConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncTxReceiptConfig.ProtoReflect.Descriptor instead.
func (*SyncTxReceiptConfig) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5}
}

func (x *SyncTxReceiptConfig) GetChain33Host() string {
	if x != nil {
		return x.Chain33Host
	}
	return ""
}

func (x *SyncTxReceiptConfig) GetPushHost() string {
	if x != nil {
		return x.PushHost
	}
	return ""
}

func (x *SyncTxReceiptConfig) GetPushName() string {
	if x != nil {
		return x.PushName
	}
	return ""
}

func (x *SyncTxReceiptConfig) GetPushBind() string {
	if x != nil {
		return x.PushBind
	}
	return ""
}

func (x *SyncTxReceiptConfig) GetStartSyncHeight() int64 {
	if x != nil {
		return x.StartSyncHeight
	}
	return 0
}

func (x *SyncTxReceiptConfig) GetStartSyncSequence() int64 {
	if x != nil {
		return x.StartSyncSequence
	}
	return 0
}

func (x *SyncTxReceiptConfig) GetStartSyncHash() string {
	if x != nil {
		return x.StartSyncHash
	}
	return ""
}

func (x *SyncTxReceiptConfig) GetContracts() []string {
	if x != nil {
		return x.Contracts
	}
	return nil
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xdc, 0x02, 0x0a, 0x0c, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x78,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x33,
	0x33, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x33, 0x33, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68,
	0x48, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x42, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x42, 0x69, 0x6e, 0x64, 0x12, 0x26, 0x0a, 0x0e,
	0x6d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x65,
	0x67, 0x72, 0x65, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x66, 0x65, 0x74, 0x63, 0x68, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x4d, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x13, 0x66, 0x65, 0x74, 0x63, 0x68, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x4d, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x2c, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63,
	0x48, 0x61, 0x73, 0x68, 0x22, 0xc1, 0x02, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x67, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x6f, 0x67, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x6f, 0x67, 0x43,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x6c, 0x6f, 0x67, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x6d, 0x61, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x6d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x85, 0x02, 0x0a, 0x0d, 0x45, 0x74, 0x68,
	0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x66, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x74,
	0x68, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x74, 0x68, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x65, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x0e, 0x65, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43,
	0x6c, 0x69, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x12, 0x26, 0x0a, 0x0e, 0x62, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x12, 0x2c, 0x0a, 0x11, 0x65, 0x74, 0x68, 0x4d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44,
	0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x65, 0x74, 0x68,
	0x4d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x12, 0x30,
	0x0a, 0x13, 0x65, 0x74, 0x68, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x65, 0x74, 0x68,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x22, 0xce, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x33, 0x33, 0x52, 0x65, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x43, 0x66, 0x67, 0x12, 0x37, 0x0a, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x54, 0x78,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x54, 0x78, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x38, 0x0a, 0x17, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x4f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x33, 0x33, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x17, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x4f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x33, 0x33, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x49, 0x44, 0x34, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x33, 0x33, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x34, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x33,
	0x33, 0x22, 0x81, 0x03, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6a, 0x72, 0x70,
	0x63, 0x42, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6a, 0x72, 0x70, 0x63, 0x42, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x12, 0x3a, 0x0a,
	0x0d, 0x65, 0x74, 0x68, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x66, 0x67, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x45, 0x74, 0x68,
	0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x66, 0x67, 0x52, 0x0d, 0x65, 0x74, 0x68, 0x52,
	0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x66, 0x67, 0x12, 0x46, 0x0a, 0x11, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x33, 0x33, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x66, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x33, 0x33, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x66, 0x67, 0x52, 0x11,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x33, 0x33, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x43, 0x66,
	0x67, 0x12, 0x1c, 0x0a, 0x03, 0x6c, 0x6f, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x62, 0x50, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x62, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x64, 0x62, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x57, 0x69, 0x74, 0x68, 0x44, 0x72, 0x61, 0x77,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x57,
	0x69, 0x74, 0x68, 0x44, 0x72, 0x61, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x6d, 0x69, 0x6e,
	0x64, 0x55, 0x72, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x6d, 0x69,
	0x6e, 0x64, 0x55, 0x72, 0x6c, 0x22, 0xa7, 0x02, 0x0a, 0x13, 0x53, 0x79, 0x6e, 0x63, 0x54, 0x78,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x33, 0x33, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x33, 0x33, 0x68, 0x6f, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x75, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x75, 0x73, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x75, 0x73, 0x68, 0x42,
	0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x75, 0x73, 0x68, 0x42,
	0x69, 0x6e, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63,
	0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2c, 0x0a,
	0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53,
	0x79, 0x6e, 0x63, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x79, 0x6e, 0x63, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_config_proto_goTypes = []interface{}{
	(*SyncTxConfig)(nil),        // 0: types.SyncTxConfig
	(*Log)(nil),                 // 1: types.Log
	(*EthRelayerCfg)(nil),       // 2: types.EthRelayerCfg
	(*Chain33RelayerCfg)(nil),   // 3: types.Chain33RelayerCfg
	(*RelayerConfig)(nil),       // 4: types.RelayerConfig
	(*SyncTxReceiptConfig)(nil), // 5: types.SyncTxReceiptConfig
}
var file_config_proto_depIdxs = []int32{
	0, // 0: types.Chain33RelayerCfg.syncTxConfig:type_name -> types.SyncTxConfig
	2, // 1: types.RelayerConfig.ethRelayerCfg:type_name -> types.EthRelayerCfg
	3, // 2: types.RelayerConfig.chain33RelayerCfg:type_name -> types.Chain33RelayerCfg
	1, // 3: types.RelayerConfig.log:type_name -> types.Log
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncTxConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EthRelayerCfg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chain33RelayerCfg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelayerConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncTxReceiptConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
