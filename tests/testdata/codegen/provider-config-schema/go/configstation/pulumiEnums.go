// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package configstation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Color string

const (
	ColorBlue = Color("blue")
	ColorRed  = Color("red")
)

func (Color) ElementType() reflect.Type {
	return reflect.TypeOf((*Color)(nil)).Elem()
}

func (e Color) ToColorOutput() ColorOutput {
	return pulumi.ToOutput(e).(ColorOutput)
}

func (e Color) ToColorOutputWithContext(ctx context.Context) ColorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ColorOutput)
}

func (e Color) ToColorPtrOutput() ColorPtrOutput {
	return e.ToColorPtrOutputWithContext(context.Background())
}

func (e Color) ToColorPtrOutputWithContext(ctx context.Context) ColorPtrOutput {
	return Color(e).ToColorOutputWithContext(ctx).ToColorPtrOutputWithContext(ctx)
}

func (e Color) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Color) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Color) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Color) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ColorOutput struct{ *pulumi.OutputState }

func (ColorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Color)(nil)).Elem()
}

func (o ColorOutput) ToColorOutput() ColorOutput {
	return o
}

func (o ColorOutput) ToColorOutputWithContext(ctx context.Context) ColorOutput {
	return o
}

func (o ColorOutput) ToColorPtrOutput() ColorPtrOutput {
	return o.ToColorPtrOutputWithContext(context.Background())
}

func (o ColorOutput) ToColorPtrOutputWithContext(ctx context.Context) ColorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Color) *Color {
		return &v
	}).(ColorPtrOutput)
}

func (o ColorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ColorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Color) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ColorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ColorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Color) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ColorPtrOutput struct{ *pulumi.OutputState }

func (ColorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Color)(nil)).Elem()
}

func (o ColorPtrOutput) ToColorPtrOutput() ColorPtrOutput {
	return o
}

func (o ColorPtrOutput) ToColorPtrOutputWithContext(ctx context.Context) ColorPtrOutput {
	return o
}

func (o ColorPtrOutput) Elem() ColorOutput {
	return o.ApplyT(func(v *Color) Color {
		if v != nil {
			return *v
		}
		var ret Color
		return ret
	}).(ColorOutput)
}

func (o ColorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ColorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Color) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ColorInput is an input type that accepts values of the Color enum
// A concrete instance of `ColorInput` can be one of the following:
//
//	ColorBlue
//	ColorRed
type ColorInput interface {
	pulumi.Input

	ToColorOutput() ColorOutput
	ToColorOutputWithContext(context.Context) ColorOutput
}

var colorPtrType = reflect.TypeOf((**Color)(nil)).Elem()

type ColorPtrInput interface {
	pulumi.Input

	ToColorPtrOutput() ColorPtrOutput
	ToColorPtrOutputWithContext(context.Context) ColorPtrOutput
}

type colorPtr string

func ColorPtr(v string) ColorPtrInput {
	return (*colorPtr)(&v)
}

func (*colorPtr) ElementType() reflect.Type {
	return colorPtrType
}

func (in *colorPtr) ToColorPtrOutput() ColorPtrOutput {
	return pulumi.ToOutput(in).(ColorPtrOutput)
}

func (in *colorPtr) ToColorPtrOutputWithContext(ctx context.Context) ColorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ColorPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ColorInput)(nil)).Elem(), Color("blue"))
	pulumi.RegisterInputType(reflect.TypeOf((*ColorPtrInput)(nil)).Elem(), Color("blue"))
	pulumi.RegisterOutputType(ColorOutput{})
	pulumi.RegisterOutputType(ColorPtrOutput{})
}