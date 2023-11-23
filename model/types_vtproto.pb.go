// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: v0.5.0
// source: types.proto

package model

import (
	binary "encoding/binary"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	math "math"
	bits "math/bits"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *MetricMetadata) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricMetadata) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *MetricMetadata) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Unit) > 0 {
		i -= len(m.Unit)
		copy(dAtA[i:], m.Unit)
		i = encodeVarint(dAtA, i, uint64(len(m.Unit)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Help) > 0 {
		i -= len(m.Help)
		copy(dAtA[i:], m.Help)
		i = encodeVarint(dAtA, i, uint64(len(m.Help)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MetricFamilyName) > 0 {
		i -= len(m.MetricFamilyName)
		copy(dAtA[i:], m.MetricFamilyName)
		i = encodeVarint(dAtA, i, uint64(len(m.MetricFamilyName)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Sample) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sample) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Sample) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Timestamp != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x10
	}
	if m.Value != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x9
	}
	return len(dAtA) - i, nil
}

func (m *Exemplar) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Exemplar) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Exemplar) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Timestamp != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if m.Value != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
		i--
		dAtA[i] = 0x11
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Labels[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Histogram) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Histogram) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Histogram) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if vtmsg, ok := m.ZeroCount.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if vtmsg, ok := m.Count.(interface {
		MarshalToSizedBufferVT([]byte) (int, error)
	}); ok {
		size, err := vtmsg.MarshalToSizedBufferVT(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
	}
	if m.Timestamp != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x78
	}
	if m.ResetHint != 0 {
		i = encodeVarint(dAtA, i, uint64(m.ResetHint))
		i--
		dAtA[i] = 0x70
	}
	if len(m.PositiveCounts) > 0 {
		for iNdEx := len(m.PositiveCounts) - 1; iNdEx >= 0; iNdEx-- {
			f1 := math.Float64bits(float64(m.PositiveCounts[iNdEx]))
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(f1))
		}
		i = encodeVarint(dAtA, i, uint64(len(m.PositiveCounts)*8))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.PositiveDeltas) > 0 {
		var pksize3 int
		for _, num := range m.PositiveDeltas {
			pksize3 += soz(uint64(num))
		}
		i -= pksize3
		j2 := i
		for _, num := range m.PositiveDeltas {
			x4 := (uint64(num) << 1) ^ uint64((num >> 63))
			for x4 >= 1<<7 {
				dAtA[j2] = uint8(uint64(x4)&0x7f | 0x80)
				j2++
				x4 >>= 7
			}
			dAtA[j2] = uint8(x4)
			j2++
		}
		i = encodeVarint(dAtA, i, uint64(pksize3))
		i--
		dAtA[i] = 0x62
	}
	if len(m.PositiveSpans) > 0 {
		for iNdEx := len(m.PositiveSpans) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.PositiveSpans[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.NegativeCounts) > 0 {
		for iNdEx := len(m.NegativeCounts) - 1; iNdEx >= 0; iNdEx-- {
			f5 := math.Float64bits(float64(m.NegativeCounts[iNdEx]))
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(f5))
		}
		i = encodeVarint(dAtA, i, uint64(len(m.NegativeCounts)*8))
		i--
		dAtA[i] = 0x52
	}
	if len(m.NegativeDeltas) > 0 {
		var pksize7 int
		for _, num := range m.NegativeDeltas {
			pksize7 += soz(uint64(num))
		}
		i -= pksize7
		j6 := i
		for _, num := range m.NegativeDeltas {
			x8 := (uint64(num) << 1) ^ uint64((num >> 63))
			for x8 >= 1<<7 {
				dAtA[j6] = uint8(uint64(x8)&0x7f | 0x80)
				j6++
				x8 >>= 7
			}
			dAtA[j6] = uint8(x8)
			j6++
		}
		i = encodeVarint(dAtA, i, uint64(pksize7))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.NegativeSpans) > 0 {
		for iNdEx := len(m.NegativeSpans) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.NegativeSpans[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x42
		}
	}
	if m.ZeroThreshold != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.ZeroThreshold))))
		i--
		dAtA[i] = 0x29
	}
	if m.Schema != 0 {
		i = encodeVarint(dAtA, i, uint64((uint32(m.Schema)<<1)^uint32((m.Schema>>31))))
		i--
		dAtA[i] = 0x20
	}
	if m.Sum != 0 {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Sum))))
		i--
		dAtA[i] = 0x19
	}
	return len(dAtA) - i, nil
}

func (m *Histogram_CountInt) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Histogram_CountInt) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarint(dAtA, i, uint64(m.CountInt))
	i--
	dAtA[i] = 0x8
	return len(dAtA) - i, nil
}
func (m *Histogram_CountFloat) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Histogram_CountFloat) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= 8
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.CountFloat))))
	i--
	dAtA[i] = 0x11
	return len(dAtA) - i, nil
}
func (m *Histogram_ZeroCountInt) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Histogram_ZeroCountInt) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i = encodeVarint(dAtA, i, uint64(m.ZeroCountInt))
	i--
	dAtA[i] = 0x30
	return len(dAtA) - i, nil
}
func (m *Histogram_ZeroCountFloat) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Histogram_ZeroCountFloat) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	i := len(dAtA)
	i -= 8
	binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.ZeroCountFloat))))
	i--
	dAtA[i] = 0x39
	return len(dAtA) - i, nil
}
func (m *BucketSpan) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BucketSpan) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *BucketSpan) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.Length != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Length))
		i--
		dAtA[i] = 0x10
	}
	if m.Offset != 0 {
		i = encodeVarint(dAtA, i, uint64((uint32(m.Offset)<<1)^uint32((m.Offset>>31))))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TimeSeries) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeSeries) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *TimeSeries) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Histograms) > 0 {
		for iNdEx := len(m.Histograms) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Histograms[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Exemplars) > 0 {
		for iNdEx := len(m.Exemplars) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Exemplars[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Samples) > 0 {
		for iNdEx := len(m.Samples) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Samples[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Labels[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Label) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Label) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Label) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarint(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Labels) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Labels) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Labels) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Labels[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LabelMatcher) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LabelMatcher) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *LabelMatcher) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarint(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Type != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReadHints) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadHints) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ReadHints) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.RangeMs != 0 {
		i = encodeVarint(dAtA, i, uint64(m.RangeMs))
		i--
		dAtA[i] = 0x38
	}
	if m.By {
		i--
		if m.By {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Grouping) > 0 {
		for iNdEx := len(m.Grouping) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Grouping[iNdEx])
			copy(dAtA[i:], m.Grouping[iNdEx])
			i = encodeVarint(dAtA, i, uint64(len(m.Grouping[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.EndMs != 0 {
		i = encodeVarint(dAtA, i, uint64(m.EndMs))
		i--
		dAtA[i] = 0x20
	}
	if m.StartMs != 0 {
		i = encodeVarint(dAtA, i, uint64(m.StartMs))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Func) > 0 {
		i -= len(m.Func)
		copy(dAtA[i:], m.Func)
		i = encodeVarint(dAtA, i, uint64(len(m.Func)))
		i--
		dAtA[i] = 0x12
	}
	if m.StepMs != 0 {
		i = encodeVarint(dAtA, i, uint64(m.StepMs))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Chunk) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Chunk) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *Chunk) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarint(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if m.Type != 0 {
		i = encodeVarint(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxTimeMs != 0 {
		i = encodeVarint(dAtA, i, uint64(m.MaxTimeMs))
		i--
		dAtA[i] = 0x10
	}
	if m.MinTimeMs != 0 {
		i = encodeVarint(dAtA, i, uint64(m.MinTimeMs))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ChunkedSeries) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChunkedSeries) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *ChunkedSeries) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if len(m.Chunks) > 0 {
		for iNdEx := len(m.Chunks) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Chunks[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Labels) > 0 {
		for iNdEx := len(m.Labels) - 1; iNdEx >= 0; iNdEx-- {
			size, err := m.Labels[iNdEx].MarshalToSizedBufferVT(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarint(dAtA, i, uint64(size))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarint(dAtA []byte, offset int, v uint64) int {
	offset -= sov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MetricMetadata) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sov(uint64(m.Type))
	}
	l = len(m.MetricFamilyName)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Help)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Unit)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *Sample) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Value != 0 {
		n += 9
	}
	if m.Timestamp != 0 {
		n += 1 + sov(uint64(m.Timestamp))
	}
	n += len(m.unknownFields)
	return n
}

func (m *Exemplar) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.Value != 0 {
		n += 9
	}
	if m.Timestamp != 0 {
		n += 1 + sov(uint64(m.Timestamp))
	}
	n += len(m.unknownFields)
	return n
}

func (m *Histogram) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if vtmsg, ok := m.Count.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if m.Sum != 0 {
		n += 9
	}
	if m.Schema != 0 {
		n += 1 + soz(uint64(m.Schema))
	}
	if m.ZeroThreshold != 0 {
		n += 9
	}
	if vtmsg, ok := m.ZeroCount.(interface{ SizeVT() int }); ok {
		n += vtmsg.SizeVT()
	}
	if len(m.NegativeSpans) > 0 {
		for _, e := range m.NegativeSpans {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.NegativeDeltas) > 0 {
		l = 0
		for _, e := range m.NegativeDeltas {
			l += soz(uint64(e))
		}
		n += 1 + sov(uint64(l)) + l
	}
	if len(m.NegativeCounts) > 0 {
		n += 1 + sov(uint64(len(m.NegativeCounts)*8)) + len(m.NegativeCounts)*8
	}
	if len(m.PositiveSpans) > 0 {
		for _, e := range m.PositiveSpans {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.PositiveDeltas) > 0 {
		l = 0
		for _, e := range m.PositiveDeltas {
			l += soz(uint64(e))
		}
		n += 1 + sov(uint64(l)) + l
	}
	if len(m.PositiveCounts) > 0 {
		n += 1 + sov(uint64(len(m.PositiveCounts)*8)) + len(m.PositiveCounts)*8
	}
	if m.ResetHint != 0 {
		n += 1 + sov(uint64(m.ResetHint))
	}
	if m.Timestamp != 0 {
		n += 1 + sov(uint64(m.Timestamp))
	}
	n += len(m.unknownFields)
	return n
}

func (m *Histogram_CountInt) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sov(uint64(m.CountInt))
	return n
}
func (m *Histogram_CountFloat) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 9
	return n
}
func (m *Histogram_ZeroCountInt) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sov(uint64(m.ZeroCountInt))
	return n
}
func (m *Histogram_ZeroCountFloat) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 9
	return n
}
func (m *BucketSpan) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Offset != 0 {
		n += 1 + soz(uint64(m.Offset))
	}
	if m.Length != 0 {
		n += 1 + sov(uint64(m.Length))
	}
	n += len(m.unknownFields)
	return n
}

func (m *TimeSeries) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.Samples) > 0 {
		for _, e := range m.Samples {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.Exemplars) > 0 {
		for _, e := range m.Exemplars {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.Histograms) > 0 {
		for _, e := range m.Histograms {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *Label) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *Labels) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func (m *LabelMatcher) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sov(uint64(m.Type))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ReadHints) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StepMs != 0 {
		n += 1 + sov(uint64(m.StepMs))
	}
	l = len(m.Func)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	if m.StartMs != 0 {
		n += 1 + sov(uint64(m.StartMs))
	}
	if m.EndMs != 0 {
		n += 1 + sov(uint64(m.EndMs))
	}
	if len(m.Grouping) > 0 {
		for _, s := range m.Grouping {
			l = len(s)
			n += 1 + l + sov(uint64(l))
		}
	}
	if m.By {
		n += 2
	}
	if m.RangeMs != 0 {
		n += 1 + sov(uint64(m.RangeMs))
	}
	n += len(m.unknownFields)
	return n
}

func (m *Chunk) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinTimeMs != 0 {
		n += 1 + sov(uint64(m.MinTimeMs))
	}
	if m.MaxTimeMs != 0 {
		n += 1 + sov(uint64(m.MaxTimeMs))
	}
	if m.Type != 0 {
		n += 1 + sov(uint64(m.Type))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sov(uint64(l))
	}
	n += len(m.unknownFields)
	return n
}

func (m *ChunkedSeries) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for _, e := range m.Labels {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	if len(m.Chunks) > 0 {
		for _, e := range m.Chunks {
			l = e.SizeVT()
			n += 1 + l + sov(uint64(l))
		}
	}
	n += len(m.unknownFields)
	return n
}

func sov(x uint64) (n int) {
	return (bits.Len64(x|1) + 6) / 7
}
func soz(x uint64) (n int) {
	return sov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
