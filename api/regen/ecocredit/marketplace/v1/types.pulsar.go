// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package marketplacev1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_AllowDenomProposal             protoreflect.MessageDescriptor
	fd_AllowDenomProposal_title       protoreflect.FieldDescriptor
	fd_AllowDenomProposal_description protoreflect.FieldDescriptor
	fd_AllowDenomProposal_denom       protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_marketplace_v1_types_proto_init()
	md_AllowDenomProposal = File_regen_ecocredit_marketplace_v1_types_proto.Messages().ByName("AllowDenomProposal")
	fd_AllowDenomProposal_title = md_AllowDenomProposal.Fields().ByName("title")
	fd_AllowDenomProposal_description = md_AllowDenomProposal.Fields().ByName("description")
	fd_AllowDenomProposal_denom = md_AllowDenomProposal.Fields().ByName("denom")
}

var _ protoreflect.Message = (*fastReflection_AllowDenomProposal)(nil)

type fastReflection_AllowDenomProposal AllowDenomProposal

func (x *AllowDenomProposal) ProtoReflect() protoreflect.Message {
	return (*fastReflection_AllowDenomProposal)(x)
}

func (x *AllowDenomProposal) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_marketplace_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_AllowDenomProposal_messageType fastReflection_AllowDenomProposal_messageType
var _ protoreflect.MessageType = fastReflection_AllowDenomProposal_messageType{}

type fastReflection_AllowDenomProposal_messageType struct{}

func (x fastReflection_AllowDenomProposal_messageType) Zero() protoreflect.Message {
	return (*fastReflection_AllowDenomProposal)(nil)
}
func (x fastReflection_AllowDenomProposal_messageType) New() protoreflect.Message {
	return new(fastReflection_AllowDenomProposal)
}
func (x fastReflection_AllowDenomProposal_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_AllowDenomProposal
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_AllowDenomProposal) Descriptor() protoreflect.MessageDescriptor {
	return md_AllowDenomProposal
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_AllowDenomProposal) Type() protoreflect.MessageType {
	return _fastReflection_AllowDenomProposal_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_AllowDenomProposal) New() protoreflect.Message {
	return new(fastReflection_AllowDenomProposal)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_AllowDenomProposal) Interface() protoreflect.ProtoMessage {
	return (*AllowDenomProposal)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_AllowDenomProposal) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Title != "" {
		value := protoreflect.ValueOfString(x.Title)
		if !f(fd_AllowDenomProposal_title, value) {
			return
		}
	}
	if x.Description != "" {
		value := protoreflect.ValueOfString(x.Description)
		if !f(fd_AllowDenomProposal_description, value) {
			return
		}
	}
	if x.Denom != nil {
		value := protoreflect.ValueOfMessage(x.Denom.ProtoReflect())
		if !f(fd_AllowDenomProposal_denom, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_AllowDenomProposal) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.title":
		return x.Title != ""
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.description":
		return x.Description != ""
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.denom":
		return x.Denom != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1.AllowDenomProposal"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1.AllowDenomProposal does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AllowDenomProposal) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.title":
		x.Title = ""
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.description":
		x.Description = ""
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.denom":
		x.Denom = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1.AllowDenomProposal"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1.AllowDenomProposal does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_AllowDenomProposal) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.title":
		value := x.Title
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.description":
		value := x.Description
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.denom":
		value := x.Denom
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1.AllowDenomProposal"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1.AllowDenomProposal does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AllowDenomProposal) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.title":
		x.Title = value.Interface().(string)
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.description":
		x.Description = value.Interface().(string)
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.denom":
		x.Denom = value.Message().Interface().(*AllowedDenom)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1.AllowDenomProposal"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1.AllowDenomProposal does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AllowDenomProposal) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.denom":
		if x.Denom == nil {
			x.Denom = new(AllowedDenom)
		}
		return protoreflect.ValueOfMessage(x.Denom.ProtoReflect())
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.title":
		panic(fmt.Errorf("field title of message regen.ecocredit.marketplace.v1.AllowDenomProposal is not mutable"))
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.description":
		panic(fmt.Errorf("field description of message regen.ecocredit.marketplace.v1.AllowDenomProposal is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1.AllowDenomProposal"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1.AllowDenomProposal does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_AllowDenomProposal) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.title":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.description":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.marketplace.v1.AllowDenomProposal.denom":
		m := new(AllowedDenom)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1.AllowDenomProposal"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1.AllowDenomProposal does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_AllowDenomProposal) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.marketplace.v1.AllowDenomProposal", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_AllowDenomProposal) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_AllowDenomProposal) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_AllowDenomProposal) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_AllowDenomProposal) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*AllowDenomProposal)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Title)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Description)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Denom != nil {
			l = options.Size(x.Denom)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*AllowDenomProposal)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Denom != nil {
			encoded, err := options.Marshal(x.Denom)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Description) > 0 {
			i -= len(x.Description)
			copy(dAtA[i:], x.Description)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Description)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Title) > 0 {
			i -= len(x.Title)
			copy(dAtA[i:], x.Title)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Title)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*AllowDenomProposal)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				wire |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			fieldNum := int32(wire >> 3)
			wireType := int(wire & 0x7)
			if wireType == 4 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AllowDenomProposal: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: AllowDenomProposal: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Title = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Description = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Denom == nil {
					x.Denom = &AllowedDenom{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Denom); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: regen/ecocredit/marketplace/v1/types.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AllowDenomProposal is a gov Content type for approving a denom for use in the
// marketplace.
//
// Deprecated (Since Revision 1): This message is no longer used and will be
// removed in the next version. See MsgAddAllowedDenom.
type AllowDenomProposal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// title is the title of the proposal.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description is the description of the proposal.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// denom contains coin denom information that will be added to the
	// list of allowed denoms for use in the marketplace.
	Denom *AllowedDenom `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty"`
}

func (x *AllowDenomProposal) Reset() {
	*x = AllowDenomProposal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_marketplace_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllowDenomProposal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllowDenomProposal) ProtoMessage() {}

// Deprecated: Use AllowDenomProposal.ProtoReflect.Descriptor instead.
func (*AllowDenomProposal) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_marketplace_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *AllowDenomProposal) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *AllowDenomProposal) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AllowDenomProposal) GetDenom() *AllowedDenom {
	if x != nil {
		return x.Denom
	}
	return nil
}

var File_regen_ecocredit_marketplace_v1_types_proto protoreflect.FileDescriptor

var file_regen_ecocredit_marketplace_v1_types_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x67, 0x6f,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2a, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65,
	0x64, 0x69, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96,
	0x01, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x50, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a,
	0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x64, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f,
	0x6d, 0x3a, 0x04, 0x98, 0xa0, 0x1f, 0x00, 0x42, 0xa3, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63,
	0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x45, 0x4d, 0xaa, 0x02, 0x1e, 0x52, 0x65, 0x67,
	0x65, 0x6e, 0x2e, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x4d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x52, 0x65,
	0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2a, 0x52,
	0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x4d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x52, 0x65, 0x67, 0x65,
	0x6e, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x3a, 0x3a, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_marketplace_v1_types_proto_rawDescOnce sync.Once
	file_regen_ecocredit_marketplace_v1_types_proto_rawDescData = file_regen_ecocredit_marketplace_v1_types_proto_rawDesc
)

func file_regen_ecocredit_marketplace_v1_types_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_marketplace_v1_types_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_marketplace_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_marketplace_v1_types_proto_rawDescData)
	})
	return file_regen_ecocredit_marketplace_v1_types_proto_rawDescData
}

var file_regen_ecocredit_marketplace_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_regen_ecocredit_marketplace_v1_types_proto_goTypes = []interface{}{
	(*AllowDenomProposal)(nil), // 0: regen.ecocredit.marketplace.v1.AllowDenomProposal
	(*AllowedDenom)(nil),       // 1: regen.ecocredit.marketplace.v1.AllowedDenom
}
var file_regen_ecocredit_marketplace_v1_types_proto_depIdxs = []int32{
	1, // 0: regen.ecocredit.marketplace.v1.AllowDenomProposal.denom:type_name -> regen.ecocredit.marketplace.v1.AllowedDenom
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_marketplace_v1_types_proto_init() }
func file_regen_ecocredit_marketplace_v1_types_proto_init() {
	if File_regen_ecocredit_marketplace_v1_types_proto != nil {
		return
	}
	file_regen_ecocredit_marketplace_v1_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_marketplace_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllowDenomProposal); i {
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
			RawDescriptor: file_regen_ecocredit_marketplace_v1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_ecocredit_marketplace_v1_types_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_marketplace_v1_types_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_marketplace_v1_types_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_marketplace_v1_types_proto = out.File
	file_regen_ecocredit_marketplace_v1_types_proto_rawDesc = nil
	file_regen_ecocredit_marketplace_v1_types_proto_goTypes = nil
	file_regen_ecocredit_marketplace_v1_types_proto_depIdxs = nil
}
