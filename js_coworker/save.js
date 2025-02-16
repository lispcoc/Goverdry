var RACE_CLASS_MAX = 36
var PC_ENTRY_MIN = 36
var RESIST_LENGTH = 15
var ATTACK_ADD_LENGTH = 5
var GoverdrySaveTweak = false

var json = JSON.parse(json_str)
var GameData = json.GameData
var PlayData = json.PlayData
var PC = json.PC
var PARTY = json.PARTY
var ITEM = json.ITEM
var ret = getSaveDataStr()
localStorage.setItem(file, ret)

function getSaveDataStr () {
  let data = ''
  data += GameData['READ_KEYWORD']
  data += '\n'
  PlayData['ShopItemListFull'] ? (data += '1') : (data += '0')
  data += '\n'
  let num = PlayData['GameFlag'].lastIndexOf(true)
  if (num < 0) {
    num = 0
  } else if (num > 9999) {
    num = 9999
  }
  num++
  for (let i = 0; i < num; i++) {
    PlayData['GameFlag'][i] ? (data += '1') : (data += '0')
  }
  data += '\n'
  num = PlayData['ItemFlag'].lastIndexOf(true)
  if (num < 0) {
    num = 0
  } else if (num > 9999) {
    num = 9999
  }
  num++
  for (let i = 0; i < num; i++) {
    PlayData['ItemFlag'][i] ? (data += '1') : (data += '0')
  }
  data += '\n'
  num = PlayData['MonsterFlag'].lastIndexOf(true)
  if (num < 0) {
    num = 0
  } else if (num > 9999) {
    num = 9999
  }
  num++
  for (let i = 0; i < num; i++) {
    PlayData['MonsterFlag'][i] ? (data += '1') : (data += '0')
  }
  data += '\n'
  data += String(PlayData['PcMax'])
  data += '\n'
  data += String(PlayData['ActiveParty'])
  data += '\n'
  let pEyMax = GameData['PC_ENTRY_MAX']
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PlayData['PcList'][i])
  }
  data += '\n'
  for (let i = 0; i < PlayData['GarbageItem'].length; i++) {
    if (i > 0) {
      data += '<>'
    }
    data += PlayData['GarbageItem'][i]
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    PC[i].SealSpell ? (data += '1') : (data += '0')
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    data += String(PC[i].Sex)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].Race)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    data += String(PC[i].Alignment)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].PcClass)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].State)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].PartyNum)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].Age)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].Days)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].HpMax0)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].Hp)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SealSpellInt)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].Level)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].Rip)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].Poison)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectRate[0])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectPlus[0])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectRate[1])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectPlus[1])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectRate[5])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectPlus[5])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectRate[2])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectPlus[2])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectRate[3])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectPlus[3])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectRate[4])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectPlus[4])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].Exp)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].Gold)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].Marks)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += '<>'
    }
    data += PC[i].Name[0]
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += '<>'
    }
    data += PC[i].Title
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += '<>'
    }
    data += PC[i].Transmigrates
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].DungeonNumber)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].Floor)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].X)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].Y)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    data += String(PARTY[i].Direction)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].Vision)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].PtSpellEffectRate[0])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].PtSpellEffectPlus[0])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectRate[0])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectPlus[0])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].Flying)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectRate[1])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectPlus[1])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectRate[5])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectPlus[5])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectRate[2])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectPlus[2])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectRate[3])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectPlus[3])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectRate[4])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectPlus[4])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].PtSpellEffectRate[1])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].PtSpellEffectPlus[1])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].PtSpellEffectRate[2])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].PtSpellEffectPlus[2])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].PtSpellEffectRate[3])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].PtSpellEffectPlus[3])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    PARTY[i].Light ? (data += '1') : (data += '0')
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += '<>'
    }
    data += PARTY[i].FlyingEffect
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 99; j++) {
      data += String(PC[i].Equip[j])
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 99; j++) {
      data += String(PC[i].ItemDecided[j])
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 99; j++) {
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PC[i].Item[j])
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 10; j++) {
      if (i > 0 || j > 0) {
        data += ','
      }
      if (j < GameData['ABILITY'].length) {
        data += String(PC[i].Ability0[j])
      } else {
        data += '-1'
      }
    }
  }
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 10; j < 36; j++) {
      data += ','
      if (j < GameData['ABILITY'].length) {
        data += String(PC[i].Ability0[j])
      } else {
        data += '-1'
      }
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 4; j++) {
      for (let k = 0; k < 10; k++) {
        if (i > 0 || j > 0 || k > 0) {
          data += ','
        }
        data += String(PC[i].Mp[j][k])
      }
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 4; j++) {
      for (let k = 0; k < 10; k++) {
        if (i > 0 || j > 0 || k > 0) {
          data += ','
        }
        data += String(PC[i].MpMax[j][k])
      }
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 4; j++) {
      for (let k = 0; k < 10; k++) {
        for (let l = 0; l < 6; l++) {
          PC[i].Spell[j][k][l] ? (data += '1') : (data += '0')
        }
      }
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 6; j++) {
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PARTY[i].PartyMember[j])
    }
  }
  data += '\n'
  for (let i = 0; i < ITEM.length; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(ITEM[i].Stock)
  }
  data += '\n'
  for (let i = 0; i < 10; i++) {
    for (let j = 0; j < 16; j++) {
      for (let k = 0; k < 20; k++) {
        for (let l = 0; l < 20; l++) {
          PlayData['MapFlag'][i][j][k][l] ? (data += '1') : (data += '0')
        }
      }
    }
  }
  data += '\n'
  if (GoverdrySaveTweak) {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        data += BoolArrayToBinStr(PlayData['SecretDoor'][i][j], 0, 64)
      }
    }
  } else {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        for (let k = 0; k < 64; k++) {
          PlayData['SecretDoor'][i][j][k] ? (data += '1') : (data += '0')
        }
      }
    }
  }
  if (GoverdrySaveTweak) {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        data += BoolArrayToBinStr(PlayData['SecretDoor'][i][j], 64)
      }
    }
  } else {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        for (let k = 64; k < 512; k++) {
          PlayData['SecretDoor'][i][j][k] ? (data += '1') : (data += '0')
        }
      }
    }
  }
  data += '\n'
  if (GoverdrySaveTweak) {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        data += BoolArrayToBinStr(PlayData['LockedDoor'][i][j], 0, 64)
      }
    }
  } else {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        for (let k = 0; k < 64; k++) {
          PlayData['LockedDoor'][i][j][k] ? (data += '1') : (data += '0')
        }
      }
    }
  }
  if (GoverdrySaveTweak) {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        data += BoolArrayToBinStr(PlayData['LockedDoor'][i][j], 64)
      }
    }
  } else {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        for (let k = 64; k < 512; k++) {
          PlayData['LockedDoor'][i][j][k] ? (data += '1') : (data += '0')
        }
      }
    }
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += '<>'
    }
    data += PARTY[i].DungeonMusic
  }
  data += '\n'
  for (let i = 0; i < 10; i++) {
    if (i > 0) {
      data += '<>'
    }
    for (let j = 0; j < 16; j++) {
      if (j > 0) {
        data += ','
      }
      data += PlayData['DungeonNewMusic'][i][j]
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 10; j++) {
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PC[i].AbiPlus[j])
    }
  }
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 10; j < 36; j++) {
      data += ',' + String(PC[i].AbiPlus[j])
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 10; j++) {
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PC[i].AbiRate[j])
    }
  }
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 10; j < 36; j++) {
      data += ',' + String(PC[i].AbiRate[j])
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 10; j++) {
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PARTY[i].AbiPlus[j])
    }
  }
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 10; j < 36; j++) {
      data += ',' + String(PARTY[i].AbiPlus[j])
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < 10; j++) {
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PARTY[i].AbiRate[j])
    }
  }
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 10; j < 36; j++) {
      data += ',' + String(PARTY[i].AbiRate[j])
    }
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].PtSpellEffectRate[4])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].PtSpellEffectPlus[4])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectRate[6])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectPlus[6])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectRate[6])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectPlus[6])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectPlus[7])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectRate[7])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectPlus[7])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectRate[7])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    PARTY[i].ViewDarkZone ? (data += '1') : (data += '0')
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectPlus[8])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].SpellEffectRate[8])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectPlus[8])
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].SpellEffectRate[8])
  }
  data += '\n'
  if (GoverdrySaveTweak) {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        data += BoolArrayToBinStr(PlayData['WallFlag'][i][j], 0, 64)
      }
    }
  } else {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        for (let k = 0; k < 64; k++) {
          PlayData['WallFlag'][i][j][k] ? (data += '1') : (data += '0')
        }
      }
    }
  }
  if (GoverdrySaveTweak) {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        data += BoolArrayToBinStr(PlayData['WallFlag'][i][j], 64)
      }
    }
  } else {
    for (let i = 0; i < 10; i++) {
      for (let j = 0; j < 16; j++) {
        for (let k = 64; k < 512; k++) {
          PlayData['WallFlag'][i][j][k] ? (data += '1') : (data += '0')
        }
      }
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < RESIST_LENGTH; j++) {
      if (j == 9) {
        continue
      }
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PC[i].ResistPlus[j])
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < RESIST_LENGTH; j++) {
      if (j == 9) {
        continue
      }
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PC[i].ResistRate[j])
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < RESIST_LENGTH; j++) {
      if (j == 9) {
        continue
      }
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PARTY[i].ResistPlus[j])
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < RESIST_LENGTH; j++) {
      if (j == 9) {
        continue
      }
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PARTY[i].ResistRate[j])
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < ATTACK_ADD_LENGTH + 1; j++) {
      PC[i].AttackAddPlus[j] ? (data += '1') : (data += '0')
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < ATTACK_ADD_LENGTH + 1; j++) {
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PC[i].AttackAddRate[j])
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < ATTACK_ADD_LENGTH + 1; j++) {
      PARTY[i].AttackAddPlus[j] ? (data += '1') : (data += '0')
    }
  }
  data += '\n'
  for (let i = 0; i < PC_ENTRY_MIN; i++) {
    for (let j = 0; j < ATTACK_ADD_LENGTH + 1; j++) {
      if (i > 0 || j > 0) {
        data += ','
      }
      data += String(PARTY[i].AttackAddRate[j])
    }
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PC[i].PoisonPlus)
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += String(PARTY[i].PoisonPlus)
  }
  data += '\n'
  for (let i = 0; i < PlayData['GameFlagS'].length; i++) {
    PlayData['GameFlagS'][i] ? (data += '1') : (data += '0')
  }
  data += '\n'
  for (let i = 0; i < pEyMax; i++) {
    if (i > 0) {
      data += ','
    }
    data += PC[i].FaceGraphic
  }
  for (let j = 0; j < 99; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      data += String(PC[i].Equip[j])
    }
  }
  for (let j = 0; j < 99; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      data += String(PC[i].ItemDecided[j])
    }
  }
  for (let j = 0; j < 99; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PC[i].Item[j])
    }
  }
  for (let j = 0; j < 36; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      if (j < GameData['ABILITY'].length) {
        data += String(PC[i].Ability0[j])
      } else {
        data += '-1'
      }
    }
  }
  for (let k = 0; k < 10; k++) {
    for (let j = 0; j < 4; j++) {
      data += '\n'
      for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
        if (i > PC_ENTRY_MIN) {
          data += ','
        }
        data += String(PC[i].Mp[j][k])
      }
    }
  }
  for (let k = 0; k < 10; k++) {
    for (let j = 0; j < 4; j++) {
      data += '\n'
      for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
        if (i > PC_ENTRY_MIN) {
          data += ','
        }
        data += String(PC[i].MpMax[j][k])
      }
    }
  }
  for (let l = 0; l < 6; l++) {
    for (let k = 0; k < 10; k++) {
      for (let j = 0; j < 4; j++) {
        data += '\n'
        for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
          PC[i].Spell[j][k][l] ? (data += '1') : (data += '0')
        }
      }
    }
  }
  for (let j = 0; j < 6; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PARTY[i].PartyMember[j])
    }
  }
  for (let j = 0; j < 36; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PC[i].AbiPlus[j])
    }
  }
  for (let j = 0; j < 36; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PC[i].AbiRate[j])
    }
  }
  for (let j = 0; j < 36; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PARTY[i].AbiPlus[j])
    }
  }
  for (let j = 0; j < 36; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PARTY[i].AbiRate[j])
    }
  }
  for (let j = 0; j < RESIST_LENGTH; j++) {
    if (j == 9) {
      continue
    }
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PC[i].ResistPlus[j])
    }
  }
  for (let j = 0; j < RESIST_LENGTH; j++) {
    if (j == 9) {
      continue
    }
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PC[i].ResistRate[j])
    }
  }
  for (let j = 0; j < RESIST_LENGTH; j++) {
    if (j == 9) {
      continue
    }
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PARTY[i].ResistPlus[j])
    }
  }
  for (let j = 0; j < RESIST_LENGTH; j++) {
    if (j == 9) {
      continue
    }
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PARTY[i].ResistRate[j])
    }
  }
  for (let j = 0; j < ATTACK_ADD_LENGTH + 1; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      PC[i].AttackAddPlus[j] ? (data += '1') : (data += '0')
    }
  }
  for (let j = 0; j < ATTACK_ADD_LENGTH + 1; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PC[i].AttackAddRate[j])
    }
  }
  for (let j = 0; j < ATTACK_ADD_LENGTH + 1; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      PARTY[i].AttackAddPlus[j] ? (data += '1') : (data += '0')
    }
  }
  for (let j = 0; j < ATTACK_ADD_LENGTH + 1; j++) {
    data += '\n'
    for (let i = PC_ENTRY_MIN; i < pEyMax; i++) {
      if (i > PC_ENTRY_MIN) {
        data += ','
      }
      data += String(PARTY[i].AttackAddRate[j])
    }
  }
  return base64.encode(data, 1)
}
