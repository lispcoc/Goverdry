class Input extends DocumentElement {
  constructor () {
    super()
    this.type = 'none'
    this.interface = null
  }
  setAttribute (a, t) {
    if (a == 'type') {
      this.type = t
    }
  }
  setSelectionRange (a, b) {}
  focus () {
    console.log(this.constructor.name, 'focus')

    if (this.type == 'text') {
      this.interface = new TextInput(this)
      this.interface.start()
    }
  }
  unfocus () {
    if (this.type == 'text') {
      OVERLAY.clear()
      this.interface = null
    }
  }
  emit (event) {
    if (this.type == 'text') {
      if (
        this.interface &&
        this.interface.processing &&
        event.id == 'keydown'
      ) {
        console.log('emit (event)')
        console.log(JSON.stringify(event))
        let key = getKey(getKeyName(event.which))
        this.interface.keyDown(key)
        if (this.interface.processing) {
          return
        }
      }
    }
    super.emit.bind(this)(event)
  }
}

class TextInput {
  constructor (parent) {
    this.parent = parent
    this.processing = true
    this.select = 0
    this.text =
      DefaultNameFemale[Math.floor(Math.random() * DefaultNameFemale.length)]
    this.charSet = KatakanaSet
  }
  start () {
    this.update()
  }
  update () {
    const ctx = OVERLAY.context
    ctx.fillStyle = 'black'
    ctx.fillRect(20, 160, GameBody.width - 40, GameBody.height - 180)

    ctx.textBaseline = 'top'
    ctx.fillStyle = 'white'
    ctx.fillText(this.text, 40, 180)

    for (var i = 0; i < this.charSet.length; i++) {
      const x = i % 10
      const y = parseInt(i / 10)
      const str = this.charSet[i]
      if (i == this.select) {
        ctx.fillStyle = 'red'
        ctx.fillRect(
          40 + x * MP.FONT_HALF_SIZE * 4,
          180 + (y + 2) * MP.FONT_SIZE,
          MP.FONT_HALF_SIZE * 2 * str.length,
          MP.FONT_SIZE
        )
      }
      ctx.fillStyle = 'white'
      ctx.fillText(
        str,
        40 + x * MP.FONT_HALF_SIZE * 4,
        180 + (y + 2) * MP.FONT_SIZE
      )
    }
  }
  keyDown (key) {
    var i = this.select
    switch (key) {
      case 'real_enter':
        if (this.charSet[i] == '決定') {
          this.parent.value = this.text
          this.processing = false
        } else if (this.charSet[i] == 'カナ') {
          this.charSet = KatakanaSet
        } else if (this.charSet[i] == 'かな') {
          this.charSet = HiraganaSet
        } else if (this.charSet[i] == '゛') {
          if (this.text.length > 0) {
            const c = this.switchDakuon(this.text[this.text.length - 1])
            this.text = this.text.slice(0, this.text.length - 1)
            this.text += c
          }
        } else if (this.charSet[i] == '゜') {
          if (this.text.length > 0) {
            const c = this.switchHandakuon(this.text[this.text.length - 1])
            this.text = this.text.slice(0, this.text.length - 1)
            this.text += c
          }
        } else if (this.charSet[i] == 'ランダム') {
          this.text =
            DefaultNameFemale[
              Math.floor(Math.random() * DefaultNameFemale.length)
            ]
        } else {
          this.text += this.charSet[i]
        }
        break
      case 'escape':
        if (this.text.length == 0) {
          this.processing = false
        } else {
          this.text = this.text.slice(0, this.text.length - 1)
        }
        break
      case 'up':
        i = i - 10
        break
      case 'down':
        i = i + 10
        break
      case 'left':
        i = parseInt(i / 10) * 10 + ((i + 9) % 10)
        break
      case 'right':
        i = parseInt(i / 10) * 10 + ((i + 1) % 10)
        break
    }
    this.select = Math.max(0, Math.min(i, HiraganaSet.length - 1))
    this.update()
  }
  switchDakuon (c) {
    if (c in Dakuon) {
      return Dakuon[c]
    }
    for (const k in Dakuon) {
      if (Dakuon[k] == c) {
        return k
      }
    }
    return c
  }
  switchHandakuon (c) {
    if (c in Handakuon) {
      return Handakuon[c]
    }
    for (const k in Handakuon) {
      if (Handakuon[k] == c) {
        return k
      }
    }
    return c
  }
}

const Dakuon = {
  か: 'が',
  き: 'ぎ',
  く: 'ぐ',
  け: 'げ',
  こ: 'ご',
  さ: 'ざ',
  し: 'じ',
  す: 'ず',
  せ: 'ぜ',
  そ: 'ぞ',
  た: 'だ',
  ち: 'ぢ',
  つ: 'づ',
  て: 'で',
  と: 'ど',
  は: 'ば',
  ひ: 'び',
  ふ: 'ぶ',
  へ: 'べ',
  ほ: 'ぼ',
  ウ: 'ヴ',
  カ: 'ガ',
  キ: 'ギ',
  ク: 'グ',
  ケ: 'ゲ',
  コ: 'ゴ',
  サ: 'ザ',
  シ: 'ジ',
  ス: 'ズ',
  セ: 'ゼ',
  ソ: 'ゾ',
  タ: 'ダ',
  チ: 'ヂ',
  ツ: 'ヅ',
  テ: 'デ',
  ト: 'ド',
  ハ: 'バ',
  ヒ: 'ビ',
  フ: 'ブ',
  ヘ: 'ベ',
  ホ: 'ボ'
}

const Handakuon = {
  は: 'ぱ',
  ひ: 'ぴ',
  ふ: 'ぷ',
  へ: 'ぺ',
  ほ: 'ぽ',
  ハ: 'パ',
  ヒ: 'ピ',
  フ: 'プ',
  ヘ: 'ペ',
  ホ: 'ポ'
}

const HiraganaSet = [
  'あ',
  'い',
  'う',
  'え',
  'お',
  'か',
  'き',
  'く',
  'け',
  'こ',
  'さ',
  'し',
  'す',
  'せ',
  'そ',
  'た',
  'ち',
  'つ',
  'て',
  'と',
  'な',
  'に',
  'ぬ',
  'ね',
  'の',
  'は',
  'ひ',
  'ふ',
  'へ',
  'ほ',
  'ま',
  'み',
  'む',
  'め',
  'も',
  'や',
  '－',
  'ゆ',
  '－',
  'よ',
  'ら',
  'り',
  'る',
  'れ',
  'ろ',
  'わ',
  '－',
  '－',
  '－',
  'を',
  'ん',
  'ー',
  'っ',
  '゛',
  '゜',
  'カナ',
  '決定',
  'ランダム'
]

const KatakanaSet = [
  'ア',
  'イ',
  'ウ',
  'エ',
  'オ',
  'カ',
  'キ',
  'ク',
  'ケ',
  'コ',
  'サ',
  'シ',
  'ス',
  'セ',
  'ソ',
  'タ',
  'チ',
  'ツ',
  'テ',
  'ト',
  'ナ',
  'ニ',
  'ヌ',
  'ネ',
  'ノ',
  'ハ',
  'ヒ',
  'フ',
  'ヘ',
  'ホ',
  'マ',
  'ミ',
  'ム',
  'メ',
  'モ',
  'ヤ',
  '－',
  'ユ',
  '－',
  'ヨ',
  'ラ',
  'リ',
  'ル',
  'レ',
  'ロ',
  'ワ',
  '－',
  '－',
  '－',
  'ヲ',
  'ン',
  'ー',
  'ッ',
  '゛',
  '゜',
  'カナ',
  '決定',
  'ランダム'
]
