package parser

var keycodeMapping = map[int]string{
	0:   "KC_NO",
	1:   "KC_ROLL_OVER",
	2:   "KC_POST_FAIL",
	3:   "KC_UNDEFINED",
	4:   "KC_A",
	5:   "KC_B",
	6:   "KC_C",
	7:   "KC_D",
	8:   "KC_E",
	9:   "KC_F",
	10:  "KC_G",
	11:  "KC_H",
	12:  "KC_I",
	13:  "KC_J",
	14:  "KC_K",
	15:  "KC_L",
	16:  "KC_M",
	17:  "KC_N",
	18:  "KC_O",
	19:  "KC_P",
	20:  "KC_Q",
	21:  "KC_R",
	22:  "KC_S",
	23:  "KC_T",
	24:  "KC_U",
	25:  "KC_V",
	26:  "KC_W",
	27:  "KC_X",
	28:  "KC_Y",
	29:  "KC_Z",
	30:  "KC_1",
	31:  "KC_2",
	32:  "KC_3",
	33:  "KC_4",
	34:  "KC_5",
	35:  "KC_6",
	36:  "KC_7",
	37:  "KC_8",
	38:  "KC_9",
	39:  "KC_0",
	40:  "KC_ENTER",
	41:  "KC_ESCAPE",
	42:  "KC_BSPACE",
	43:  "KC_TAB",
	44:  "KC_SPACE",
	45:  "KC_MINUS",
	46:  "KC_EQUAL",
	47:  "KC_LBRACKET",
	48:  "KC_RBRACKET",
	49:  "KC_BSLASH",
	50:  "KC_NONUS_HASH",
	51:  "KC_SCOLON",
	52:  "KC_QUOTE",
	53:  "KC_GRAVE",
	54:  "KC_COMMA",
	55:  "KC_DOT",
	56:  "KC_SLASH",
	57:  "KC_CAPSLOCK",
	58:  "KC_F1",
	59:  "KC_F2",
	60:  "KC_F3",
	61:  "KC_F4",
	62:  "KC_F5",
	63:  "KC_F6",
	64:  "KC_F7",
	65:  "KC_F8",
	66:  "KC_F9",
	67:  "KC_F10",
	68:  "KC_F11",
	69:  "KC_F12",
	70:  "KC_PSCREEN",
	71:  "KC_SCROLLLOCK",
	72:  "KC_PAUSE",
	73:  "KC_INSERT",
	74:  "KC_HOME",
	75:  "KC_PGUP",
	76:  "KC_DELETE",
	77:  "KC_END",
	78:  "KC_PGDOWN",
	79:  "KC_RIGHT",
	80:  "KC_LEFT",
	81:  "KC_DOWN",
	82:  "KC_UP",
	83:  "KC_NUMLOCK",
	84:  "KC_KP_SLASH",
	85:  "KC_KP_ASTERISK",
	86:  "KC_KP_MINUS",
	87:  "KC_KP_PLUS",
	88:  "KC_KP_ENTER",
	89:  "KC_KP_1",
	90:  "KC_KP_2",
	91:  "KC_KP_3",
	92:  "KC_KP_4",
	93:  "KC_KP_5",
	94:  "KC_KP_6",
	95:  "KC_KP_7",
	96:  "KC_KP_8",
	97:  "KC_KP_9",
	98:  "KC_KP_0",
	99:  "KC_KP_DOT",
	100: "KC_NONUS_BSLASH",
	101: "KC_APPLICATION",
	102: "KC_POWER",
	103: "KC_KP_EQUAL",
	104: "KC_F13",
	105: "KC_F14",
	106: "KC_F15",
	107: "KC_F16",
	108: "KC_F17",
	109: "KC_F18",
	110: "KC_F19",
	111: "KC_F20",
	112: "KC_F21",
	113: "KC_F22",
	114: "KC_F23",
	115: "KC_F24",
	116: "KC_EXECUTE",
	117: "KC_HELP",
	118: "KC_MENU",
	119: "KC_SELECT",
	120: "KC_STOP",
	121: "KC_AGAIN",
	122: "KC_UNDO",
	123: "KC_CUT",
	124: "KC_COPY",
	125: "KC_PASTE",
	126: "KC_FIND",
	127: "KC__MUTE",
	128: "KC__VOLUP",
	129: "KC__VOLDOWN",
	130: "KC_LOCKING_CAPS",
	131: "KC_LOCKING_NUM",
	132: "KC_LOCKING_SCROLL",
	133: "KC_KP_COMMA",
	134: "KC_KP_EQUAL_AS400",
	135: "KC_INT1",
	136: "KC_INT2",
	137: "KC_INT3",
	138: "KC_INT4",
	139: "KC_INT5",
	140: "KC_INT6",
	141: "KC_INT7",
	142: "KC_INT8",
	143: "KC_INT9",
	144: "KC_LANG1",
	145: "KC_LANG2",
	146: "KC_LANG3",
	147: "KC_LANG4",
	148: "KC_LANG5",
	149: "KC_LANG6",
	150: "KC_LANG7",
	151: "KC_LANG8",
	152: "KC_LANG9",
	153: "KC_ALT_ERASE",
	154: "KC_SYSREQ",
	155: "KC_CANCEL",
	156: "KC_CLEAR",
	157: "KC_PRIOR",
	158: "KC_RETURN",
	159: "KC_SEPARATOR",
	160: "KC_OUT",
	161: "KC_OPER",
	162: "KC_CLEAR_AGAIN",
	163: "KC_CRSEL",
	164: "KC_EXSEL",
	224: "KC_LCTRL",
	225: "KC_LSHIFT",
	226: "KC_LALT",
	227: "KC_LGUI",
	228: "KC_RCTRL",
	229: "KC_RSHIFT",
	230: "KC_RALT",
	231: "KC_RGUI",
	165: "KC_SYSTEM_POWER",
	166: "KC_SYSTEM_SLEEP",
	167: "KC_SYSTEM_WAKE",
	168: "KC_AUDIO_MUTE",
	169: "KC_AUDIO_VOL_UP",
	170: "KC_AUDIO_VOL_DOWN",
	171: "KC_MEDIA_NEXT_TRACK",
	172: "KC_MEDIA_PREV_TRACK",
	173: "KC_MEDIA_STOP",
	174: "KC_MEDIA_PLAY_PAUSE",
	175: "KC_MEDIA_SELECT",
	176: "KC_MEDIA_EJECT",
	177: "KC_CALCULATOR",
	178: "KC_MY_COMPUTER",
	179: "KC_WWW_SEARCH",
	180: "KC_WWW_HOME",
	181: "KC_WWW_BACK",
	182: "KC_WWW_FORWARD",
	183: "KC_WWW_STOP",
	184: "KC_WWW_REFRESH",
	185: "KC_WWW_FAVORITES",
	186: "KC_MEDIA_FAST_FORWARD",
	187: "KC_MEDIA_REWIND",
	188: "KC_BRIGHTNESS_UP",
	189: "KC_BRIGHTNESS_DOWN",
	240: "KC_MS_UP",
	241: "KC_MS_DOWN",
	242: "KC_MS_LEFT",
	243: "KC_MS_RIGHT",
	244: "KC_MS_BTN1",
	245: "KC_MS_BTN2",
	246: "KC_MS_BTN3",
	247: "KC_MS_BTN4",
	248: "KC_MS_BTN5",
	249: "KC_MS_WH_UP",
	250: "KC_MS_WH_DOWN",
	251: "KC_MS_WH_LEFT",
	252: "KC_MS_WH_RIGHT",
	253: "KC_MS_ACCEL0",
	254: "KC_MS_ACCEL1",
	255: "KC_MS_ACCEL2",
}

const (
	UnknownKeycode = "UNKNOWN"
)

func mapKeycode(keycode int) string {
	if val, ok := keycodeMapping[keycode]; ok {
		return val
	}
	return UnknownKeycode
}