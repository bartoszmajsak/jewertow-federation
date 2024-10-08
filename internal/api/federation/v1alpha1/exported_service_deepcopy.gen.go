// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1alpha1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using ExportedService within kubernetes types, where deepcopy-gen is used.
func (in *ExportedService) DeepCopyInto(out *ExportedService) {
	p := proto.Clone(in).(*ExportedService)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportedService. Required by controller-gen.
func (in *ExportedService) DeepCopy() *ExportedService {
	if in == nil {
		return nil
	}
	out := new(ExportedService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ExportedService. Required by controller-gen.
func (in *ExportedService) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using ServicePort within kubernetes types, where deepcopy-gen is used.
func (in *ServicePort) DeepCopyInto(out *ServicePort) {
	p := proto.Clone(in).(*ServicePort)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePort. Required by controller-gen.
func (in *ServicePort) DeepCopy() *ServicePort {
	if in == nil {
		return nil
	}
	out := new(ServicePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new ServicePort. Required by controller-gen.
func (in *ServicePort) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
