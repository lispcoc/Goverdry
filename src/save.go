package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"unsafe"

	"github.com/buke/quickjs-go"
)

var SaveRuntime quickjs.Runtime
var SaveContext *quickjs.Context

const PC_ENTRY_MIN = 36
const RESIST_LENGTH = 15
const ATTACK_ADD_LENGTH = 5
const GoverdrySaveTweak = false

type Data struct {
	GameData GameData
	PlayData PlayData
	PC       []PC
	PARTY    []PARTY
	ITEM     []ITEM
}

type PlayData struct {
	ShopItemListFull bool
	GameFlag         []bool
	ItemFlag         []bool
	MonsterFlag      []bool
	PcMax            int
	ActiveParty      int
	PcList           []int
	GarbageItem      []string
	MapFlag          [10][16][20][20]bool
	SecretDoor       [10][16][512]bool
	LockedDoor       [10][16][512]bool
	DungeonNewMusic  [10][16]string
	WallFlag         [10][16][512]bool
	GameFlagS        []bool
}

type GameData struct {
	READ_KEYWORD string
	CastleTown   string
	PC_ENTRY_MAX int
	ABILITY      []int
}

type PC struct {
	SealSpell       bool
	Sex             int
	Race            int
	Alignment       int
	PcClass         int
	Ability         []int
	State           int
	PartyNum        int
	Age             int
	Days            int
	HpMax0          int
	Hp              int
	SealSpellInt    int
	Level           int
	Rip             int
	Poison          int
	Exp             int
	Gold            int
	Marks           int
	PoisonPlus      int
	SpellEffectRate []int
	SpellEffectPlus []int
	Equip           []int
	ItemDecided     []int
	Item            []int
	Ability0        []int
	AbiPlus         []int
	AbiRate         []int
	ResistPlus      []int
	ResistRate      []int
	AttackAddRate   []int
	AttackAddPlus   []bool
	Mp              [][]int
	MpMax           [][]int
	Spell           [][][]bool
	Title           string
	Transmigrates   string
	FaceGraphic     string
	Name            []string
}

type PARTY struct {
	Light             bool
	ViewDarkZone      bool
	FlyingEffect      string
	DungeonMusic      string
	DungeonNumber     int
	Floor             int
	X                 int
	Y                 int
	Direction         int
	Vision            int
	Flying            int
	PoisonPlus        int
	PtSpellEffectRate []int
	PtSpellEffectPlus []int
	SpellEffectRate   []int
	SpellEffectPlus   []int
	PartyMember       []int
	AbiPlus           []int
	AbiRate           []int
	ResistPlus        []int
	ResistRate        []int
	AttackAddRate     []int
	AttackAddPlus     []bool
}

type ITEM struct {
	Stock int
}

func lastTrue(GameFlag []bool) int {
	r := 0
	println(len(GameFlag))
	for i := 0; i < len(GameFlag); i++ {
		if GameFlag[i] {
			println(i)
			r = i
		}
	}
	return r
}

func initSave(ctx *quickjs.Context) {
	save := ctx.Object()
	save.Set("getSaveDataStr", ctx.Function(func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		println("getSaveDataStr start")
		s := getSaveDataStr(args[0].String())
		println("getSaveDataStr end")
		return ctx.String(s)
	}))
	ctx.Globals().Set("SAVE", save)
}

func appendStr(a []byte, b string) []byte {
	s := *(*[]byte)(unsafe.Pointer(&b))
	for i := 0; i < len(b); i++ {
		a = append(a, s[i])
	}
	return a
}

func getSaveDataStr(json_str string) string {
	var data_in Data
	test := 0
	fmt.Printf("%d\n", test)

	json.Unmarshal([]byte(json_str), &data_in)

	GameData := data_in.GameData
	PlayData := data_in.PlayData
	PC := data_in.PC
	PARTY := data_in.PARTY
	ITEM := data_in.ITEM

	data := make([]byte, 0, 2*1024*1024)

	data = appendStr(data, GameData.READ_KEYWORD)
	data = appendStr(data, "\n")
	if PlayData.ShopItemListFull {
		data = appendStr(data, "1")
	} else {
		data = appendStr(data, "0")
	}
	data = appendStr(data, "\n")
	var num = lastTrue(PlayData.GameFlag)
	if num < 0 {
		num = 0
	} else if num > 9999 {
		num = 9999
	}
	var i, j, k, l int
	for i = 0; i < num; i++ {
		if i < len(PlayData.GameFlag) && PlayData.GameFlag[i] {
			data = appendStr(data, "1")
		} else {
			data = appendStr(data, "0")
		}
	}
	data = appendStr(data, "\n")
	num = lastTrue(PlayData.ItemFlag)
	if num < 0 {
		num = 0
	} else if num > 9999 {
		num = 9999
	}
	num++
	for i = 0; i < num; i++ {
		if i < len(PlayData.ItemFlag) && PlayData.ItemFlag[i] {
			data = appendStr(data, "1")
		} else {
			data = appendStr(data, "0")
		}
	}
	data = appendStr(data, "\n")
	num = lastTrue(PlayData.MonsterFlag)
	if num < 0 {
		num = 0
	} else if num > 9999 {
		num = 9999
	}
	num++
	for i = 0; i < num; i++ {
		if i < len(PlayData.MonsterFlag) && PlayData.MonsterFlag[i] {
			data = appendStr(data, "1")
		} else {
			data = appendStr(data, "0")
		}
	}
	data = appendStr(data, "\n")
	data = appendStr(data, strconv.Itoa(PlayData.PcMax))
	data = appendStr(data, "\n")
	data = appendStr(data, strconv.Itoa(PlayData.ActiveParty))
	data = appendStr(data, "\n")
	var pEyMax = GameData.PC_ENTRY_MAX
	for i = 0; (i < pEyMax) && i < len(PlayData.PcList); i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PlayData.PcList[i]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < len(PlayData.GarbageItem); i++ {
		if i > 0 {
			data = appendStr(data, "<>")
		}
		data = appendStr(data, PlayData.GarbageItem[i])
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if PC[i].SealSpell {
			data = appendStr(data, "1")
		} else {
			data = appendStr(data, "0")
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		data = appendStr(data, strconv.Itoa(PC[i].Sex))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].Race))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		data = appendStr(data, strconv.Itoa(PC[i].Alignment))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].PcClass))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].State))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].PartyNum))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].Age))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].Days))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].HpMax0))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].Hp))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SealSpellInt))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].Level))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].Rip))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].Poison))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectRate[0]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectPlus[0]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectRate[1]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectPlus[1]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectRate[5]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectPlus[5]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectRate[2]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectPlus[2]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectRate[3]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectPlus[3]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectRate[4]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectPlus[4]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].Exp))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].Gold))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].Marks))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, "<>")
		}
		data = appendStr(data, PC[i].Name[0])
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, "<>")
		}
		data = appendStr(data, PC[i].Title)
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, "<>")
		}
		data = appendStr(data, PC[i].Transmigrates)
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].DungeonNumber))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].Floor))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].X))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].Y))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		data = appendStr(data, strconv.Itoa(PARTY[i].Direction))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].Vision))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].PtSpellEffectRate[0]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].PtSpellEffectPlus[0]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectRate[0]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectPlus[0]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].Flying))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectRate[1]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectPlus[1]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectRate[5]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectPlus[5]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectRate[2]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectPlus[2]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectRate[3]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectPlus[3]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectRate[4]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectPlus[4]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].PtSpellEffectRate[1]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].PtSpellEffectPlus[1]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].PtSpellEffectRate[2]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].PtSpellEffectPlus[2]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].PtSpellEffectRate[3]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].PtSpellEffectPlus[3]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if PARTY[i].Light {
			data = appendStr(data, "1")
		} else {
			data = appendStr(data, "0")
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, "<>")
		}
		data = appendStr(data, PARTY[i].FlyingEffect)
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 99; j++ {
			data = appendStr(data, strconv.Itoa(PC[i].Equip[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 99; j++ {
			data = appendStr(data, strconv.Itoa(PC[i].ItemDecided[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 99; j++ {
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].Item[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 10; j++ {
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			if j < len(GameData.ABILITY) {
				data = appendStr(data, strconv.Itoa(PC[i].Ability0[j]))
			} else {
				data = appendStr(data, "-1")
			}
		}
	}
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 10; j < 36; j++ {
			data = appendStr(data, ",")
			if j < len(GameData.ABILITY) {
				data = appendStr(data, strconv.Itoa(PC[i].Ability0[j]))
			} else {
				data = appendStr(data, "-1")
			}
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 4; j++ {
			for k = 0; k < 10; k++ {
				if i > 0 || j > 0 || k > 0 {
					data = appendStr(data, ",")
				}
				data = appendStr(data, strconv.Itoa(PC[i].Mp[j][k]))
			}
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 4; j++ {
			for k = 0; k < 10; k++ {
				if i > 0 || j > 0 || k > 0 {
					data = appendStr(data, ",")
				}
				data = appendStr(data, strconv.Itoa(PC[i].MpMax[j][k]))
			}
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 4; j++ {
			for k = 0; k < 10; k++ {
				for l = 0; l < 6; l++ {
					if PC[i].Spell[j][k][l] {
						data = appendStr(data, "1")
					} else {
						data = appendStr(data, "0")
					}
				}
			}
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 6; j++ {
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].PartyMember[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < len(ITEM); i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(ITEM[i].Stock))
	}
	data = appendStr(data, "\n")
	for i = 0; i < 10; i++ {
		for j = 0; j < 16; j++ {
			for k = 0; k < 20; k++ {
				for l = 0; l < 20; l++ {
					if PlayData.MapFlag[i][j][k][l] {
						data = appendStr(data, "1")
					} else {
						data = appendStr(data, "0")
					}
				}
			}
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < 10; i++ {
		for j = 0; j < 16; j++ {
			for k = 0; k < 64; k++ {
				if PlayData.SecretDoor[i][j][k] {
					data = appendStr(data, "1")
				} else {
					data = appendStr(data, "0")
				}
			}
		}
	}
	for i = 0; i < 10; i++ {
		for j = 0; j < 16; j++ {
			for k = 64; k < 512; k++ {
				if PlayData.SecretDoor[i][j][k] {
					data = appendStr(data, "1")
				} else {
					data = appendStr(data, "0")
				}
			}
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < 10; i++ {
		for j = 0; j < 16; j++ {
			for k = 0; k < 64; k++ {
				if PlayData.LockedDoor[i][j][k] {
					data = appendStr(data, "1")
				} else {
					data = appendStr(data, "0")
				}
			}
		}
	}
	for i = 0; i < 10; i++ {
		for j = 0; j < 16; j++ {
			for k = 64; k < 512; k++ {
				if PlayData.LockedDoor[i][j][k] {
					data = appendStr(data, "1")
				} else {
					data = appendStr(data, "0")
				}
			}
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, "<>")
		}
		data = appendStr(data, PARTY[i].DungeonMusic)
	}
	data = appendStr(data, "\n")
	for i = 0; i < 10; i++ {
		if i > 0 {
			data = appendStr(data, "<>")
		}
		for j = 0; j < 16; j++ {
			if j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, PlayData.DungeonNewMusic[i][j])
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 10; j++ {
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].AbiPlus[j]))
		}
	}
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 10; j < 36; j++ {
			data = appendStr(data, ","+strconv.Itoa(PC[i].AbiPlus[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 10; j++ {
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].AbiRate[j]))
		}
	}
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 10; j < 36; j++ {
			data = appendStr(data, ","+strconv.Itoa(PC[i].AbiRate[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 10; j++ {
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].AbiPlus[j]))
		}
	}
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 10; j < 36; j++ {
			data = appendStr(data, ","+strconv.Itoa(PARTY[i].AbiPlus[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < 10; j++ {
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].AbiRate[j]))
		}
	}
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 10; j < 36; j++ {
			data = appendStr(data, ","+strconv.Itoa(PARTY[i].AbiRate[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].PtSpellEffectRate[4]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].PtSpellEffectPlus[4]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectRate[6]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectPlus[6]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectRate[6]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectPlus[6]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectPlus[7]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectRate[7]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectPlus[7]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectRate[7]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if PARTY[i].ViewDarkZone {
			data = appendStr(data, "1")
		} else {
			data = appendStr(data, "0")
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectPlus[8]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].SpellEffectRate[8]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectPlus[8]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].SpellEffectRate[8]))
	}
	data = appendStr(data, "\n")
	for i = 0; i < 10; i++ {
		for j = 0; j < 16; j++ {
			for k = 0; k < 64; k++ {
				if PlayData.WallFlag[i][j][k] {
					data = appendStr(data, "1")
				} else {
					data = appendStr(data, "0")
				}
			}
		}
	}
	for i = 0; i < 10; i++ {
		for j = 0; j < 16; j++ {
			for k = 64; k < 512; k++ {
				if PlayData.WallFlag[i][j][k] {
					data = appendStr(data, "1")
				} else {
					data = appendStr(data, "0")
				}
			}
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < RESIST_LENGTH; j++ {
			if j == 9 {
				continue
			}
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].ResistPlus[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < RESIST_LENGTH; j++ {
			if j == 9 {
				continue
			}
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].ResistRate[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < RESIST_LENGTH; j++ {
			if j == 9 {
				continue
			}
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].ResistPlus[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < RESIST_LENGTH; j++ {
			if j == 9 {
				continue
			}
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].ResistRate[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < ATTACK_ADD_LENGTH+1; j++ {
			if PC[i].AttackAddPlus[j] {
				data = appendStr(data, "1")
			} else {
				data = appendStr(data, "0")
			}
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < ATTACK_ADD_LENGTH+1; j++ {
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].AttackAddRate[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < ATTACK_ADD_LENGTH+1; j++ {
			if PARTY[i].AttackAddPlus[j] {
				data = appendStr(data, "1")
			} else {
				data = appendStr(data, "0")
			}
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < PC_ENTRY_MIN; i++ {
		for j = 0; j < ATTACK_ADD_LENGTH+1; j++ {
			if i > 0 || j > 0 {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].AttackAddRate[j]))
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PC[i].PoisonPlus))
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, strconv.Itoa(PARTY[i].PoisonPlus))
	}
	data = appendStr(data, "\n")
	for i = 0; i < len(PlayData.GameFlagS); i++ {
		if PlayData.GameFlagS[i] {
			data = appendStr(data, "1")
		} else {
			data = appendStr(data, "0")
		}
	}
	data = appendStr(data, "\n")
	for i = 0; i < pEyMax; i++ {
		if i > 0 {
			data = appendStr(data, ",")
		}
		data = appendStr(data, PC[i].FaceGraphic)
	}
	for j = 0; j < 99; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			data = appendStr(data, strconv.Itoa(PC[i].Equip[j]))
		}
	}
	for j = 0; j < 99; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			data = appendStr(data, strconv.Itoa(PC[i].ItemDecided[j]))
		}
	}
	for j = 0; j < 99; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].Item[j]))
		}
	}
	for j = 0; j < 36; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			if j < len(GameData.ABILITY) {
				data = appendStr(data, strconv.Itoa(PC[i].Ability0[j]))
			} else {
				data = appendStr(data, "-1")
			}
		}
	}
	for k = 0; k < 10; k++ {
		for j = 0; j < 4; j++ {
			data = appendStr(data, "\n")
			for i = PC_ENTRY_MIN; i < pEyMax; i++ {
				if i > PC_ENTRY_MIN {
					data = appendStr(data, ",")
				}
				data = appendStr(data, strconv.Itoa(PC[i].Mp[j][k]))
			}
		}
	}
	for k = 0; k < 10; k++ {
		for j = 0; j < 4; j++ {
			data = appendStr(data, "\n")
			for i = PC_ENTRY_MIN; i < pEyMax; i++ {
				if i > PC_ENTRY_MIN {
					data = appendStr(data, ",")
				}
				data = appendStr(data, strconv.Itoa(PC[i].MpMax[j][k]))
			}
		}
	}
	for l = 0; l < 6; l++ {
		for k = 0; k < 10; k++ {
			for j = 0; j < 4; j++ {
				data = appendStr(data, "\n")
				for i = PC_ENTRY_MIN; i < pEyMax; i++ {
					if PC[i].Spell[j][k][l] {
						data = appendStr(data, "1")
					} else {
						data = appendStr(data, "0")
					}
				}
			}
		}
	}
	for j = 0; j < 6; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].PartyMember[j]))
		}
	}
	for j = 0; j < 36; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].AbiPlus[j]))
		}
	}
	for j = 0; j < 36; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].AbiRate[j]))
		}
	}
	for j = 0; j < 36; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].AbiPlus[j]))
		}
	}
	for j = 0; j < 36; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].AbiRate[j]))
		}
	}
	for j = 0; j < RESIST_LENGTH; j++ {
		if j == 9 {
			continue
		}
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].ResistPlus[j]))
		}
	}
	for j = 0; j < RESIST_LENGTH; j++ {
		if j == 9 {
			continue
		}
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].ResistRate[j]))
		}
	}
	for j = 0; j < RESIST_LENGTH; j++ {
		if j == 9 {
			continue
		}
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].ResistPlus[j]))
		}
	}
	for j = 0; j < RESIST_LENGTH; j++ {
		if j == 9 {
			continue
		}
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].ResistRate[j]))
		}
	}
	for j = 0; j < ATTACK_ADD_LENGTH+1; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if PC[i].AttackAddPlus[j] {
				data = appendStr(data, "1")
			} else {
				data = appendStr(data, "0")
			}
		}
	}
	for j = 0; j < ATTACK_ADD_LENGTH+1; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PC[i].AttackAddRate[j]))
		}
	}
	for j = 0; j < ATTACK_ADD_LENGTH+1; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if PARTY[i].AttackAddPlus[j] {
				data = appendStr(data, "1")
			} else {
				data = appendStr(data, "0")
			}
		}
	}
	for j = 0; j < ATTACK_ADD_LENGTH+1; j++ {
		data = appendStr(data, "\n")
		for i = PC_ENTRY_MIN; i < pEyMax; i++ {
			if i > PC_ENTRY_MIN {
				data = appendStr(data, ",")
			}
			data = appendStr(data, strconv.Itoa(PARTY[i].AttackAddRate[j]))
		}
	}

	return base64.StdEncoding.EncodeToString([]byte(data))
}
