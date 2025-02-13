class SDLSoundData {
  constructor (handle, file) {
    console.log(this.constructor.name, 'constructor', handle)
    this.handle = handle
    this.src = { file: file, loop: true }
    this.muted = false
    this.volume = 0
    this.currentTime = 0
  }
  clone () {
    console.log(this.constructor.name, 'clone')
    return this
  }
  play () {
    console.log(this.constructor.name, 'play', this.src.file, this.volume)
    MIX.PlayChannel(this.handle, this.volume)
    if (this.isMusic()) {
      console.log('isMusic')
    }
  }
  pause () {
    console.log(this.constructor.name, 'pause')
    console.log('[WIP]not implemented.')
  }
  stop () {
    console.log(this.constructor.name, 'pause')
    console.log('[WIP]not implemented.')
  }
  isMusic () {
    return DirName['music'] && this.src.file.indexOf(DirName['music']) > -1
  }
}
