// *** WARNING: this file was generated by the Lumi IDL Compiler (CLIDL). ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package lambda

import (
    "errors"

    pbempty "github.com/golang/protobuf/ptypes/empty"
    pbstruct "github.com/golang/protobuf/ptypes/struct"
    "golang.org/x/net/context"

    "github.com/pulumi/lumi/pkg/resource"
    "github.com/pulumi/lumi/pkg/tokens"
    "github.com/pulumi/lumi/pkg/util/contract"
    "github.com/pulumi/lumi/pkg/util/mapper"
    "github.com/pulumi/lumi/sdk/go/pkg/lumirpc"

    __aws "github.com/pulumi/lumi/lib/aws/rpc"
)

/* RPC stubs for Permission resource provider */

// PermissionToken is the type token corresponding to the Permission package type.
const PermissionToken = tokens.Type("aws:lambda/permission:Permission")

// PermissionProviderOps is a pluggable interface for Permission-related management functionality.
type PermissionProviderOps interface {
    Check(ctx context.Context, obj *Permission) ([]mapper.FieldError, error)
    Create(ctx context.Context, obj *Permission) (resource.ID, error)
    Get(ctx context.Context, id resource.ID) (*Permission, error)
    InspectChange(ctx context.Context,
        id resource.ID, old *Permission, new *Permission, diff *resource.ObjectDiff) ([]string, error)
    Update(ctx context.Context,
        id resource.ID, old *Permission, new *Permission, diff *resource.ObjectDiff) error
    Delete(ctx context.Context, id resource.ID) error
}

// PermissionProvider is a dynamic gRPC-based plugin for managing Permission resources.
type PermissionProvider struct {
    ops PermissionProviderOps
}

// NewPermissionProvider allocates a resource provider that delegates to a ops instance.
func NewPermissionProvider(ops PermissionProviderOps) lumirpc.ResourceProviderServer {
    contract.Assert(ops != nil)
    return &PermissionProvider{ops: ops}
}

func (p *PermissionProvider) Check(
    ctx context.Context, req *lumirpc.CheckRequest) (*lumirpc.CheckResponse, error) {
    contract.Assert(req.GetType() == string(PermissionToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr == nil || len(decerr.Failures()) == 0 {
        failures, err := p.ops.Check(ctx, obj)
        if err != nil {
            return nil, err
        }
        if len(failures) > 0 {
            decerr = mapper.NewDecodeErr(failures)
        }
    }
    return resource.NewCheckResponse(decerr), nil
}

func (p *PermissionProvider) Name(
    ctx context.Context, req *lumirpc.NameRequest) (*lumirpc.NameResponse, error) {
    contract.Assert(req.GetType() == string(PermissionToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    if obj.Name == "" {
        return nil, errors.New("Name property cannot be empty")
    }
    return &lumirpc.NameResponse{Name: obj.Name}, nil
}

func (p *PermissionProvider) Create(
    ctx context.Context, req *lumirpc.CreateRequest) (*lumirpc.CreateResponse, error) {
    contract.Assert(req.GetType() == string(PermissionToken))
    obj, _, decerr := p.Unmarshal(req.GetProperties())
    if decerr != nil {
        return nil, decerr
    }
    id, err := p.ops.Create(ctx, obj)
    if err != nil {
        return nil, err
    }
    return &lumirpc.CreateResponse{
        Id:   string(id),
    }, nil
}

func (p *PermissionProvider) Get(
    ctx context.Context, req *lumirpc.GetRequest) (*lumirpc.GetResponse, error) {
    contract.Assert(req.GetType() == string(PermissionToken))
    id := resource.ID(req.GetId())
    obj, err := p.ops.Get(ctx, id)
    if err != nil {
        return nil, err
    }
    return &lumirpc.GetResponse{
        Properties: resource.MarshalProperties(
            nil, resource.NewPropertyMap(obj), resource.MarshalOptions{}),
    }, nil
}

func (p *PermissionProvider) InspectChange(
    ctx context.Context, req *lumirpc.ChangeRequest) (*lumirpc.InspectChangeResponse, error) {
    contract.Assert(req.GetType() == string(PermissionToken))
    id := resource.ID(req.GetId())
    old, oldprops, decerr := p.Unmarshal(req.GetOlds())
    if decerr != nil {
        return nil, decerr
    }
    new, newprops, decerr := p.Unmarshal(req.GetNews())
    if decerr != nil {
        return nil, decerr
    }
    var replaces []string
    diff := oldprops.Diff(newprops)
    if diff != nil {
        if diff.Changed("name") {
            replaces = append(replaces, "name")
        }
        if diff.Changed("action") {
            replaces = append(replaces, "action")
        }
        if diff.Changed("function") {
            replaces = append(replaces, "function")
        }
        if diff.Changed("principal") {
            replaces = append(replaces, "principal")
        }
        if diff.Changed("sourceAccount") {
            replaces = append(replaces, "sourceAccount")
        }
        if diff.Changed("sourceARN") {
            replaces = append(replaces, "sourceARN")
        }
    }
    more, err := p.ops.InspectChange(ctx, id, old, new, diff)
    if err != nil {
        return nil, err
    }
    return &lumirpc.InspectChangeResponse{
        Replaces: append(replaces, more...),
    }, err
}

func (p *PermissionProvider) Update(
    ctx context.Context, req *lumirpc.ChangeRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(PermissionToken))
    id := resource.ID(req.GetId())
    old, oldprops, err := p.Unmarshal(req.GetOlds())
    if err != nil {
        return nil, err
    }
    new, newprops, err := p.Unmarshal(req.GetNews())
    if err != nil {
        return nil, err
    }
    diff := oldprops.Diff(newprops)
    if err := p.ops.Update(ctx, id, old, new, diff); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *PermissionProvider) Delete(
    ctx context.Context, req *lumirpc.DeleteRequest) (*pbempty.Empty, error) {
    contract.Assert(req.GetType() == string(PermissionToken))
    id := resource.ID(req.GetId())
    if err := p.ops.Delete(ctx, id); err != nil {
        return nil, err
    }
    return &pbempty.Empty{}, nil
}

func (p *PermissionProvider) Unmarshal(
    v *pbstruct.Struct) (*Permission, resource.PropertyMap, mapper.DecodeError) {
    var obj Permission
    props := resource.UnmarshalProperties(v)
    result := mapper.MapIU(props.Mappable(), &obj)
    return &obj, props, result
}

/* Marshalable Permission structure(s) */

// Permission is a marshalable representation of its corresponding IDL type.
type Permission struct {
    Name string `json:"name"`
    Action string `json:"action"`
    Function resource.ID `json:"function"`
    Principal string `json:"principal"`
    SourceAccount *string `json:"sourceAccount,omitempty"`
    SourceARN *__aws.ARN `json:"sourceARN,omitempty"`
}

// Permission's properties have constants to make dealing with diffs and property bags easier.
const (
    Permission_Name = "name"
    Permission_Action = "action"
    Permission_Function = "function"
    Permission_Principal = "principal"
    Permission_SourceAccount = "sourceAccount"
    Permission_SourceARN = "sourceARN"
)


