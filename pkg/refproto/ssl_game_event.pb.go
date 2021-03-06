// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ssl_game_event.proto

package refproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Game_Event_GameEventType int32

const (
	// not set
	Game_Event_UNKNOWN Game_Event_GameEventType = 0
	// an event that is not listed in this enum yet.
	// Give further details in the message below
	Game_Event_CUSTOM Game_Event_GameEventType = 1
	// Law 3: Number of players
	Game_Event_NUMBER_OF_PLAYERS Game_Event_GameEventType = 2
	// Law 9: Ball out of play
	Game_Event_BALL_LEFT_FIELD Game_Event_GameEventType = 3
	// Law 10: Team scored a goal
	Game_Event_GOAL Game_Event_GameEventType = 4
	// Law 9.3: lack of progress while bringing the ball into play
	Game_Event_KICK_TIMEOUT Game_Event_GameEventType = 5
	// Law ?: There is no progress in game for (10|15)? seconds
	Game_Event_NO_PROGRESS_IN_GAME Game_Event_GameEventType = 6
	// Law 12: Pushing / Substantial Contact
	Game_Event_BOT_COLLISION Game_Event_GameEventType = 7
	// Law 12.2: Defender is completely inside penalty area
	Game_Event_MULTIPLE_DEFENDER Game_Event_GameEventType = 8
	// Law 12: Defender is partially inside penalty area
	Game_Event_MULTIPLE_DEFENDER_PARTIALLY Game_Event_GameEventType = 9
	// Law 12.3: Attacker in defense area
	Game_Event_ATTACKER_IN_DEFENSE_AREA Game_Event_GameEventType = 10
	// Law 12: Icing (kicking over midline and opponent goal line)
	Game_Event_ICING Game_Event_GameEventType = 11
	// Law 12: Ball speed
	Game_Event_BALL_SPEED Game_Event_GameEventType = 12
	// Law 12: Robot speed during STOP
	Game_Event_ROBOT_STOP_SPEED Game_Event_GameEventType = 13
	// Law 12: Maximum dribbling distance
	Game_Event_BALL_DRIBBLING Game_Event_GameEventType = 14
	// Law 12: Touching the opponent goalkeeper
	Game_Event_ATTACKER_TOUCH_KEEPER Game_Event_GameEventType = 15
	// Law 12: Double touch
	Game_Event_DOUBLE_TOUCH Game_Event_GameEventType = 16
	// Law 13-17: Attacker not too close to the opponent's penalty area when ball enters play
	Game_Event_ATTACKER_TO_DEFENCE_AREA Game_Event_GameEventType = 17
	// Law 13-17: Keeping the correct distance to the ball during opponents freekicks
	Game_Event_DEFENDER_TO_KICK_POINT_DISTANCE Game_Event_GameEventType = 18
	// Law 12: A robot holds the ball deliberately
	Game_Event_BALL_HOLDING Game_Event_GameEventType = 19
	// Law 12: The ball entered the goal directly after an indirect kick was performed
	Game_Event_INDIRECT_GOAL Game_Event_GameEventType = 20
	// Law 9.2: Ball placement
	Game_Event_BALL_PLACEMENT_FAILED Game_Event_GameEventType = 21
	// Law 10: A goal is only scored if the ball has not exceeded a robot height (150mm) between the last
	// kick of an attacker and the time the ball crossed the goal line.
	Game_Event_CHIP_ON_GOAL Game_Event_GameEventType = 22
)

var Game_Event_GameEventType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CUSTOM",
	2:  "NUMBER_OF_PLAYERS",
	3:  "BALL_LEFT_FIELD",
	4:  "GOAL",
	5:  "KICK_TIMEOUT",
	6:  "NO_PROGRESS_IN_GAME",
	7:  "BOT_COLLISION",
	8:  "MULTIPLE_DEFENDER",
	9:  "MULTIPLE_DEFENDER_PARTIALLY",
	10: "ATTACKER_IN_DEFENSE_AREA",
	11: "ICING",
	12: "BALL_SPEED",
	13: "ROBOT_STOP_SPEED",
	14: "BALL_DRIBBLING",
	15: "ATTACKER_TOUCH_KEEPER",
	16: "DOUBLE_TOUCH",
	17: "ATTACKER_TO_DEFENCE_AREA",
	18: "DEFENDER_TO_KICK_POINT_DISTANCE",
	19: "BALL_HOLDING",
	20: "INDIRECT_GOAL",
	21: "BALL_PLACEMENT_FAILED",
	22: "CHIP_ON_GOAL",
}

var Game_Event_GameEventType_value = map[string]int32{
	"UNKNOWN":                         0,
	"CUSTOM":                          1,
	"NUMBER_OF_PLAYERS":               2,
	"BALL_LEFT_FIELD":                 3,
	"GOAL":                            4,
	"KICK_TIMEOUT":                    5,
	"NO_PROGRESS_IN_GAME":             6,
	"BOT_COLLISION":                   7,
	"MULTIPLE_DEFENDER":               8,
	"MULTIPLE_DEFENDER_PARTIALLY":     9,
	"ATTACKER_IN_DEFENSE_AREA":        10,
	"ICING":                           11,
	"BALL_SPEED":                      12,
	"ROBOT_STOP_SPEED":                13,
	"BALL_DRIBBLING":                  14,
	"ATTACKER_TOUCH_KEEPER":           15,
	"DOUBLE_TOUCH":                    16,
	"ATTACKER_TO_DEFENCE_AREA":        17,
	"DEFENDER_TO_KICK_POINT_DISTANCE": 18,
	"BALL_HOLDING":                    19,
	"INDIRECT_GOAL":                   20,
	"BALL_PLACEMENT_FAILED":           21,
	"CHIP_ON_GOAL":                    22,
}

func (x Game_Event_GameEventType) Enum() *Game_Event_GameEventType {
	p := new(Game_Event_GameEventType)
	*p = x
	return p
}

func (x Game_Event_GameEventType) String() string {
	return proto.EnumName(Game_Event_GameEventType_name, int32(x))
}

func (x *Game_Event_GameEventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Game_Event_GameEventType_value, data, "Game_Event_GameEventType")
	if err != nil {
		return err
	}
	*x = Game_Event_GameEventType(value)
	return nil
}

func (Game_Event_GameEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d7737834ddab500, []int{0, 0}
}

// a team
type Game_Event_Team int32

const (
	Game_Event_TEAM_UNKNOWN Game_Event_Team = 0
	Game_Event_TEAM_YELLOW  Game_Event_Team = 1
	Game_Event_TEAM_BLUE    Game_Event_Team = 2
)

var Game_Event_Team_name = map[int32]string{
	0: "TEAM_UNKNOWN",
	1: "TEAM_YELLOW",
	2: "TEAM_BLUE",
}

var Game_Event_Team_value = map[string]int32{
	"TEAM_UNKNOWN": 0,
	"TEAM_YELLOW":  1,
	"TEAM_BLUE":    2,
}

func (x Game_Event_Team) Enum() *Game_Event_Team {
	p := new(Game_Event_Team)
	*p = x
	return p
}

func (x Game_Event_Team) String() string {
	return proto.EnumName(Game_Event_Team_name, int32(x))
}

func (x *Game_Event_Team) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Game_Event_Team_value, data, "Game_Event_Team")
	if err != nil {
		return err
	}
	*x = Game_Event_Team(value)
	return nil
}

func (Game_Event_Team) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d7737834ddab500, []int{0, 1}
}

// a game event that caused a referee command
type Game_Event struct {
	// the game event type that happened
	GameEventType *Game_Event_GameEventType `protobuf:"varint,1,req,name=gameEventType,enum=Game_Event_GameEventType" json:"gameEventType,omitempty"`
	// the team and optionally a designated robot that is the originator of the game event
	Originator *Game_Event_Originator `protobuf:"bytes,2,opt,name=originator" json:"originator,omitempty"`
	// a message describing further details of this game event
	Message              *string  `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Game_Event) Reset()         { *m = Game_Event{} }
func (m *Game_Event) String() string { return proto.CompactTextString(m) }
func (*Game_Event) ProtoMessage()    {}
func (*Game_Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d7737834ddab500, []int{0}
}

func (m *Game_Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Game_Event.Unmarshal(m, b)
}
func (m *Game_Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Game_Event.Marshal(b, m, deterministic)
}
func (m *Game_Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game_Event.Merge(m, src)
}
func (m *Game_Event) XXX_Size() int {
	return xxx_messageInfo_Game_Event.Size(m)
}
func (m *Game_Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Game_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Game_Event proto.InternalMessageInfo

func (m *Game_Event) GetGameEventType() Game_Event_GameEventType {
	if m != nil && m.GameEventType != nil {
		return *m.GameEventType
	}
	return Game_Event_UNKNOWN
}

func (m *Game_Event) GetOriginator() *Game_Event_Originator {
	if m != nil {
		return m.Originator
	}
	return nil
}

func (m *Game_Event) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

// information about an originator
type Game_Event_Originator struct {
	Team                 *Game_Event_Team `protobuf:"varint,1,req,name=team,enum=Game_Event_Team" json:"team,omitempty"`
	BotId                *uint32          `protobuf:"varint,2,opt,name=botId" json:"botId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Game_Event_Originator) Reset()         { *m = Game_Event_Originator{} }
func (m *Game_Event_Originator) String() string { return proto.CompactTextString(m) }
func (*Game_Event_Originator) ProtoMessage()    {}
func (*Game_Event_Originator) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d7737834ddab500, []int{0, 0}
}

func (m *Game_Event_Originator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Game_Event_Originator.Unmarshal(m, b)
}
func (m *Game_Event_Originator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Game_Event_Originator.Marshal(b, m, deterministic)
}
func (m *Game_Event_Originator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game_Event_Originator.Merge(m, src)
}
func (m *Game_Event_Originator) XXX_Size() int {
	return xxx_messageInfo_Game_Event_Originator.Size(m)
}
func (m *Game_Event_Originator) XXX_DiscardUnknown() {
	xxx_messageInfo_Game_Event_Originator.DiscardUnknown(m)
}

var xxx_messageInfo_Game_Event_Originator proto.InternalMessageInfo

func (m *Game_Event_Originator) GetTeam() Game_Event_Team {
	if m != nil && m.Team != nil {
		return *m.Team
	}
	return Game_Event_TEAM_UNKNOWN
}

func (m *Game_Event_Originator) GetBotId() uint32 {
	if m != nil && m.BotId != nil {
		return *m.BotId
	}
	return 0
}

func init() {
	proto.RegisterEnum("Game_Event_GameEventType", Game_Event_GameEventType_name, Game_Event_GameEventType_value)
	proto.RegisterEnum("Game_Event_Team", Game_Event_Team_name, Game_Event_Team_value)
	proto.RegisterType((*Game_Event)(nil), "Game_Event")
	proto.RegisterType((*Game_Event_Originator)(nil), "Game_Event.Originator")
}

func init() { proto.RegisterFile("ssl_game_event.proto", fileDescriptor_9d7737834ddab500) }

var fileDescriptor_9d7737834ddab500 = []byte{
	// 532 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x4f, 0x6f, 0x1a, 0x31,
	0x10, 0xc5, 0x0b, 0x21, 0x7f, 0x18, 0xb2, 0xc4, 0x99, 0x90, 0x94, 0xb4, 0x95, 0x12, 0xa5, 0x3d,
	0xe4, 0x94, 0x43, 0x0e, 0x55, 0x6f, 0x95, 0x77, 0x77, 0x00, 0x0b, 0xaf, 0xbd, 0xf2, 0x7a, 0x15,
	0x71, 0xb2, 0xa8, 0xba, 0x42, 0x91, 0x4a, 0x88, 0x00, 0x55, 0xea, 0xa7, 0xe9, 0xe7, 0xec, 0xad,
	0xb2, 0xa1, 0xfc, 0x51, 0x6f, 0x3b, 0x6f, 0xde, 0xf3, 0xfc, 0x3c, 0x6b, 0xe8, 0x2c, 0x16, 0x3f,
	0xdc, 0x64, 0x3c, 0xad, 0x5c, 0xf5, 0xb3, 0x7a, 0x59, 0x3e, 0xbc, 0xce, 0x67, 0xcb, 0xd9, 0xdd,
	0xef, 0x23, 0x80, 0xbe, 0x17, 0xc9, 0x8b, 0xf8, 0x15, 0x22, 0x6f, 0x09, 0x85, 0xfd, 0xf5, 0x5a,
	0x75, 0x6b, 0xb7, 0xf5, 0xfb, 0xf6, 0xe3, 0xf5, 0xc3, 0xd6, 0x13, 0x3e, 0x37, 0x06, 0xb3, 0xef,
	0xc7, 0xcf, 0x00, 0xb3, 0xf9, 0xf3, 0xe4, 0xf9, 0x65, 0xbc, 0x9c, 0xcd, 0xbb, 0xf5, 0xdb, 0xda,
	0x7d, 0xeb, 0xf1, 0x6a, 0x37, 0xad, 0x37, 0x5d, 0xb3, 0xe3, 0xc4, 0x2e, 0x1c, 0x4f, 0xab, 0xc5,
	0x62, 0x3c, 0xa9, 0xba, 0x07, 0xb7, 0xb5, 0xfb, 0xa6, 0xf9, 0x57, 0xbe, 0x1b, 0x00, 0x6c, 0x33,
	0xf8, 0x09, 0x1a, 0xcb, 0x6a, 0x3c, 0x5d, 0x73, 0xb1, 0xdd, 0x93, 0x6d, 0x35, 0x9e, 0x9a, 0xd0,
	0xc5, 0x0e, 0x1c, 0x7e, 0x9b, 0x2d, 0xc5, 0xf7, 0x00, 0x10, 0x99, 0x55, 0x71, 0xf7, 0xe7, 0x00,
	0xa2, 0x3d, 0x78, 0x6c, 0xc1, 0x71, 0xa9, 0x86, 0x4a, 0x3f, 0x29, 0xf6, 0x06, 0x01, 0x8e, 0x92,
	0xb2, 0xb0, 0x3a, 0x63, 0x35, 0xbc, 0x84, 0x73, 0x55, 0x66, 0x31, 0x19, 0xa7, 0x7b, 0x2e, 0x97,
	0x7c, 0x44, 0xa6, 0x60, 0x75, 0xbc, 0x80, 0xb3, 0x98, 0x4b, 0xe9, 0x24, 0xf5, 0xac, 0xeb, 0x09,
	0x92, 0x29, 0x3b, 0xc0, 0x13, 0x68, 0xf4, 0x35, 0x97, 0xac, 0x81, 0x0c, 0x4e, 0x87, 0x22, 0x19,
	0x3a, 0x2b, 0x32, 0xd2, 0xa5, 0x65, 0x87, 0xf8, 0x16, 0x2e, 0x94, 0x76, 0xb9, 0xd1, 0x7d, 0x43,
	0x45, 0xe1, 0x84, 0x72, 0x7d, 0x9e, 0x11, 0x3b, 0xc2, 0x73, 0x88, 0x62, 0x6d, 0x5d, 0xa2, 0xa5,
	0x14, 0x85, 0xd0, 0x8a, 0x1d, 0xfb, 0x99, 0x59, 0x29, 0xad, 0xc8, 0x25, 0xb9, 0x94, 0x7a, 0xa4,
	0x52, 0x32, 0xec, 0x04, 0x6f, 0xe0, 0xfd, 0x7f, 0xb2, 0xcb, 0xb9, 0xb1, 0x82, 0x4b, 0x39, 0x62,
	0x4d, 0xfc, 0x00, 0x5d, 0x6e, 0x2d, 0x4f, 0x86, 0x64, 0xfc, 0x80, 0xe0, 0x29, 0xc8, 0x71, 0x43,
	0x9c, 0x01, 0x36, 0xe1, 0x50, 0x24, 0x42, 0xf5, 0x59, 0x0b, 0xdb, 0x00, 0x81, 0xbe, 0xc8, 0x89,
	0x52, 0x76, 0x8a, 0x1d, 0x60, 0x46, 0x7b, 0x8a, 0xc2, 0xea, 0x7c, 0xad, 0x46, 0x88, 0xd0, 0x0e,
	0xae, 0xd4, 0x88, 0x38, 0x96, 0x3e, 0xd9, 0xc6, 0x6b, 0xb8, 0xdc, 0x8c, 0xb0, 0xba, 0x4c, 0x06,
	0x6e, 0x48, 0x94, 0x93, 0x61, 0x67, 0xfe, 0xce, 0xa9, 0x2e, 0x63, 0x49, 0xab, 0x06, 0x63, 0x7b,
	0x3c, 0x56, 0xaf, 0x78, 0x92, 0x35, 0xcf, 0x39, 0x7e, 0x84, 0x9b, 0xcd, 0x2d, 0xac, 0x76, 0x61,
	0x5f, 0xb9, 0x16, 0xca, 0xba, 0x54, 0x14, 0x96, 0xab, 0x84, 0x18, 0xfa, 0x43, 0x03, 0xc3, 0x40,
	0xcb, 0xd4, 0x13, 0x5c, 0xf8, 0x7d, 0x09, 0x95, 0x0a, 0x43, 0x89, 0x75, 0x61, 0xdb, 0x1d, 0x0f,
	0x15, 0x4c, 0xb9, 0xe4, 0x09, 0x65, 0xa4, 0xac, 0xeb, 0x71, 0x21, 0x29, 0x65, 0x97, 0x3e, 0x9f,
	0x0c, 0x44, 0xee, 0xb4, 0x5a, 0x99, 0xaf, 0xee, 0xbe, 0x40, 0xc3, 0xbf, 0x0f, 0xdf, 0xb1, 0xc4,
	0x33, 0xb7, 0xfd, 0xed, 0x67, 0xd0, 0x0a, 0xca, 0x88, 0xa4, 0xd4, 0x4f, 0xac, 0x86, 0x11, 0x34,
	0x83, 0x10, 0xcb, 0x92, 0x58, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x1c, 0x24, 0xc5,
	0x38, 0x03, 0x00, 0x00,
}
