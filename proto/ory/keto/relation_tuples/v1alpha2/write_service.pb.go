// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: ory/keto/relation_tuples/v1alpha2/write_service.proto

package rts

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RelationTupleDelta_Action int32

const (
	// Unspecified.
	// The `TransactRelationTuples` RPC ignores this
	// RelationTupleDelta if an action was unspecified.
	RelationTupleDelta_ACTION_UNSPECIFIED RelationTupleDelta_Action = 0
	// Insertion of a new RelationTuple.
	// It is ignored if already existing.
	RelationTupleDelta_ACTION_INSERT RelationTupleDelta_Action = 1
	// Deletion of the RelationTuple.
	// It is ignored if it does not exist.
	RelationTupleDelta_ACTION_DELETE RelationTupleDelta_Action = 2
)

// Enum value maps for RelationTupleDelta_Action.
var (
	RelationTupleDelta_Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "ACTION_INSERT",
		2: "ACTION_DELETE",
	}
	RelationTupleDelta_Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"ACTION_INSERT":      1,
		"ACTION_DELETE":      2,
	}
)

func (x RelationTupleDelta_Action) Enum() *RelationTupleDelta_Action {
	p := new(RelationTupleDelta_Action)
	*p = x
	return p
}

func (x RelationTupleDelta_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelationTupleDelta_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_enumTypes[0].Descriptor()
}

func (RelationTupleDelta_Action) Type() protoreflect.EnumType {
	return &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_enumTypes[0]
}

func (x RelationTupleDelta_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelationTupleDelta_Action.Descriptor instead.
func (RelationTupleDelta_Action) EnumDescriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{1, 0}
}

// The request of a WriteService.TransactRelationTuples RPC.
type TransactRelationTuplesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The write delta for the relationships operated in one single transaction.
	// Either all actions succeed or no change takes effect on error.
	RelationTupleDeltas []*RelationTupleDelta `protobuf:"bytes,1,rep,name=relation_tuple_deltas,json=relationTupleDeltas,proto3" json:"relation_tuple_deltas,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *TransactRelationTuplesRequest) Reset() {
	*x = TransactRelationTuplesRequest{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactRelationTuplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactRelationTuplesRequest) ProtoMessage() {}

func (x *TransactRelationTuplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactRelationTuplesRequest.ProtoReflect.Descriptor instead.
func (*TransactRelationTuplesRequest) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{0}
}

func (x *TransactRelationTuplesRequest) GetRelationTupleDeltas() []*RelationTupleDelta {
	if x != nil {
		return x.RelationTupleDeltas
	}
	return nil
}

// Write-delta for a TransactRelationTuplesRequest.
type RelationTupleDelta struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The action to do on the RelationTuple.
	Action RelationTupleDelta_Action `protobuf:"varint,1,opt,name=action,proto3,enum=ory.keto.relation_tuples.v1alpha2.RelationTupleDelta_Action" json:"action,omitempty"`
	// The target RelationTuple.
	RelationTuple *RelationTuple `protobuf:"bytes,2,opt,name=relation_tuple,json=relationTuple,proto3" json:"relation_tuple,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RelationTupleDelta) Reset() {
	*x = RelationTupleDelta{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RelationTupleDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelationTupleDelta) ProtoMessage() {}

func (x *RelationTupleDelta) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelationTupleDelta.ProtoReflect.Descriptor instead.
func (*RelationTupleDelta) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{1}
}

func (x *RelationTupleDelta) GetAction() RelationTupleDelta_Action {
	if x != nil {
		return x.Action
	}
	return RelationTupleDelta_ACTION_UNSPECIFIED
}

func (x *RelationTupleDelta) GetRelationTuple() *RelationTuple {
	if x != nil {
		return x.RelationTuple
	}
	return nil
}

// The response of a WriteService.TransactRelationTuples rpc.
type TransactRelationTuplesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This field is not implemented yet and has no effect.
	// <!--
	// The list of the new latest snapshot tokens of the affected RelationTuple,
	// with the same index as specified in the `relation_tuple_deltas` field of
	// the TransactRelationTuplesRequest request.
	//
	// If the RelationTupleDelta_Action was DELETE
	// the snaptoken is empty at the same index.
	// -->
	Snaptokens    []string `protobuf:"bytes,1,rep,name=snaptokens,proto3" json:"snaptokens,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactRelationTuplesResponse) Reset() {
	*x = TransactRelationTuplesResponse{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactRelationTuplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactRelationTuplesResponse) ProtoMessage() {}

func (x *TransactRelationTuplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactRelationTuplesResponse.ProtoReflect.Descriptor instead.
func (*TransactRelationTuplesResponse) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{2}
}

func (x *TransactRelationTuplesResponse) GetSnaptokens() []string {
	if x != nil {
		return x.Snaptokens
	}
	return nil
}

type DeleteRelationTuplesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/write_service.proto.
	Query         *DeleteRelationTuplesRequest_Query `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	RelationQuery *RelationQuery                     `protobuf:"bytes,2,opt,name=relation_query,json=relationQuery,proto3" json:"relation_query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRelationTuplesRequest) Reset() {
	*x = DeleteRelationTuplesRequest{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRelationTuplesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationTuplesRequest) ProtoMessage() {}

func (x *DeleteRelationTuplesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationTuplesRequest.ProtoReflect.Descriptor instead.
func (*DeleteRelationTuplesRequest) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{3}
}

// Deprecated: Marked as deprecated in ory/keto/relation_tuples/v1alpha2/write_service.proto.
func (x *DeleteRelationTuplesRequest) GetQuery() *DeleteRelationTuplesRequest_Query {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *DeleteRelationTuplesRequest) GetRelationQuery() *RelationQuery {
	if x != nil {
		return x.RelationQuery
	}
	return nil
}

type DeleteRelationTuplesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRelationTuplesResponse) Reset() {
	*x = DeleteRelationTuplesResponse{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRelationTuplesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationTuplesResponse) ProtoMessage() {}

func (x *DeleteRelationTuplesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationTuplesResponse.ProtoReflect.Descriptor instead.
func (*DeleteRelationTuplesResponse) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{4}
}

// The query for deleting relationships
type DeleteRelationTuplesRequest_Query struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The namespace to query.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Optional. The object to query for.
	Object string `protobuf:"bytes,2,opt,name=object,proto3" json:"object,omitempty"`
	// Optional. The relation to query for.
	Relation string `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
	// Optional. The subject to query for.
	Subject       *Subject `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRelationTuplesRequest_Query) Reset() {
	*x = DeleteRelationTuplesRequest_Query{}
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRelationTuplesRequest_Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRelationTuplesRequest_Query) ProtoMessage() {}

func (x *DeleteRelationTuplesRequest_Query) ProtoReflect() protoreflect.Message {
	mi := &file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRelationTuplesRequest_Query.ProtoReflect.Descriptor instead.
func (*DeleteRelationTuplesRequest_Query) Descriptor() ([]byte, []int) {
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *DeleteRelationTuplesRequest_Query) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DeleteRelationTuplesRequest_Query) GetObject() string {
	if x != nil {
		return x.Object
	}
	return ""
}

func (x *DeleteRelationTuplesRequest_Query) GetRelation() string {
	if x != nil {
		return x.Relation
	}
	return ""
}

func (x *DeleteRelationTuplesRequest_Query) GetSubject() *Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

var File_ory_keto_relation_tuples_v1alpha2_write_service_proto protoreflect.FileDescriptor

const file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDesc = "" +
	"\n" +
	"5ory/keto/relation_tuples/v1alpha2/write_service.proto\x12!ory.keto.relation_tuples.v1alpha2\x1a7ory/keto/relation_tuples/v1alpha2/relation_tuples.proto\"\x8a\x01\n" +
	"\x1dTransactRelationTuplesRequest\x12i\n" +
	"\x15relation_tuple_deltas\x18\x01 \x03(\v25.ory.keto.relation_tuples.v1alpha2.RelationTupleDeltaR\x13relationTupleDeltas\"\x8b\x02\n" +
	"\x12RelationTupleDelta\x12T\n" +
	"\x06action\x18\x01 \x01(\x0e2<.ory.keto.relation_tuples.v1alpha2.RelationTupleDelta.ActionR\x06action\x12W\n" +
	"\x0erelation_tuple\x18\x02 \x01(\v20.ory.keto.relation_tuples.v1alpha2.RelationTupleR\rrelationTuple\"F\n" +
	"\x06Action\x12\x16\n" +
	"\x12ACTION_UNSPECIFIED\x10\x00\x12\x11\n" +
	"\rACTION_INSERT\x10\x01\x12\x11\n" +
	"\rACTION_DELETE\x10\x02\"@\n" +
	"\x1eTransactRelationTuplesResponse\x12\x1e\n" +
	"\n" +
	"snaptokens\x18\x01 \x03(\tR\n" +
	"snaptokens\"\xf8\x02\n" +
	"\x1bDeleteRelationTuplesRequest\x12^\n" +
	"\x05query\x18\x01 \x01(\v2D.ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest.QueryB\x02\x18\x01R\x05query\x12W\n" +
	"\x0erelation_query\x18\x02 \x01(\v20.ory.keto.relation_tuples.v1alpha2.RelationQueryR\rrelationQuery\x1a\x9f\x01\n" +
	"\x05Query\x12\x1c\n" +
	"\tnamespace\x18\x01 \x01(\tR\tnamespace\x12\x16\n" +
	"\x06object\x18\x02 \x01(\tR\x06object\x12\x1a\n" +
	"\brelation\x18\x03 \x01(\tR\brelation\x12D\n" +
	"\asubject\x18\x04 \x01(\v2*.ory.keto.relation_tuples.v1alpha2.SubjectR\asubject\"\x1e\n" +
	"\x1cDeleteRelationTuplesResponse2\xc8\x02\n" +
	"\fWriteService\x12\x9d\x01\n" +
	"\x16TransactRelationTuples\x12@.ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesRequest\x1aA.ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesResponse\x12\x97\x01\n" +
	"\x14DeleteRelationTuples\x12>.ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest\x1a?.ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesResponseB\xc2\x01\n" +
	"$sh.ory.keto.relation_tuples.v1alpha2B\x11WriteServiceProtoP\x01Z?github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2;rts\xaa\x02 Ory.Keto.RelationTuples.v1alpha2\xca\x02 Ory\\Keto\\RelationTuples\\v1alpha2b\x06proto3"

var (
	file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescOnce sync.Once
	file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescData []byte
)

func file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescGZIP() []byte {
	file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescOnce.Do(func() {
		file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDesc), len(file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDesc)))
	})
	return file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDescData
}

var file_ory_keto_relation_tuples_v1alpha2_write_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ory_keto_relation_tuples_v1alpha2_write_service_proto_goTypes = []any{
	(RelationTupleDelta_Action)(0),            // 0: ory.keto.relation_tuples.v1alpha2.RelationTupleDelta.Action
	(*TransactRelationTuplesRequest)(nil),     // 1: ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesRequest
	(*RelationTupleDelta)(nil),                // 2: ory.keto.relation_tuples.v1alpha2.RelationTupleDelta
	(*TransactRelationTuplesResponse)(nil),    // 3: ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesResponse
	(*DeleteRelationTuplesRequest)(nil),       // 4: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest
	(*DeleteRelationTuplesResponse)(nil),      // 5: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesResponse
	(*DeleteRelationTuplesRequest_Query)(nil), // 6: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest.Query
	(*RelationTuple)(nil),                     // 7: ory.keto.relation_tuples.v1alpha2.RelationTuple
	(*RelationQuery)(nil),                     // 8: ory.keto.relation_tuples.v1alpha2.RelationQuery
	(*Subject)(nil),                           // 9: ory.keto.relation_tuples.v1alpha2.Subject
}
var file_ory_keto_relation_tuples_v1alpha2_write_service_proto_depIdxs = []int32{
	2, // 0: ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesRequest.relation_tuple_deltas:type_name -> ory.keto.relation_tuples.v1alpha2.RelationTupleDelta
	0, // 1: ory.keto.relation_tuples.v1alpha2.RelationTupleDelta.action:type_name -> ory.keto.relation_tuples.v1alpha2.RelationTupleDelta.Action
	7, // 2: ory.keto.relation_tuples.v1alpha2.RelationTupleDelta.relation_tuple:type_name -> ory.keto.relation_tuples.v1alpha2.RelationTuple
	6, // 3: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest.query:type_name -> ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest.Query
	8, // 4: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest.relation_query:type_name -> ory.keto.relation_tuples.v1alpha2.RelationQuery
	9, // 5: ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest.Query.subject:type_name -> ory.keto.relation_tuples.v1alpha2.Subject
	1, // 6: ory.keto.relation_tuples.v1alpha2.WriteService.TransactRelationTuples:input_type -> ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesRequest
	4, // 7: ory.keto.relation_tuples.v1alpha2.WriteService.DeleteRelationTuples:input_type -> ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesRequest
	3, // 8: ory.keto.relation_tuples.v1alpha2.WriteService.TransactRelationTuples:output_type -> ory.keto.relation_tuples.v1alpha2.TransactRelationTuplesResponse
	5, // 9: ory.keto.relation_tuples.v1alpha2.WriteService.DeleteRelationTuples:output_type -> ory.keto.relation_tuples.v1alpha2.DeleteRelationTuplesResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ory_keto_relation_tuples_v1alpha2_write_service_proto_init() }
func file_ory_keto_relation_tuples_v1alpha2_write_service_proto_init() {
	if File_ory_keto_relation_tuples_v1alpha2_write_service_proto != nil {
		return
	}
	file_ory_keto_relation_tuples_v1alpha2_relation_tuples_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDesc), len(file_ory_keto_relation_tuples_v1alpha2_write_service_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ory_keto_relation_tuples_v1alpha2_write_service_proto_goTypes,
		DependencyIndexes: file_ory_keto_relation_tuples_v1alpha2_write_service_proto_depIdxs,
		EnumInfos:         file_ory_keto_relation_tuples_v1alpha2_write_service_proto_enumTypes,
		MessageInfos:      file_ory_keto_relation_tuples_v1alpha2_write_service_proto_msgTypes,
	}.Build()
	File_ory_keto_relation_tuples_v1alpha2_write_service_proto = out.File
	file_ory_keto_relation_tuples_v1alpha2_write_service_proto_goTypes = nil
	file_ory_keto_relation_tuples_v1alpha2_write_service_proto_depIdxs = nil
}
