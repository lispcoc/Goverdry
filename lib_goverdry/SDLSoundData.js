class SDLSoundData {
  constructor (handle) {
    console.log(this.constructor.name, 'constructor', handle)
    this.handle = handle
  }
  clone() {
    console.log(this.constructor.name, 'clone')
    return this
  }
  play() {
    console.log(this.constructor.name, 'play')
    MIX.PlayChannel(this.handle)
  }
}
