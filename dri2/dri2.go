package dri2

/*
	This file was generated by dri2.xml on May 10 2012 4:20:27pm EDT.
	This file is automatically generated. Edit at your peril!
*/

import (
	"github.com/BurntSushi/xgb"

	"github.com/BurntSushi/xgb/xproto"
)

// Init must be called before using the DRI2 extension.
func Init(c *xgb.Conn) error {
	reply, err := xproto.QueryExtension(c, 4, "DRI2").Reply()
	switch {
	case err != nil:
		return err
	case !reply.Present:
		return xgb.Errorf("No extension named DRI2 could be found on on the server.")
	}

	xgb.ExtLock.Lock()
	c.Extensions["DRI2"] = reply.MajorOpcode
	for evNum, fun := range xgb.NewExtEventFuncs["DRI2"] {
		xgb.NewEventFuncs[int(reply.FirstEvent)+evNum] = fun
	}
	for errNum, fun := range xgb.NewExtErrorFuncs["DRI2"] {
		xgb.NewErrorFuncs[int(reply.FirstError)+errNum] = fun
	}
	xgb.ExtLock.Unlock()

	return nil
}

func init() {
	xgb.NewExtEventFuncs["DRI2"] = make(map[int]xgb.NewEventFun)
	xgb.NewExtErrorFuncs["DRI2"] = make(map[int]xgb.NewErrorFun)
}

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Card32'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

const (
	AttachmentBufferFrontLeft      = 0
	AttachmentBufferBackLeft       = 1
	AttachmentBufferFrontRight     = 2
	AttachmentBufferBackRight      = 3
	AttachmentBufferDepth          = 4
	AttachmentBufferStencil        = 5
	AttachmentBufferAccum          = 6
	AttachmentBufferFakeFrontLeft  = 7
	AttachmentBufferFakeFrontRight = 8
	AttachmentBufferDepthStencil   = 9
	AttachmentBufferHiz            = 10
)

const (
	DriverTypeDri   = 0
	DriverTypeVdpau = 1
)

const (
	EventTypeExchangeComplete = 1
	EventTypeBlitComplete     = 2
	EventTypeFlipComplete     = 3
)

// 'DRI2Buffer' struct definition
// Size: 20
type DRI2Buffer struct {
	Attachment uint32
	Name       uint32
	Pitch      uint32
	Cpp        uint32
	Flags      uint32
}

// Struct read DRI2Buffer
func DRI2BufferRead(buf []byte, v *DRI2Buffer) int {
	b := 0

	v.Attachment = xgb.Get32(buf[b:])
	b += 4

	v.Name = xgb.Get32(buf[b:])
	b += 4

	v.Pitch = xgb.Get32(buf[b:])
	b += 4

	v.Cpp = xgb.Get32(buf[b:])
	b += 4

	v.Flags = xgb.Get32(buf[b:])
	b += 4

	return b
}

// Struct list read DRI2Buffer
func DRI2BufferReadList(buf []byte, dest []DRI2Buffer) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = DRI2Buffer{}
		b += DRI2BufferRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Struct write DRI2Buffer
func (v DRI2Buffer) Bytes() []byte {
	buf := make([]byte, 20)
	b := 0

	xgb.Put32(buf[b:], v.Attachment)
	b += 4

	xgb.Put32(buf[b:], v.Name)
	b += 4

	xgb.Put32(buf[b:], v.Pitch)
	b += 4

	xgb.Put32(buf[b:], v.Cpp)
	b += 4

	xgb.Put32(buf[b:], v.Flags)
	b += 4

	return buf
}

// Write struct list DRI2Buffer
func DRI2BufferListBytes(buf []byte, list []DRI2Buffer) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

// 'AttachFormat' struct definition
// Size: 8
type AttachFormat struct {
	Attachment uint32
	Format     uint32
}

// Struct read AttachFormat
func AttachFormatRead(buf []byte, v *AttachFormat) int {
	b := 0

	v.Attachment = xgb.Get32(buf[b:])
	b += 4

	v.Format = xgb.Get32(buf[b:])
	b += 4

	return b
}

// Struct list read AttachFormat
func AttachFormatReadList(buf []byte, dest []AttachFormat) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = AttachFormat{}
		b += AttachFormatRead(buf[b:], &dest[i])
	}
	return xgb.Pad(b)
}

// Struct write AttachFormat
func (v AttachFormat) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	xgb.Put32(buf[b:], v.Attachment)
	b += 4

	xgb.Put32(buf[b:], v.Format)
	b += 4

	return buf
}

// Write struct list AttachFormat
func AttachFormatListBytes(buf []byte, list []AttachFormat) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += xgb.Pad(len(structBytes))
	}
	return b
}

// Event definition BufferSwapComplete (0)
// Size: 32

const BufferSwapComplete = 0

type BufferSwapCompleteEvent struct {
	Sequence uint16
	// padding: 1 bytes
	EventType uint16
	// padding: 2 bytes
	Drawable xproto.Drawable
	UstHi    uint32
	UstLo    uint32
	MscHi    uint32
	MscLo    uint32
	Sbc      uint32
}

// Event read BufferSwapComplete
func BufferSwapCompleteEventNew(buf []byte) xgb.Event {
	v := BufferSwapCompleteEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.EventType = xgb.Get16(buf[b:])
	b += 2

	b += 2 // padding

	v.Drawable = xproto.Drawable(xgb.Get32(buf[b:]))
	b += 4

	v.UstHi = xgb.Get32(buf[b:])
	b += 4

	v.UstLo = xgb.Get32(buf[b:])
	b += 4

	v.MscHi = xgb.Get32(buf[b:])
	b += 4

	v.MscLo = xgb.Get32(buf[b:])
	b += 4

	v.Sbc = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Event write BufferSwapComplete
func (v BufferSwapCompleteEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	xgb.Put16(buf[b:], v.EventType)
	b += 2

	b += 2 // padding

	xgb.Put32(buf[b:], uint32(v.Drawable))
	b += 4

	xgb.Put32(buf[b:], v.UstHi)
	b += 4

	xgb.Put32(buf[b:], v.UstLo)
	b += 4

	xgb.Put32(buf[b:], v.MscHi)
	b += 4

	xgb.Put32(buf[b:], v.MscLo)
	b += 4

	xgb.Put32(buf[b:], v.Sbc)
	b += 4

	return buf
}

func (v BufferSwapCompleteEvent) ImplementsEvent() {}

func (v BufferSwapCompleteEvent) SequenceId() uint16 {
	return v.Sequence
}

func (v BufferSwapCompleteEvent) String() string {
	fieldVals := make([]string, 0, 9)
	fieldVals = append(fieldVals, xgb.Sprintf("Sequence: %d", v.Sequence))
	fieldVals = append(fieldVals, xgb.Sprintf("EventType: %d", v.EventType))
	fieldVals = append(fieldVals, xgb.Sprintf("Drawable: %d", v.Drawable))
	fieldVals = append(fieldVals, xgb.Sprintf("UstHi: %d", v.UstHi))
	fieldVals = append(fieldVals, xgb.Sprintf("UstLo: %d", v.UstLo))
	fieldVals = append(fieldVals, xgb.Sprintf("MscHi: %d", v.MscHi))
	fieldVals = append(fieldVals, xgb.Sprintf("MscLo: %d", v.MscLo))
	fieldVals = append(fieldVals, xgb.Sprintf("Sbc: %d", v.Sbc))
	return "BufferSwapComplete {" + xgb.StringsJoin(fieldVals, ", ") + "}"
}

func init() {
	xgb.NewExtEventFuncs["DRI2"][0] = BufferSwapCompleteEventNew
}

// Event definition InvalidateBuffers (1)
// Size: 32

const InvalidateBuffers = 1

type InvalidateBuffersEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Drawable xproto.Drawable
}

// Event read InvalidateBuffers
func InvalidateBuffersEventNew(buf []byte) xgb.Event {
	v := InvalidateBuffersEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Drawable = xproto.Drawable(xgb.Get32(buf[b:]))
	b += 4

	return v
}

// Event write InvalidateBuffers
func (v InvalidateBuffersEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 1
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	xgb.Put32(buf[b:], uint32(v.Drawable))
	b += 4

	return buf
}

func (v InvalidateBuffersEvent) ImplementsEvent() {}

func (v InvalidateBuffersEvent) SequenceId() uint16 {
	return v.Sequence
}

func (v InvalidateBuffersEvent) String() string {
	fieldVals := make([]string, 0, 2)
	fieldVals = append(fieldVals, xgb.Sprintf("Sequence: %d", v.Sequence))
	fieldVals = append(fieldVals, xgb.Sprintf("Drawable: %d", v.Drawable))
	return "InvalidateBuffers {" + xgb.StringsJoin(fieldVals, ", ") + "}"
}

func init() {
	xgb.NewExtEventFuncs["DRI2"][1] = InvalidateBuffersEventNew
}

// Request QueryVersion
// size: 12
type QueryVersionCookie struct {
	*xgb.Cookie
}

func QueryVersion(c *xgb.Conn, MajorVersion uint32, MinorVersion uint32) QueryVersionCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(queryVersionRequest(c, MajorVersion, MinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

func QueryVersionUnchecked(c *xgb.Conn, MajorVersion uint32, MinorVersion uint32) QueryVersionCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(queryVersionRequest(c, MajorVersion, MinorVersion), cookie)
	return QueryVersionCookie{cookie}
}

// Request reply for QueryVersion
// size: 16
type QueryVersionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	MajorVersion uint32
	MinorVersion uint32
}

// Waits and reads reply data from request QueryVersion
func (cook QueryVersionCookie) Reply() (*QueryVersionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return queryVersionReply(buf), nil
}

// Read reply into structure from buffer for QueryVersion
func queryVersionReply(buf []byte) *QueryVersionReply {
	v := new(QueryVersionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.MajorVersion = xgb.Get32(buf[b:])
	b += 4

	v.MinorVersion = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for QueryVersion
func queryVersionRequest(c *xgb.Conn, MajorVersion uint32, MinorVersion uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], MajorVersion)
	b += 4

	xgb.Put32(buf[b:], MinorVersion)
	b += 4

	return buf
}

// Request Connect
// size: 12
type ConnectCookie struct {
	*xgb.Cookie
}

func Connect(c *xgb.Conn, Window xproto.Window, DriverType uint32) ConnectCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(connectRequest(c, Window, DriverType), cookie)
	return ConnectCookie{cookie}
}

func ConnectUnchecked(c *xgb.Conn, Window xproto.Window, DriverType uint32) ConnectCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(connectRequest(c, Window, DriverType), cookie)
	return ConnectCookie{cookie}
}

// Request reply for Connect
// size: (((32 + xgb.Pad((int(DriverNameLength) * 1))) + xgb.Pad(((((int(DriverNameLength) + 3) & -4) - int(DriverNameLength)) * 1))) + xgb.Pad((int(DeviceNameLength) * 1)))
type ConnectReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	DriverNameLength uint32
	DeviceNameLength uint32
	// padding: 16 bytes
	DriverName   string // size: xgb.Pad((int(DriverNameLength) * 1))
	AlignmentPad []byte // size: xgb.Pad(((((int(DriverNameLength) + 3) & -4) - int(DriverNameLength)) * 1))
	DeviceName   string // size: xgb.Pad((int(DeviceNameLength) * 1))
}

// Waits and reads reply data from request Connect
func (cook ConnectCookie) Reply() (*ConnectReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return connectReply(buf), nil
}

// Read reply into structure from buffer for Connect
func connectReply(buf []byte) *ConnectReply {
	v := new(ConnectReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.DriverNameLength = xgb.Get32(buf[b:])
	b += 4

	v.DeviceNameLength = xgb.Get32(buf[b:])
	b += 4

	b += 16 // padding

	{
		byteString := make([]byte, v.DriverNameLength)
		copy(byteString[:v.DriverNameLength], buf[b:])
		v.DriverName = string(byteString)
		b += xgb.Pad(int(v.DriverNameLength))
	}

	v.AlignmentPad = make([]byte, (((int(v.DriverNameLength) + 3) & -4) - int(v.DriverNameLength)))
	copy(v.AlignmentPad[:(((int(v.DriverNameLength)+3)&-4)-int(v.DriverNameLength))], buf[b:])
	b += xgb.Pad(int((((int(v.DriverNameLength) + 3) & -4) - int(v.DriverNameLength))))

	{
		byteString := make([]byte, v.DeviceNameLength)
		copy(byteString[:v.DeviceNameLength], buf[b:])
		v.DeviceName = string(byteString)
		b += xgb.Pad(int(v.DeviceNameLength))
	}

	return v
}

// Write request to wire for Connect
func connectRequest(c *xgb.Conn, Window xproto.Window, DriverType uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	xgb.Put32(buf[b:], DriverType)
	b += 4

	return buf
}

// Request Authenticate
// size: 12
type AuthenticateCookie struct {
	*xgb.Cookie
}

func Authenticate(c *xgb.Conn, Window xproto.Window, Magic uint32) AuthenticateCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(authenticateRequest(c, Window, Magic), cookie)
	return AuthenticateCookie{cookie}
}

func AuthenticateUnchecked(c *xgb.Conn, Window xproto.Window, Magic uint32) AuthenticateCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(authenticateRequest(c, Window, Magic), cookie)
	return AuthenticateCookie{cookie}
}

// Request reply for Authenticate
// size: 12
type AuthenticateReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Authenticated uint32
}

// Waits and reads reply data from request Authenticate
func (cook AuthenticateCookie) Reply() (*AuthenticateReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return authenticateReply(buf), nil
}

// Read reply into structure from buffer for Authenticate
func authenticateReply(buf []byte) *AuthenticateReply {
	v := new(AuthenticateReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Authenticated = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for Authenticate
func authenticateRequest(c *xgb.Conn, Window xproto.Window, Magic uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Window))
	b += 4

	xgb.Put32(buf[b:], Magic)
	b += 4

	return buf
}

// Request CreateDrawable
// size: 8
type CreateDrawableCookie struct {
	*xgb.Cookie
}

// Write request to wire for CreateDrawable
func CreateDrawable(c *xgb.Conn, Drawable xproto.Drawable) CreateDrawableCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(createDrawableRequest(c, Drawable), cookie)
	return CreateDrawableCookie{cookie}
}

func CreateDrawableChecked(c *xgb.Conn, Drawable xproto.Drawable) CreateDrawableCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(createDrawableRequest(c, Drawable), cookie)
	return CreateDrawableCookie{cookie}
}

func (cook CreateDrawableCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for CreateDrawable
func createDrawableRequest(c *xgb.Conn, Drawable xproto.Drawable) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// Request DestroyDrawable
// size: 8
type DestroyDrawableCookie struct {
	*xgb.Cookie
}

// Write request to wire for DestroyDrawable
func DestroyDrawable(c *xgb.Conn, Drawable xproto.Drawable) DestroyDrawableCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(destroyDrawableRequest(c, Drawable), cookie)
	return DestroyDrawableCookie{cookie}
}

func DestroyDrawableChecked(c *xgb.Conn, Drawable xproto.Drawable) DestroyDrawableCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(destroyDrawableRequest(c, Drawable), cookie)
	return DestroyDrawableCookie{cookie}
}

func (cook DestroyDrawableCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for DestroyDrawable
func destroyDrawableRequest(c *xgb.Conn, Drawable xproto.Drawable) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// Request GetBuffers
// size: xgb.Pad((12 + xgb.Pad((len(Attachments) * 4))))
type GetBuffersCookie struct {
	*xgb.Cookie
}

func GetBuffers(c *xgb.Conn, Drawable xproto.Drawable, Count uint32, Attachments []uint32) GetBuffersCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(getBuffersRequest(c, Drawable, Count, Attachments), cookie)
	return GetBuffersCookie{cookie}
}

func GetBuffersUnchecked(c *xgb.Conn, Drawable xproto.Drawable, Count uint32, Attachments []uint32) GetBuffersCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(getBuffersRequest(c, Drawable, Count, Attachments), cookie)
	return GetBuffersCookie{cookie}
}

// Request reply for GetBuffers
// size: (32 + xgb.Pad((int(Count) * 20)))
type GetBuffersReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Width  uint32
	Height uint32
	Count  uint32
	// padding: 12 bytes
	Buffers []DRI2Buffer // size: xgb.Pad((int(Count) * 20))
}

// Waits and reads reply data from request GetBuffers
func (cook GetBuffersCookie) Reply() (*GetBuffersReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getBuffersReply(buf), nil
}

// Read reply into structure from buffer for GetBuffers
func getBuffersReply(buf []byte) *GetBuffersReply {
	v := new(GetBuffersReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Width = xgb.Get32(buf[b:])
	b += 4

	v.Height = xgb.Get32(buf[b:])
	b += 4

	v.Count = xgb.Get32(buf[b:])
	b += 4

	b += 12 // padding

	v.Buffers = make([]DRI2Buffer, v.Count)
	b += DRI2BufferReadList(buf[b:], v.Buffers)

	return v
}

// Write request to wire for GetBuffers
func getBuffersRequest(c *xgb.Conn, Drawable xproto.Drawable, Count uint32, Attachments []uint32) []byte {
	size := xgb.Pad((12 + xgb.Pad((len(Attachments) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put32(buf[b:], Count)
	b += 4

	for i := 0; i < int(len(Attachments)); i++ {
		xgb.Put32(buf[b:], Attachments[i])
		b += 4
	}
	b = xgb.Pad(b)

	return buf
}

// Request CopyRegion
// size: 20
type CopyRegionCookie struct {
	*xgb.Cookie
}

func CopyRegion(c *xgb.Conn, Drawable xproto.Drawable, Region uint32, Dest uint32, Src uint32) CopyRegionCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(copyRegionRequest(c, Drawable, Region, Dest, Src), cookie)
	return CopyRegionCookie{cookie}
}

func CopyRegionUnchecked(c *xgb.Conn, Drawable xproto.Drawable, Region uint32, Dest uint32, Src uint32) CopyRegionCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(copyRegionRequest(c, Drawable, Region, Dest, Src), cookie)
	return CopyRegionCookie{cookie}
}

// Request reply for CopyRegion
// size: 8
type CopyRegionReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
}

// Waits and reads reply data from request CopyRegion
func (cook CopyRegionCookie) Reply() (*CopyRegionReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return copyRegionReply(buf), nil
}

// Read reply into structure from buffer for CopyRegion
func copyRegionReply(buf []byte) *CopyRegionReply {
	v := new(CopyRegionReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	return v
}

// Write request to wire for CopyRegion
func copyRegionRequest(c *xgb.Conn, Drawable xproto.Drawable, Region uint32, Dest uint32, Src uint32) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put32(buf[b:], Region)
	b += 4

	xgb.Put32(buf[b:], Dest)
	b += 4

	xgb.Put32(buf[b:], Src)
	b += 4

	return buf
}

// Request GetBuffersWithFormat
// size: xgb.Pad((12 + xgb.Pad((len(Attachments) * 8))))
type GetBuffersWithFormatCookie struct {
	*xgb.Cookie
}

func GetBuffersWithFormat(c *xgb.Conn, Drawable xproto.Drawable, Count uint32, Attachments []AttachFormat) GetBuffersWithFormatCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(getBuffersWithFormatRequest(c, Drawable, Count, Attachments), cookie)
	return GetBuffersWithFormatCookie{cookie}
}

func GetBuffersWithFormatUnchecked(c *xgb.Conn, Drawable xproto.Drawable, Count uint32, Attachments []AttachFormat) GetBuffersWithFormatCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(getBuffersWithFormatRequest(c, Drawable, Count, Attachments), cookie)
	return GetBuffersWithFormatCookie{cookie}
}

// Request reply for GetBuffersWithFormat
// size: (32 + xgb.Pad((int(Count) * 20)))
type GetBuffersWithFormatReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	Width  uint32
	Height uint32
	Count  uint32
	// padding: 12 bytes
	Buffers []DRI2Buffer // size: xgb.Pad((int(Count) * 20))
}

// Waits and reads reply data from request GetBuffersWithFormat
func (cook GetBuffersWithFormatCookie) Reply() (*GetBuffersWithFormatReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getBuffersWithFormatReply(buf), nil
}

// Read reply into structure from buffer for GetBuffersWithFormat
func getBuffersWithFormatReply(buf []byte) *GetBuffersWithFormatReply {
	v := new(GetBuffersWithFormatReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.Width = xgb.Get32(buf[b:])
	b += 4

	v.Height = xgb.Get32(buf[b:])
	b += 4

	v.Count = xgb.Get32(buf[b:])
	b += 4

	b += 12 // padding

	v.Buffers = make([]DRI2Buffer, v.Count)
	b += DRI2BufferReadList(buf[b:], v.Buffers)

	return v
}

// Write request to wire for GetBuffersWithFormat
func getBuffersWithFormatRequest(c *xgb.Conn, Drawable xproto.Drawable, Count uint32, Attachments []AttachFormat) []byte {
	size := xgb.Pad((12 + xgb.Pad((len(Attachments) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put32(buf[b:], Count)
	b += 4

	b += AttachFormatListBytes(buf[b:], Attachments)

	return buf
}

// Request SwapBuffers
// size: 32
type SwapBuffersCookie struct {
	*xgb.Cookie
}

func SwapBuffers(c *xgb.Conn, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) SwapBuffersCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(swapBuffersRequest(c, Drawable, TargetMscHi, TargetMscLo, DivisorHi, DivisorLo, RemainderHi, RemainderLo), cookie)
	return SwapBuffersCookie{cookie}
}

func SwapBuffersUnchecked(c *xgb.Conn, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) SwapBuffersCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(swapBuffersRequest(c, Drawable, TargetMscHi, TargetMscLo, DivisorHi, DivisorLo, RemainderHi, RemainderLo), cookie)
	return SwapBuffersCookie{cookie}
}

// Request reply for SwapBuffers
// size: 16
type SwapBuffersReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	SwapHi uint32
	SwapLo uint32
}

// Waits and reads reply data from request SwapBuffers
func (cook SwapBuffersCookie) Reply() (*SwapBuffersReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return swapBuffersReply(buf), nil
}

// Read reply into structure from buffer for SwapBuffers
func swapBuffersReply(buf []byte) *SwapBuffersReply {
	v := new(SwapBuffersReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.SwapHi = xgb.Get32(buf[b:])
	b += 4

	v.SwapLo = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for SwapBuffers
func swapBuffersRequest(c *xgb.Conn, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put32(buf[b:], TargetMscHi)
	b += 4

	xgb.Put32(buf[b:], TargetMscLo)
	b += 4

	xgb.Put32(buf[b:], DivisorHi)
	b += 4

	xgb.Put32(buf[b:], DivisorLo)
	b += 4

	xgb.Put32(buf[b:], RemainderHi)
	b += 4

	xgb.Put32(buf[b:], RemainderLo)
	b += 4

	return buf
}

// Request GetMSC
// size: 8
type GetMSCCookie struct {
	*xgb.Cookie
}

func GetMSC(c *xgb.Conn, Drawable xproto.Drawable) GetMSCCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(getMSCRequest(c, Drawable), cookie)
	return GetMSCCookie{cookie}
}

func GetMSCUnchecked(c *xgb.Conn, Drawable xproto.Drawable) GetMSCCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(getMSCRequest(c, Drawable), cookie)
	return GetMSCCookie{cookie}
}

// Request reply for GetMSC
// size: 32
type GetMSCReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	UstHi uint32
	UstLo uint32
	MscHi uint32
	MscLo uint32
	SbcHi uint32
	SbcLo uint32
}

// Waits and reads reply data from request GetMSC
func (cook GetMSCCookie) Reply() (*GetMSCReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return getMSCReply(buf), nil
}

// Read reply into structure from buffer for GetMSC
func getMSCReply(buf []byte) *GetMSCReply {
	v := new(GetMSCReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.UstHi = xgb.Get32(buf[b:])
	b += 4

	v.UstLo = xgb.Get32(buf[b:])
	b += 4

	v.MscHi = xgb.Get32(buf[b:])
	b += 4

	v.MscLo = xgb.Get32(buf[b:])
	b += 4

	v.SbcHi = xgb.Get32(buf[b:])
	b += 4

	v.SbcLo = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for GetMSC
func getMSCRequest(c *xgb.Conn, Drawable xproto.Drawable) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 9 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// Request WaitMSC
// size: 32
type WaitMSCCookie struct {
	*xgb.Cookie
}

func WaitMSC(c *xgb.Conn, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) WaitMSCCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(waitMSCRequest(c, Drawable, TargetMscHi, TargetMscLo, DivisorHi, DivisorLo, RemainderHi, RemainderLo), cookie)
	return WaitMSCCookie{cookie}
}

func WaitMSCUnchecked(c *xgb.Conn, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) WaitMSCCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(waitMSCRequest(c, Drawable, TargetMscHi, TargetMscLo, DivisorHi, DivisorLo, RemainderHi, RemainderLo), cookie)
	return WaitMSCCookie{cookie}
}

// Request reply for WaitMSC
// size: 32
type WaitMSCReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	UstHi uint32
	UstLo uint32
	MscHi uint32
	MscLo uint32
	SbcHi uint32
	SbcLo uint32
}

// Waits and reads reply data from request WaitMSC
func (cook WaitMSCCookie) Reply() (*WaitMSCReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return waitMSCReply(buf), nil
}

// Read reply into structure from buffer for WaitMSC
func waitMSCReply(buf []byte) *WaitMSCReply {
	v := new(WaitMSCReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.UstHi = xgb.Get32(buf[b:])
	b += 4

	v.UstLo = xgb.Get32(buf[b:])
	b += 4

	v.MscHi = xgb.Get32(buf[b:])
	b += 4

	v.MscLo = xgb.Get32(buf[b:])
	b += 4

	v.SbcHi = xgb.Get32(buf[b:])
	b += 4

	v.SbcLo = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for WaitMSC
func waitMSCRequest(c *xgb.Conn, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 10 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put32(buf[b:], TargetMscHi)
	b += 4

	xgb.Put32(buf[b:], TargetMscLo)
	b += 4

	xgb.Put32(buf[b:], DivisorHi)
	b += 4

	xgb.Put32(buf[b:], DivisorLo)
	b += 4

	xgb.Put32(buf[b:], RemainderHi)
	b += 4

	xgb.Put32(buf[b:], RemainderLo)
	b += 4

	return buf
}

// Request WaitSBC
// size: 16
type WaitSBCCookie struct {
	*xgb.Cookie
}

func WaitSBC(c *xgb.Conn, Drawable xproto.Drawable, TargetSbcHi uint32, TargetSbcLo uint32) WaitSBCCookie {
	cookie := c.NewCookie(true, true)
	c.NewRequest(waitSBCRequest(c, Drawable, TargetSbcHi, TargetSbcLo), cookie)
	return WaitSBCCookie{cookie}
}

func WaitSBCUnchecked(c *xgb.Conn, Drawable xproto.Drawable, TargetSbcHi uint32, TargetSbcLo uint32) WaitSBCCookie {
	cookie := c.NewCookie(false, true)
	c.NewRequest(waitSBCRequest(c, Drawable, TargetSbcHi, TargetSbcLo), cookie)
	return WaitSBCCookie{cookie}
}

// Request reply for WaitSBC
// size: 32
type WaitSBCReply struct {
	Sequence uint16
	Length   uint32
	// padding: 1 bytes
	UstHi uint32
	UstLo uint32
	MscHi uint32
	MscLo uint32
	SbcHi uint32
	SbcLo uint32
}

// Waits and reads reply data from request WaitSBC
func (cook WaitSBCCookie) Reply() (*WaitSBCReply, error) {
	buf, err := cook.Cookie.Reply()
	if err != nil {
		return nil, err
	}
	if buf == nil {
		return nil, nil
	}
	return waitSBCReply(buf), nil
}

// Read reply into structure from buffer for WaitSBC
func waitSBCReply(buf []byte) *WaitSBCReply {
	v := new(WaitSBCReply)
	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = xgb.Get16(buf[b:])
	b += 2

	v.Length = xgb.Get32(buf[b:]) // 4-byte units
	b += 4

	v.UstHi = xgb.Get32(buf[b:])
	b += 4

	v.UstLo = xgb.Get32(buf[b:])
	b += 4

	v.MscHi = xgb.Get32(buf[b:])
	b += 4

	v.MscLo = xgb.Get32(buf[b:])
	b += 4

	v.SbcHi = xgb.Get32(buf[b:])
	b += 4

	v.SbcLo = xgb.Get32(buf[b:])
	b += 4

	return v
}

// Write request to wire for WaitSBC
func waitSBCRequest(c *xgb.Conn, Drawable xproto.Drawable, TargetSbcHi uint32, TargetSbcLo uint32) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 11 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put32(buf[b:], TargetSbcHi)
	b += 4

	xgb.Put32(buf[b:], TargetSbcLo)
	b += 4

	return buf
}

// Request SwapInterval
// size: 12
type SwapIntervalCookie struct {
	*xgb.Cookie
}

// Write request to wire for SwapInterval
func SwapInterval(c *xgb.Conn, Drawable xproto.Drawable, Interval uint32) SwapIntervalCookie {
	cookie := c.NewCookie(false, false)
	c.NewRequest(swapIntervalRequest(c, Drawable, Interval), cookie)
	return SwapIntervalCookie{cookie}
}

func SwapIntervalChecked(c *xgb.Conn, Drawable xproto.Drawable, Interval uint32) SwapIntervalCookie {
	cookie := c.NewCookie(true, false)
	c.NewRequest(swapIntervalRequest(c, Drawable, Interval), cookie)
	return SwapIntervalCookie{cookie}
}

func (cook SwapIntervalCookie) Check() error {
	return cook.Cookie.Check()
}

// Write request to wire for SwapInterval
func swapIntervalRequest(c *xgb.Conn, Drawable xproto.Drawable, Interval uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = c.Extensions["DRI2"]
	b += 1

	buf[b] = 12 // request opcode
	b += 1

	xgb.Put16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	xgb.Put32(buf[b:], uint32(Drawable))
	b += 4

	xgb.Put32(buf[b:], Interval)
	b += 4

	return buf
}