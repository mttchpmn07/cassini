package graphics

import "github.com/go-gl/glfw/v3.3/glfw"

type Key uint

const (
	MOUSE_BUTTON_LEFT   = 0
	MOUSE_BUTTON_RIGHT  = 1
	MOUSE_BUTTON_MIDDLE = 2
	MOUSE_BUTTON_4      = 3
	MOUSE_BUTTON_5      = 4
	MOUSE_BUTTON_6      = 5
	MOUSE_BUTTON_7      = 6
	MOUSE_BUTTON_LAST   = 7
	KEY_UNKNOWN         = -1
	KEY_SPACE           = 32
	KEY_APOSTROPHE      = 39
	KEY_COMMA           = 44
	KEY_MINUS           = 45
	KEY_PERIOD          = 46
	KEY_SLASH           = 47
	KEY_0               = 48
	KEY_1               = 49
	KEY_2               = 50
	KEY_3               = 51
	KEY_4               = 52
	KEY_5               = 53
	KEY_6               = 54
	KEY_7               = 55
	KEY_8               = 56
	KEY_9               = 57
	KEY_SEMICOLON       = 59
	KEY_EQUAL           = 61
	KEY_A               = 65
	KEY_B               = 66
	KEY_C               = 67
	KEY_D               = 68
	KEY_E               = 69
	KEY_F               = 70
	KEY_G               = 71
	KEY_H               = 72
	KEY_I               = 73
	KEY_J               = 74
	KEY_K               = 75
	KEY_L               = 76
	KEY_M               = 77
	KEY_N               = 78
	KEY_O               = 79
	KEY_P               = 80
	KEY_Q               = 81
	KEY_R               = 82
	KEY_S               = 83
	KEY_T               = 84
	KEY_U               = 85
	KEY_V               = 86
	KEY_W               = 87
	KEY_X               = 88
	KEY_Y               = 89
	KEY_Z               = 90
	KEY_LEFT_BRACKET    = 91
	KEY_BACKSLASH       = 92
	KEY_RIGHT_BRACKET   = 93
	KEY_GRAVE_ACCENT    = 96
	KEY_WORLD_1         = 161
	KEY_WORLD_2         = 162
	KEY_ESCAPE          = 256
	KEY_ENTER           = 257
	KEY_TAB             = 258
	KEY_BACKSPACE       = 259
	KEY_INSERT          = 260
	KEY_DELETE          = 261
	KEY_RIGHT           = 262
	KEY_LEFT            = 263
	KEY_DOWN            = 264
	KEY_UP              = 265
	KEY_PAGE_UP         = 266
	KEY_PAGE_DOWN       = 267
	KEY_HOME            = 268
	KEY_END             = 269
	KEY_CAPS_LOCK       = 280
	KEY_SCROLL_LOCK     = 281
	KEY_NUM_LOCK        = 282
	KEY_PRINT_SCREEN    = 283
	KEY_PAUSE           = 284
	KEY_F1              = 290
	KEY_F2              = 291
	KEY_F3              = 292
	KEY_F4              = 293
	KEY_F5              = 294
	KEY_F6              = 295
	KEY_F7              = 296
	KEY_F8              = 297
	KEY_F9              = 298
	KEY_F10             = 299
	KEY_F11             = 300
	KEY_F12             = 301
	KEY_F13             = 302
	KEY_F14             = 303
	KEY_F15             = 304
	KEY_F16             = 305
	KEY_F17             = 306
	KEY_F18             = 307
	KEY_F19             = 308
	KEY_F20             = 309
	KEY_F21             = 310
	KEY_F22             = 311
	KEY_F23             = 312
	KEY_F24             = 313
	KEY_F25             = 314
	KEY_KP_0            = 320
	KEY_KP_1            = 321
	KEY_KP_2            = 322
	KEY_KP_3            = 323
	KEY_KP_4            = 324
	KEY_KP_5            = 325
	KEY_KP_6            = 326
	KEY_KP_7            = 327
	KEY_KP_8            = 328
	KEY_KP_9            = 329
	KEY_KP_DECIMAL      = 330
	KEY_KP_DIVIDE       = 331
	KEY_KP_MULTIPLY     = 332
	KEY_KP_SUBTRACT     = 333
	KEY_KP_ADD          = 334
	KEY_KP_ENTER        = 335
	KEY_KP_EQUAL        = 336
	KEY_LEFT_SHIFT      = 340
	KEY_LEFT_CONTROL    = 341
	KEY_LEFT_ALT        = 342
	KEY_LEFT_SUPER      = 343
	KEY_RIGHT_SHIFT     = 344
	KEY_RIGHT_CONTROL   = 345
	KEY_RIGHT_ALT       = 346
	KEY_RIGHT_SUPER     = 347
	KEY_LAST            = 348
)

var GLFWkeyMap = map[glfw.Key]Key{
	glfw.Key(glfw.MouseButton1): Key(glfw.MouseButton1),
}

var KeyMap = map[string]int{
	"MOUSE_BUTTON_LEFT":   MOUSE_BUTTON_LEFT,
	"MOUSE_BUTTON_RIGHT":  MOUSE_BUTTON_RIGHT,
	"MOUSE_BUTTON_MIDDLE": MOUSE_BUTTON_MIDDLE,
	"MOUSE_BUTTON_4":      MOUSE_BUTTON_4,
	"MOUSE_BUTTON_5":      MOUSE_BUTTON_5,
	"MOUSE_BUTTON_6":      MOUSE_BUTTON_6,
	"MOUSE_BUTTON_7":      MOUSE_BUTTON_7,
	"MOUSE_BUTTON_LAST":   MOUSE_BUTTON_LAST,
	"KEY_SPACE":           KEY_SPACE,
	"KEY_APOSTROPHE":      KEY_APOSTROPHE,
	"KEY_COMMA":           KEY_COMMA,
	"KEY_MINUS":           KEY_MINUS,
	"KEY_PERIOD":          KEY_PERIOD,
	"KEY_SLASH":           KEY_SLASH,
	"KEY_0":               KEY_0,
	"KEY_1":               KEY_1,
	"KEY_2":               KEY_2,
	"KEY_3":               KEY_3,
	"KEY_4":               KEY_4,
	"KEY_5":               KEY_5,
	"KEY_6":               KEY_6,
	"KEY_7":               KEY_7,
	"KEY_8":               KEY_8,
	"KEY_9":               KEY_9,
	"KEY_SEMICOLON":       KEY_SEMICOLON,
	"KEY_EQUAL":           KEY_EQUAL,
	"KEY_A":               KEY_A,
	"KEY_B":               KEY_B,
	"KEY_C":               KEY_C,
	"KEY_D":               KEY_D,
	"KEY_E":               KEY_E,
	"KEY_F":               KEY_F,
	"KEY_G":               KEY_G,
	"KEY_H":               KEY_H,
	"KEY_I":               KEY_I,
	"KEY_J":               KEY_J,
	"KEY_K":               KEY_K,
	"KEY_L":               KEY_L,
	"KEY_M":               KEY_M,
	"KEY_N":               KEY_N,
	"KEY_O":               KEY_O,
	"KEY_P":               KEY_P,
	"KEY_Q":               KEY_Q,
	"KEY_R":               KEY_R,
	"KEY_S":               KEY_S,
	"KEY_T":               KEY_T,
	"KEY_U":               KEY_U,
	"KEY_V":               KEY_V,
	"KEY_W":               KEY_W,
	"KEY_X":               KEY_X,
	"KEY_Y":               KEY_Y,
	"KEY_Z":               KEY_Z,
	"KEY_LEFT_BRACKET":    KEY_LEFT_BRACKET,
	"KEY_BACKSLASH":       KEY_BACKSLASH,
	"KEY_RIGHT_BRACKET":   KEY_RIGHT_BRACKET,
	"KEY_GRAVE_ACCENT":    KEY_GRAVE_ACCENT,
	"KEY_WORLD_1":         KEY_WORLD_1,
	"KEY_WORLD_2":         KEY_WORLD_2,
	"KEY_ESCAPE":          KEY_ESCAPE,
	"KEY_ENTER":           KEY_ENTER,
	"KEY_TAB":             KEY_TAB,
	"KEY_BACKSPACE":       KEY_BACKSPACE,
	"KEY_INSERT":          KEY_INSERT,
	"KEY_DELETE":          KEY_DELETE,
	"KEY_RIGHT":           KEY_RIGHT,
	"KEY_LEFT":            KEY_LEFT,
	"KEY_DOWN":            KEY_DOWN,
	"KEY_UP":              KEY_UP,
	"KEY_PAGE_UP":         KEY_PAGE_UP,
	"KEY_PAGE_DOWN":       KEY_PAGE_DOWN,
	"KEY_HOME":            KEY_HOME,
	"KEY_END":             KEY_END,
	"KEY_CAPS_LOCK":       KEY_CAPS_LOCK,
	"KEY_SCROLL_LOCK":     KEY_SCROLL_LOCK,
	"KEY_NUM_LOCK":        KEY_NUM_LOCK,
	"KEY_PRINT_SCREEN":    KEY_PRINT_SCREEN,
	"KEY_PAUSE":           KEY_PAUSE,
	"KEY_F1":              KEY_F1,
	"KEY_F2":              KEY_F2,
	"KEY_F3":              KEY_F3,
	"KEY_F4":              KEY_F4,
	"KEY_F5":              KEY_F5,
	"KEY_F6":              KEY_F6,
	"KEY_F7":              KEY_F7,
	"KEY_F8":              KEY_F8,
	"KEY_F9":              KEY_F9,
	"KEY_F10":             KEY_F10,
	"KEY_F11":             KEY_F11,
	"KEY_F12":             KEY_F12,
	"KEY_F13":             KEY_F13,
	"KEY_F14":             KEY_F14,
	"KEY_F15":             KEY_F15,
	"KEY_F16":             KEY_F16,
	"KEY_F17":             KEY_F17,
	"KEY_F18":             KEY_F18,
	"KEY_F19":             KEY_F19,
	"KEY_F20":             KEY_F20,
	"KEY_F21":             KEY_F21,
	"KEY_F22":             KEY_F22,
	"KEY_F23":             KEY_F23,
	"KEY_F24":             KEY_F24,
	"KEY_F25":             KEY_F25,
	"KEY_KP_0":            KEY_KP_0,
	"KEY_KP_1":            KEY_KP_1,
	"KEY_KP_2":            KEY_KP_2,
	"KEY_KP_3":            KEY_KP_3,
	"KEY_KP_4":            KEY_KP_4,
	"KEY_KP_5":            KEY_KP_5,
	"KEY_KP_6":            KEY_KP_6,
	"KEY_KP_7":            KEY_KP_7,
	"KEY_KP_8":            KEY_KP_8,
	"KEY_KP_9":            KEY_KP_9,
	"KEY_KP_DECIMAL":      KEY_KP_DECIMAL,
	"KEY_KP_DIVIDE":       KEY_KP_DIVIDE,
	"KEY_KP_MULTIPLY":     KEY_KP_MULTIPLY,
	"KEY_KP_SUBTRACT":     KEY_KP_SUBTRACT,
	"KEY_KP_ADD":          KEY_KP_ADD,
	"KEY_KP_ENTER":        KEY_KP_ENTER,
	"KEY_KP_EQUAL":        KEY_KP_EQUAL,
	"KEY_LEFT_SHIFT":      KEY_LEFT_SHIFT,
	"KEY_LEFT_CONTROL":    KEY_LEFT_CONTROL,
	"KEY_LEFT_ALT":        KEY_LEFT_ALT,
	"KEY_LEFT_SUPER":      KEY_LEFT_SUPER,
	"KEY_RIGHT_SHIFT":     KEY_RIGHT_SHIFT,
	"KEY_RIGHT_CONTROL":   KEY_RIGHT_CONTROL,
	"KEY_RIGHT_ALT":       KEY_RIGHT_ALT,
	"KEY_RIGHT_SUPER":     KEY_RIGHT_SUPER,
	"KEY_LAST":            KEY_LAST,
}
